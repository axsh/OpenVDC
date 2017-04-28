// Code generated by go-bindata.
// sources:
// ../schema/none.json
// ../schema/v1.json
// ../schema/vm/lxc.json
// ../schema/vm/null.json
// DO NOT EDIT!

package registry

import (
	"github.com/elazarl/go-bindata-assetfs"
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

var _schemaNoneJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x8e\xb1\x4e\x03\x31\x0c\x86\xf7\x3c\x85\x15\x18\x40\x6a\x39\x40\x4c\x59\x41\xac\x48\x08\xb1\xa0\x0a\x85\x9c\xb9\xa6\x6a\xec\xe0\xf8\x86\xaa\xea\xbb\xa3\x5c\xe0\x54\x26\x3a\x64\xc8\xe7\xef\xf7\xef\xbd\x01\xb0\xe7\x25\xac\x31\x79\xeb\xc0\xae\x55\xb3\xeb\xba\x4d\x61\x5a\x36\x7a\xc5\x32\x74\xbd\xf8\x4f\x5d\x5e\xdf\x75\x8d\x9d\xd9\x45\xcd\xf5\x58\x82\xc4\xac\x91\xa9\x66\x9f\x32\xd2\xeb\xc3\x3d\x3c\x63\xe1\x51\x02\xc2\x0b\xa6\xbc\xf5\x8a\x0e\x88\x09\xe1\xe2\x91\x05\x14\x8b\x46\x1a\x80\x69\xbb\xbb\x6c\x6b\x74\x97\xb1\xe6\xf9\x63\x83\x41\x1b\x13\xfc\x1a\xa3\x60\x6f\x1d\xbc\x19\x80\x5f\xcb\x00\xac\xa6\x79\x16\xce\x28\x1a\xb1\x58\x07\xfb\x66\xbc\x07\x4e\x09\x49\x67\x72\xb4\xbb\xa8\x44\x1a\xec\x84\x0f\x0b\x73\x3c\x9b\x5d\xa4\x31\xcd\x7d\x13\xa9\x67\xdb\x9f\xef\xea\x4f\x36\x7b\xf1\xe9\xe6\xd4\xa6\xc9\xbe\xfd\xd7\x36\xf5\x1d\xcc\x77\x00\x00\x00\xff\xff\x57\x3a\x39\x38\x94\x01\x00\x00")

func schemaNoneJsonBytes() ([]byte, error) {
	return bindataRead(
		_schemaNoneJson,
		"schema/none.json",
	)
}

func schemaNoneJson() (*asset, error) {
	bytes, err := schemaNoneJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/none.json", size: 404, mode: os.FileMode(420), modTime: time.Unix(1492164609, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaV1Json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x91\x41\x4b\xc3\x40\x10\x85\xef\xf9\x15\xcb\xea\xb1\xcd\x2a\x78\xca\x55\xef\x05\x15\x2f\x52\x24\xdd\x4c\x9b\x2d\xd9\x99\x75\x76\x52\x2b\x25\xff\x5d\x36\x49\x43\x44\xc1\x62\x4e\xbb\x6f\xde\xf7\x36\x8f\x39\x65\x4a\x69\x57\xe9\x42\xe9\x5a\x24\xc4\xc2\x18\x2e\x3f\xf2\x9d\x93\xba\xdd\xb4\x11\xd8\x12\x0a\xa0\xe4\x96\xbc\x29\x8f\xb1\x36\x14\x00\x0f\x95\x35\xbe\x8c\x02\x6c\xa2\xad\xc1\x97\xe6\x70\x9b\xef\x23\xe1\x95\x5e\xa4\xc0\xeb\x41\x3d\xa7\x16\xc6\xa4\xe1\x72\x50\x73\xe2\x9d\xa9\xb8\xdc\xca\xf2\xe6\x6e\xe4\x47\xae\x82\x68\xd9\x05\x71\x84\x89\x5d\x05\xc0\x97\x87\x7b\xf5\x08\x91\x5a\xb6\xa0\x9e\xc1\x87\xa6\x14\x50\x4f\x43\x7e\x0f\xc9\x67\x80\xe4\xa6\xcd\x1e\xac\x0c\x1a\xc3\x7b\xeb\x18\x52\xaf\xd7\x4c\xa9\xe4\x72\xd2\x40\x3f\x4c\x97\x31\x47\x67\x4a\xad\x7b\x20\x30\x05\x60\x71\x10\x75\xa1\x4e\x83\xeb\xcd\x92\xf7\x80\x32\x29\xb3\xc7\xa2\xb0\xc3\x9d\xee\xe5\x6e\x31\x7f\xe2\x42\xf3\xf7\xaa\x97\xe5\x9f\xff\xfa\x17\xff\xac\x7c\xaf\x13\xc2\x6a\x3b\xb5\x4f\xdf\x69\x3a\xa5\x05\x31\xa4\xa9\xce\x0d\x12\xc2\xb8\xba\xc9\xd0\x2d\xfe\xa2\x0e\xde\x34\x47\xfb\x2f\x0e\xdb\xa6\xf9\x01\x8e\xa7\x75\x76\xbe\x75\x59\x97\x7d\x05\x00\x00\xff\xff\x60\x1a\x39\xfd\x9c\x02\x00\x00")

func schemaV1JsonBytes() ([]byte, error) {
	return bindataRead(
		_schemaV1Json,
		"schema/v1.json",
	)
}

func schemaV1Json() (*asset, error) {
	bytes, err := schemaV1JsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/v1.json", size: 668, mode: os.FileMode(420), modTime: time.Unix(1492164609, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaVmLxcJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x56\xcd\x6e\xdb\x3c\x10\xbc\xeb\x29\x08\x25\xb7\xcf\xfe\x94\x14\xb9\xd4\xb7\xa2\x45\x81\x9e\x5a\x14\x45\x2f\x81\x2a\x50\xd4\x4a\x66\x2a\xfe\x64\x49\xba\x36\x0c\xbd\x7b\x41\x5b\xb6\xfe\x13\xa7\x4e\x82\x1c\x0c\x58\xb3\xdc\xd9\x25\x77\x34\xd4\x36\x20\x24\xbc\x34\x6c\x09\x82\x86\x0b\x12\x2e\xad\xd5\x8b\x28\xba\x33\x4a\xce\xf7\xe8\xff\x0a\x8b\x28\x43\x9a\xdb\xf9\xd5\x4d\xb4\xc7\x2e\xc2\x99\xcf\xcb\xc0\x30\xe4\xda\x72\x25\x7d\xee\x57\x0d\xf2\xe7\xa7\x8f\xe4\x3b\x18\xe5\x90\x01\xf9\x01\x42\x97\xd4\xc2\x82\xac\x44\x54\xae\xd9\x3e\xcb\x6e\x34\xf8\xe5\x2a\xbd\x03\x66\xf7\x18\xc2\xbd\xe3\x08\x59\xb8\x20\xb7\x01\x21\x87\x55\x01\x21\xf1\x2e\xae\x51\x69\x40\xcb\xc1\x84\x0b\xb2\xdd\xaf\x48\x98\x12\x02\xa4\x3d\x22\x2d\x6e\x63\x91\xcb\x22\xdc\xc1\xd5\x2c\x68\xc7\x8e\x6b\x41\x3a\x71\xac\xb7\x43\xea\x2e\x6b\x20\xee\x64\x0b\x2e\x93\x15\xd3\x6e\xac\x1a\x97\x16\x0a\xc0\x70\x76\x08\x64\x90\x53\x57\xfa\xce\xae\x07\x24\x02\x84\xc2\x4d\x52\xa4\x67\x31\x9d\xdf\xca\xb3\xb4\x21\x55\x06\x49\x81\xca\x69\x33\xc6\x43\x11\xe9\xa6\x61\x71\x92\xdf\x3b\xf8\x62\x41\xf8\xd5\x16\x1d\x1c\x43\xbc\x06\x9b\x69\x6c\xfb\xc3\xac\x46\xe7\xe2\x5b\xc5\x9c\x32\x38\xa5\x81\x43\x95\x6d\x33\xf3\x11\x35\xd6\x91\x81\x26\xdb\x19\x47\x20\x6e\x65\x8c\xa8\xb4\x5b\xa5\x8d\x0d\xd5\x3a\xeb\x46\x07\xfa\x3c\xa8\x14\xec\xb2\xb7\x76\x37\x09\x09\x43\x14\x84\xb6\x9b\x21\xbc\x2a\xa9\x1c\xa2\x82\xb2\xf1\x80\x5e\x6e\x4c\xd8\x01\xe3\xd6\x53\xd5\x5e\xef\x49\x68\x96\xe1\x53\x37\xab\xa9\xb5\x80\x3b\x1b\xf9\x75\x7b\x35\x7f\x4f\xe7\xf9\x87\xf9\xe7\x78\xfb\xae\x6a\x9e\x16\xf1\x7f\x97\xe1\x64\x61\xae\x57\x37\xff\x52\x39\x57\x28\xa8\xdd\xa9\x5e\xaf\x6e\x3a\xfc\x41\xff\x5f\xd5\x91\x5e\xb9\x66\x89\xad\x2d\x6e\x4c\x7c\x3d\x49\xf9\x97\x88\x4b\xee\xdd\xb2\xa7\xc1\x94\x1a\x3e\x4a\xf5\x10\x1d\x79\x50\x71\xe3\xfe\x38\x75\x24\x9d\x70\xd5\x3b\x20\x84\x12\xa8\x19\xca\xf7\x89\x34\x14\xd9\xf2\x7c\x8e\x62\xb8\xd1\xc7\x38\x1e\x9e\xe8\x71\x3a\x82\xae\xbf\xb5\x4f\xf3\xba\x89\x70\x39\x11\x99\x38\xfe\x30\x53\x7f\x64\xa9\x68\xf6\x94\x51\x4e\xd8\x8d\x67\xe3\xc6\xa2\xea\xeb\xf6\x30\x96\x16\x1a\xbf\xba\x36\xea\xd6\xde\x86\xc2\x56\x14\x39\x3d\x7f\x4f\x6f\x4c\xa8\xad\xfb\xc5\x77\x56\x72\xe9\xd6\x7d\x5d\x5d\x22\xe4\x9e\xf6\x22\x6a\x86\x1e\xb5\xed\x29\x6a\x79\x4f\xd4\xb3\x9b\xd1\x52\xa9\x33\x9b\x54\xbd\x42\x21\x06\xd2\xaa\xc1\x7d\xf9\xfc\x75\x32\x48\x39\x95\x2f\x5f\x27\x87\x4c\x21\x7d\xf9\x3a\x2e\x75\xd2\xba\x17\xaa\x33\x79\xdd\x71\x41\x8b\x93\xee\xba\x29\x67\x9c\x30\x9e\xc9\x17\xa3\x33\xc2\xda\x55\x13\x87\xe5\x23\xf9\x1d\x23\x6c\x6e\x78\x87\x7c\x42\x85\x4b\x60\xbf\x8d\x13\xc9\xc8\x87\xda\x49\xbd\x1d\x08\x4e\xce\x1d\xdc\x3f\xa3\xfe\xdf\xdd\x73\xf7\x03\x38\xf0\xbf\x2a\xf8\x1b\x00\x00\xff\xff\xe7\x65\x42\x4c\xc2\x0d\x00\x00")

func schemaVmLxcJsonBytes() ([]byte, error) {
	return bindataRead(
		_schemaVmLxcJson,
		"schema/vm/lxc.json",
	)
}

func schemaVmLxcJson() (*asset, error) {
	bytes, err := schemaVmLxcJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/vm/lxc.json", size: 3522, mode: os.FileMode(420), modTime: time.Unix(1493359667, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaVmNullJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x91\x41\x4b\xf3\x40\x10\x86\xef\xf9\x15\xc3\x7e\xdf\x41\xa1\x35\x0a\x9e\x72\x55\x04\x4f\x82\x88\x17\x29\x21\x4d\xa6\xe9\x96\xec\xcc\x76\x76\xb6\x10\x4a\xff\xbb\xa4\x49\xd3\x56\x7b\xab\x87\xbd\x3c\xf3\xce\xb3\x2f\xcc\x36\x01\x30\xff\x43\xb9\x44\x57\x98\x0c\xcc\x52\xd5\x67\x69\xba\x0a\x4c\xd3\x9e\xde\xb1\xd4\x69\x25\xc5\x42\xa7\xf7\x8f\x69\xcf\xfe\x99\x49\xb7\x57\x61\x28\xc5\x7a\xb5\x4c\xdd\xee\x9b\x47\xfa\x7c\x7e\x82\x77\x0c\x1c\xa5\x44\xf8\x40\xe7\x9b\x42\x31\x83\x8d\x4b\x29\x36\x0d\xdc\xbc\xb0\x80\x62\x50\x4b\x35\x30\x35\xed\x6d\x6f\xd2\xd6\x63\xa7\xe0\xf9\x0a\x4b\xed\x99\xe0\x3a\x5a\xc1\xca\x64\xf0\x95\x00\x1c\x52\x09\xc0\x6c\x3f\xf7\xc2\x1e\x45\x2d\x06\x93\xc1\xb6\x4f\xe4\x25\x3b\x87\xa4\x23\x39\x71\x07\x15\x4b\xb5\xd9\xe3\xdd\x24\x39\x9d\x8d\x59\xa4\xe8\xc6\xff\xf6\x64\x68\x6e\x06\x32\x3b\x5b\x77\x96\xf2\x4d\xe9\xe3\xa5\xef\x2c\x29\xd6\x28\x66\x72\x18\x54\xb8\x28\x62\xd3\x55\x7b\xf8\x25\x71\xe8\x58\xda\xbc\x9e\x5f\x65\xba\xbe\xca\x9f\xd4\x20\xae\x30\xaf\x85\xa3\x0f\x97\x3c\x85\x48\xd1\x1e\x2d\x91\xec\x3a\xe2\xab\xa2\xeb\xd2\x2a\x11\xc7\x91\x1d\xe0\xf1\x1c\xdb\x9f\xd7\xdc\x9d\xdf\x25\xe9\xde\x2e\xf9\x0e\x00\x00\xff\xff\x71\xee\x14\x09\xd6\x02\x00\x00")

func schemaVmNullJsonBytes() ([]byte, error) {
	return bindataRead(
		_schemaVmNullJson,
		"schema/vm/null.json",
	)
}

func schemaVmNullJson() (*asset, error) {
	bytes, err := schemaVmNullJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/vm/null.json", size: 726, mode: os.FileMode(420), modTime: time.Unix(1493359678, 0)}
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
	"schema/none.json": schemaNoneJson,
	"schema/v1.json": schemaV1Json,
	"schema/vm/lxc.json": schemaVmLxcJson,
	"schema/vm/null.json": schemaVmNullJson,
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
	"schema": &bintree{nil, map[string]*bintree{
		"none.json": &bintree{schemaNoneJson, map[string]*bintree{}},
		"v1.json": &bintree{schemaV1Json, map[string]*bintree{}},
		"vm": &bintree{nil, map[string]*bintree{
			"lxc.json": &bintree{schemaVmLxcJson, map[string]*bintree{}},
			"null.json": &bintree{schemaVmNullJson, map[string]*bintree{}},
		}},
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


func assetFS() *assetfs.AssetFS {
	assetInfo := func(path string) (os.FileInfo, error) {
		return os.Stat(path)
	}
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: assetInfo, Prefix: k}
	}
	panic("unreachable")
}
