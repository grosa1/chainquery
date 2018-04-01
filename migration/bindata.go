// Code generated by go-bindata.
// sources:
// migration/000_init_schema.sql
// DO NOT EDIT!

package migration

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

var _migration000_init_schemaSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5b\x4d\x73\xdb\xbc\x11\xbe\xfb\x57\xe0\x16\x69\x6a\xcf\x58\x7e\xd3\xb7\xe9\x74\x72\xa0\x29\xd8\xe6\x1b\x89\x72\x29\x2a\x4d\x7a\x81\x20\x12\x92\x50\x53\xa0\x42\x80\xb1\xdd\x5f\xdf\x21\x40\x8a\xe0\x97\x44\x49\x4c\x1b\x67\x7a\x8b\xc5\xc5\x2e\x3e\x76\x9f\x7d\xb0\x8b\x5c\x5d\x81\x3f\x6d\xe8\x2a\xc2\x82\x80\xd9\xf6\xe2\x42\xff\x7b\x2a\xb0\x20\x1b\xc2\xc4\x2d\x59\x51\x76\x61\x3a\xd0\x70\x21\x70\x8d\xdb\x11\x04\xd6\x1d\xb0\x27\x2e\x80\x5f\xac\xa9\x3b\x05\xf3\x45\x10\x7a\x4f\xf3\x8b\xde\x05\x00\x00\xcc\xa9\x3f\x07\x53\xe8\x58\xc6\xe8\xf2\x42\xfd\xb2\xa0\x82\xcf\xc1\x67\xc3\x31\x1f\x0c\xa7\x77\x73\xdd\x97\xa3\xed\xd9\x68\x74\xa9\x04\xbc\x35\xa6\xec\x39\x8c\x9e\x72\xa9\xbf\x5c\xf7\x41\xf2\x0f\xc3\x74\xa1\x03\xa6\xd0\x05\x01\x16\x94\x0d\x80\x39\x19\x8d\x92\x99\xa8\x3f\xd1\x8a\x30\x12\xe1\x00\x79\xb4\xa2\x34\x64\x4b\x1a\x6d\xb0\xa0\x21\xe3\x73\x60\xd9\x2e\xbc\x87\x0e\x98\xd9\x53\xeb\xde\x86\xc3\xb2\xb8\x4f\x97\x4b\xea\xc5\x81\x78\x9d\x83\xe1\x64\x76\x3b\x82\xbd\xc1\x87\xcb\x0f\x95\xb9\xae\x31\x5f\x77\x32\x4d\x90\x29\x24\x74\xb5\x16\x73\x70\x6b\xdd\x5b\xb6\xdb\x38\xbf\x0d\xf1\x29\x66\x48\xd0\x0d\x69\x21\x1b\x3d\x05\x04\x45\x61\x28\xba\xdc\x51\x86\x37\x04\x79\x01\xa6\x9b\xee\x55\x87\xcc\x3b\xbc\xae\x6d\x44\xbe\xd3\x30\xe6\x48\x7a\x1c\x3a\xf7\x28\x32\xdb\xe4\x45\x74\xab\x51\x29\xe3\xf4\xdf\x87\x97\x24\x70\xb4\x22\x9d\x6e\xa5\x32\xde\xca\x4f\xbe\x93\x88\xd3\x90\xb5\x95\x43\x6b\xf2\x92\xcf\x74\x70\xf6\x4c\x45\x84\x19\xc7\x9e\x90\xaa\x31\x5f\x13\x3e\x07\x2e\xfc\xe2\x56\x3f\x73\xb4\x8d\x42\x8f\x70\x4e\xfc\x39\x70\x2d\xfb\xab\x65\xbb\xbd\x41\x1f\x0c\xe1\x9d\x31\x1b\xb9\xe0\x5a\x53\x9d\x46\x7f\x44\xb0\x48\xa4\x87\x86\x0b\x5d\x6b\x0c\xf3\xb8\xcb\x06\x99\x33\xc7\x81\xb6\x8b\x92\xaf\x53\xd7\x18\x3f\x66\xd1\x13\xfa\x74\x49\x8f\x1c\x0b\x26\x36\x98\x3d\x26\x03\xea\xf4\x4a\xc5\x8f\x8e\x35\x36\x9c\xaf\xe0\x13\xfc\x0a\xe6\x8f\x9f\xd0\xad\x84\x4d\xd0\x4b\x10\xb3\xaf\x6c\xcf\x6c\xeb\xef\x33\xa8\x24\x2c\xff\x45\x89\x3c\x48\xa7\xec\x29\xe4\x49\x05\xcd\x89\x3d\x75\x1d\x23\x39\xb4\xb9\xc9\x04\x72\xf3\xbd\x7a\x90\x3b\xf9\x19\x07\xd4\xff\x83\x27\x87\x6b\x3e\x40\xf3\x53\xaf\x76\xb7\xad\xa9\x5a\xd6\xc4\x01\x7f\x4c\x27\x36\xfa\x6c\x8c\xac\x61\xad\x68\x3f\x35\x6c\xd9\x43\xf8\x45\x9f\x5c\x8a\x5e\xbd\x0c\xc7\x1a\xe4\x5c\xe9\x8f\x3d\xdd\x3b\xab\x92\x63\x89\x71\x99\xa8\x8e\x78\x55\xd9\xc7\x14\x0b\x0a\x5b\x54\x07\x10\x0d\x13\x32\x33\x07\xe9\xed\x7c\xa5\x41\x72\xbc\xf3\x87\x5e\xee\x1b\xfd\x8b\x3e\x80\xf6\xbd\x65\xc3\x8f\x16\x63\xe1\xf0\x36\x77\x8d\x07\xc3\x99\x42\xf7\x63\x2c\x96\x1f\x36\x8b\xf7\xbb\x70\x48\xff\x46\x31\xa3\x5e\xe8\x93\x24\x1e\x9c\xc9\x3f\xd0\xdd\xc4\x19\x1b\xee\x47\x73\x32\x7e\x74\xe0\x74\x0a\x87\xc9\xe9\xa3\xdb\xd1\xc4\xfc\x84\xa6\xd6\x3f\xe1\xc7\xf7\x7f\xab\xcf\xcb\x90\xf9\x27\x67\x6c\xed\x80\x6b\xf3\xb6\x06\x24\x8b\x57\xb9\x8d\x28\xf9\x7c\x36\x30\x52\xb6\x8d\x05\xf2\xc2\x98\x89\xc3\x19\x39\x8c\xc5\x11\xd2\xdf\x71\x10\x93\xfd\xa9\x7b\x49\xca\x02\x75\xf8\x51\x81\xa6\x5a\x28\xad\x91\x6b\x85\xf7\x9d\xd1\x87\x32\x84\x67\xfb\x53\xfa\xae\x25\x83\x43\x3b\x18\xe1\xe7\x02\xfe\xfe\x5c\x10\xaa\x4f\x29\x5d\xd0\xd1\xf3\x2a\x23\xb0\x86\x9a\x45\x1c\xbe\x9b\x38\xd0\xba\xb7\x95\xe0\x5d\x41\xb0\x80\x37\x95\x10\xe9\x03\x07\xde\x41\x07\xda\x26\xdc\x11\xe3\x1d\x74\x27\x0b\x1c\xc2\x11\x4c\x16\x68\x4c\x4d\x63\x08\xb5\x25\xdb\x13\x60\x98\xae\x35\xb1\x41\x7d\x2a\x28\x21\x7c\x39\x21\x68\xa0\xa5\x49\x66\x50\x5a\x71\xe7\xbd\xa3\x52\x68\xcc\x06\x17\x36\xbd\xcd\xc0\xfd\x98\xaa\xc9\xff\x82\xc8\x8a\x7d\x3f\x22\x9c\x37\xa3\x6a\x26\xb0\xc3\x80\xf7\xc7\x63\x40\xea\x1b\x65\x78\xa3\x11\x17\x88\x13\xc2\xf2\xd0\xd8\xb1\xcc\x55\x6e\xf0\xb7\xeb\x7e\xaa\x21\xff\x8a\xe2\x28\xd0\xaf\x68\xd7\xfd\x9f\x16\x04\xca\x41\x6c\x64\x1b\xba\x8f\x48\xa5\x42\x9a\x6c\x76\x0e\x99\x83\x32\x9f\xbc\x14\x64\xdd\x64\xcf\x7a\x72\xeb\xaa\x4e\x9c\xca\xb4\x72\xf8\x54\xf6\x17\x74\x76\x99\xce\x9b\x5d\x5d\x47\x9d\xe4\xe3\xa1\xab\x50\x89\x75\x76\x99\x26\x15\xf1\x48\xcf\xbc\x6e\x32\x99\x1c\x47\x5e\x48\xd9\x02\x73\x72\xf0\xa2\x91\x56\x19\x32\xe9\xb3\xc9\x51\x42\x5e\xc3\x58\x74\x75\x09\xcd\xd4\xb1\x6a\xe6\x2f\x49\xf0\x2d\x61\x3e\x8a\xb7\xbe\xf2\xe5\x36\xcb\xe6\xe4\x5b\x4c\xe4\x95\xfd\x04\x5e\x96\xe9\xf0\x22\xba\x15\x88\xd3\x15\xc2\x7c\xa3\xa8\xc7\x69\x4b\xd5\x34\xc9\x2b\xea\x49\x9a\xde\x0a\xde\x59\x32\xe8\xf6\xd2\x15\x29\xa2\xe3\x62\xd9\xfb\x8b\x34\x05\x17\x11\xb4\x1d\x4d\xd9\x63\xb9\x44\xab\x4a\x30\x50\xb4\x2d\xaa\x14\xec\x18\xfb\x1a\xcc\x4a\xd3\x9f\x95\xb7\xf5\x52\xb7\xab\xbf\x36\x86\xb1\xd0\x2f\x8c\xbb\x98\xab\x4a\x4b\x9d\xad\x40\x5e\x4a\xfe\xaa\x10\x8f\x2a\xac\x46\xfe\xda\x02\xd3\xf7\x20\x6e\x49\xb4\xd6\xcd\x2b\x3e\x4c\xfd\xf9\x65\x41\x6b\x5d\x04\xec\xce\x23\xcb\xf7\x84\xe7\x41\x93\xa9\x29\xba\x21\xd5\x83\xea\xe4\x00\xa8\x31\x5c\x65\x1c\x1d\xc5\xdf\x9b\xf6\x2a\x75\xb5\x7f\x13\xcc\xa1\x39\x7f\x25\xc0\x71\x38\xff\x89\xd7\x2d\x29\x36\x3f\xce\x49\x71\xdb\x78\x81\x9e\xc8\x6b\x47\x09\x33\xd3\x76\x7a\xd2\x54\x85\x03\xf2\x2d\xa6\x11\xf1\x93\x04\xcc\xb0\x88\x23\x52\xd3\x6f\xd1\xea\x1f\x83\xdf\xaf\xf3\x1d\xf9\xf3\xa9\x3b\x92\x85\x53\x40\xb9\x38\x67\xf6\x94\x4b\x02\x24\x5a\x12\x9f\x44\x34\xb9\xf1\x37\x42\xe0\x9b\xb9\x38\x4d\x54\x10\x36\xde\x9b\x66\xba\x48\x25\xc4\x2e\x95\xff\x37\x30\x10\x35\xf0\xbf\x48\x04\xea\x27\x30\x4d\x4e\xeb\xf6\x75\x87\xfe\xd5\xd3\xeb\x28\x0d\x94\xeb\xf1\x3b\xfc\xaf\xd6\xe1\x8b\x7e\xdb\x50\x81\x2f\x08\xd5\xd4\xde\xd5\xf2\x0e\xb2\x9d\x49\xbc\x0d\x29\x93\x2b\x92\xa7\x75\x59\x3d\xc7\x3e\x30\x27\xe3\x31\xb4\x5d\xf0\x8e\x11\xe2\x13\x1f\x2c\xc3\x08\x44\x64\x49\xa2\x84\xe0\x73\x40\x19\x10\x6b\xca\x81\x17\x06\xf1\x86\x81\x30\xf2\x49\xf4\xae\xce\x92\x68\x49\x95\xd4\xdc\x7f\x41\xae\x94\x16\xac\xcb\x64\x29\xfd\xf9\x87\xb2\x25\xb5\xa7\x1a\xd9\xc8\x8d\xb6\xe3\x4b\x4a\x41\xce\x5b\xf2\xc8\xcf\x35\x15\x63\x25\x2c\xc0\xc7\x59\x9c\xa9\x6c\xfc\xff\xa4\xe9\x40\xd3\xa6\xe2\x63\x47\x12\xa6\xf6\x8e\x36\xf7\xc9\x82\x0a\x84\x37\xaa\x0b\x73\xa0\x7b\x92\x63\xc9\x34\xde\x80\x70\x09\xc4\x9a\x00\x09\xaa\x1c\x88\x50\xe1\x48\x6a\x5b\xc2\x4c\xf2\x59\xbc\xbc\xcb\x53\xa6\x7f\xa6\x2d\xe5\x94\x6d\x8c\x05\x58\x10\x2e\x50\xb5\xd1\xd3\x65\x57\x41\x73\xe4\xd2\x09\xb5\x8b\x4a\x4d\x55\x1e\x1d\xff\xab\xac\xda\x38\xa1\x1f\x16\xae\xfb\x7a\x07\xb9\xf9\x91\x3c\xc9\x9a\xae\x47\xd3\x11\xef\xed\x49\xe4\x7a\x87\x89\xe7\x27\x7a\x0a\x21\xd0\x72\xb0\x29\x7d\x39\x4d\x84\x9a\x57\xbf\xed\x0c\x27\x1f\xff\xb4\xbb\xb6\x75\xda\x37\x6e\x77\xd5\x62\x78\x43\xf4\xf7\x29\x37\xef\xab\x2f\xcd\xe4\xeb\xa5\x64\x4e\xa7\xf6\x5e\x6a\x15\xaa\x4b\x9e\x76\x7f\xd8\x49\x81\xab\x2b\x30\x00\x57\xc0\x24\x91\xa0\x4b\xea\x61\x41\xdc\xd7\x2d\xb9\x04\x37\xe0\x0a\x4c\x45\x44\xf0\x26\xf9\x1b\xf0\x75\x18\x07\x3e\x60\xa1\x00\x0b\x02\x04\x65\xaf\x94\x09\x10\xd0\x0d\x15\x1c\x60\x0e\x16\x61\x18\x00\xfe\x2d\x58\x84\x34\x20\x51\x5a\xbf\x8d\x17\x01\xe5\x6b\x12\x9d\xb5\xa0\x1d\x8c\x6a\x9c\x0f\x03\x33\x59\x97\xe5\x83\x67\x2a\xd6\xe5\xc9\x67\x10\x9a\xdb\xe7\x74\x55\xec\x1e\x9d\x76\xcc\x5e\x6e\xa7\xe9\x25\xd0\xde\x6e\xfc\xae\x19\xde\xdd\x1b\x25\xc9\xac\x11\xe6\xea\x8e\x3c\x86\x43\x6b\x36\x96\x77\xcd\x26\xb9\x7f\x49\xae\x9f\x0b\xa6\x85\xe5\xab\x2b\x60\xf8\x3e\x4d\x96\x80\x03\xb0\xa4\x24\xf0\x55\x56\x22\x98\xbf\x02\xca\x7c\xf2\x42\xd9\x2a\x49\x63\x5c\x3a\x05\x48\x5c\x8a\xa7\xd9\x3a\x16\xeb\x30\xd2\xae\xcc\x83\x9b\xfe\x2e\x3b\xab\xab\x3c\x2d\x1b\x4d\x9b\x13\x4c\x24\x17\x9e\x62\x0d\x62\xf0\xfb\xcd\xa9\xcf\x37\x38\x62\x88\xa3\x25\x7a\x6e\x77\x59\x0e\x30\x5b\xc5\x78\xd5\x45\xfd\x43\xac\xe3\xcd\x82\x61\x1a\xa8\x4e\xa5\xee\x1c\x54\x04\x45\x77\x39\xe2\x9d\xc7\x92\x10\xe4\xc5\x51\xe2\xf6\xaf\x69\x08\xfd\xd6\xcf\x57\xbb\xa4\x81\x20\x51\xdb\xa6\x08\xe5\x69\x13\xe5\x4d\xbe\x51\xab\xa3\x32\x12\x05\xf6\xf7\x1a\xa4\xc8\x1e\x4a\xd2\xf8\x38\xa2\x44\x4d\x8e\x79\x22\xb1\x67\x26\x8f\x19\x26\xc9\xea\xbe\x0e\x90\x45\xeb\x5e\xb6\xb2\x5d\x46\x38\xca\x76\xb9\xbf\x2c\x6d\xcf\x18\xfd\x16\x57\xde\x5b\x68\x1b\x70\x99\x66\xb2\x4b\x2d\x11\x35\x3c\xe8\x93\x0a\x4d\x1d\x0f\xd3\xf2\x41\x01\x23\x1b\xaa\x07\xba\x4c\xbf\x2f\x73\x90\xf6\x93\x84\x96\x32\x89\x31\xab\xfb\x51\x21\x3a\xe5\xa3\x3e\xe2\x75\x89\x5a\x4e\x9b\xea\x80\x94\xac\x2f\x0e\xa4\x71\x53\x16\x36\x52\x78\xec\x65\x40\xd9\x1b\xfc\x75\x50\x53\x33\x51\x73\x50\x90\xe8\x4a\x44\xec\x15\x11\xb2\x61\xc8\x68\x87\x62\xbd\x1c\xd1\x9a\x76\x47\x81\x51\x2f\x45\x25\x35\x93\x37\x4d\xf9\x62\xf6\xc4\xc2\x67\x86\x0e\x50\xbf\xa3\xb8\x57\x9e\x89\xaa\xff\x17\xa0\x3d\x88\x82\xfc\x91\x62\x47\x2d\xf3\x2e\xfb\x08\x47\x31\xd7\xf6\x25\xa2\x0e\xf8\x48\x1d\xd0\xa7\xa7\xec\x55\xf1\xbe\x0c\x4c\xb2\xe6\x58\xad\x6a\x96\xec\x35\x00\x53\x51\xaa\xdf\x90\x52\xf4\xc9\x84\x3f\xac\x0e\xa5\x85\xee\x4c\x1a\x94\x01\x5c\xf3\xbc\x2f\xef\x0d\xd7\x0f\xa9\x2d\x95\x35\x8b\xbb\x2f\x99\xfa\x6a\x2d\xf6\xa7\x40\x8a\xff\x04\x00\x00\xff\xff\x29\xd8\x3b\x09\x16\x34\x00\x00")

func migration000_init_schemaSqlBytes() ([]byte, error) {
	return bindataRead(
		_migration000_init_schemaSql,
		"migration/000_init_schema.sql",
	)
}

func migration000_init_schemaSql() (*asset, error) {
	bytes, err := migration000_init_schemaSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migration/000_init_schema.sql", size: 13334, mode: os.FileMode(420), modTime: time.Unix(1522546476, 0)}
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
	"migration/000_init_schema.sql": migration000_init_schemaSql,
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
	"migration": &bintree{nil, map[string]*bintree{
		"000_init_schema.sql": &bintree{migration000_init_schemaSql, map[string]*bintree{}},
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

