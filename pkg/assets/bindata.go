// Code generated by go-bindata.
// sources:
// assets/schema/building.graphql
// assets/schema/resident.graphql
// assets/schema/root.graphql
// assets/schema/unit.graphql
// DO NOT EDIT!

package assets

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

var _assetsSchemaBuildingGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\x70\x2a\xcd\xcc\x49\xc9\xcc\x4b\x57\xa8\xe6\x52\x50\x50\x50\xc8\x4c\xb1\x52\xf0\x74\x51\x04\xb3\xf3\x12\x73\x53\xad\x14\x82\x4b\x8a\x32\xf3\xd2\x21\x22\x89\x29\x29\x45\xa9\xc5\xc5\xa8\x82\xa5\x79\x99\x25\xc5\x56\x0a\xd1\xa1\x79\x99\x25\xb1\x8a\x5c\xb5\x5c\x80\x00\x00\x00\xff\xff\xe7\x81\x86\xea\x58\x00\x00\x00")

func assetsSchemaBuildingGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_assetsSchemaBuildingGraphql,
		"assets/schema/building.graphql",
	)
}

func assetsSchemaBuildingGraphql() (*asset, error) {
	bytes, err := assetsSchemaBuildingGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/schema/building.graphql", size: 88, mode: os.FileMode(420), modTime: time.Unix(1498812156, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsSchemaResidentGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\x08\x4a\x2d\xce\x4c\x49\xcd\x2b\x51\xa8\xe6\x52\x50\x50\x50\xc8\x4c\xb1\x52\xf0\x74\x51\x04\xb3\xf3\x12\x73\x53\xad\x14\x82\x4b\x8a\x32\xf3\xd2\x21\x22\xa5\x79\x99\x25\x56\x0a\xa1\x79\x99\x25\x8a\x5c\xb5\x5c\x80\x00\x00\x00\xff\xff\x2d\xa1\x37\x12\x40\x00\x00\x00")

func assetsSchemaResidentGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_assetsSchemaResidentGraphql,
		"assets/schema/resident.graphql",
	)
}

func assetsSchemaResidentGraphql() (*asset, error) {
	bytes, err := assetsSchemaResidentGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/schema/resident.graphql", size: 64, mode: os.FileMode(420), modTime: time.Unix(1498812156, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsSchemaRootGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x4e\xce\x48\xcd\x4d\x54\xa8\xe6\x52\x50\x50\x50\x28\x2c\x4d\x2d\xaa\xb4\x52\x08\x04\x51\x5c\xb5\x5c\x5c\x25\x95\x05\xa9\x10\x1e\x54\x81\x53\x69\x66\x4e\x4a\x66\x5e\x7a\xb1\x95\x42\x34\x8c\x1d\xab\x88\x22\xa5\x01\x63\x78\xba\x58\x29\x78\xba\x28\x6a\x5a\xc1\xa5\x14\xb9\x6a\xb9\x00\x01\x00\x00\xff\xff\x00\x19\x81\xad\x70\x00\x00\x00")

func assetsSchemaRootGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_assetsSchemaRootGraphql,
		"assets/schema/root.graphql",
	)
}

func assetsSchemaRootGraphql() (*asset, error) {
	bytes, err := assetsSchemaRootGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/schema/root.graphql", size: 112, mode: os.FileMode(420), modTime: time.Unix(1498812156, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsSchemaUnitGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\x08\xcd\xcb\x2c\x51\xa8\xe6\x52\x50\x50\x50\xc8\x4c\xb1\x52\xf0\x74\x51\x04\xb3\xf3\x12\x73\x53\xad\x14\x82\x4b\x8a\x32\xf3\xd2\x21\x22\x49\xa5\x99\x39\x29\x99\x79\xe9\x56\x0a\x4e\x50\x16\x44\xbc\x28\xb5\x38\x33\x25\x35\xaf\xa4\xd8\x4a\x21\x3a\x08\xca\x8e\x55\xe4\xaa\xe5\x02\x04\x00\x00\xff\xff\x79\x26\xc7\xf9\x5f\x00\x00\x00")

func assetsSchemaUnitGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_assetsSchemaUnitGraphql,
		"assets/schema/unit.graphql",
	)
}

func assetsSchemaUnitGraphql() (*asset, error) {
	bytes, err := assetsSchemaUnitGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/schema/unit.graphql", size: 95, mode: os.FileMode(420), modTime: time.Unix(1498812156, 0)}
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
	"assets/schema/building.graphql": assetsSchemaBuildingGraphql,
	"assets/schema/resident.graphql": assetsSchemaResidentGraphql,
	"assets/schema/root.graphql": assetsSchemaRootGraphql,
	"assets/schema/unit.graphql": assetsSchemaUnitGraphql,
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
	"assets": &bintree{nil, map[string]*bintree{
		"schema": &bintree{nil, map[string]*bintree{
			"building.graphql": &bintree{assetsSchemaBuildingGraphql, map[string]*bintree{}},
			"resident.graphql": &bintree{assetsSchemaResidentGraphql, map[string]*bintree{}},
			"root.graphql": &bintree{assetsSchemaRootGraphql, map[string]*bintree{}},
			"unit.graphql": &bintree{assetsSchemaUnitGraphql, map[string]*bintree{}},
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

