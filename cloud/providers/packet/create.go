package packet

import (
	"fmt"
	"time"

	proto "github.com/appscode/api/kubernetes/v1beta1"
	"github.com/appscode/errors"
	"github.com/appscode/pharmer/api"
	"github.com/appscode/pharmer/cloud"
	"github.com/appscode/pharmer/storage"
	"github.com/cenkalti/backoff"
	"github.com/packethost/packngo"
)

func (cm *clusterManager) create(req *proto.ClusterCreateRequest) error {
	err := cm.initContext(req)
	if err != nil {
		cm.ctx.StatusCause = err.Error()
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	cm.ins, err = cloud.NewInstances(cm.ctx)
	if err != nil {
		cm.ctx.StatusCause = err.Error()
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	cm.conn, err = NewConnector(cm.ctx)
	if err != nil {
		cm.ctx.StatusCause = err.Error()
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}

	defer func(releaseReservedIp bool) {
		if cm.ctx.Status == storage.KubernetesStatus_Pending {
			cm.ctx.Status = storage.KubernetesStatus_Failing
		}
		cm.ctx.Save()
		cm.ins.Save()
		cm.ctx.Logger().Infof("Cluster %v is %v", cm.ctx.Name, cm.ctx.Status)
		if cm.ctx.Status != storage.KubernetesStatus_Ready {
			cm.ctx.Logger().Infof("Cluster %v is deleting", cm.ctx.Name)
			cm.delete(&proto.ClusterDeleteRequest{
				Name:              cm.ctx.Name,
				ReleaseReservedIp: releaseReservedIp,
			})
		}
	}(cm.ctx.MasterReservedIP == "auto")

	cm.ctx.InstanceImage = "debian_8"

	err = cm.importPublicKey()
	if err != nil {
		cm.ctx.StatusCause = err.Error()
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}

	// -------------------------------------------------------------------ASSETS
	im := &instanceManager{ctx: cm.ctx, conn: cm.conn}

	cm.ctx.Logger().Info("Creating master instance")
	masterDroplet, err := im.createInstance(cm.ctx.KubernetesMasterName, api.RoleKubernetesMaster, cm.ctx.MasterSKU)
	if err != nil {
		cm.ctx.StatusCause = err.Error()
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	if err = cm.conn.waitForInstance(masterDroplet.ID, "active"); err != nil {
		cm.ctx.StatusCause = err.Error()
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	// Need to reload te master server Object to get the IP address
	masterInstance, err := im.newKubeInstance(masterDroplet.ID)
	if err != nil {
		cm.ctx.StatusCause = err.Error()
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	masterInstance.Role = api.RoleKubernetesMaster
	cm.ctx.MasterExternalIP = masterInstance.ExternalIP
	cm.ctx.MasterInternalIP = masterInstance.InternalIP
	fmt.Println("Master EXTERNAL IP ================", cm.ctx.MasterExternalIP)
	cm.ins.Instances = append(cm.ins.Instances, masterInstance)

	if err = cloud.GenClusterCerts(cm.ctx); err != nil {
		cm.ctx.StatusCause = err.Error()
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	err = cloud.EnsureARecord(cm.ctx, masterInstance) // works for reserved or non-reserved mode
	if err != nil {
		cm.ctx.StatusCause = err.Error()
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	cm.ctx.DetectApiServerURL()
	// needed to get master_internal_ip
	if err = cm.ctx.Save(); err != nil {
		cm.ctx.StatusCause = err.Error()
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	cm.UploadStartupConfig()

	// -------------------------- NODES

	// reboot master to use cert with internal_ip as SANS
	time.Sleep(60 * time.Second)

	cm.ctx.Logger().Info("Rebooting master instance")
	if err = im.reboot(masterDroplet.ID); err != nil {
		cm.ctx.StatusCause = err.Error()
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	cm.ctx.Logger().Info("Rebooted master instance")

	// start nodes
	for _, ng := range req.NodeGroups {
		for i := int64(0); i < ng.Count; i++ {
			droplet, err := im.createInstance(cm.namer.GenNodeName(), api.RoleKubernetesPool, ng.Sku)
			if err != nil {
				cm.ctx.StatusCause = err.Error()
				return errors.FromErr(err).WithContext(cm.ctx).Err()
			}

			// record nodes
			cm.conn.waitForInstance(droplet.ID, "active")
			node, err := im.newKubeInstanceFromServer(droplet)
			if err != nil {
				cm.ctx.StatusCause = err.Error()
				return errors.FromErr(err).WithContext(cm.ctx).Err()
			}
			node.Role = api.RoleKubernetesPool

			cm.ins.Instances = append(cm.ins.Instances, node)
		}
	}

	cm.ctx.Logger().Info("Waiting for cluster initialization")

	// Wait for master A record to propagate
	if err := cloud.EnsureDnsIPLookup(cm.ctx); err != nil {
		cm.ctx.StatusCause = err.Error()
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}

	// wait for nodes to start
	if err := cloud.ProbeKubeAPI(cm.ctx); err != nil {
		cm.ctx.StatusCause = err.Error()
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	// check all components are ok
	if err = cloud.CheckComponentStatuses(cm.ctx); err != nil {
		cm.ctx.StatusCause = err.Error()
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}
	// Make sure nodes are connected to master and are ready
	if err = cloud.WaitForReadyNodes(cm.ctx); err != nil {
		cm.ctx.StatusCause = err.Error()
		return errors.FromErr(err).WithContext(cm.ctx).Err()
	}

	cm.ctx.Status = storage.KubernetesStatus_Ready
	return nil
}

func (cm *clusterManager) importPublicKey() error {
	cm.ctx.Logger().Debugln("Adding SSH public key")
	backoff.Retry(func() error {
		sk, _, err := cm.conn.client.SSHKeys.Create(&packngo.SSHKeyCreateRequest{
			Key:       string(cm.ctx.SSHKey.PublicKey),
			Label:     cm.ctx.Name,
			ProjectID: cm.ctx.Project,
		})
		cm.ctx.SSHKeyExternalID = sk.ID
		return err
	}, backoff.NewExponentialBackOff())
	cm.ctx.Logger().Debugf("Created new ssh key with fingerprint=%v", cm.ctx.SSHKey.OpensshFingerprint)
	return nil
}