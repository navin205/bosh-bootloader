// Code generated by go-bindata.
// sources:
// templates/network_security_group_template.tf
// templates/network_template.tf
// templates/output_template.tf
// templates/resource_group_template.tf
// templates/storage_template.tf
// templates/vars_template.tf
// DO NOT EDIT!

package azure

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

var _templatesNetwork_security_group_templateTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x95\xcd\x8e\x9b\x30\x10\x80\xef\x7e\x8a\x91\xd5\xd3\x4a\x89\xb6\x6c\x58\xed\x65\x0f\x3d\xf6\xde\x3b\xf2\x9a\x81\xb5\x4a\x3c\x68\x6c\x92\xb6\x11\xef\x5e\x99\x9f\x28\x50\xd2\x94\xa8\xad\x44\x14\xae\x7c\xf3\xff\x49\x66\x74\x54\xb1\x46\x90\xea\x47\xc5\xc8\xdb\xc4\xa2\xdf\x13\x7f\x4d\x1c\xea\x8a\x8d\xff\x9e\xe4\x4c\x55\x29\x41\xbe\x91\x7b\x97\x70\x10\x00\x56\x6d\x11\x46\xdf\x2b\xc8\x0f\x87\x9d\xe2\x35\xda\x5d\x62\xd2\x7a\xd5\xe0\x02\xa0\x20\xad\xbc\x21\x3b\x09\xf7\x3f\xeb\x40\xf6\xbd\xb4\x15\x93\xa6\x4a\x43\xf6\xad\x0d\x81\x75\xa8\xb0\x0e\x54\x2d\x85\x00\xf0\x2a\x77\x4d\x7b\x00\x68\x77\x86\xc9\x6e\xd1\xfa\x5f\x1a\x0b\x95\x6a\x51\x0b\x31\x63\x74\x9d\xcd\x18\x5c\x67\x4b\x1f\x9b\xab\x02\x25\x48\xf7\xbb\x7b\x9f\x9b\xde\xb5\x57\x2f\xd9\x50\xc8\x35\x19\x12\x3d\x3e\x0a\x80\xd4\x30\xea\xf1\x8a\x8e\x59\x3f\xdb\x37\xaa\x6c\x1a\x72\x29\xad\xd1\xb9\x73\xe5\x3f\x15\x05\xed\xdb\x92\xe4\x49\x53\x31\x8d\x7d\xd1\x65\x80\xba\x4d\x96\xc4\x3e\x61\x65\x73\x1c\x40\x0f\x01\x49\xd1\x79\x63\x9b\xfb\x8c\xb9\x57\x90\x51\x74\x92\x46\xa5\x29\xa3\x73\x49\xc9\x98\x99\x6f\xe7\xd3\x8c\xb8\x1e\x99\xba\xfc\x60\xb1\x7f\x60\x00\xc0\xb4\xb5\x13\x1e\x4d\x83\x83\x6c\xb3\xfc\x08\x81\x2b\x95\xa3\xf5\xf3\x35\x39\x89\xbd\x6c\xcb\xc7\xc5\xda\xf2\xfc\xf2\xfc\x72\xf7\xe5\xd4\x97\xf6\x8e\xc4\x57\x2a\x73\x0c\xbf\x6c\x4d\xb4\x58\x6b\xa2\x38\x8e\xe3\xbb\x36\x9d\x36\xa9\x75\xf3\x65\x09\x41\x97\x15\x79\xfa\xef\x8a\x3c\xfc\x15\x41\xe2\xa7\xbb\x1d\x9d\x1d\x3a\x5b\xbd\x7b\x5f\xfe\x33\x45\x96\xfb\xf6\x6c\x36\x37\x67\x89\xce\xae\x75\xa4\xa0\x7c\xbe\x21\x5d\xdc\x2d\x3f\x35\x9b\x1b\xb7\xe4\x67\x00\x00\x00\xff\xff\x3b\xed\x87\x70\xe0\x0e\x00\x00")

func templatesNetwork_security_group_templateTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesNetwork_security_group_templateTf,
		"templates/network_security_group_template.tf",
	)
}

func templatesNetwork_security_group_templateTf() (*asset, error) {
	bytes, err := templatesNetwork_security_group_templateTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/network_security_group_template.tf", size: 3808, mode: os.FileMode(420), modTime: time.Unix(1506449713, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesNetwork_templateTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x91\xcb\x4e\xc5\x20\x10\x40\xf7\x7c\xc5\x64\xe2\xb6\x68\x37\xee\xfc\x12\x63\x08\x6d\x47\x25\xb6\xd0\x0c\x0f\x8d\x0d\xff\x6e\x68\x52\xb5\xdc\x36\xf7\xc2\x92\xc3\xe4\x9c\x0c\x93\x77\x91\x7b\x02\xd4\xdf\x91\x89\x27\x95\x0c\x87\xa8\x47\x65\x29\x7c\x3a\xfe\x40\xc0\xce\xf9\x77\x84\x45\x00\x58\x3d\x11\x54\xe7\x09\xf0\x6e\x49\x9a\x25\xd9\xa4\xcc\x90\x9b\x82\x37\xc9\xa2\x00\xd0\xc3\xc0\xe4\xbd\xf2\xb3\xee\xe9\x97\x7f\xc6\xf6\x41\xae\xf7\xbe\x7d\xc4\x17\x01\x30\xba\x5e\x07\xe3\xec\xe1\xdc\xed\x31\x97\x89\x9b\xaf\x7a\x63\x17\x67\xb5\x0a\xad\xe4\xa6\xbf\x07\x64\x91\x91\x85\xca\x28\xb2\x10\x97\xb9\x3e\x76\x96\xc2\xd5\xca\x93\x4c\xbf\xcb\x9c\x99\x5e\xcd\xd7\xdf\x87\xff\x99\x27\xee\x37\xcb\x03\x54\x8b\x39\x68\xaf\x88\x2a\xfe\x27\x00\x00\xff\xff\x03\x56\xce\x9c\xeb\x01\x00\x00")

func templatesNetwork_templateTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesNetwork_templateTf,
		"templates/network_template.tf",
	)
}

func templatesNetwork_templateTf() (*asset, error) {
	bytes, err := templatesNetwork_templateTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/network_template.tf", size: 491, mode: os.FileMode(420), modTime: time.Unix(1506449671, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesOutput_templateTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xd1\x4d\x6a\x03\x31\x0c\x05\xe0\x75\xe7\x14\x66\xe8\x3a\x81\xc2\x6c\x02\x3d\x8b\xd0\x78\xd4\xc4\xd4\xb1\x8c\x2c\xa5\x3f\x21\x77\x2f\x26\x9d\x82\x17\xf5\x64\xed\xf7\x3e\x3f\x10\x9b\x66\x53\x37\xce\x5c\x4e\x90\x48\x3f\x58\xde\x21\xe1\x99\x46\x77\x1d\x9c\x73\xee\x82\xd1\xc8\xbd\xba\xf1\xf9\x8a\xdf\x26\x24\x67\xb8\x04\x51\xc3\xb8\xc6\x77\xb5\xbb\xab\x9d\xdb\x38\xdc\x86\xa1\x21\x8b\xcd\x89\x74\x4b\xbc\xa7\xba\x90\x50\x61\x13\x4f\x70\x14\xb6\xbc\x05\xb6\xe9\xfe\x42\x65\xc1\x23\x01\x7a\xcf\x96\xb6\xa7\xb6\xf1\x2e\xbd\xd0\x1b\x5a\x54\x28\xe4\x4d\x82\x7e\xdd\xd7\x74\xf0\xf5\x00\x6d\xe1\xbf\x3f\xe8\x53\x49\x12\x46\x08\x3d\x33\xdb\x1c\x83\x87\xf0\xcb\x84\x0c\xb8\x2c\x42\xa5\xb4\xd8\x12\x84\xbc\xb2\xac\xaf\x55\x7c\xfa\xe3\x4e\xaa\xb9\x1c\xf6\xfb\x47\xd8\xc3\xcb\x34\x4d\x53\xc5\x7f\x02\x00\x00\xff\xff\x5c\x02\xb8\xbf\x5d\x02\x00\x00")

func templatesOutput_templateTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesOutput_templateTf,
		"templates/output_template.tf",
	)
}

func templatesOutput_templateTf() (*asset, error) {
	bytes, err := templatesOutput_templateTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/output_template.tf", size: 605, mode: os.FileMode(420), modTime: time.Unix(1506449731, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesResource_group_templateTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x90\x41\x6a\xc5\x30\x0c\x44\xf7\x3e\xc5\x60\xba\x6d\x6e\xd0\xb3\x18\xc5\x11\xa9\x21\xb1\x83\x6c\x67\xd1\xe0\xbb\x17\x1b\xd2\x26\x24\xff\x7f\xf8\x5a\x8a\xd1\x68\xe6\x09\xc7\x90\xc5\x32\x34\xfd\x64\x61\x99\xcd\xbe\x31\xa3\x84\xbc\x68\xe8\x3e\xc4\x6f\x8d\x4d\x01\x9e\x66\x46\x9d\x2f\xe8\x8f\x6d\x25\xe9\xd8\xaf\xc6\x0d\xe5\xb3\x69\x14\x30\x05\x4b\xc9\x05\xff\xaf\xd8\x37\x45\x2b\x05\x24\x1a\x63\xb3\x02\xd8\xaf\x4e\x82\x9f\xd9\xa7\x8b\x5f\xb5\x2a\xaa\x28\x75\x8d\xb7\xe4\x7e\x72\xd6\xb8\x07\xc9\xee\xe6\x75\xda\xa7\x57\x87\x06\xc0\x99\x8e\x39\xff\x6d\x27\xf7\x1c\xbb\xfa\xb3\xab\xf2\x66\xf3\xd7\xc2\xd0\x30\x08\xc7\x68\x68\x3a\xb2\x8b\x89\x92\xb3\xef\x20\xfb\x0d\x00\x00\xff\xff\xe9\x3c\x7f\x17\xd1\x01\x00\x00")

func templatesResource_group_templateTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesResource_group_templateTf,
		"templates/resource_group_template.tf",
	)
}

func templatesResource_group_templateTf() (*asset, error) {
	bytes, err := templatesResource_group_templateTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/resource_group_template.tf", size: 465, mode: os.FileMode(420), modTime: time.Unix(1506449653, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesStorage_templateTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x91\x41\x6e\x83\x40\x0c\x45\xf7\x73\x0a\x6b\xd4\x75\x6e\xd0\x75\xf7\xcd\x01\x46\x66\xb0\xe8\x48\x60\x23\xdb\x50\xb5\x11\x77\xaf\x18\x05\xda\x20\x45\xc9\xb2\x6c\xf9\xfe\x7e\xcf\xa3\x64\x32\x69\x26\x88\xf8\x3d\x29\xe9\x90\xcc\x45\xb1\xa3\x84\x39\xcb\xc4\x1e\x21\x36\x62\x1f\x11\x2e\x01\x80\x71\x20\x38\x7c\xaf\x10\x5f\x2e\x33\xea\xc9\xca\x30\xf6\x94\x88\xe7\x54\xda\x25\x06\x80\xad\x3c\x75\x2a\xd3\x98\xea\x74\x8d\x6f\xbb\x6e\x03\xa7\x75\xd1\x69\x4d\x2d\x31\x04\x80\x5e\x32\x7a\x11\xde\xd6\x7c\x92\xf9\x64\x6b\xf1\x95\x2d\xf9\xd7\x58\x1b\xcf\x8e\xdc\xa2\xb6\xe9\xed\xfd\x5c\x47\x1d\x3b\xab\xc4\x00\xc4\x73\x51\xe1\x81\xd8\x7f\x59\xff\x40\x2e\x61\x09\xe1\xfe\x19\xb2\xb0\x63\x61\xd2\x87\x87\xa8\x8c\x35\x72\x47\x1d\x9e\x96\x07\x38\xbc\xc2\xb5\xe0\x66\xfe\x10\x39\x14\xec\xdc\xeb\x7f\x32\xdb\x6f\x35\x6a\x99\xd1\x29\x3e\xaf\x6d\x4e\x43\xa6\xbe\x7f\xa0\xbe\xc7\xfe\xb5\x7e\xd3\x4b\xb3\xba\xff\x04\x00\x00\xff\xff\xcf\x0a\x89\xf7\xf9\x02\x00\x00")

func templatesStorage_templateTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesStorage_templateTf,
		"templates/storage_template.tf",
	)
}

func templatesStorage_templateTf() (*asset, error) {
	bytes, err := templatesStorage_templateTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/storage_template.tf", size: 761, mode: os.FileMode(420), modTime: time.Unix(1506449688, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesVars_templateTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\xcf\xd1\xaa\xc2\x30\x0c\x06\xe0\xeb\xd3\xa7\x08\xe5\x5c\xfb\x06\x3e\xcb\xc8\xba\x20\x81\xae\x2b\x69\x56\xd0\xb1\x77\x97\x3a\xa9\x32\x15\x6b\x6e\xff\x2f\xe1\x4f\x46\x61\xec\x3d\x81\xa5\x90\x3b\x1e\x2c\x2c\xe6\x4f\xcf\x91\xe0\x08\x36\xa9\x70\x38\x59\xb3\x1a\xf3\x70\x7e\x72\xa8\x3c\x85\xef\x32\xf1\x18\x3d\x75\xad\x87\xd3\xdc\x27\x27\x1c\xcb\xf1\xa6\x05\xa5\x80\x41\x9b\xa8\xf3\x4c\xbf\xd1\x44\x4e\x48\x3f\xf1\x28\x53\xe6\x81\x04\x2c\x5e\x66\x21\x19\x0b\x04\xd8\xbd\x00\x65\xeb\x7f\xc9\x28\x87\x5d\xb2\x5a\x03\x50\xfb\xc3\x7d\xaa\xae\xc9\xcd\xd5\xf2\x2f\xae\x26\xcf\x6e\x6b\xfe\xce\x6d\xc9\x5a\x3e\xb8\x06\x00\x00\xff\xff\xb7\xcd\x6b\x24\xf8\x01\x00\x00")

func templatesVars_templateTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesVars_templateTf,
		"templates/vars_template.tf",
	)
}

func templatesVars_templateTf() (*asset, error) {
	bytes, err := templatesVars_templateTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/vars_template.tf", size: 504, mode: os.FileMode(420), modTime: time.Unix(1506449632, 0)}
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
	"templates/network_security_group_template.tf": templatesNetwork_security_group_templateTf,
	"templates/network_template.tf": templatesNetwork_templateTf,
	"templates/output_template.tf": templatesOutput_templateTf,
	"templates/resource_group_template.tf": templatesResource_group_templateTf,
	"templates/storage_template.tf": templatesStorage_templateTf,
	"templates/vars_template.tf": templatesVars_templateTf,
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
	"templates": &bintree{nil, map[string]*bintree{
		"network_security_group_template.tf": &bintree{templatesNetwork_security_group_templateTf, map[string]*bintree{}},
		"network_template.tf": &bintree{templatesNetwork_templateTf, map[string]*bintree{}},
		"output_template.tf": &bintree{templatesOutput_templateTf, map[string]*bintree{}},
		"resource_group_template.tf": &bintree{templatesResource_group_templateTf, map[string]*bintree{}},
		"storage_template.tf": &bintree{templatesStorage_templateTf, map[string]*bintree{}},
		"vars_template.tf": &bintree{templatesVars_templateTf, map[string]*bintree{}},
	}},
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

