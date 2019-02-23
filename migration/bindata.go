// Code generated by go-bindata.
// sources:
// migration/000_init_schema.sql
// migration/001_supports.sql
// migration/002_decimals.sql
// migration/003_dht_tracking.sql
// migration/004_new_indices.sql
// migration/005_remove_foreign_keys.sql
// migration/006_add_height_index.sql
// migration/007_more_decimals.sql
// migration/008_schema_fix.sql
// migration/009_certificate_validation.sql
// migration/010_triggers.sql
// migration/011_add_license.sql
// migration/012_store_last_height.sql
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

var _migration000_init_schemaSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5b\x5f\x73\xdb\xb8\x11\x7f\xf7\xa7\xc0\xdc\xcb\x51\x53\x79\xc6\xf2\x25\xd7\x74\x3a\x79\xa0\x29\x3a\xe1\x45\xa2\x5c\x91\x4a\x93\xbe\x40\x10\x09\x49\xa8\x25\x90\x21\xc0\xd8\xee\xa7\xef\x10\xe0\x1f\xf0\x9f\x44\xc9\x4c\x2f\xf6\xf4\xd1\xc6\x12\x0b\x2c\x76\x7f\xfb\x5b\x2c\x74\x79\x09\xfe\xb2\x27\x9b\x08\x71\x0c\x16\xe1\xc5\x85\xfa\xb7\xc3\x11\xc7\x7b\x4c\xf9\x0d\xde\x10\x7a\x61\xcc\x4d\xdd\x35\x81\x63\x7c\x34\xa7\x3a\xb0\x6e\x81\x3d\x73\x81\xf9\xc5\x72\x5c\x07\x2c\xbd\x2d\x22\xf4\x5b\x8c\xa3\xa7\x25\x18\x9b\xb7\xfa\x62\xe2\x02\xe3\xa3\x3e\xd7\x0d\xd7\x9c\x03\xc7\x74\xc1\x0e\x71\x42\x47\xc0\x98\x4d\x26\xc9\x34\xf2\x4f\xb8\xc1\x14\x47\x68\x07\x3d\x02\xfe\xde\xac\xdc\xa4\x7e\x97\x65\xb9\xfa\xcd\xc4\xac\xae\x6a\xb5\x0b\xbc\xfb\xe5\x85\x76\x01\x00\x00\x4b\xe2\x2f\x81\x63\xce\x2d\x7d\x32\xbc\x90\xff\x59\x11\xce\x96\xe0\xb3\x3e\x4f\x96\xaa\x5d\x5f\x0d\xc4\xd7\xf6\x62\x32\x19\x4a\x01\xb1\xad\x87\x20\xba\x2f\xa4\xfe\x7a\x35\x38\x75\x67\xd5\x49\x03\xba\x26\xd1\x1e\x71\x12\x50\xb6\x04\x96\xed\x9a\x1f\xcc\x39\x58\xd8\x8e\xf5\xc1\x36\xc7\x55\x71\x9f\xac\xd7\xc4\x8b\x77\x3c\x31\xed\x6c\x71\x33\x31\xb5\xd1\xbb\xe1\xbb\xda\x5a\xb7\x88\x6d\x7b\x59\x26\xc8\x26\xc4\x64\xb3\xe5\x4b\x70\x63\x7d\xb0\x6c\xb7\x75\x7d\x7b\xec\x13\x44\x21\x27\x7b\xdc\x41\x36\xba\xdf\x61\x18\x05\x01\xef\xd3\xa2\x14\xed\x31\xf4\x76\x88\xec\xfb\x9f\x3a\xa0\xde\xf1\x7d\x85\x11\xfe\x4e\x82\x98\x41\xe1\x71\xf0\xb9\x47\x91\xe9\xc6\x8f\xbc\xdf\x19\xe5\x64\x8c\xfc\xe7\xf8\x96\x38\x8a\x36\xb8\x57\x53\x4a\xe5\x9d\xfc\xe4\x3b\x8e\x18\x09\x68\x57\x39\xb8\xc5\x8f\xc5\x4a\x47\xcf\x5e\x29\x8f\x10\x65\xc8\xe3\x62\x6a\xc4\xb6\x98\x2d\x81\x6b\x7e\x71\xeb\xc3\x0c\x86\x51\xe0\x61\xc6\xb0\xbf\x04\xae\x65\x7f\xb5\x6c\x57\x1b\x0d\x72\x0c\xbc\x52\xa6\x4e\xa3\x3f\xc2\x88\x27\xd2\x63\xdd\x35\x5d\x6b\x6a\x16\x71\x97\x03\xe7\x62\x3e\x37\x6d\x17\x26\xa3\x8e\xab\x4f\xef\xb2\xe8\x09\x7c\xb2\x26\x27\x7e\x0b\x66\x36\x58\xdc\x25\x1f\x34\xcd\x2b\x26\xbe\x9b\x5b\x53\x7d\xfe\x15\x7c\x32\xbf\x82\xe5\xdd\x27\x78\x23\x60\x13\x68\x09\x62\x0e\xa4\xee\x85\x6d\xfd\x63\x61\x4a\x09\xcb\x7f\x94\x22\x1f\x85\x53\x6a\x12\x79\x52\x41\x63\x66\x3b\xee\x5c\x4f\x0e\x6d\x69\x50\x0e\xdd\xc2\x56\x1f\x85\x25\x3f\xa3\x1d\xf1\xff\x60\xc9\xe1\x1a\x1f\x4d\xe3\x93\xd6\x68\x6d\xcb\x91\xdb\x9a\xcd\xc1\x1f\xce\xcc\x86\x9f\xf5\x89\x35\x6e\x14\x1d\xa4\x8a\x2d\x7b\x6c\x7e\x51\x17\x97\xa2\x97\x96\xe1\x58\x8b\x9c\x2b\xfc\x51\x53\xbd\xb3\x2e\x39\x15\x18\x97\x89\xaa\x88\x57\x97\xbd\x4b\xb1\xa0\x64\xa2\x26\x80\x68\x59\x90\x91\x39\x88\x96\xfb\x4a\x8b\xe4\x34\xf7\x07\xad\xf0\x8d\xc1\xc5\x00\x98\xf6\x07\xcb\x36\xdf\x5b\x94\x06\xe3\x9b\x52\x3e\x76\x4c\xf7\x7d\xcc\xd7\xef\xf6\xab\x37\x79\x38\xa4\x7f\xc3\x98\x12\x2f\xf0\x71\x12\x0f\xf3\xd9\x3f\xe1\xed\x6c\x3e\xd5\xdd\xf7\xc6\x6c\x7a\x37\x37\x1d\xc7\x1c\x27\xa7\x0f\x6f\x26\x33\xe3\x13\x74\xac\x7f\x99\xef\xdf\xfc\x80\x8c\xad\x1c\x70\x63\xde\x56\x80\x64\xf5\x24\xcc\x08\x93\xe1\x67\x03\x23\xa1\x61\xcc\xa1\x17\xc4\x94\x1f\xcf\xc8\x41\xcc\x4f\x90\xfe\x8e\x76\x31\x3e\x9c\xba\xd7\xb8\x2a\xd0\x84\x1f\x35\x68\x6a\x84\xd2\x06\xb9\x4e\x78\xdf\x1b\x7d\xa8\x42\x78\x66\x9f\xca\xb8\x92\x0c\x8e\x59\x30\x42\x0f\x25\xfc\xfd\xb9\x20\x54\x5d\x52\xba\xa1\x93\xd7\x55\x45\x60\x05\x35\xcb\x38\x7c\x3b\x9b\x9b\xd6\x07\x5b\x0a\xde\x96\x04\x4b\x78\x53\x0b\x91\x01\x98\x9b\xb7\xe6\xdc\xb4\x0d\x33\x27\xc6\x39\x74\x27\x1b\x1c\x9b\x13\x33\xd9\xa0\xee\x18\xfa\xd8\x54\xb6\x6c\xcf\x80\x6e\xb8\xd6\xcc\x06\xcd\xa9\xa0\x82\xf0\xd5\x84\xa0\x80\x96\x22\x99\x41\x69\xcd\x9d\x0f\x7e\x95\x42\x63\xf6\x71\xc9\xe8\x5d\x3e\x3c\x8c\xa9\x8a\xfc\x2b\x44\x56\xe4\xfb\x11\x66\xac\x1d\x55\x33\x81\x1c\x03\xde\x9c\x8e\x01\xa9\x6f\x54\xe1\x8d\x44\x8c\x43\x86\x31\x2d\x42\x23\x67\x99\x9b\x42\xe1\x6f\x57\x83\x74\x86\x62\x14\xc6\xd1\x4e\x2d\xd1\xae\x06\x3f\x2d\x08\x54\x83\x58\xcf\x0c\x7a\x88\x48\xa5\x42\x8a\x6c\x76\x0e\x99\x83\x52\x1f\x3f\x96\x64\xdd\xc4\x66\x9a\x30\x5d\xdd\x89\x53\x99\x4e\x0e\x9f\xca\xbe\x42\x67\x17\xe9\xbc\xdd\xd5\x55\xd4\x49\x06\x8f\x95\x42\x15\xd6\xd9\x67\x9a\x94\xc4\x23\x3d\xf3\xa6\xc5\x64\x72\x0c\x7a\x01\xa1\x2b\xc4\xf0\xd1\x42\x23\xbd\x65\xc8\xa4\x9f\x4d\x8e\x12\xf2\x1a\xc4\xbc\xaf\x22\x34\x9b\x8e\xd6\x33\x7f\x45\x82\x85\x98\xfa\x30\x0e\x7d\xe9\xcb\x5d\xb6\xcd\xf0\xb7\x18\x8b\x92\xfd\x0c\x5e\x96\xcd\xe1\x45\x24\xe4\x90\x91\x0d\x44\x6c\x2f\xa9\xc7\x79\x5b\x55\x66\x12\x25\xea\x59\x33\xbd\x14\xbc\xb3\x44\xd0\x1d\xa4\x2b\x42\x44\xc5\xc5\xaa\xf7\x97\x69\x0a\x2a\x23\x68\x37\x9a\x72\x40\x73\x85\x56\x55\x60\xa0\xac\x9b\xd7\x29\xd8\x29\xfa\x15\x98\x15\xaa\x3f\x4b\x6f\xd3\x52\xb7\x6b\x2e\x1b\x83\x98\xab\x05\x63\x1e\x73\x75\x69\x31\x67\x27\x90\x17\x92\xcd\x10\xdf\x2c\xdb\xc0\xe7\x6a\x00\xf8\x0a\xb2\x03\xac\x11\x22\xf1\xdf\x0e\xe9\xe0\x00\x58\x57\x44\x1b\x23\xa4\xe6\xfe\xc4\x5f\x0e\x4b\xb3\x36\x05\x4f\x7e\x3c\x19\x55\xc0\xac\x88\xb7\x6c\x9a\xb2\x07\x13\x35\x1e\xcf\x8e\x9d\x06\xc5\x75\xb2\xd2\x53\xe8\xbe\x68\xaf\x92\xb7\x02\x2f\x82\x74\xb4\xa7\xbe\x04\x73\x8e\xa7\x4e\xfe\x14\xe2\x72\xdf\xe4\x39\xd9\x31\x8c\x57\xf0\x1e\x3f\xf5\x94\x6b\xb3\xd9\xce\xcf\xb7\xf2\xce\x01\x7f\x8b\x49\x84\xfd\x24\x77\x53\xc4\xe3\x08\x37\xb4\x6a\x94\xab\x93\xd1\xef\x57\x85\x45\xde\x9e\x6b\x91\x2c\x9c\x76\x84\xf1\xe7\xac\x9e\x30\xc1\x9d\x78\x47\xce\x94\x88\xc2\xd5\x13\x6c\x85\xc0\x17\x53\x73\xcd\x64\x10\xb6\x96\x5c\x0b\x55\xa4\x16\x62\x43\xe9\xff\x2d\xe4\x45\x7e\xf8\x3f\xe4\x10\xcd\x0b\x70\x92\xd3\xba\x79\xca\xd1\xbf\x7e\x7a\x3d\xa5\x81\xea\x55\x7e\x8e\xff\xf5\x2b\xfc\xb2\xdf\xb6\x5c\xde\x97\x84\x1a\xae\xed\xe5\xf6\x8e\x12\xa5\x59\x1c\x06\x84\x8a\x1d\x89\xd3\x1a\x36\xd0\x13\x60\xcc\xa6\x53\xd3\x76\xc1\xaf\x14\x63\x1f\xfb\x60\x1d\x44\x20\xc2\x6b\x1c\x25\xb5\x01\x03\x84\x02\xbe\x25\x0c\x78\xc1\x2e\xde\x53\x10\x44\x3e\x8e\x7e\xad\xd7\xc7\xce\x74\x09\xb4\x3a\x4a\x69\xd7\x6f\xdf\x0e\xea\x3a\x08\x05\x1e\x62\x98\x81\x87\x2d\x8e\x30\xe0\x5b\x0c\x74\x67\x0a\x92\x51\x06\x78\x00\x56\x18\x84\x28\x62\xd8\x07\x0f\x84\x6f\x01\x0b\x31\xf6\xeb\x4a\x67\x71\xc8\x3b\x52\x3b\x69\xb0\x57\x58\xbe\xa7\x17\xec\x55\x86\x96\xfe\xfb\x87\x52\x34\x69\x53\x85\xe1\x14\x4a\xbb\x91\x34\x39\x41\x41\x96\x0a\xb8\x29\x66\x2a\x07\x68\x50\xc2\xac\x67\x11\xb5\xaa\xf2\xff\x33\xb5\x23\x4d\xa6\x9a\x8f\x9d\xc8\xd2\xba\x3b\xda\xd2\xc7\x2b\xc2\x21\xda\xcb\xae\xd1\x91\x6e\x4f\x01\x2e\x4e\xbc\x07\xc1\x5a\xa0\x89\x40\x72\x01\x25\x02\xbc\x52\xdd\x02\xdb\x92\x61\xfe\xf8\x6b\x91\xa7\xfd\x67\xea\x92\x4e\xd9\x45\xd9\x0e\x71\xcc\x38\xac\x37\xa6\xfa\xec\x82\x28\x8e\x5c\x39\xa1\x6e\x51\xa9\x4c\x55\x44\xc7\x9f\x95\xca\x5b\x17\xf4\xc3\xc2\xf5\x50\xaf\xa3\x50\x3f\x11\x27\xd9\xd0\xa5\x69\x3b\xe2\x83\x3d\x94\x62\xde\x71\xe2\xf9\xc9\x3c\xa5\x10\xe8\xf8\xb1\x21\x7c\x39\x4d\x84\x8a\x57\xbf\xec\x0c\x27\x1e\x2b\x75\xab\x15\x7b\xed\x73\x77\xab\xef\x28\xda\x63\xf5\x3d\xcd\xf5\x9b\xfa\xcb\x38\xf1\xda\x2a\x59\xd3\xb9\xbd\xa2\xc6\x09\x65\x65\x99\x15\x2d\xd7\x8a\x5a\x70\x79\x09\x46\xe0\x12\x18\x38\xe2\x64\x4d\x3c\xc4\xb1\xfb\x14\xe2\x21\xb8\x06\x97\xc0\xe1\x11\x46\xfb\xe4\xef\xf4\x02\x39\x5e\xed\x08\xdb\xe2\xe8\x59\x2b\xcc\x71\x51\x61\x8e\x08\x18\xc9\x42\xad\x94\xc4\x55\x56\x93\x61\x62\xa1\x9f\x91\x4d\xb9\x7d\x75\xde\xb9\x79\x85\x9e\x52\x2b\x9c\xf9\x95\xcb\x81\xd1\xd9\xd5\x78\xd7\x97\x05\x79\x63\xbf\xbf\xf7\x56\x82\xea\x43\xc4\x64\xd1\x3e\x35\xc7\xd6\x62\x2a\x8a\xdf\x36\xb9\x7f\x8b\xe2\xa3\x10\x2c\xc6\x89\x0f\x11\x87\xd9\xd3\xc5\x63\x8e\xde\x55\x0e\xaf\xd7\xd8\xe3\xe4\x3b\xee\x9a\x53\xd3\x5b\xfb\xcb\x4b\xa0\xfb\x3e\x49\x6c\x8a\x76\x60\x4d\xf0\xce\x97\x29\x14\x23\xf6\x04\x08\xf5\xf1\x23\xa1\x9b\x24\xe7\x32\xe1\xc1\x20\xf1\x7f\x96\x52\x8b\x98\x6f\x83\x48\xb9\x54\x18\x5d\x0f\x72\x2a\x21\x8b\x12\xd2\x68\x05\x2f\xa0\x3c\x29\x09\xcb\xb7\x34\xa3\xdf\xaf\xcf\x7d\x1b\xc3\x20\x85\x0c\xae\xe1\x43\xb7\xeb\x84\x1d\xa2\x9b\x18\x6d\xfa\xb8\x21\xe2\xdb\x78\xbf\xa2\x88\xec\x64\x1b\x58\x7d\x83\x47\xf8\xae\x1c\x0a\x27\x3c\xa2\x59\x63\x0c\xbd\x38\x4a\x42\xfa\x29\x85\x87\xdf\xf2\xc6\x72\x32\xd8\x47\x2f\xbc\xda\xe8\x63\x70\x4d\x76\x1c\x47\x5d\x3b\x59\x2b\xe2\x43\xc6\x45\xc0\xab\x76\xcc\xe4\x7f\xd1\x3d\x0f\x87\x1c\xfb\xbf\xfc\xe4\x4f\x0c\x9b\x98\x9d\xc0\xd0\xc3\xad\x22\x21\x72\x80\xa1\xb5\xbe\x6d\xa9\x30\xb5\x53\x5e\xb8\x1c\x58\xc9\x5d\x86\xe8\xa2\x39\xa3\xa6\x97\xb2\x76\x2f\xdb\x59\x9e\x20\x4f\xd2\x5d\x7d\x1e\x20\x74\x2f\x28\xf9\x16\xd7\x9e\xcb\x28\x06\x18\xa6\x89\x7d\xa8\xe4\xe5\x96\xf7\x98\x62\x42\x43\xcd\x26\xe9\x15\x4e\x29\xc3\xb4\xdc\xe0\xa8\x32\x83\x81\x48\xc9\xca\xbf\x04\x78\x55\x39\x9d\x51\xb7\x47\x8d\xf7\x55\x8f\xfa\x84\xc7\x41\x72\x3b\x5d\x2e\x4b\x84\x64\xc7\x3e\x98\x90\x15\xb7\x5c\x3a\x2f\xde\x91\x56\xb3\x4b\xcb\x77\x37\xc4\x77\x64\xd8\x6a\x4a\x0c\xb7\x08\xdb\x48\x6e\x55\x30\x2e\x79\xaf\x34\xbc\xb8\x68\x14\xd5\xd3\x7c\xa0\x65\x99\x41\x1b\xfd\x6d\xd4\x70\x8d\x26\x4d\x22\x73\x80\x2b\x52\x80\x56\x4e\x09\x2d\x9f\x4c\x72\xd8\xd6\x0a\x08\x6f\x3b\x2c\x89\xbe\x5a\x0a\xc3\x72\x25\x2f\x9a\x90\xb3\x38\x0c\x83\xe8\x40\xfb\x26\x15\xc0\x3e\xec\x9d\xfa\xa6\x53\x77\xe7\x16\x67\xe7\x87\xde\x3b\x4b\x5d\xcb\x8a\xc6\x2c\x90\xd9\xbc\x3d\x0f\xac\xef\x61\x6e\xf7\x1c\x59\x1b\x4e\xa2\x57\x0c\x56\x9c\x9d\x75\x09\xe4\xf2\x0a\x89\xdf\xb6\xc6\xda\x87\xbc\x3d\xbb\xb5\xb4\xfd\xa5\x9d\xb5\x52\x9b\x42\x19\x0e\x62\x9e\xdf\x8d\xb7\x36\x37\x5e\x74\x98\xc6\xf4\x9e\x06\x0f\x14\x1e\xa9\x9f\x4f\x2a\x60\x0b\x86\x5c\xff\x01\x18\x61\xe9\xf3\xa3\x8e\x9c\xad\xcf\x1f\xeb\xfc\x19\x71\x7a\xe2\x3d\x7b\x4f\x85\x5b\x15\x18\xd2\x53\xf6\xea\x2c\xb1\x4a\x67\x44\xb7\xa8\xde\x8f\xaa\xe8\x6b\xa1\x33\x65\xa9\x41\x0b\x11\x55\x17\x13\xfc\xb0\xcb\x7c\x25\x8a\x17\x42\xa1\xe4\x12\xf5\x37\xdd\x05\x32\x34\x7f\xd2\xd8\x6f\x68\x17\x77\x1f\x5f\xe9\x23\x1f\x14\x86\xbb\x84\x96\x8a\x1f\x40\x70\xc4\xe3\x43\x4f\x9f\xc3\x10\xaa\xbf\x57\xa8\x5d\xdd\x23\x8e\x72\x81\x86\x71\x14\x92\x96\xe1\x46\xf7\x96\x29\x82\x88\x2e\xdd\x4f\x60\xe0\xff\x06\x00\x00\xff\xff\xfa\xf5\x06\x12\x04\x3c\x00\x00")

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

	info := bindataFileInfo{name: "migration/000_init_schema.sql", size: 15364, mode: os.FileMode(420), modTime: time.Unix(1525307810, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migration001_supportsSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x92\x4f\x6f\xda\x30\x18\xc6\xef\xf9\x14\xef\xad\xa0\x15\x29\x48\x1c\xa6\x4d\x1c\x8c\xf3\xd2\x5a\x4d\x1c\x64\x3b\x5b\xd9\xc5\x09\xc1\xa0\xb4\x24\x41\xb1\x73\xd8\xb7\x9f\x1c\x10\x74\xd5\x36\x69\x55\x1b\xf9\x90\xd7\x7f\x9e\xe7\xe7\xc7\xef\x64\x02\x9f\xea\x6a\xdf\x15\xce\x40\x76\x0c\x82\x97\xb5\x74\x85\x33\xb5\x69\xdc\xc2\xec\xab\x26\x20\xb1\x42\x01\x8a\x2c\x62\x84\xb6\x77\xc7\xde\x01\x89\x22\xa0\x69\x9c\x25\x1c\xf2\xf2\x50\x54\xb5\xae\xb6\x39\xd0\x7b\x22\x46\xb3\x70\x3c\xfc\x10\xea\x4f\x49\x54\x70\x28\x5c\xd5\x4c\xfd\xfe\x98\x28\x3c\x97\x7a\x6f\x1a\xd3\x15\x07\x5d\x56\x5f\xff\x6c\x8e\xcd\xf6\x2d\x58\xcb\x54\x20\xbb\xe3\xf0\x80\x6b\xc8\x77\xcf\x7a\xc0\xcb\x61\x74\xe5\x1c\x83\xc0\x25\x0a\xe4\x14\xe5\x19\xff\xd5\x7a\xca\x21\xc2\x18\x15\x02\x25\x92\x92\x08\xfd\x4c\xb6\x8a\x3c\x3e\x4f\x81\x50\xc5\x52\xfe\x0e\xd8\xb6\x3f\x1e\xdb\xee\x55\x9c\x9d\x29\x9c\xd9\xe6\xe0\xdd\x14\x4b\xbc\xa3\x02\x9e\xc5\x31\x44\xb8\x24\x59\xac\x80\x66\x42\x20\x57\xda\xaf\x4a\x45\x92\xd5\x07\xa1\xd4\xed\xb6\xda\x55\xff\xc9\xf2\x22\xaa\xf7\xe2\xa4\x02\xbd\xdc\x09\x94\x2d\x07\x08\x7c\x64\x52\x49\xc8\x9f\xda\x8d\xb6\xae\x70\xbd\xcd\x83\x51\x00\xa7\x89\xa6\xa8\x4d\x0e\xc3\xf7\x8d\x88\x37\xb6\xe5\xe5\xae\xb7\x5e\xf6\x50\x58\xa7\xed\xcf\xa6\x3c\xe9\xfe\x3d\x90\x9b\x30\x0c\xa7\x93\x61\x40\x18\x7e\x19\xc6\xcd\x20\x51\x59\x6d\xfb\xb2\x34\xd6\x0e\x1a\x8a\xf1\x35\xe3\x6a\x34\x1d\x5f\xce\x86\xbf\x7b\x9a\xae\x6b\x3b\x5d\x1b\x6b\x8b\xbd\xc9\x41\xe1\xa3\xba\x0d\x02\x80\x95\x60\x09\x11\xeb\x53\x8b\x1f\x9f\xf5\x53\xbb\x39\x47\x00\xa3\xeb\xfd\xc7\x41\x30\x46\x7e\xc7\x38\xce\x59\xd3\xb4\xd1\xe2\xfa\x66\xf7\x44\x48\x54\xf3\xde\xed\x3e\xd7\x9b\xd9\x25\x82\x73\xad\xfb\xa6\x2a\xdb\xad\xf1\x19\x88\xf4\xbb\x5e\xa6\x22\x21\x6a\x4e\xd3\x64\x25\x50\x4a\x8c\xbc\xb1\x5e\xc4\x29\x7d\xd0\x92\xfd\xc0\xf9\xec\x1f\xef\xfa\x2b\x00\x00\xff\xff\x84\xc3\x3d\xac\x68\x04\x00\x00")

func migration001_supportsSqlBytes() ([]byte, error) {
	return bindataRead(
		_migration001_supportsSql,
		"migration/001_supports.sql",
	)
}

func migration001_supportsSql() (*asset, error) {
	bytes, err := migration001_supportsSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migration/001_supports.sql", size: 1128, mode: os.FileMode(420), modTime: time.Unix(1525738256, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migration002_decimalsSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x8d\x41\xcb\x82\x40\x14\x45\xf7\xfe\x8a\xbb\x53\xf9\x3e\x63\x08\x02\xc1\xd5\xa8\x93\x09\xaf\x27\xd4\xbc\x75\x8a\x3c\x43\xc8\x8a\xb0\x7e\x7f\x1b\x77\x05\xdd\xc5\x5d\x9c\xc5\x39\x49\x82\xbf\x69\x3c\x3f\xba\x59\x21\xf7\xc0\x92\x77\x07\x78\x9b\x93\x43\x7f\xe9\xc6\x09\xc5\xce\x72\xe5\x50\x34\x24\x7b\x46\xab\xc3\xa0\xfd\x3c\xbe\xf4\xd4\x4d\xb7\xe7\x75\x6e\xbf\xa1\xbc\xae\x6a\xf6\xd1\xda\xc4\x10\x3e\xd6\x15\xbb\x12\xdc\x78\xb0\x10\xa1\x74\x5b\x2b\xe4\x61\x90\xfd\xce\x0d\xaa\xed\xf2\x65\x23\x39\xb9\x68\x93\xfe\xa7\xf1\xa7\x2d\x34\x2b\xb3\x2c\x44\x16\x04\xef\x00\x00\x00\xff\xff\xde\x43\xaf\xbd\xd8\x00\x00\x00")

func migration002_decimalsSqlBytes() ([]byte, error) {
	return bindataRead(
		_migration002_decimalsSql,
		"migration/002_decimals.sql",
	)
}

func migration002_decimalsSql() (*asset, error) {
	bytes, err := migration002_decimalsSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migration/002_decimals.sql", size: 216, mode: os.FileMode(420), modTime: time.Unix(1526446354, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migration003_dht_trackingSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x95\xcf\x6f\xdb\x20\x14\xc7\xef\xfc\x15\xef\x88\xb5\x56\x6a\xa5\xdd\x76\xa2\x0e\x55\x59\x1d\x5c\x01\x99\x92\x13\x66\x36\x4a\x51\x1c\x62\xc5\x68\xdb\x9f\x3f\xd9\x4e\x63\xe7\x87\x92\x66\x6b\xb4\xdd\xf2\xc8\x97\xf7\x78\x1f\xbe\x0f\xdf\xde\xc2\xa7\xa5\x9b\xaf\x4d\xb0\x30\xa9\xd0\x30\x94\xc1\x04\xbb\xb4\x3e\x3c\xd8\xb9\xf3\x28\x16\x94\x28\x0a\x8a\x3c\x24\x14\xd8\x23\xf0\x54\x01\x9d\x32\xa9\x24\x64\x79\x69\xdc\x52\xe7\xaf\x36\x5f\x54\x2b\xe7\x43\x86\x30\x02\xc8\x5c\x91\x81\xa4\x82\x91\xe4\xa6\x09\x3b\x55\xb3\x18\x3f\x11\x81\x3f\xdf\x45\xed\x0f\x12\x2b\x2a\x40\x52\x05\xa5\x09\xce\xdf\x43\x9c\x26\x49\x53\xaa\x0b\xf5\xdc\x7a\xbb\x36\xa5\xce\x5d\x5b\x93\x4f\x92\x4d\xba\xbe\x1c\x8c\x88\xa2\x8a\x8d\xe9\x56\x01\x23\xfa\x48\x26\x89\x82\x78\x22\x04\xe5\x4a\x37\xff\x4a\x45\xc6\x2f\xed\x5e\x57\x6b\xf3\xc3\xb8\xd2\x7c\x2f\x6d\x06\x8a\xf1\x19\xe3\x0a\xdf\x47\xbb\x15\x5e\xad\x29\xde\xa3\xab\xf5\x79\x19\x02\x78\x11\x6c\x4c\xc4\x0c\x9e\xe9\x0c\xb2\x6a\xa1\x0f\xa8\x01\x6e\x90\x45\x4d\x4e\xc6\x47\x74\xda\x10\xfc\xa5\x7b\x6c\xb8\x47\x78\x20\xda\xc9\x32\x88\xf6\x85\xbb\x9d\xe3\x5d\x12\x11\x8a\xbe\x1c\xf7\x00\xf5\x05\xfa\x53\x77\x54\xd6\xae\x37\x8e\xf0\xab\xc2\xf6\x0e\xb8\xbf\xfb\x4b\x0b\x2c\xfc\xea\xa7\xd7\x4e\x57\xba\x74\x75\xc8\x40\xd1\xa9\xba\x30\xe1\xd1\x9b\xe9\x60\xbf\x1d\xb6\x45\x18\xa7\x5c\x2a\x41\x18\x57\x90\xe5\x3e\xe8\xfd\xd2\xf1\x13\x8d\x9f\xf1\xc1\x89\x98\xec\xec\x98\x0a\xf8\x2a\x53\xae\xbf\x91\x84\x8d\x0e\x64\xd1\xf5\xc0\x77\xf6\xd9\xe0\x6f\x17\x3e\x0e\xff\x47\x0f\xf4\xda\x9a\x60\x8b\xcb\xa7\xb9\x34\x75\xd0\xb5\xb5\xfe\xc2\xad\x47\x6e\x7e\x80\x0c\xf0\x96\xd7\xcd\x89\xc1\x7b\xd7\x74\xbe\x75\x86\xb7\x4d\xee\x4b\x06\x3d\xe0\x41\x43\x11\xba\xb2\x33\xce\xbe\xda\xff\xb9\x67\xae\xfb\x11\x38\xe9\x90\xf3\x4f\xf7\x96\x5d\x6f\xa5\x7f\xfa\xba\x9f\xf6\x12\xfa\x1d\x00\x00\xff\xff\x70\xd5\x5a\x52\x0d\x08\x00\x00")

func migration003_dht_trackingSqlBytes() ([]byte, error) {
	return bindataRead(
		_migration003_dht_trackingSql,
		"migration/003_dht_tracking.sql",
	)
}

func migration003_dht_trackingSql() (*asset, error) {
	bytes, err := migration003_dht_trackingSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migration/003_dht_tracking.sql", size: 2061, mode: os.FileMode(420), modTime: time.Unix(1526786378, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migration004_new_indicesSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x92\xbf\x6e\xc2\x30\x10\x87\x77\x9e\xe2\x36\x40\x2d\x52\x2b\x75\xeb\x64\x12\xa3\x22\x99\x44\x22\x4e\xd5\xcd\x36\x89\x49\x5c\x11\x1b\xc5\x97\xfe\x79\xfb\x2a\x21\xa8\x20\x68\xd5\xa1\xd9\xee\x2e\x39\x7f\xfa\x4e\xbf\xd9\x0c\x6e\x2a\x53\xd4\x0a\x35\xa4\xfb\xd1\xe8\xb4\x4f\x50\xa1\xae\xb4\xc5\xb9\x2e\x8c\x1d\x11\xc6\xe9\x1a\x38\x99\x33\x0a\xd2\x35\xb8\x6f\x50\x02\x09\x43\x58\x46\x21\x7d\x01\xb9\xcc\x3f\xc4\xd2\x27\x7b\x6d\x51\xc2\x44\x1a\x2f\xfc\xa1\x26\x49\x30\x7d\xbc\xfe\x32\xb5\xf9\xdf\x99\xd9\x4e\x99\xea\x02\xb9\xd0\x9a\xe4\x79\xad\xbd\x6f\xa9\x5b\xad\x85\x3a\xb6\xff\x05\xfe\x41\xb6\x53\x8d\xfb\x6f\x13\x89\xb5\xb2\x5e\x65\x68\x9c\x15\xa5\xf2\x65\xc7\xbf\x05\xf9\xe6\x1a\x3c\xd6\xe7\x47\x81\x20\x5e\xad\x68\xc4\x61\xdc\x78\x9d\xc3\xd6\xd5\x50\xd4\x6a\xb3\x31\xb6\x80\xee\x37\x38\x90\x3d\xbc\x1b\x2c\xe1\xd5\x19\xeb\xc7\xc3\x5d\x32\x68\xa7\xad\x8f\x33\xf6\xc2\x68\xf3\xd9\x49\x09\x93\x5f\x78\x5d\xf3\xa8\x14\x66\x25\x74\x1c\x40\xd7\x7b\x0c\xa8\x11\xc4\x2c\x5d\x45\xfd\xe8\x3b\x01\xcf\x64\x1d\x3c\x91\xf5\xe4\xe1\x6e\x0a\x6d\x41\x82\x76\x3d\xa1\x1c\x76\x0a\x8d\xbd\x6f\xf7\x18\xe1\xb4\x6f\x45\xa1\xad\xae\xd5\x4e\x64\x06\xa2\x98\x43\x94\x32\x06\x21\x5d\x90\x94\x71\x18\x0f\x7d\xfa\x93\x18\x9f\x6b\xfc\x16\xe2\xaf\x00\x00\x00\xff\xff\xc1\x98\xdc\xea\xc1\x03\x00\x00")

func migration004_new_indicesSqlBytes() ([]byte, error) {
	return bindataRead(
		_migration004_new_indicesSql,
		"migration/004_new_indices.sql",
	)
}

func migration004_new_indicesSql() (*asset, error) {
	bytes, err := migration004_new_indicesSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migration/004_new_indices.sql", size: 961, mode: os.FileMode(420), modTime: time.Unix(1527384337, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migration005_remove_foreign_keysSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x92\xcd\x6e\xea\x30\x14\x84\xf7\x79\x8a\xd9\x01\xba\x20\x5d\xee\x5d\x74\xd1\x95\x71\x0e\x3f\xaa\x6b\x23\xc7\xa9\xd4\x95\xe5\x42\x00\x0b\x30\x51\x30\xef\x5f\x85\x50\x95\xaa\xe9\xcf\x82\xe5\x39\x33\x39\xdf\x44\xe3\xc1\x00\x7f\xf6\x7e\x5d\xb9\x58\x20\x2f\x93\xe4\x7a\xce\xa2\x8b\xc5\xbe\x08\x71\x54\xac\x7d\x48\x98\x30\xa4\x61\xd8\x48\x10\x16\x3b\xe7\xf7\x09\x90\x6a\x35\xc7\x58\x69\x9a\x4d\x24\x1e\xe8\xb9\x11\xac\x7f\x59\x6d\xed\xbf\xfb\xf6\x6b\x14\x96\xbf\xe6\xf8\x50\x9e\x62\x1b\xe7\x2c\x34\x9c\xe1\x0d\x38\xc7\x53\x59\x1e\xaa\x56\xd2\x45\xba\xb0\xfa\x09\xc0\xa7\x4c\x4e\x08\x5c\x89\xfc\x51\x22\x56\x2e\x1c\xdd\x22\xfa\x43\xb0\x1b\x77\xdc\x7c\x5a\x58\xbf\xc4\x13\xd3\x7c\xca\x74\xf7\xee\x6f\xaf\xfe\x5c\x33\x5e\xb3\x33\x32\xe8\xec\x5c\xf4\x61\xd8\xa9\xcf\x09\x66\xe8\x6d\x61\xd7\x45\x28\x2a\xb7\xb3\x0b\xdf\x81\xcc\x85\x40\x4a\x63\x96\x0b\x73\x1e\xea\x18\x2c\x4d\xc1\x95\xcc\x8c\x66\x33\x69\xb0\xda\xda\x2b\x72\x02\xe0\xc3\x6f\x74\x5b\x62\xf5\xce\x2e\x4d\x63\xd2\x24\x39\x65\xd7\xd1\xd1\xad\x4d\x8d\x03\x50\x12\x29\x09\x32\x04\xce\x32\xce\x52\x7a\xdf\xe7\xf3\xb4\x8e\x2d\x15\x18\x37\x33\x25\x6f\xd0\xc6\xe1\x14\xbf\xa8\xbd\x51\x9a\x2e\xfe\xf7\x7f\x72\x7c\xf3\x02\x5f\x03\x00\x00\xff\xff\x12\x29\xde\x2e\xf8\x02\x00\x00")

func migration005_remove_foreign_keysSqlBytes() ([]byte, error) {
	return bindataRead(
		_migration005_remove_foreign_keysSql,
		"migration/005_remove_foreign_keys.sql",
	)
}

func migration005_remove_foreign_keysSql() (*asset, error) {
	bytes, err := migration005_remove_foreign_keysSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migration/005_remove_foreign_keys.sql", size: 760, mode: os.FileMode(420), modTime: time.Unix(1530938161, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migration006_add_height_indexSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd5\x55\xd0\xce\xcd\x4c\x2f\x4a\x2c\x49\x55\x08\x2d\xe0\xe2\x42\xe6\x07\x97\x24\x96\xa4\xe6\xa6\xe6\x95\x38\xa5\xa6\x67\xe6\x71\x39\xfa\x84\xb8\x06\x29\x84\x38\x3a\xf9\xb8\x2a\x24\x24\xe7\x24\x66\xe6\x26\x28\x38\xba\xb8\x28\x78\xfa\xb9\xb8\x46\x28\x24\x78\xa6\x54\xc4\x7b\xa4\x66\xa6\x67\x94\x24\x28\x68\x24\x64\x40\x59\x8e\xc1\xce\x9a\xd6\xd8\x4d\x75\xcd\x4b\x01\x04\x00\x00\xff\xff\xd4\x45\xc4\xac\x7e\x00\x00\x00")

func migration006_add_height_indexSqlBytes() ([]byte, error) {
	return bindataRead(
		_migration006_add_height_indexSql,
		"migration/006_add_height_index.sql",
	)
}

func migration006_add_height_indexSql() (*asset, error) {
	bytes, err := migration006_add_height_indexSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migration/006_add_height_index.sql", size: 126, mode: os.FileMode(420), modTime: time.Unix(1530911689, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migration007_more_decimalsSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd5\x55\xd0\xce\xcd\x4c\x2f\x4a\x2c\x49\x55\x08\x2d\xe0\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x48\xca\xc9\x4f\xce\x56\x70\xf6\x70\xf4\x73\x77\x55\x70\xf6\xf7\x09\xf5\xf5\x53\x48\xc9\x4c\x4b\xcb\x4c\x2e\xcd\x29\xa9\x44\x66\xba\xf8\x87\x3a\xf9\xb8\x6a\x98\x1a\xe8\x58\x68\x2a\xf8\xf9\x87\x28\xf8\x85\xfa\xf8\x28\x58\x03\x02\x00\x00\xff\xff\x9f\xa4\x9c\x44\x5c\x00\x00\x00")

func migration007_more_decimalsSqlBytes() ([]byte, error) {
	return bindataRead(
		_migration007_more_decimalsSql,
		"migration/007_more_decimals.sql",
	)
}

func migration007_more_decimalsSql() (*asset, error) {
	bytes, err := migration007_more_decimalsSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migration/007_more_decimals.sql", size: 92, mode: os.FileMode(420), modTime: time.Unix(1534082603, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migration008_schema_fixSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x55\x4f\x8f\x9b\x3e\x10\xbd\xf3\x29\xe6\x96\x44\xbf\x5d\x69\x73\xf9\xf5\xd0\x93\x17\xbc\xbb\xa8\xc4\xac\x8c\xa9\xb4\x27\xcb\x01\x27\xb1\x02\x06\x81\x69\xd4\x6f\x5f\x39\x90\x06\x96\x1c\xda\x86\xaa\x7f\xb6\x27\x60\x06\xde\x30\xef\xbd\x19\xdf\xde\xc2\x7f\xb9\xda\x56\xc2\x48\x88\x4b\xc7\xe9\x3f\x47\x46\x18\x99\x4b\x6d\xee\xe5\x56\x69\xc7\xa3\xe1\x33\x30\x74\x1f\x60\x28\xa5\xac\x78\x92\x09\x95\xf3\x64\x27\x93\x7d\x59\x28\x6d\xde\x5f\xfe\x18\xeb\xf4\x5b\x61\xa7\x47\x3c\xff\xe8\x34\x58\xd7\xa3\x28\x5d\x36\x86\x8b\x34\xad\x64\x5d\x5f\x0f\x57\x34\xe6\x7a\x3c\x14\x30\x4c\x3b\xc0\x0e\xc9\x01\x38\x56\x71\xc3\x20\x5e\x11\x30\x62\x7b\x33\x0e\xf1\xa6\xca\x6c\xd8\x7d\x42\xe4\x11\x9f\x12\x49\x25\x85\x91\xe9\xe9\xca\x85\x01\x0f\x31\xcc\xfc\x15\x06\x12\x32\x20\x71\x10\x80\x87\x1f\x50\x1c\x30\x70\x63\x4a\x31\x61\xdc\x66\x23\x86\x56\xcf\x63\xbc\xbc\x48\xd5\x46\xc9\xf4\xeb\xcd\x77\x22\x42\x48\x20\x7e\xb6\x1f\x8c\x73\x13\x10\xd6\x2a\xe0\x00\x0c\xe9\xd9\x89\x7a\xb7\xfc\xff\xee\xcd\xd3\x63\x2a\xa1\x6b\x91\x18\x55\xe8\x11\x47\x9f\x44\xd6\xc8\x7f\x0c\x9d\x19\xe2\xe7\xe9\x1b\x32\x95\x09\x23\x6b\xc3\xfb\xaf\x1a\x95\xcb\x09\xaa\xaf\xb3\x22\xd9\x8f\xea\xe5\x32\x55\xa2\xad\x31\x9e\xfb\x6a\x2b\xcd\x8d\x33\xad\x66\x7f\x94\x62\x8d\xde\xeb\xe2\xa0\xdb\x83\xc5\xa1\x98\xa0\x15\x06\x16\x02\x88\xb5\x2e\xaa\x5c\x64\x6d\xc6\x32\x87\x3c\xef\x15\x45\x3f\x60\xe7\x1e\xc8\xef\x49\x48\x4b\x44\xe7\x93\x87\x90\x62\xff\x91\xc0\x07\xfc\xd2\x9d\xe8\x6a\xbd\xd9\xf3\xe5\x84\x75\x86\x6e\xe9\x4f\xc5\xfa\x33\xb7\x9b\x97\xab\x74\x10\x3e\xc5\x3e\x22\xea\x3e\x21\x3a\x7f\x77\xb7\xb0\x20\x14\xb9\x16\x3c\xc2\x0c\x66\x99\x30\x4a\x2f\x67\x16\x34\xb0\x3c\x75\x01\xbe\x95\x5a\x56\x56\x52\x35\x1b\x92\x7c\x7c\x18\x6f\x1b\x55\x73\xcd\x6b\xbe\xe1\x87\xe3\x6d\xbd\x39\x00\xf3\xc9\x8b\x4f\xd8\x7c\xb9\x18\x4b\x35\xbb\x9b\x5d\x00\x79\x63\x2b\xf0\xa4\x6b\x6b\x74\x12\x31\x8a\x7c\xc2\x06\xee\x71\xfa\xb6\x9a\x5f\xd0\x76\xe1\x50\xfc\x80\x29\x26\x2e\x8e\xfa\xda\xc3\xdc\xbe\xb0\x70\xc0\xf6\xe0\xe1\x00\xdb\x1e\x50\xe4\x22\x0f\xb7\xb1\xae\x2f\x12\x02\x72\x99\x1f\x92\xc9\x0f\xbd\xd1\x54\xf4\xff\x7e\xb2\xd9\x18\xd6\x1c\xca\x7f\x5c\xf2\xfd\xd9\x68\x03\x3f\x6d\x2a\x26\x6f\xe7\x95\x31\xc6\x04\x0e\xed\x31\x68\x6f\x60\x8c\x63\xe6\x57\x58\xa2\x6e\xca\xb2\xa8\xcc\x5f\x3a\xe9\x5f\x02\x00\x00\xff\xff\x60\xb2\x7b\x1f\x47\x0e\x00\x00")

func migration008_schema_fixSqlBytes() ([]byte, error) {
	return bindataRead(
		_migration008_schema_fixSql,
		"migration/008_schema_fix.sql",
	)
}

func migration008_schema_fixSql() (*asset, error) {
	bytes, err := migration008_schema_fixSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migration/008_schema_fix.sql", size: 3655, mode: os.FileMode(420), modTime: time.Unix(1534092768, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migration009_certificate_validationSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\xd0\x41\x4b\xc3\x30\x14\x07\xf0\x7b\x3e\xc5\x3b\xb6\xb8\x0d\x3c\xcb\x0e\x99\x79\xc3\x48\x4c\xa5\x4d\x50\x4f\x25\x36\x61\x7b\xd0\xc6\xd2\x46\xf1\xe3\x0b\x9b\xd3\xc9\xea\x0e\xbb\xe5\x25\xff\x17\x7e\xef\xcd\xe7\x70\xd5\xd1\x66\x70\x29\x80\xed\x19\x3b\xae\xab\xe4\x52\xe8\x42\x4c\xab\xb0\xa1\xc8\xb8\x32\x58\x82\xe1\x2b\x85\xd0\xb4\x8e\x3a\x06\xc0\x85\x80\xdb\x42\xd9\x07\x0d\x34\xd6\x4d\x18\x52\xfd\xe1\x5a\xf2\x60\xa4\x7e\x91\xda\x64\xd7\x39\xe8\xc2\x80\xb6\x4a\xcd\xa6\xf3\xfd\xf0\xd6\x84\x71\x0c\x93\x3d\x37\xd3\x20\x8c\xfe\x22\xaa\xd4\x02\x9f\x81\xfc\xe7\x31\x35\xfb\x23\xcf\x67\xd3\xd9\x5f\x66\x76\x22\xcf\x2f\x64\xda\x47\xc1\xcd\x41\xa8\x70\x6d\xe0\xbe\x90\x7a\x5f\x43\xb3\x75\x31\x86\x16\x0a\x7d\x38\x2e\x76\x0f\x35\x79\x58\xee\x33\x8b\xfe\xfd\xb5\xa5\x71\x1b\x86\x9a\x3c\xab\xd0\x7c\x5f\x9f\xae\x76\x09\xa6\xb4\xc8\x9e\xee\xb0\xc4\x9f\xef\xc8\x83\xac\x76\x6b\x66\x5c\x8b\x33\xbd\x6b\xae\x2a\xfc\x7f\xc4\xaf\x00\x00\x00\xff\xff\x81\x19\x77\xe9\x44\x02\x00\x00")

func migration009_certificate_validationSqlBytes() ([]byte, error) {
	return bindataRead(
		_migration009_certificate_validationSql,
		"migration/009_certificate_validation.sql",
	)
}

func migration009_certificate_validationSql() (*asset, error) {
	bytes, err := migration009_certificate_validationSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migration/009_certificate_validation.sql", size: 580, mode: os.FileMode(420), modTime: time.Unix(1541981688, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migration010_triggersSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x95\x51\xab\xda\x30\x1c\xc5\xdf\xfb\x29\xce\xdb\xad\xb8\x96\xfb\x32\xb8\x30\x7c\xa8\xe6\xef\xbd\x42\xd6\x8e\x98\xe2\xde\x4a\xae\x09\x52\xd0\x28\x6d\xdc\xe7\x1f\xd5\xe8\x5a\xea\xc6\x36\xaf\x62\xdf\x72\xfa\xe7\xe4\xe4\xc7\x69\x13\x45\x18\x6e\xca\x55\xa5\x9c\x41\xbe\x0b\xda\xcb\xb9\x53\xce\x6c\x8c\x75\x63\xb3\x2a\x6d\x90\x70\x49\x02\x32\x19\x73\x82\xd2\xba\x32\x75\x8d\x84\x31\x4c\x32\x9e\x7f\x4d\xf1\xae\xd6\xca\x2e\x0d\x58\x96\x8f\x39\x85\x9f\x5f\x3e\xbd\x0c\x90\x66\x12\x69\xce\x39\x18\x4d\x93\x9c\x4b\x3c\x3d\xc7\xcf\xfe\x79\xc2\x97\xcb\xdb\x91\xd5\xc1\xdf\x06\x71\x95\xb2\xb5\x5a\xba\x72\x6b\xdb\x61\x7e\xa8\xf5\xfe\x1e\x51\x26\x82\x12\x49\x90\x62\xf6\xfa\x4a\x02\x6e\x55\x94\xb6\x36\x95\x2b\x8e\x01\x92\x69\x93\x74\x96\xce\x49\x48\x64\x69\x3b\x6d\xe1\x11\x06\xc0\x34\x13\xa0\x64\xf2\x06\x91\x2d\x02\x20\xff\xc6\x1a\xcf\xd6\x6c\x00\xcc\x49\xb6\x95\xf8\xe8\x3f\xba\xa0\x0d\x91\xd2\x22\x5e\x56\x46\x97\xae\x50\x9b\xed\xde\xba\x00\x58\xbc\x91\xe8\x78\xc6\xa5\xc6\xe8\x30\xda\x0e\x55\xea\x0f\x03\xb1\xdf\x69\xe5\x4c\x07\x84\x3f\xd9\x9d\x40\x44\xc8\x38\xeb\x82\x78\x18\x38\xbe\x25\xa7\x6f\xe6\x9a\x9e\xfc\x7a\xdf\xa0\xf1\xab\xf8\xe4\x3c\xea\x29\x43\x84\x3d\x06\x88\x0e\x87\xd5\xe6\xfd\x2c\x0d\xce\x5c\x4e\x06\x67\x26\x5e\xb8\x41\x59\xba\x3c\xfe\xaf\x2e\xff\xca\x23\x42\xd8\xef\xc9\xb1\x3b\x1d\x1e\x8f\x09\x4e\x9b\xb5\xe9\x81\x63\xc4\xe9\x71\xc0\x5d\xe6\xd1\x8c\x5d\xcd\xe3\x77\x57\xd2\x2c\x65\xf4\x1d\xa1\x4f\x3a\xb8\xc1\x3d\xe3\x77\x38\xfc\x69\xfe\xe0\xff\x33\x00\x00\xff\xff\xd9\x39\x35\x55\x5e\x07\x00\x00")

func migration010_triggersSqlBytes() ([]byte, error) {
	return bindataRead(
		_migration010_triggersSql,
		"migration/010_triggers.sql",
	)
}

func migration010_triggersSql() (*asset, error) {
	bytes, err := migration010_triggersSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migration/010_triggers.sql", size: 1886, mode: os.FileMode(420), modTime: time.Unix(1546042864, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migration011_add_licenseSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd5\x55\xd0\xce\xcd\x4c\x2f\x4a\x2c\x49\x55\x08\x2d\xe0\xe2\x42\xe6\x07\x97\x24\x96\xa4\xe6\xa6\xe6\x95\x38\xa5\xa6\x67\xe6\x71\x39\xfa\x84\xb8\x06\x29\x84\x38\x3a\xf9\xb8\x2a\x24\xe7\x24\x66\xe6\x72\x29\x28\x38\xba\xb8\x28\x38\xfb\xfb\x84\xfa\xfa\x29\xe4\x64\x26\xa7\xe6\x15\xa7\x2a\xf8\x85\x39\x06\x39\x7b\x38\x06\x69\x18\x99\x9a\x6a\xea\x60\x55\x13\x5f\x5a\x94\x83\x57\x5d\x41\x51\x6a\x59\x66\x6a\x39\xaa\x1a\x6b\xec\x8e\x73\xcd\x4b\x21\xd9\xd9\x10\xcb\x3c\xfd\x5c\x5c\x23\x14\x3c\x53\x2a\xe2\x7d\xa0\x6e\xd7\x80\x3a\x10\xec\x1e\x1c\x8a\x42\x83\x7c\xe0\xea\x40\x1e\xc1\xaa\x36\x00\xea\x01\x0d\xa8\x4f\xf0\x38\x1e\x10\x00\x00\xff\xff\xda\x9e\xbc\xf1\x82\x01\x00\x00")

func migration011_add_licenseSqlBytes() ([]byte, error) {
	return bindataRead(
		_migration011_add_licenseSql,
		"migration/011_add_license.sql",
	)
}

func migration011_add_licenseSql() (*asset, error) {
	bytes, err := migration011_add_licenseSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migration/011_add_license.sql", size: 386, mode: os.FileMode(420), modTime: time.Unix(1549472714, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migration012_store_last_heightSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\xd5\x55\xd0\xce\xcd\x4c\x2f\x4a\x2c\x49\x55\x08\x2d\xe0\xe2\x42\xe6\x07\x97\x24\x96\xa4\xe6\xa6\xe6\x95\x38\xa5\xa6\x67\xe6\x71\x39\xfa\x84\xb8\x06\x29\x84\x38\x3a\xf9\xb8\x2a\x64\xe5\x27\xc5\x17\x97\x24\x96\x94\x16\x2b\x38\xba\xb8\x28\x38\xfb\xfb\x84\xfa\xfa\x29\x80\x44\x52\x15\xbc\x82\xfd\xfd\xac\xb1\x1b\xe4\x9a\x97\x02\x08\x00\x00\xff\xff\x00\xee\x2c\xf4\x71\x00\x00\x00")

func migration012_store_last_heightSqlBytes() ([]byte, error) {
	return bindataRead(
		_migration012_store_last_heightSql,
		"migration/012_store_last_height.sql",
	)
}

func migration012_store_last_heightSql() (*asset, error) {
	bytes, err := migration012_store_last_heightSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migration/012_store_last_height.sql", size: 113, mode: os.FileMode(420), modTime: time.Unix(1550937583, 0)}
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
	"migration/000_init_schema.sql":            migration000_init_schemaSql,
	"migration/001_supports.sql":               migration001_supportsSql,
	"migration/002_decimals.sql":               migration002_decimalsSql,
	"migration/003_dht_tracking.sql":           migration003_dht_trackingSql,
	"migration/004_new_indices.sql":            migration004_new_indicesSql,
	"migration/005_remove_foreign_keys.sql":    migration005_remove_foreign_keysSql,
	"migration/006_add_height_index.sql":       migration006_add_height_indexSql,
	"migration/007_more_decimals.sql":          migration007_more_decimalsSql,
	"migration/008_schema_fix.sql":             migration008_schema_fixSql,
	"migration/009_certificate_validation.sql": migration009_certificate_validationSql,
	"migration/010_triggers.sql":               migration010_triggersSql,
	"migration/011_add_license.sql":            migration011_add_licenseSql,
	"migration/012_store_last_height.sql":      migration012_store_last_heightSql,
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
		"000_init_schema.sql":            &bintree{migration000_init_schemaSql, map[string]*bintree{}},
		"001_supports.sql":               &bintree{migration001_supportsSql, map[string]*bintree{}},
		"002_decimals.sql":               &bintree{migration002_decimalsSql, map[string]*bintree{}},
		"003_dht_tracking.sql":           &bintree{migration003_dht_trackingSql, map[string]*bintree{}},
		"004_new_indices.sql":            &bintree{migration004_new_indicesSql, map[string]*bintree{}},
		"005_remove_foreign_keys.sql":    &bintree{migration005_remove_foreign_keysSql, map[string]*bintree{}},
		"006_add_height_index.sql":       &bintree{migration006_add_height_indexSql, map[string]*bintree{}},
		"007_more_decimals.sql":          &bintree{migration007_more_decimalsSql, map[string]*bintree{}},
		"008_schema_fix.sql":             &bintree{migration008_schema_fixSql, map[string]*bintree{}},
		"009_certificate_validation.sql": &bintree{migration009_certificate_validationSql, map[string]*bintree{}},
		"010_triggers.sql":               &bintree{migration010_triggersSql, map[string]*bintree{}},
		"011_add_license.sql":            &bintree{migration011_add_licenseSql, map[string]*bintree{}},
		"012_store_last_height.sql":      &bintree{migration012_store_last_heightSql, map[string]*bintree{}},
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
