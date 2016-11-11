// Code generated by go-bindata.
// sources:
// NAME-service/NAME-cli-client/client_main.gotemplate
// NAME-service/NAME-server/server_main.gotemplate
// NAME-service/generated/cli/handlers/client_handler.gotemplate
// NAME-service/generated/client/grpc/client.gotemplate
// NAME-service/generated/client/http/client.gotemplate
// NAME-service/generated/endpoints.gotemplate
// NAME-service/generated/transport_grpc.gotemplate
// NAME-service/generated/transport_http.gotemplate
// NAME-service/handlers/server/server_handler.gotemplate
// DO NOT EDIT!

package template

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

var _nameServiceNameCliClientClient_mainGotemplate = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x56\x4d\x6f\xdb\x38\x10\x3d\x4b\xbf\x62\x2a\x24\x80\x5c\x38\x74\x17\xd8\xbd\x18\xf0\x21\xcd\xb6\xdb\x00\xdb\xc2\x48\x02\xf4\x52\x60\x41\x53\x23\x89\x1b\x8a\x54\x49\x5a\x49\x40\xe8\xbf\x2f\x86\x92\x65\xbb\x89\x91\xfd\x0a\x90\x48\x1a\xcd\xbc\xf9\x7a\x8f\x4a\xcb\xc5\x3d\xaf\x10\x1a\x2e\x75\x9a\xca\xa6\x35\xd6\x43\x9e\x26\x19\x6a\x61\x0a\xa9\xab\xc5\x9f\xce\xe8\x2c\x4d\xb2\x52\xf1\x2a\x5e\x1b\x4f\x17\xe3\xe8\xaf\xf3\x56\x18\xdd\x8d\xb7\x52\x57\xd1\xea\x65\x83\x59\x9a\x26\x59\x65\x14\xd7\x15\x33\xb6\x5a\x3c\x2e\x34\xfa\x85\x30\xda\xe3\x63\x04\xa8\x8c\xa9\x14\xb2\x03\x97\xca\xb6\x62\x08\x93\xbe\xde\x6e\x98\x30\xcd\xa2\xbd\xaf\x16\x68\xad\xb1\x8e\xde\x2c\x16\x70\x57\x4b\x07\xb7\x68\x3b\x29\x30\x4d\x84\x92\xa8\xfd\x27\xae\x0b\x85\x16\xb2\x10\xd8\x75\x6c\x61\xcd\x7d\x0d\x17\x7d\x0f\x8b\x0a\x35\x5a\xee\xb1\x58\x08\x25\x17\xf5\xe0\x49\x55\x52\xb6\x21\xfc\xf5\x38\xd4\x7e\xac\x2e\xa9\xbd\x6f\xff\x49\x18\xf9\x67\x69\xd2\x6e\xa2\xfb\xfa\xfd\x71\x40\x96\xce\xd2\xb4\xe3\x96\x46\xfe\x07\xac\x60\x9c\x27\x5b\x73\xeb\xf0\x5a\xfb\xc9\x4a\xa3\x65\xb7\xad\x92\xa3\x89\xb6\xc2\xae\x4c\xd3\x72\x31\x5a\x86\x29\xb1\xaf\x96\xb7\xe5\x60\x69\x37\xec\x06\x2b\xe9\x3c\xda\x10\xd8\x38\x33\xf6\x85\x37\xd8\xf7\xf4\x84\x96\xb2\xa7\xe5\x56\x8b\xb8\xff\x7c\x06\x61\x9c\x31\x02\x2f\x0a\xa1\x24\xb4\x16\xdd\xb6\x41\x07\xda\x80\x1b\x10\xa0\x90\x4e\x98\x0e\xed\x13\xb8\x27\xe7\xb1\x99\x03\xd7\x05\xe0\x63\x8b\xc2\x3b\xd8\x3a\xb4\x0e\xbc\x89\x48\xad\x35\x9d\x2c\x10\x7c\x4d\x61\x16\x85\x27\x60\x8b\xce\x81\x29\x81\xeb\x1d\x26\x1b\xf6\x3a\x64\x6b\xbd\x34\x1a\xa4\x03\x8b\xa5\x42\xe1\xb1\x00\xa9\x23\x1c\xc1\x50\x55\x1b\xa9\xb9\x7d\x8a\x69\xc9\x34\x9a\x69\x27\x23\x9f\xdd\x32\x1a\x2f\xbc\xe5\xda\xd1\xc0\x19\xa5\x05\xe2\xb0\x8b\x48\x14\xda\x71\x2b\xcd\xd6\xed\x42\x85\xd1\xce\xdb\xad\xf0\xc6\x3a\xd8\x18\x5f\x8f\x2d\x41\x6d\x9c\x5f\x46\x61\xec\x16\x91\xa6\xc9\xb8\xb4\xc8\x87\x4b\xc2\x1e\x7e\x56\x31\x07\xbb\x8d\x8e\x79\x46\x6f\x63\xea\x6c\x0e\x19\xfd\x7e\xba\xbb\x5b\x1f\x8d\xa0\x28\x5c\x27\xb2\x59\x9a\x44\x42\x9e\x46\xa2\xb7\x13\xd2\xf2\x97\x77\x3f\xbf\xa3\x9b\xea\x66\x7d\x05\x39\x81\xce\x4e\xa0\x36\xe8\x6b\x53\x00\xbc\x8c\x3a\xbc\x25\xa4\x10\x2c\xd7\x15\xc2\x99\xd4\x05\x3e\xce\xe1\x4c\xc2\x72\x05\x13\x6b\x3e\x47\x47\xd7\xf7\x21\xc8\x72\x74\xa2\x07\x54\x0e\xe9\x7a\x67\x7e\x37\x0f\x68\xe1\x4c\x8e\x04\x0b\x01\x75\x31\x5d\xb2\x79\xf2\xef\x32\xcc\x27\x9c\x53\x19\xa8\xcb\xd9\xc1\x42\x42\x60\x57\x71\xa1\x97\xb6\x72\xec\x52\xa9\x8f\xb4\xf4\xbe\x27\xaf\x24\x36\x1f\xc5\x95\x13\xf5\xa7\xa0\x1d\xb5\xdb\x0d\x0b\xe1\x37\x43\x09\xe0\x65\xc5\x24\x09\xda\x61\x45\x51\x70\x11\x55\x96\xf0\x76\xe2\xc1\x9b\x15\x64\x19\x09\x69\x07\x3a\x27\x4f\x58\xc1\xfe\xe4\x60\x5f\xf0\x21\x9f\x22\x66\x69\xd2\x03\xcd\x11\x08\x67\x62\xc1\x1e\x47\x18\xad\x07\x90\xe5\x0a\x22\x0f\x7e\x95\x5c\xe5\x93\xeb\x7c\x30\x7e\x95\xbe\xbe\xd6\x0e\xc5\xd6\x62\x3e\x3b\x30\xde\xc9\x06\xcd\xd6\xe7\x74\x2c\xb3\x5b\x14\x46\x17\x33\xa2\x86\x2c\x23\xe8\x9b\x15\x68\xa9\x62\xa6\xa4\x6c\x3c\xfb\xd8\x5a\xa9\x7d\x99\x1b\xc7\x6e\x7d\x81\xd6\xce\x21\xfb\x40\xbd\xc2\x43\x2d\x15\xe9\x98\x2b\xa9\xab\x88\x4f\xa2\xd1\x28\x48\xb0\x4b\x38\xef\xb2\x58\x26\x61\x27\xc6\xb1\x0f\x8f\xd2\xe7\x3f\xd1\x53\x9f\x26\x49\x81\x25\xda\xe8\xcf\xae\x94\x89\x0b\x78\x36\xa2\xfd\x99\x1c\x47\x44\xce\xfb\xe9\x50\x81\xa7\xea\x8b\xbb\x58\xd2\x21\x65\xb1\x31\x1e\x27\x31\xb8\x16\x85\x2c\x25\x16\xdf\x74\x94\xc3\x61\x59\x7d\xfa\xc2\x08\x5e\xc9\x70\xde\x7d\xd3\xfb\x2e\x8f\xd1\xd2\xc4\x3d\x48\x2f\x6a\x78\x3b\x8a\x2e\xa4\x21\xc0\x83\xf4\x35\x9c\x79\x8c\x64\x27\x1a\xee\x75\x40\xa6\x33\x8f\xcf\x25\x40\x4b\xe7\x0e\x49\x95\xcf\x68\x9f\x2d\x69\xbc\x21\x5c\x0c\xc8\x51\x27\x11\xe5\x80\xf5\x04\x44\x37\xbb\x28\xfa\xd4\x50\x54\xe4\x3b\x35\x3c\x72\x37\x02\xb1\xcf\xdc\xba\x9a\xef\x75\x32\xa0\x47\x71\xd1\x83\xc5\xef\x5b\x74\x7e\x62\xe0\xd1\x37\x97\x85\x30\x55\x96\x87\xf0\xf7\x4b\x22\xfd\xb2\x2b\xae\x14\x19\x27\x31\x47\xea\xbc\xc0\xcb\x57\x88\x29\xb8\x8a\x94\x3c\x59\xda\x8f\x8b\xfb\x81\x9f\x71\x79\x49\xd2\x4d\x3d\xee\x3e\x4b\x87\xdd\x8d\xff\xbb\xb0\xf7\x5c\xdc\x57\xd6\x6c\x75\x41\x2a\x1b\xa7\xf3\x1f\x0b\x7f\x21\xdf\xeb\x25\x8f\xe8\x6b\x02\x57\x3a\xcf\x86\x69\xc3\xcd\x50\x11\x16\x91\x20\xcb\x48\xfb\x23\xc7\xff\x69\x4d\x47\xb9\x87\xd3\x11\x6e\xd0\xb5\x46\x17\xa7\x73\x77\xb3\xf4\x80\x5d\x07\x3c\x2b\xb0\xe4\x5b\xe5\x97\xaf\x4b\x50\xea\x8e\x2b\x59\xc0\x28\xb2\xf3\xef\x71\x4a\xc3\xd3\x73\x51\xf6\x69\xfa\x57\x00\x00\x00\xff\xff\x6c\xf0\x67\x0a\xe4\x0a\x00\x00")

func nameServiceNameCliClientClient_mainGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_nameServiceNameCliClientClient_mainGotemplate,
		"NAME-service/NAME-cli-client/client_main.gotemplate",
	)
}

func nameServiceNameCliClientClient_mainGotemplate() (*asset, error) {
	bytes, err := nameServiceNameCliClientClient_mainGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "NAME-service/NAME-cli-client/client_main.gotemplate", size: 2788, mode: os.FileMode(436), modTime: time.Unix(1478668572, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _nameServiceNameServerServer_mainGotemplate = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x56\xdf\x6f\xdb\x36\x10\x7e\x16\xff\x0a\x56\xe8\x00\x69\x50\xa8\x0e\x5b\xf7\x10\x34\x0f\xad\xdb\xa5\xc1\x92\xc0\x88\x3d\xec\x99\x96\xce\x34\x11\x89\x14\x48\xca\x75\x20\xf8\x7f\x1f\x8e\xa4\x1c\xb9\x75\x9c\x3d\xb4\x40\x2d\x89\xf7\xf1\xe3\x7d\xc7\xfb\x91\x8e\x57\x8f\x5c\x00\x6d\xb9\x54\x84\xc8\xb6\xd3\xc6\xd1\x8c\x24\xe9\xba\xe1\x22\xc5\x67\xeb\xf0\xa1\x60\x7c\x94\x1b\xe7\xba\xe9\x7b\xd9\x75\x46\xaf\x71\x45\xdb\xf0\x5b\x5a\x29\x14\x6f\xf0\xc3\x3e\xd9\x8a\x37\x4d\x4a\x48\x52\x96\xf4\xf7\x9a\xce\xb9\x71\x4f\x24\x49\x85\x6e\xb8\x12\x4c\x1b\x51\xee\x4a\xa4\xaa\xb4\x72\xb0\xf3\xa7\x08\xad\x45\x03\x6c\x02\x11\xa6\xab\x22\xc7\xb5\xa6\x7f\x4b\x87\x28\xe9\x36\xfd\x8a\x55\xba\x2d\x85\xbe\x78\x94\xae\xc4\xff\xa0\xea\x4e\x4b\x15\x78\x4e\x22\x1a\x2d\x22\xd5\x72\x23\x2d\x5d\x80\xd9\xca\x0a\x48\xb2\xe1\xaa\x6e\xc0\xd0\x74\x18\xd8\x8d\x0f\xc4\x9c\xbb\x0d\xbd\xd8\xef\x69\x19\x6d\xb6\xb4\x60\xb6\x60\x52\x92\xd8\x6d\x75\x12\x29\x40\x81\xe1\x0e\xea\x94\x24\xdd\xca\x43\xe6\x9f\x8e\x41\x29\x21\x39\x21\xeb\x5e\x55\x3e\xec\x59\x4e\x07\x92\x6c\xb9\xc1\xb8\x27\x35\xac\x7a\xf1\xb1\xae\x0d\xf5\xff\xae\x28\x5e\x04\x5b\x38\x23\x95\xc8\x52\x6f\x65\xbc\xae\x4d\x5a\xd0\xf4\xf2\xfd\xbb\x3f\xdf\xe1\xcb\x67\x5c\xa6\x5c\xd5\xb4\x05\x67\x64\x65\x69\x23\xad\x03\x45\x11\x09\xd6\xa6\x39\x49\x12\xbc\xab\x67\xe2\xef\x99\xd1\x3a\x25\x7e\xef\x89\xbf\x2e\x97\xf3\x53\x5c\x78\x1f\x2f\x73\xa1\x75\xca\xf5\x87\xe7\x12\x0f\xf3\x19\xcd\x90\x31\x3f\x41\x99\x93\xc4\x73\xcc\xb9\xb1\x90\xe5\xe1\x82\x6e\xb5\x10\x52\x09\x5a\x6b\x0c\x14\x0b\x51\x6a\xb4\x10\xe0\x1f\xec\xd6\xbf\x92\x64\x20\x49\x12\x97\xaf\xbc\xe1\x1e\xbe\xdd\x6a\xb1\x6e\x5d\x40\x64\xda\xb2\x85\xab\x75\xef\xf2\x13\xc8\x59\xc8\xbc\x2c\xac\xe7\xec\x5f\xe9\x36\x59\xea\x6c\x5a\x78\xc4\x67\x58\xf3\xbe\x71\x4b\xd9\x82\x75\xbc\xed\xfe\x59\xce\xfe\x3f\x0b\x26\x3f\x98\x63\xa6\x99\x5f\xcb\x49\xb2\x27\x91\x05\x85\x64\x69\x6b\x05\xc6\x69\x03\x4d\xa3\x31\x24\x35\xac\x61\x94\x7b\x84\x10\x5a\xd7\xab\x27\x48\x63\x94\x3e\xf5\x56\x2a\xb0\xf6\x38\x4c\x36\xe4\x35\xed\x56\x6c\x18\xae\xf5\x3d\x6f\x81\xb2\x98\xec\x0c\xbf\xf6\xfb\x85\x4f\xe6\x10\xbe\x11\x7e\x45\x63\xae\xa3\xa6\x08\xcf\x50\x6f\x59\xe2\x75\x1d\x68\x9b\x78\x35\xd3\xa4\xdb\x80\x01\x54\xe5\xbd\xfa\x12\x0b\xf1\xd9\xab\x61\x30\x5c\x09\xa0\x6f\x25\xbd\xbc\x7a\xf6\xe5\x0e\xdc\x46\xd7\x76\xbf\x0f\x7e\x0f\xc3\x52\xdf\xea\x6f\x60\xe8\x5b\x19\xfd\x3c\x50\x8d\xc5\xcd\xc6\x95\xe0\xfb\xb9\x2d\x57\xd4\x6e\x2b\x76\xc7\x1f\x61\x18\x7e\xb0\x66\x51\x4d\xd4\xf7\xb1\xae\x0f\x47\x50\x67\x78\x25\x95\x28\xa8\x54\xd6\x99\xbe\x05\xe5\xb8\x93\x5a\x79\xc5\xa3\xfa\x51\x71\x32\x0c\xa0\x6a\x94\x30\xee\xb7\xa8\x11\x8f\x1e\xcf\xb2\xc3\x6b\x11\xc0\xde\xe0\xd5\xfc\xe0\xe7\x25\x16\xd9\x19\x95\xc5\xc4\x81\x18\xfe\x3b\xa8\x36\x5c\xc9\x8a\x37\xcf\x17\x00\xc6\x54\x78\x70\xcb\x1f\x21\x43\x33\x05\x63\x34\x66\x62\xe5\x76\x68\x88\x3d\x98\x7d\xe2\xd5\xa3\x30\xba\x57\xf5\x58\x89\x37\xca\x81\x31\x7d\xe7\x0e\xe9\x41\x12\xa1\x29\x36\xb1\xd0\xbf\x92\xef\x98\xb1\xe2\xfc\x14\x28\xe8\x6f\x18\xde\x30\x12\xd8\xbd\x76\x72\xfd\x94\x55\x05\x8d\x93\x81\x2d\x6e\xae\x6f\xee\x97\x47\xdf\xcb\x2f\x0f\x77\xb8\xc7\xfb\xfb\xe1\x82\xae\x5b\xc7\xbe\xa0\xa7\xeb\x2c\xfd\x05\xcb\xf2\xc3\x45\x85\xe5\x33\x3a\x17\xda\x5f\xe8\x29\x27\x3c\x8b\x95\x7a\xf9\x5a\xc1\x1b\xae\x2c\xf6\x69\x2c\x31\xdf\x68\x7d\x81\x25\x2d\xee\xf4\xed\x31\x56\x04\xdc\xf5\x3b\x5f\x12\x2d\xfb\xea\x83\x91\xa5\xa5\xc7\x87\x49\x58\xa6\x45\x80\x07\xa3\xf9\x0b\x3d\xf1\x16\x76\xa3\x6a\xd8\xe5\x67\xb6\x56\x6d\xdd\x48\x05\x2f\x33\xcc\x02\xe0\x1c\x07\xfe\xc8\xe6\x0c\xc7\x3c\x00\xce\x71\xd8\xa7\x76\xa5\x9b\x97\x29\x16\xde\x7e\x8e\x01\xcb\xe7\x8c\x0f\x4b\x34\xe7\x3e\xbe\xd3\x06\x17\x87\xc6\xaf\x87\x29\x38\x4d\x03\x4f\x75\xeb\x6f\xf9\xa3\xaa\xfd\x4d\x64\xcf\xc8\x82\xb6\xd3\x9c\xf0\x93\xeb\x70\xa5\x3f\x25\x27\x90\x32\x4c\xd2\xb1\xb6\xb1\xad\xe0\x6a\xd4\x97\x55\x6e\x57\x1c\x3a\x88\x2d\x62\xef\xce\x5f\x10\x39\x0e\xe4\x57\x35\x8e\xc0\x82\x6e\xa6\x12\xfd\x40\xfd\xb9\x12\x91\x32\x64\x7d\xa3\x0a\xec\x0e\xb8\x5d\x81\x8b\x2e\x65\xa9\xab\x3a\x74\x7d\x9c\xff\xe8\xba\x5c\x7b\xe0\x9b\x2b\xaa\x64\xe3\x4f\x3e\xa8\x01\x63\xf0\xd3\x80\xeb\x8d\x22\x89\xef\x4d\x89\x35\xdb\x69\xfc\xae\x1f\xe6\xb3\x30\x88\xbe\x0b\x9f\xef\x1a\x88\xf4\x7f\x4e\x8c\xb5\x67\x7c\xe5\x75\x2b\xf6\x00\x02\x7d\x32\xc3\x70\x72\xaa\x65\xb6\xa0\xd6\x6c\x5f\xcc\xb0\xa9\x82\xd1\x5d\xcb\x42\xc0\x1b\x35\x0d\xf3\x43\xaf\xde\x1c\x0f\x6a\xd8\x49\xe7\xbb\x10\x6e\xcc\xc9\x9e\xfc\x17\x00\x00\xff\xff\xf4\x3e\x62\x2e\x46\x0b\x00\x00")

func nameServiceNameServerServer_mainGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_nameServiceNameServerServer_mainGotemplate,
		"NAME-service/NAME-server/server_main.gotemplate",
	)
}

func nameServiceNameServerServer_mainGotemplate() (*asset, error) {
	bytes, err := nameServiceNameServerServer_mainGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "NAME-service/NAME-server/server_main.gotemplate", size: 2886, mode: os.FileMode(436), modTime: time.Unix(1478654064, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _nameServiceGeneratedCliHandlersClient_handlerGotemplate = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x52\x4d\x8f\xda\x30\x10\x3d\xdb\xbf\x62\x84\x22\x94\x54\x60\xee\x48\x3d\x40\xd5\x22\x0e\xad\x10\xed\x1f\x30\xc9\x90\x58\x4d\x1c\xaf\xe3\xec\x87\x46\xfe\xef\x2b\x1b\x43\x96\xbd\xb0\x97\x64\xf2\xde\x9b\xe7\xf1\x9b\x18\x59\xfe\x97\x35\x42\xd9\x2a\xd4\xae\x91\xba\x6a\xd1\x72\xae\x3a\xd3\x5b\x07\x39\x67\xe6\x04\x33\xa2\x4c\x1c\xb6\xfb\x88\x1d\xa4\x6b\x60\xe9\xfd\x8c\x17\x9c\x13\xc1\x8b\x72\x0d\x64\x0e\x61\xfd\x1d\x84\xf7\x9c\x11\x59\xa9\x6b\x84\x4c\x05\x28\x73\x28\xfe\xa2\x7d\x56\x25\x8a\xdf\xe8\x9a\xbe\x1a\x82\x68\xb5\x02\xa2\x4c\x89\x3f\xb2\x43\xef\x41\x75\xa6\xc5\x0e\xb5\x1b\xe0\x2a\xe6\xec\x3c\xea\xf2\xa3\x2a\x27\x8a\x87\x29\x5d\xe1\x6b\x34\xfe\x11\x87\xde\xd8\x7a\x88\xde\xa1\x80\x9b\x9c\x68\xd7\x87\x0a\xc4\xaf\x51\x97\x4e\xf5\x3a\xf0\x01\x47\x5d\x79\x5f\x40\xfe\xcd\x9c\xc4\x4d\x95\x29\x71\xc4\xa7\x11\x07\xf7\xef\xcd\x60\xf2\x58\x00\x5a\xdb\xdb\x82\x38\x63\xf6\xc2\x86\x4b\x3d\xec\x0b\x7a\xa2\x65\x0a\xa7\x43\xd7\x84\xb6\x2f\x0d\x1e\xb2\xe5\x8c\x4d\x31\x1a\x69\x65\x17\xa3\x0c\x3e\x22\x6a\x93\x26\x9e\xa1\xce\x20\x75\x05\x79\x7c\xe8\xde\xa5\x0e\xb1\x1f\xb6\x72\xc0\x30\x54\x71\x87\x1f\xd1\xa0\x74\x58\x15\xf7\xf0\x4f\x3d\x76\x45\xb2\x65\xd3\xed\x2e\x64\x5a\xd3\x1a\xe6\x9f\x99\x5d\xbf\xb1\xb5\xf7\x8b\xdb\x38\xd8\x0e\x08\x0f\x7d\x1e\xdb\xe8\x6a\x4a\x62\xfa\x48\x54\xac\x7d\x5c\x8a\x1b\xad\x86\x79\x5a\xce\x02\xb4\x6a\x39\x8b\x7f\xe1\x45\x76\x7d\xbf\x07\x00\x00\xff\xff\x50\xcc\x31\x90\xe8\x02\x00\x00")

func nameServiceGeneratedCliHandlersClient_handlerGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_nameServiceGeneratedCliHandlersClient_handlerGotemplate,
		"NAME-service/generated/cli/handlers/client_handler.gotemplate",
	)
}

func nameServiceGeneratedCliHandlersClient_handlerGotemplate() (*asset, error) {
	bytes, err := nameServiceGeneratedCliHandlersClient_handlerGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "NAME-service/generated/cli/handlers/client_handler.gotemplate", size: 744, mode: os.FileMode(436), modTime: time.Unix(1478668572, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _nameServiceGeneratedClientGrpcClientGotemplate = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x55\xcd\x6e\xe3\x36\x10\x3e\x8b\x4f\x31\x15\x82\x42\x0a\x14\xea\x9e\xc2\x97\xf5\x6e\x17\x5b\x34\xa9\x91\x0d\xda\xc3\x62\x51\xd0\xd4\x58\x26\x6c\x93\x2a\x49\x3b\x36\x04\xbd\x7b\x31\x14\xe9\xc8\xde\x64\x0f\x41\xc4\x99\x6f\xfe\xbe\xf9\x71\x5d\xc3\x42\xc8\x8d\x68\x11\x5a\xdb\x49\xe8\xac\x39\xa8\x06\x1d\x08\x68\x9f\x16\x73\x90\x5b\x85\xda\xc3\xca\x58\xf0\x6b\x84\xbe\xe7\x5f\xd1\x1e\x94\x44\xfe\x28\x76\x38\x0c\xe0\xe2\x93\x75\x13\x37\x8c\xa9\x5d\x67\xac\x87\x82\x65\x79\x6b\xb6\x42\xb7\xdc\xd8\xb6\x3e\xd6\x1a\x7d\x2d\x8d\xf6\x78\xf4\x79\xd0\x99\x76\x8b\x7c\x02\x21\xf3\xf7\x35\xf5\x0e\xbd\x68\x84\x17\x01\xa2\xfc\x7a\xbf\xe4\xd2\xec\xea\x6e\xd3\xd6\x68\xad\xb1\x2e\x67\x97\x9a\xd6\xdc\x6d\x94\xaf\xe9\x0f\x75\xd3\x19\xa5\x29\x30\xf9\xf2\x56\x68\x17\xb2\x7c\x07\x7f\x06\xc4\xa4\x58\x56\xd7\xf0\xbc\x56\x0e\x22\x07\x2c\x73\x07\x09\x79\xdf\xf3\x2f\xa1\xdc\x85\xf0\x6b\xb8\x1b\x06\xa8\x5b\xd4\x68\x85\xc7\x26\x67\x59\xb7\x0c\x90\xc5\x87\x4b\x50\xce\x4a\xc6\xea\x1a\x1e\xf1\x05\x2c\xfa\xbd\xd5\x0e\x84\x4e\x7c\xc2\x52\xc8\x0d\x36\xb0\x3c\x5d\x75\x42\x1a\xad\x51\x7a\x65\x34\x87\x2f\x1e\x94\xa3\xbe\xb0\xd5\x5e\x4b\xf2\x54\x90\x1a\x6e\x29\x5f\x3e\x0f\x06\x73\xa3\x75\x05\xa6\x23\x0b\x07\x9c\x47\xf1\x5f\x41\x50\x42\xd1\x2d\x79\xdf\x7f\x36\xd4\x4d\xb8\xea\x2d\xbd\xd0\x56\x10\x88\x2d\xa1\x67\xd9\x41\x58\x90\x32\xa6\x32\x37\x7a\xa5\x5a\xc6\x32\x1a\x8e\x7f\x2b\x58\xc1\xfd\x0c\xac\xd0\x2d\x9e\xc3\xf5\x2c\xcb\xd0\x5a\x52\xac\x8a\x5f\xa5\x2c\x59\x96\xa9\x15\x39\x84\x5f\x66\xa0\xd5\x36\x20\xb2\xb1\x7c\x7a\xc7\x60\x8e\xff\x63\x45\x57\xa0\xb5\x15\xe4\x52\x68\x6d\x3c\x88\xae\xdb\x9e\xa2\xe7\x9c\x1c\x0d\x2c\x1b\x18\xcb\xe4\xa4\x1e\x47\x91\xbe\x7d\xbf\xe8\xee\x45\xc1\x14\xee\x2d\xed\x07\x5c\x19\x8b\x05\x25\x13\xa7\xf3\x6f\xb1\xdd\xa3\x7b\x36\x9f\x9f\x16\xf3\x87\x38\x74\x85\x94\x7c\x8d\xa2\x41\xeb\xca\xb2\xa2\xf0\x59\xdf\xdf\xc1\x8b\xf2\x6b\xb8\xf1\x48\xc1\xf9\x30\xb0\x6c\x22\xed\x36\x6d\xa0\xf6\x7e\x46\x08\x1e\x77\x6d\xe4\x97\xa2\x11\x72\xe4\xec\x46\x25\x50\xea\xc2\x03\xfa\xb5\x69\xdc\x08\x0c\xdc\xf7\xfd\xb3\xf9\xd3\xbc\xa0\x85\x1b\x15\x9b\xf4\x29\x0e\x35\xa4\xe9\xe6\x49\x12\xac\x02\xbf\x14\xe6\x7d\xc3\x19\x5c\x32\xf2\x88\x2f\x23\x29\xc5\x68\x4b\x8c\xe8\x2a\x7e\xe7\x7d\x9f\x6a\x1a\x06\xde\xf7\xd3\x7c\x47\x61\x3e\x85\xaa\x6b\xa1\x3b\x48\xfe\x49\x4b\xd3\x20\x11\x3b\x41\x3c\xe1\x7f\x7b\x74\x7e\x8a\xfb\x88\x6f\xe2\x5c\x67\xb4\xc3\x04\x9c\xce\xef\x8d\xe2\x49\xfd\x7c\xea\x52\x42\xfd\x90\xb0\x17\xa3\xc2\x39\x8f\xf2\xf2\x4c\x59\x51\x06\x49\xec\x0c\xea\x26\x76\x33\x7e\xa5\x0f\x96\x26\x76\xac\x66\xb4\x75\x3d\x01\xa6\xbd\xbc\x6e\x24\x6d\x7d\x70\xf7\x43\x0f\xee\x01\xe0\x67\xcd\xad\x5e\x63\x67\x43\x45\x8b\xc2\x06\xc6\xfc\xa9\xc3\x8b\x5d\x04\xe7\xed\x5e\x7a\x5a\xaa\x38\xa6\xf0\xed\xbb\xf3\x56\xe9\x96\xf0\x75\x0d\xd3\x5d\xa0\xdb\x21\x80\x2e\x47\x78\xf9\xb5\xf0\xb0\x33\x8d\x5a\x29\x0c\x47\x65\x72\x71\x68\xcf\x43\xb4\x0b\x7b\x32\x2d\x6e\xa7\x09\x94\xe3\xfa\xb2\xf1\x1e\xcd\xfd\x31\x6d\xd1\x57\xd4\x4d\xb1\xc1\x53\xb8\x40\x63\x46\xe5\xa5\xb3\xfe\x4c\x6a\x70\x6b\xe0\x2d\xc7\xe1\x5c\x98\xb4\x83\x30\x03\x72\xc9\xa6\x07\x84\x96\x72\x88\xf1\x7f\xb6\xc9\x21\x97\x44\x4e\x79\xb5\x01\x71\x16\x7f\x27\x27\x57\x79\x49\x7f\x4c\x7e\xf9\x7c\xfc\x5f\xc1\xae\x81\xdb\xf4\xc3\xc4\x1f\x3e\x96\xd7\x88\x90\x36\xed\x6f\x27\xd4\xb4\x27\x59\x3a\x9d\x9b\xd7\xd3\x19\x12\x0b\x5b\xab\x56\x70\xa8\xc0\x04\x9d\xf4\x47\x1e\xea\x28\x36\x25\x2f\x62\xd6\xbf\x91\x72\x5c\xf0\xd1\xf1\x8c\x8e\x24\x31\x1d\x9e\x15\x6c\x2a\x38\x84\x89\x1e\xc2\xb9\x1c\x8f\xef\x08\x9d\x9e\xdf\xdb\x5d\x03\x33\x38\x17\xf0\x87\x51\xba\xb8\xdd\x35\xd5\xab\x68\x41\x36\xa3\x57\xce\x79\x59\x26\x77\x91\x19\xe9\x8f\x23\xef\xff\x07\x00\x00\xff\xff\x3a\x03\xa3\xb4\x4e\x08\x00\x00")

func nameServiceGeneratedClientGrpcClientGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_nameServiceGeneratedClientGrpcClientGotemplate,
		"NAME-service/generated/client/grpc/client.gotemplate",
	)
}

func nameServiceGeneratedClientGrpcClientGotemplate() (*asset, error) {
	bytes, err := nameServiceGeneratedClientGrpcClientGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "NAME-service/generated/client/grpc/client.gotemplate", size: 2126, mode: os.FileMode(436), modTime: time.Unix(1478654064, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _nameServiceGeneratedClientHttpClientGotemplate = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x2c\xcb\xb1\x0d\xc2\x30\x10\x05\xd0\x3e\x53\x5c\x1d\x09\xdf\x10\x34\x29\x91\xc8\x02\x56\xf8\x98\x88\xc3\x67\x9d\x3f\x95\xe5\xdd\x69\x18\xe0\x8d\xa1\xab\xdc\x01\x29\x7e\x61\x7c\x7b\xd7\x82\x5a\xfc\x7d\x52\x5f\x64\x63\xe4\xda\x9b\x07\x95\xf8\x34\xcb\x44\x2a\x2e\x4f\x0f\x39\xfc\x01\x59\x75\xce\x65\x8c\x23\x9b\x49\xda\xf6\xfd\xb6\xc1\x1a\x22\x5d\xed\x44\xe5\xfe\x27\x92\xe6\x5c\x7e\x01\x00\x00\xff\xff\x0b\x3c\x4c\x9e\x69\x00\x00\x00")

func nameServiceGeneratedClientHttpClientGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_nameServiceGeneratedClientHttpClientGotemplate,
		"NAME-service/generated/client/http/client.gotemplate",
	)
}

func nameServiceGeneratedClientHttpClientGotemplate() (*asset, error) {
	bytes, err := nameServiceGeneratedClientHttpClientGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "NAME-service/generated/client/http/client.gotemplate", size: 105, mode: os.FileMode(436), modTime: time.Unix(1478654064, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _nameServiceGeneratedEndpointsGotemplate = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x55\xdd\x6e\xdc\x38\x0f\xbd\x1e\x3f\x05\xbf\x20\x40\x67\x3e\x4c\x34\xf7\x0b\xf4\x66\x8b\xfd\xc9\x45\x8b\x62\xdb\x17\xd0\xc8\xf4\x98\x88\x2c\xa9\x12\x3d\xc9\xac\xe1\x77\x5f\x50\xb2\x1d\x67\x63\x2c\x7a\x11\x38\x43\x53\x87\x87\xe7\x90\x72\xd0\xe6\x49\x5f\x10\xd2\xd5\x54\xd5\xe9\x04\xdf\x5b\x4a\xd0\x90\x45\x30\xde\xb1\x26\x97\xa0\x43\x6e\x7d\x9d\x80\x3d\x74\xfa\x09\x81\x5c\x4d\x57\xaa\x7b\x6d\x01\x5d\x1d\x3c\x39\x4e\xd0\x44\xdf\x41\xc2\x78\x25\x83\xe9\x28\x48\x11\x7f\xf4\x98\x18\xb4\xab\x21\x62\x0a\xde\x25\x04\xbe\x05\xcc\x48\x92\x8a\xc0\xad\x4f\xf8\x8a\x72\x04\x9d\xe0\x19\xad\x95\x27\x3a\xe3\x6b\x8c\x49\x00\x04\xaf\xc6\xe9\x77\xe3\xe3\x74\x30\xa3\x1d\x73\x40\x5b\x0b\xbe\x01\xdf\x47\x48\x7d\x08\x3e\x32\xd6\xc0\x51\xbb\x24\xff\x4b\x39\xd2\x96\xfe\xd6\x4c\xde\x09\x5a\xe3\x63\xa7\x39\x29\x78\x64\xd0\x36\x79\x20\x67\x6c\x5f\x63\x5a\xd8\x40\x47\x75\x6d\xf1\x59\x47\x4c\xaa\xaa\xa8\xcb\x40\xfb\x6a\x77\x77\xf1\x56\xbb\x8b\xf2\xf1\x72\x7a\x39\x39\xe4\x93\x48\x85\x2f\x7c\x57\xc9\x4b\xe2\xb6\x3f\x2b\xe3\xbb\xd3\xc5\x3f\x3c\x11\x9f\xe4\x6f\x06\x95\x94\x70\x86\xbb\x61\x50\x5f\x7f\x7d\xcc\x90\x5f\x35\xb7\xf0\x30\x8e\x77\xd5\x21\x3b\xf0\xdb\xa2\xa9\xf1\xd6\xa2\xe1\x34\x37\xc7\xed\x4a\x2b\xe0\x56\x33\x18\xdf\x05\x51\x42\x3b\xd0\x75\x3d\x1b\x20\x5d\x7d\x48\x02\xd6\xa1\x76\x2c\x7a\x9f\x11\xfa\x84\xb5\x08\xab\xa1\x45\x1b\x30\x42\xe2\xd8\x1b\x3e\xca\xeb\xa9\xd4\x76\x25\x72\xec\x41\x0b\x5c\x22\x77\xb1\x08\x41\x47\xdd\x21\x63\x54\xd5\xe9\x24\xf1\x47\x07\xba\x58\x1a\x8f\x40\xfc\x21\x49\xb1\xa6\xb7\xd9\x9a\xa6\x77\x46\x64\x9f\x28\x3b\x14\x67\x3c\xf8\x80\x51\x33\x82\x97\xb3\x01\xe3\xc3\x5c\x50\x00\xcf\x3a\x51\x52\xf0\xbb\x8f\x80\x2f\xba\x0b\x16\x8f\x70\xf3\x3d\x74\x74\x69\x19\x82\x4e\x32\x16\x2b\xa9\x84\xe0\x52\xa8\xd4\x09\xd1\xd7\xbd\xc1\x2c\x83\x76\xd0\x32\x07\xf5\xa7\x76\xb5\x15\x8e\xcf\xc4\x2d\xa0\x36\xed\x34\xdd\xb0\x9f\xab\x1f\xe0\x99\x22\xd6\xd0\x87\x02\x9a\x02\x1a\x6a\xc8\x40\xd0\xdc\x2a\xd8\x3f\x66\x7e\x94\x04\xff\xac\xcf\xf6\x06\x1a\x3a\x4a\x5c\x36\x03\x6a\x4c\x74\x71\x72\x94\xdc\xd5\x3f\x61\x96\xf2\x5b\xb1\x65\xd9\xa4\x4c\x11\xdf\x9a\x5d\xcc\x10\x88\x59\x49\x75\x58\xab\x6b\x2c\xa1\xe3\xb7\xea\xae\x8c\x7b\x5d\x4a\x7b\x93\xd5\x2d\x70\x58\xff\x97\x8d\xb2\x3e\x45\x2b\x12\x85\x3b\x2c\x63\xf5\xca\x97\x1c\x63\x6c\xb4\x0c\xd4\xb6\x13\x02\xb6\x14\xdb\xbe\x18\x7a\x29\xf6\xba\x89\xa7\xec\xc3\x17\x7c\xfe\x34\xf5\x63\x7c\x77\x26\x97\x75\xea\xb2\xb2\x99\xe5\xca\xdb\xe3\x74\x83\x70\x1f\x1d\x50\x1e\x66\xe1\x68\xb4\xb5\x18\xcb\x3c\x4f\x7c\x55\x95\x3b\x7a\xa7\xe9\x50\x0d\x43\xd4\xee\x82\x70\x4f\xf0\xcb\x47\x50\x73\xfe\xe7\xe2\xc7\x38\x56\xbb\x61\xb8\x27\xf5\x45\x77\x38\x8e\xf3\x79\x00\x58\xfa\x50\x73\xb0\x1a\x86\x07\x89\x8e\x63\x35\xbe\x5d\xd7\x9f\x28\x22\x03\x0a\xfb\x15\xc3\x03\xac\xea\xee\x0d\xbf\xc0\x74\x95\xa8\x4f\xe5\x79\x94\x81\xf8\x7f\x38\xab\x61\xf8\xc3\x4b\x1a\xdc\x93\xfa\xab\xdc\xac\xdf\x6f\x01\xa7\xa3\x07\xd8\xbf\x4f\x2a\x57\xee\x2a\xeb\x08\x18\xa3\x8f\x07\x18\xaa\xdd\x6e\xbe\x92\x73\x50\x08\xa3\xda\xd0\x40\x38\x09\x87\x43\xb5\xdb\x51\x93\x53\xff\xf7\x11\x1c\xd9\x8c\xb1\x9b\x5c\x71\x64\x33\x4c\xb5\xdb\x8d\xd5\x12\x9d\x2b\xa8\x9f\xe1\x76\x38\x0a\x4a\xb5\x1b\xab\x61\x28\xf2\x8a\xb8\x9f\x65\xab\xd6\x0a\xe7\xbd\xbd\x67\xcc\x0a\x17\xdf\xd6\xa2\xdf\x33\x6e\xe9\x5e\x84\x17\xb0\xad\x16\x13\x64\x7a\xeb\xb3\x25\xe3\x5b\x5e\xc3\xc3\xfb\x21\x78\xd3\xbc\x60\x6f\x5b\x37\x7f\x01\x97\x35\x1a\xc4\xa8\xe5\x5b\xb8\x0a\x17\x13\x56\xee\x08\xfa\x0f\xe9\x68\xc2\xd8\xd2\xf0\xdd\x10\xe4\x73\xd7\xc5\xd0\xa4\xfe\x35\x5c\x99\x51\xc9\xda\xf0\x72\xcb\xcd\xe2\xe7\xf2\xe6\x3a\x99\x54\xc2\x59\xfd\xe2\xd5\xfc\xfc\x27\x00\x00\xff\xff\x5a\xa4\x38\xcd\x4e\x08\x00\x00")

func nameServiceGeneratedEndpointsGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_nameServiceGeneratedEndpointsGotemplate,
		"NAME-service/generated/endpoints.gotemplate",
	)
}

func nameServiceGeneratedEndpointsGotemplate() (*asset, error) {
	bytes, err := nameServiceGeneratedEndpointsGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "NAME-service/generated/endpoints.gotemplate", size: 2126, mode: os.FileMode(436), modTime: time.Unix(1478654064, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _nameServiceGeneratedTransport_grpcGotemplate = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x57\x4f\x6f\xe3\xb6\x13\x3d\x8b\x9f\x62\x7e\x46\xf0\x83\x1d\x38\x52\xcf\x01\x72\xd9\xec\x76\xb3\x68\xb3\x0d\xd2\xa0\x3d\x2c\x16\x0b\x5a\x1a\x4b\x84\x29\x52\x21\x69\x27\xae\xa0\xef\x5e\x0c\x29\xc9\x4a\xec\xc8\xde\xe6\x10\x38\x16\xe7\xcf\x9b\xf7\x66\x46\x74\xc5\xd3\x15\xcf\x11\xec\x26\x65\x2c\x49\xe0\xa1\x10\x16\x96\x42\x22\x54\x46\x6f\x44\x86\x16\x2c\x9a\x0d\x9a\x0b\x2b\x32\x84\x85\x50\x99\x50\xb9\x85\xa5\x36\xe0\x0a\x84\xfc\xfe\xee\x1a\x9c\xe1\xca\x56\xda\xb8\x98\x42\x7c\x71\xb0\x76\x42\x8a\x7f\xd0\x7a\x93\xfe\x34\xc9\x4d\x95\xc6\x7f\xfa\x70\x31\x63\xa2\xa4\x87\x30\x65\xd1\x44\xa1\x4b\x0a\xe7\xaa\x09\x63\xd1\x24\xd7\x92\xab\x3c\xd6\x26\x4f\x9e\x13\x3a\x49\xb5\x72\xf8\xec\x26\xfe\x4c\xe7\x12\xe3\x81\x09\xc5\x4c\x4a\x74\x3c\xe3\x8e\x93\x3f\x3d\xe8\x53\xc2\x24\x17\xae\x58\x2f\xe2\x54\x97\x49\xae\x2f\x56\xc2\x25\xf4\xf7\x12\x13\xb9\x75\xb5\x13\x3c\x91\x22\x8b\xaa\x05\x4c\xea\x3a\xbe\xfb\xf0\xc5\xe3\xbc\xe3\xae\x80\x8b\xa6\x99\xb0\x19\xf3\x4c\xdd\xf2\x15\x7e\xbe\xbf\xbb\x0e\xf5\x40\xc9\x57\x68\x81\x83\x45\x07\x7a\x09\xa8\xb2\x4a\x0b\xe5\x2c\xf0\x0d\x17\x92\x2f\x24\x02\xa7\x73\x4f\x58\x5d\x7f\xd6\x5f\x79\x89\x10\xb7\xe9\x62\xfa\xd6\x34\x1d\x37\xcb\xb5\x4a\x5f\x25\x98\xa6\xee\x19\x5a\x26\xe2\xeb\xf0\x39\x1f\xa4\xf9\xd4\xfd\x37\x83\x6a\x11\x8f\x27\x80\x9a\x45\x41\xd5\x3f\x2a\x27\xb4\xb2\x70\x79\x05\xdf\xbe\xbf\x60\xae\xd5\x29\x18\xd4\x2c\x8a\xe0\xd0\xf1\x07\x5c\x6a\x83\xd3\x8e\xff\x07\xdd\x22\x9b\xcd\x59\xd4\xb0\xc8\xa0\x5b\x1b\x05\xff\x27\xd7\xe0\x50\x7b\xa6\xeb\x1a\x1e\xf4\xef\xfa\x09\xcd\x4b\x80\xd0\x34\x2c\xaa\x6b\xc3\x55\x8e\x70\x26\x08\x56\x7f\x7e\x8b\xae\xd0\x99\x25\x8b\xa8\xae\x3b\xf7\x33\xd1\x56\x76\xf9\x0a\xdf\x57\x7c\x6a\x89\x63\x51\x14\xa5\xee\x79\x4e\x9f\x3d\x5f\x71\x5d\xf7\xae\x1d\x75\xde\xe2\x23\xa6\x3a\xf3\xbc\x0f\x2c\xee\xf1\x71\x8d\x36\x18\x7c\x52\x07\x0d\x6c\xa5\x95\x45\x6f\xf1\x82\xda\x38\x8e\xe9\x21\x11\x52\xd7\x17\x24\x18\x55\xd0\xb0\xc6\x37\xd1\x8e\x18\x10\x65\x25\xb1\x44\xd2\x92\xa6\xe6\x88\x82\x42\x39\x34\x4b\x9e\x22\x73\xdb\x0a\x87\x71\xac\x33\xeb\xd4\x41\xcd\x8e\xf3\x78\x80\x46\x80\x57\x3c\xde\x70\x95\x49\x34\x6c\x07\x3e\x20\x6f\xc3\xf8\x45\x30\xc8\xee\xf4\xae\x90\xd3\x6b\x38\x0a\xd5\x0f\xc4\xd4\xc2\xf9\x2e\xd5\x6c\x17\xbe\x47\x7f\x78\x48\x0c\x3e\xc2\xf9\x70\x28\xce\x44\xdc\x2a\xfa\xb0\xad\x3a\x50\x33\x98\xee\x1b\x05\x55\x07\x56\x73\x40\x63\x34\x25\x67\xd1\x0f\x0a\x5d\xf9\x27\x04\x9b\x7a\x6a\x8f\xcf\x30\x27\xd4\x2d\x84\xcd\x63\x99\xb1\x48\x2c\xbd\xd3\xff\xae\x40\x09\x49\xa1\xba\x49\x51\x42\xfa\x78\xc3\xe9\x31\x58\xc5\xa7\x40\x9b\xcd\xc9\x9d\x35\xac\xae\x83\x50\x24\x53\x4b\x75\xe8\xea\xe3\x3c\x27\x09\x8c\x0d\x00\x08\x5a\x61\xaf\x16\x7a\x70\x68\x2d\x7e\x25\xa1\x5c\xc1\x1d\xc9\xb0\x41\x43\x0b\xd0\x37\x7a\x58\x7b\xfb\xfd\x66\xda\xc8\x4e\x03\x87\xb5\x45\x73\x91\xe9\x92\x0b\x35\x66\x1c\xc3\x9d\x11\x25\x37\x42\x6e\xc9\x65\xb9\x96\x20\x94\xdf\xbd\x83\xf5\x39\x56\xc7\xf4\xc7\x7e\x97\x50\x2d\xf7\xf8\xb8\xeb\xca\x9a\x5a\x62\xf0\x6d\x28\x3d\xb5\xd4\xe5\x55\xe7\x73\x48\x9e\xbd\xf6\x1a\xe8\xf9\xb8\xa7\x14\x51\x74\x2d\x05\x0d\xcd\xfb\xa5\x0a\x9d\x31\xaa\x55\x30\xf9\x0f\x62\x55\x72\x7b\xaa\x54\x21\xc7\x5b\x5a\xa5\xbe\xda\x71\xad\x42\x84\xb7\xc5\x22\x30\x27\xca\x45\xa6\xbd\x60\x95\xdc\x9e\x36\x51\xc3\x19\x94\xdb\x91\xf9\x0a\x2f\x85\x93\x44\x1b\x7d\x7f\x1c\x14\x2d\x78\x1c\x13\xed\x54\x41\x82\x7c\xe3\x12\x9f\x34\x60\xa3\x85\x1c\x12\xad\x47\x70\xa2\x66\xb6\x22\x16\xfb\x46\xfa\x59\xc5\x6c\x35\x36\x66\xef\x57\xec\xed\x8d\xd8\x09\x36\xba\x11\x4f\xdc\x75\x47\xe5\x1a\xdd\x88\x2f\xa6\x6c\xac\x8e\xc3\x7a\xb5\x25\xfe\xc4\x46\xec\xf0\xbc\x7b\x23\x26\x09\xdc\xa0\xac\xd0\x58\x16\xd0\xef\xdd\x31\x0f\xbf\xec\xcb\x0c\xce\x3b\xd3\xf8\xf6\xe3\xec\xb5\x05\x61\xa5\x3b\xcb\x6a\x0e\x1b\x0f\xd8\xcb\x7f\x5e\x66\xfe\x35\x2c\x96\xb0\x19\xbe\x96\xc3\xcf\x02\x84\x15\x6e\xbd\xd2\x59\x86\x19\x2c\xb4\x2b\x88\xde\x2e\x0d\xdd\x81\x4a\xee\x60\xba\x9a\xc1\x53\x21\xd2\xc2\x9b\x4a\x09\x92\xc4\x6a\xa3\x70\x95\xf9\x7b\x1d\xfd\xcc\x89\xaf\xb9\xd2\x4a\xa4\x5c\xde\x20\xcf\xd0\xfc\x86\x5b\xfa\xcd\xe0\xda\x44\x56\x87\x7e\x11\x0e\x52\xae\x60\x81\x5d\x88\x34\x45\x6b\x31\xa3\xdc\x28\x5c\x81\xa6\xcd\xdc\xde\x70\xe1\xaa\x2f\xf6\x6f\xe1\x8a\xbf\xb8\x5c\x63\xb8\x75\x50\xb1\xdf\x7e\xf9\x3e\x3b\x6a\xf8\x06\xba\xe9\x6a\xb6\x8b\xe0\xaf\xaf\xbd\x76\xa9\x7b\x66\x0d\xfb\x37\x00\x00\xff\xff\x09\x00\x51\xc2\x47\x0e\x00\x00")

func nameServiceGeneratedTransport_grpcGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_nameServiceGeneratedTransport_grpcGotemplate,
		"NAME-service/generated/transport_grpc.gotemplate",
	)
}

func nameServiceGeneratedTransport_grpcGotemplate() (*asset, error) {
	bytes, err := nameServiceGeneratedTransport_grpcGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "NAME-service/generated/transport_grpc.gotemplate", size: 3655, mode: os.FileMode(436), modTime: time.Unix(1478654064, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _nameServiceGeneratedTransport_httpGotemplate = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x2c\xcb\xb1\x0d\xc2\x30\x10\x05\xd0\x3e\x53\x5c\x1d\x09\xdf\x1a\x29\x91\xe2\x05\xac\xf0\x31\x08\x93\xb3\xce\x1f\x1a\xeb\x76\xa7\x61\x80\x37\xa7\xae\xb2\x03\x52\xed\x42\xff\x8c\xa1\x15\x67\xb5\xd7\x93\xfa\x20\x3b\xbd\x9c\xa3\x9b\x53\x89\x77\x6f\x85\x18\xa9\x9a\xdc\xcd\xe5\xb0\x1b\x64\xd5\x88\x65\xce\xa3\xb4\x26\x69\xcb\xf9\xba\xa1\x75\x78\xda\xe1\x5f\x78\xfe\x1b\x49\x11\xcb\x2f\x00\x00\xff\xff\xdd\x3a\x4a\x8f\x6a\x00\x00\x00")

func nameServiceGeneratedTransport_httpGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_nameServiceGeneratedTransport_httpGotemplate,
		"NAME-service/generated/transport_http.gotemplate",
	)
}

func nameServiceGeneratedTransport_httpGotemplate() (*asset, error) {
	bytes, err := nameServiceGeneratedTransport_httpGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "NAME-service/generated/transport_http.gotemplate", size: 106, mode: os.FileMode(436), modTime: time.Unix(1478821739, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _nameServiceHandlersServerServer_handlerGotemplate = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x04\xc0\xc1\x0d\x80\x20\x0c\x05\xd0\xbb\x53\xf4\x4c\xa2\x1d\xc6\x09\x88\x7c\xab\x11\xa9\x69\xeb\x89\xb0\x3b\xaf\x77\x4e\xb4\x03\x24\xba\x86\xfd\xee\x2c\x68\xa2\xcf\x1d\x7c\xe5\x56\x2a\x8c\x03\xef\x57\x73\xc0\x37\x51\x3a\xd5\xe8\xd0\x02\x4a\x3c\xc6\x32\x03\x00\x00\xff\xff\xd6\x21\xab\x2e\x3e\x00\x00\x00")

func nameServiceHandlersServerServer_handlerGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_nameServiceHandlersServerServer_handlerGotemplate,
		"NAME-service/handlers/server/server_handler.gotemplate",
	)
}

func nameServiceHandlersServerServer_handlerGotemplate() (*asset, error) {
	bytes, err := nameServiceHandlersServerServer_handlerGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "NAME-service/handlers/server/server_handler.gotemplate", size: 62, mode: os.FileMode(436), modTime: time.Unix(1478821735, 0)}
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
	"NAME-service/NAME-cli-client/client_main.gotemplate": nameServiceNameCliClientClient_mainGotemplate,
	"NAME-service/NAME-server/server_main.gotemplate": nameServiceNameServerServer_mainGotemplate,
	"NAME-service/generated/cli/handlers/client_handler.gotemplate": nameServiceGeneratedCliHandlersClient_handlerGotemplate,
	"NAME-service/generated/client/grpc/client.gotemplate": nameServiceGeneratedClientGrpcClientGotemplate,
	"NAME-service/generated/client/http/client.gotemplate": nameServiceGeneratedClientHttpClientGotemplate,
	"NAME-service/generated/endpoints.gotemplate": nameServiceGeneratedEndpointsGotemplate,
	"NAME-service/generated/transport_grpc.gotemplate": nameServiceGeneratedTransport_grpcGotemplate,
	"NAME-service/generated/transport_http.gotemplate": nameServiceGeneratedTransport_httpGotemplate,
	"NAME-service/handlers/server/server_handler.gotemplate": nameServiceHandlersServerServer_handlerGotemplate,
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
	"NAME-service": &bintree{nil, map[string]*bintree{
		"NAME-cli-client": &bintree{nil, map[string]*bintree{
			"client_main.gotemplate": &bintree{nameServiceNameCliClientClient_mainGotemplate, map[string]*bintree{}},
		}},
		"NAME-server": &bintree{nil, map[string]*bintree{
			"server_main.gotemplate": &bintree{nameServiceNameServerServer_mainGotemplate, map[string]*bintree{}},
		}},
		"generated": &bintree{nil, map[string]*bintree{
			"cli": &bintree{nil, map[string]*bintree{
				"handlers": &bintree{nil, map[string]*bintree{
					"client_handler.gotemplate": &bintree{nameServiceGeneratedCliHandlersClient_handlerGotemplate, map[string]*bintree{}},
				}},
			}},
			"client": &bintree{nil, map[string]*bintree{
				"grpc": &bintree{nil, map[string]*bintree{
					"client.gotemplate": &bintree{nameServiceGeneratedClientGrpcClientGotemplate, map[string]*bintree{}},
				}},
				"http": &bintree{nil, map[string]*bintree{
					"client.gotemplate": &bintree{nameServiceGeneratedClientHttpClientGotemplate, map[string]*bintree{}},
				}},
			}},
			"endpoints.gotemplate": &bintree{nameServiceGeneratedEndpointsGotemplate, map[string]*bintree{}},
			"transport_grpc.gotemplate": &bintree{nameServiceGeneratedTransport_grpcGotemplate, map[string]*bintree{}},
			"transport_http.gotemplate": &bintree{nameServiceGeneratedTransport_httpGotemplate, map[string]*bintree{}},
		}},
		"handlers": &bintree{nil, map[string]*bintree{
			"server": &bintree{nil, map[string]*bintree{
				"server_handler.gotemplate": &bintree{nameServiceHandlersServerServer_handlerGotemplate, map[string]*bintree{}},
			}},
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

