// Code generated by go-bindata.
// sources:
// cloud.json
// DO NOT EDIT!

package packet

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _cloudJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x95\xdf\x6f\xda\x30\x10\xc7\xdf\xf9\x2b\x4e\x79\xce\x58\x13\x28\x42\xbc\xa5\x69\x36\xd1\xae\x5d\x54\xa8\x2a\x34\x55\x91\x49\xae\x9d\x4b\x62\x7b\xb6\xa1\x65\x15\xff\xfb\xe4\x00\x31\x0d\x29\xfd\xa1\xbd\xf0\xe3\x2e\xbe\xef\xe7\x1b\xfb\xce\xcf\x2d\x00\x87\x91\x02\x9d\x01\x38\x82\xa4\x33\xd4\x8e\x6b\x62\xc8\x16\xca\x19\xc0\xaf\x16\x00\x80\x93\xe1\xa2\x0c\x03\x38\x7f\xc8\xf6\x97\x90\x3c\x73\x5a\x00\xb7\xe5\x02\x89\xf7\x94\x33\xbb\xe6\xb9\xfc\x04\x70\x72\x9e\x12\x4d\x39\x33\x0a\x97\x93\x10\x2e\x50\x4b\xee\xc2\xe5\xc4\x85\xeb\x51\xb0\x29\x56\x15\x30\x4f\x5d\x8f\x20\x22\x4a\xdb\xd4\x5f\xce\xd0\x56\x2e\x43\xf8\x28\x3d\x67\xf3\xf7\xb6\xfc\x5e\xb9\xaf\xeb\x8e\x08\x83\x33\xae\xd0\x85\x30\x38\x24\x7b\x83\x6f\xc8\xaa\x87\xf4\x03\xb2\x41\xa1\x34\xca\x8c\x14\x2e\x5c\xfe\x68\x92\x8c\xae\xe1\x26\x1a\x8d\x0f\x4a\x92\x42\x7d\x40\x72\xcc\x67\x4b\xee\xc2\x59\xdc\x24\x17\x8c\x86\x01\x44\xc1\x1b\x82\x4c\xea\xba\x60\xb5\xcb\x94\x29\x4d\x58\x8a\xe3\xa5\xc0\x86\xbd\x56\xb3\xb9\xd1\x99\x12\x89\x05\x6a\x92\x27\x47\x56\x29\x43\x95\x4a\x2a\x2a\xd0\x49\x1c\xc1\xd1\x00\xba\xf0\xd4\xef\x41\xaf\x3b\xa5\x1a\x42\x2e\x51\xb9\xd0\xff\x7e\x02\xa7\xa7\x57\x1d\xb8\x0a\x2e\xec\xfa\x94\x68\xbc\xe7\x72\x69\x16\x9f\x10\x89\xe6\x20\x91\x7c\x27\x2f\x8c\x76\xd7\xda\x26\x85\x33\x80\xbe\xd5\xa7\x6a\x66\x02\x47\x8d\xef\x70\x8f\xdc\x3b\x4c\xee\x35\x92\x77\xfc\xff\x88\xde\xf1\xeb\xec\x9e\xff\x4e\x78\xff\x30\xbc\x3f\x00\xbf\x81\xde\x3f\xee\xad\xf1\xbb\x10\x85\xe1\x67\x2c\xf8\x75\x0f\xfe\x71\xaf\x6e\xc2\xef\xbf\x77\x0b\x3a\x87\x5d\x74\x06\xe0\xf5\xf6\x5d\x78\x7e\x7f\xeb\xe2\x13\x0e\xbc\x5e\xcd\x81\xe7\xef\x1d\xa1\x6a\x1b\xaa\xb6\x48\x25\x66\xc8\x34\x25\x79\x43\x53\x08\xc9\x17\x34\x43\x69\x84\x63\x3b\x62\xb7\x15\x45\x4e\x96\xdf\xb8\x2c\x88\x36\x0f\xdc\x51\xcc\x33\x9b\x27\x8c\x71\x5d\x36\xb7\x29\xfc\x6c\x9b\x54\xfc\x26\xb2\x40\xd9\x26\x42\xa8\x94\x67\xd8\x4e\x79\xf1\x35\xcd\xe7\x66\xe0\x7c\xb1\x38\xa6\xe4\xb6\x97\x57\x55\xd5\x52\xe4\x65\xd7\xdb\xd2\xeb\xe1\x9f\x72\x76\x47\xef\x4b\xe4\x20\x3c\x8f\xc6\x49\x7c\xf5\xf3\x2c\x0a\xc7\xc9\xf0\xb4\xa2\x5b\xd7\xe2\xb2\xb0\x97\x47\x22\x24\x7f\xc0\x54\x27\x34\x7b\xf9\xd8\x83\x5a\xef\xdc\x26\x5f\xaf\x92\x93\x29\x96\xb0\xf1\x3a\x0f\xc3\xda\x7a\xca\xc4\xbc\x7c\x41\x1a\x9f\xb4\x53\x65\x56\xee\xfb\x1d\x04\xf1\x30\x39\x8f\x26\x07\xf1\x89\xa0\xc9\x0c\x97\xcd\xec\x44\xd0\xf3\x7a\xae\x02\x0f\xe2\x21\xec\x65\x2b\x6a\x41\x94\x7a\xe4\x32\xdb\x21\x7f\x65\xc2\xce\xe6\x53\x94\x0c\x75\xd3\x78\x5d\xa0\x54\x9b\x0e\xf0\xda\xfd\xf6\xeb\xc3\xb5\x96\xdd\x5c\xe6\x3b\xc7\xc7\x5c\xe8\x03\xd0\x72\x8e\x96\xd7\x5c\xed\x7b\xb1\xf2\x92\x5f\x47\x5b\xbb\xe0\x25\x70\x6b\xd5\xfa\x17\x00\x00\xff\xff\xfd\x56\x3f\xd9\x41\x08\x00\x00")

func cloudJsonBytes() ([]byte, error) {
	return bindataRead(
		_cloudJson,
		"cloud.json",
	)
}

func cloudJson() (*asset, error) {
	bytes, err := cloudJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cloud.json", size: 2113, mode: os.FileMode(420), modTime: time.Unix(1453795200, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"cloud.json": cloudJson,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"cloud.json": {cloudJson, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
