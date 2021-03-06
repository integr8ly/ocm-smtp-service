// Code generated by go-bindata.
// sources:
// ../../openapi/ocm-smtp-service.yaml
// DO NOT EDIT!

package openapi

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

var _ocmSmtpServiceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x59\x4d\x6f\x1b\x37\x13\xbe\xeb\x57\x0c\xf0\xbe\x85\x5a\xc0\x96\xe4\xa4\x3d\x54\x40\x0e\xf9\x68\x80\x04\x71\x9c\xc6\x4e\x7b\x28\x0a\x7b\xb4\x1c\x69\x19\xef\x92\x9b\xe1\xac\x6d\x05\xfd\xf1\x05\x49\xad\xb4\x5a\xad\xbe\x5c\x17\x4a\xd0\xe8\x24\x52\xc3\xe1\x33\x5f\x0f\x39\x94\x2d\xc8\x60\xa1\x87\xf0\xb8\x37\xe8\x0d\x3a\xda\x8c\xed\xb0\x03\x20\x5a\x32\x1a\xc2\xd9\xf3\x53\x38\x3f\xbd\x78\x07\xe7\xc4\x37\x3a\x21\x78\xfa\xee\x55\x07\x40\x91\x4b\x58\x17\xa2\xad\x19\xc2\x29\x1a\x9c\x90\x8b\x72\x09\x93\x22\x23\x1a\x33\xd7\x01\xb8\x21\x76\x41\x66\xd0\x1b\xf4\x4e\x3a\x8e\xd8\xcf\x78\xfd\xc7\x50\x72\x36\x84\x54\xa4\x70\xc3\x7e\x1f\x0b\xdd\xf3\x48\x5c\xaa\xc7\xd2\x4b\x6c\xde\x01\x58\xd9\x46\x1b\xf8\xbe\x60\xab\xca\xc4\xcf\xfc\x00\x51\x5d\xbb\x32\x27\x38\xa1\x6d\x2a\xcf\x05\x27\xda\x4c\x56\x14\xf5\x57\x45\x93\x92\x99\x8c\x80\xb2\x39\x6a\xd3\x29\x50\xd2\x60\x87\xdf\xac\x6f\x93\xfc\xd8\xe5\x52\x1c\xbb\xe8\xa5\xfe\xcd\x49\xdf\x8f\x87\x41\xcf\x84\x24\x7e\x01\x70\x65\x9e\x23\x4f\x87\xf0\x9e\xa4\x64\xe3\x00\x21\xd3\x4e\xc0\x8e\x83\xf7\x2a\x29\x4a\x4a\xd6\x32\xad\x56\x79\x5c\xcf\x08\x99\x78\x08\x7f\xfc\x39\x9b\x64\x72\x85\x35\x8e\xdc\x42\xaa\xfb\x68\x30\xe8\x2e\x86\x0d\x0b\x9e\xc2\xeb\xf3\xb3\xb7\x80\xcc\x38\xad\x36\x04\x3b\xfa\x48\x89\xb8\xda\x9a\xc4\x1a\x21\x23\x75\x35\x00\x58\x14\x99\x4e\xd0\x2b\xea\x7f\x74\xd6\x2c\xff\x0a\xe0\x92\x94\x72\x6c\xce\x02\xfc\x9f\x69\x3c\x84\xee\xff\xfa\x89\xcd\x0b\x6b\xc8\x88\xeb\x47\x59\xd7\xf7\xfb\xbf\xd1\x4e\xba\x0b\xfc\x3f\x0e\x4e\x36\xe0\x2f\x25\x05\xb1\xd7\x64\x40\x3b\xd0\xe6\x06\x33\xad\x0e\x01\xfc\x17\x66\xcb\x4b\xa8\x1f\xaf\x47\xfd\xc1\x60\x29\xa9\x65\xfd\x99\x14\x88\x85\x82\x78\x6c\x39\x07\x5b\x10\x07\x58\x5f\x82\x05\x3f\x6d\xca\x9b\x0f\x86\xee\x0a\x4a\x84\x14\x90\x5f\x07\x36\x09\xc5\x70\x78\xdf\x17\xc8\x98\x93\xcc\x18\x25\x7e\x8e\x5b\x97\x2f\x24\xfb\x05\x4e\xa8\xbb\xbb\xb8\xd3\x9f\xf7\x12\x27\xe4\x24\xdd\x63\x81\x65\x45\xfc\x6c\x1a\x57\x14\xd6\xad\x52\xc5\x73\x26\x14\x02\x04\x43\xb7\x4d\x8e\xdd\x8f\x30\x3e\x95\xe4\xe4\x99\x55\x35\xb9\xa5\x48\x3f\xcf\x4a\x27\xc4\xa7\x24\x08\x0a\x05\xe7\x52\x7e\xa9\x66\x52\x43\x10\x2e\xa9\xb3\x21\xea\x9b\x63\xde\x1e\xf1\x6d\x24\xd1\xdd\xc8\x78\x1b\x18\x23\x7a\xee\x20\x79\x5a\xc7\x1d\x28\x62\x43\x81\xfd\xe6\x89\x2c\x6c\x1f\x0b\xcc\x7d\x39\x15\xf6\x8d\x93\x0f\x68\xc1\xcf\xeb\x2d\x68\xd0\x00\x60\xc6\x84\x6a\x0a\x74\xa7\xdd\x61\x4e\xf3\xbd\x8e\x94\xa7\x06\xca\x75\xa7\x8a\x37\x0b\xc5\xdf\xca\x24\xa5\x35\x7c\x77\x28\xcb\x14\x65\x24\xb4\xc2\xd0\x2f\xc2\x34\xa0\x59\x89\xcb\xd8\x32\x20\x24\x91\x57\xff\x13\x64\x1d\x7d\xf1\x3e\xc2\xdf\xcc\xdc\x8f\xd6\x27\x48\xd4\x72\x10\x26\xa9\xf9\xf5\x1b\x81\x7f\x23\xf0\x7f\x62\xc1\x3d\x08\x5c\x1d\x2e\xf1\x1f\x8e\xc1\x83\x11\x15\x83\xcf\xea\xe9\x0b\x62\xf2\xc5\x2f\x7e\x79\x45\xc7\xe7\x5e\xa8\x62\xa8\x19\x1f\xcf\xb4\xcb\xb4\xa0\xf8\xba\xd0\xa9\x6d\x4e\x43\x18\x05\xb1\xd9\x64\x1c\xbc\xb4\x9c\xa3\x0c\xe1\xf5\xef\x17\x9d\x0a\xe5\x4c\xe9\x59\xe8\xb5\xdf\xd3\x98\x98\x4c\x42\xcb\xda\x63\x23\x5e\x35\x36\xec\x33\x59\x74\x9d\x30\xb5\xaa\x1b\x1b\x17\x39\x61\x6d\x26\xf3\xe9\x6b\x6d\xb6\x0b\xa5\xde\x41\x9b\x84\x7c\x4f\xbe\x27\xb6\x9d\x36\xf6\x8d\xd7\xaa\x90\x36\x42\x93\xb9\x0f\x01\x7c\xbf\xb5\x5d\x4a\xac\x60\xb6\x4d\x6c\x7e\x10\xd6\x0e\x5a\x8f\xb4\x36\xf4\x98\x6a\x43\xbf\x79\x6d\x18\x76\xa9\x8d\xb5\x50\x1e\x2f\x57\x21\x93\x2a\xbd\x98\x65\x67\xe3\x6d\x1d\x68\x95\x83\x8d\x24\xa8\x77\x8a\x2d\xce\x5e\xe7\x70\x08\x65\xa3\xa8\x99\xff\xad\x8e\x8f\xbe\xc0\x96\x22\x5a\x2b\x3e\x27\xd2\xcb\xe5\xb4\x6b\x5d\x14\x9c\x51\xcf\x9a\xbd\x1c\xb2\xfc\x04\xb4\xb7\x17\x42\x4c\xda\x21\x86\x97\xae\xc6\x2f\xad\xe2\x3b\xf3\x62\xed\x6e\x70\xe8\xe0\x47\x24\xaf\x5e\x6c\x8d\x8e\xe7\xdd\xaf\x06\x6c\xf5\x49\x6b\x0f\x21\x5b\x85\x0b\xcb\xbb\x0b\x4b\xb6\x26\x5b\x5a\x64\x4b\x47\x6c\x30\xdf\xbd\xca\x0a\x74\xee\xd6\xf2\xf6\x92\x99\x7b\x26\xbe\x50\x5c\xe2\xee\x06\x80\x6f\x2a\xc2\x11\xa3\x50\xe8\x58\x74\x4e\xcb\xa0\x0b\xf5\x50\x2a\x57\xee\xf3\x7b\x1e\x0b\xad\x71\x6f\x4d\xd0\xaf\x9b\x3d\x1a\x88\x17\x9f\xcd\xd8\xab\x37\xa2\xe6\x33\xe6\x82\x72\x43\xf2\xc1\xfc\x42\xae\xcd\x10\x0a\x94\x74\x36\x5c\xba\x90\x5d\xa4\x04\x5a\x81\x1d\x03\x53\x62\x59\x35\xcf\xc0\x7a\x33\xd8\xbc\x3b\xad\x84\xa4\x7e\x54\x47\x0c\xb5\x83\xd2\xa3\xf8\x54\x12\x4f\xdb\x60\xbc\xc3\x09\x81\x29\xf3\x11\xf1\x02\x4b\xfc\x8f\xe3\x36\x25\xb3\x34\x41\x77\x09\x91\x72\xe0\x0a\x4a\xf4\x58\x93\x0a\xbb\xd4\x0f\xe1\x76\xa0\xcd\xcb\x80\xa2\x31\x96\x99\x0c\xe1\x64\x3e\x95\x6b\xa3\xf3\x32\x5f\x4c\x2d\xfc\x30\xc6\xcc\x45\xfd\xf5\xab\x46\xb4\xb2\xb6\xf5\x46\x2b\x4f\xf1\xce\xab\x5f\x31\xd4\xf9\x66\x84\xc3\x5f\x3b\xf7\xb4\x60\x30\x58\xb5\x61\xb0\xc9\x86\xf0\xde\xdc\xb0\x22\xcc\xad\xb1\xa3\x4d\x49\xc3\xba\xbf\x8e\xe7\x18\xce\x67\xa1\x71\xe1\x2a\x1f\x15\x43\xc2\x5a\x88\x35\xf6\x42\xd2\xb9\xa9\x11\xbc\xf3\x3e\x90\x54\xbb\x45\x32\x83\x76\xb5\x4b\x5d\xae\x33\x64\xef\x1d\x69\x2c\x21\xb8\xbc\x4d\x89\xe9\x12\x92\x0c\x4b\x47\x7e\x16\x0d\x9c\xff\xfa\x06\x9c\xa0\x50\x4e\x46\x8e\xe6\x8a\x4a\x57\xb5\x15\xde\x54\x57\xa9\xf0\x0d\x02\xa0\x08\xeb\x51\x29\xe4\xa0\x0f\x89\xcd\xca\xdc\x2c\x4b\x61\x92\xd8\xd2\x48\x0f\xe6\xea\x5e\x5a\x06\xba\xc3\xbc\xc8\xe8\x08\xb4\x81\xf0\x18\x3f\x8b\x21\x6b\xba\x21\x5f\xda\xf5\xb5\x0e\x6e\xb5\xa4\x80\xf3\xa3\x61\x61\xa2\x20\x87\x9e\x27\x08\x5c\xe5\xd3\xab\x45\xd0\xe7\x5f\xae\xae\xae\xdc\xa7\xac\x66\x4d\x54\x02\x99\xbe\x26\xe8\xe6\xd3\xef\xba\x75\xd1\xd5\xf5\x17\xab\x41\x80\x04\x0d\x60\xe6\x2c\x8c\x28\x36\x4d\xa4\xc0\xfa\x42\xcb\xfc\x31\x00\x4c\xce\x96\x9c\x50\xef\x1e\x46\xbb\x72\x34\x4f\x0b\x07\x19\x8e\x28\x23\x05\xa3\x29\x5c\x8d\xad\x7d\x32\x42\xbe\x3a\xda\x6a\x63\x5d\xc7\x65\x50\xe1\x7a\xd7\x34\x85\x27\xd0\x1d\x5b\xdb\x05\x34\xaa\x55\xe6\x06\xb3\x92\xbc\xd4\x08\x79\x8b\x57\x5e\xc5\xf0\xd6\x33\xcf\x74\xc5\x73\xfe\x8d\x56\xa4\x8e\xc0\x32\xe8\x28\x13\xb5\x6a\x07\x94\x17\x32\x3d\xf2\x73\x8b\xb7\x82\x95\x58\x4b\x8a\x12\x66\x7c\xa0\x20\x45\x07\x05\x71\xae\x9d\xd3\xd6\x78\x87\x39\x22\xb8\xd5\x59\x06\xa3\x45\x1e\xc4\xea\x27\xd5\xdb\x95\x6b\x67\x7f\x00\x2d\x97\xf0\x6c\xf2\x5f\xa8\xe1\x18\xed\xd1\xf4\xc1\xab\xb8\x52\xbc\x5b\x21\x8f\x4a\xd9\xbb\x98\x1b\x65\xbc\x67\x42\xcf\xa3\x1a\x7e\x8e\x79\xbc\x7a\xc1\xdb\x5a\xaa\xe8\x92\xcd\xd9\x78\xc6\xf7\xc3\x00\x97\x68\xd4\x25\x8c\x35\x3b\x81\xfd\x41\x1d\xc5\x95\x6f\x77\xc2\xf8\x50\x15\x63\x2c\xd0\x5d\x91\xe9\x44\x4b\x34\x29\x12\x60\xa8\x88\x8a\x8c\x76\x2a\x84\xbf\x03\x00\x00\xff\xff\xf4\xb4\x2b\x7b\x85\x22\x00\x00")

func ocmSmtpServiceYamlBytes() ([]byte, error) {
	return bindataRead(
		_ocmSmtpServiceYaml,
		"ocm-smtp-service.yaml",
	)
}

func ocmSmtpServiceYaml() (*asset, error) {
	bytes, err := ocmSmtpServiceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ocm-smtp-service.yaml", size: 8837, mode: os.FileMode(420), modTime: time.Unix(1586965608, 0)}
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
	"ocm-smtp-service.yaml": ocmSmtpServiceYaml,
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
	"ocm-smtp-service.yaml": &bintree{ocmSmtpServiceYaml, map[string]*bintree{}},
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
