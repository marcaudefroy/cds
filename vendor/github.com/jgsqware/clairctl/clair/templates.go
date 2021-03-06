// Code generated by go-bindata.
// sources:
// templates/analysis-template.html
// DO NOT EDIT!

package clair

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

var _templatesAnalysisTemplateHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x5c\x7d\x6f\xe3\xb6\x19\xff\x3f\x9f\xe2\xa9\x73\x45\x5e\x16\xc9\xb1\x93\x4b\x53\xc5\xf6\xd0\xe5\x2e\xe8\x80\xb4\x1d\x7a\x5d\x81\xa1\x38\x14\xb4\x44\x4b\x44\x68\x51\xa3\xe8\x24\x9e\xe1\xef\x3e\x90\x7a\xb1\x24\x52\x12\xed\xcb\x61\xc0\x9c\x4b\x62\x4b\xe4\xef\x79\xe3\xf3\x46\x31\x37\xf9\xe6\xc3\x2f\xf7\xbf\xfd\xeb\x1f\x1f\x21\x12\x4b\x3a\x3b\x9a\xc8\x5f\x40\x51\x1c\x4e\x07\x38\x1e\xcc\x8e\x8e\x26\x11\x46\xc1\xec\x08\x00\x60\x22\x88\xa0\x78\x76\x4f\x11\xe1\x70\xcf\x62\xc1\x19\x05\x8e\x13\xc6\x05\x78\xb0\xd9\xb8\x7f\x5f\xa2\x10\xff\x8c\x96\x78\xbb\x9d\x0c\xb3\xc1\x47\xd9\x4c\x4a\xe2\x27\x88\x38\x5e\x4c\x4f\x22\x21\x92\xd4\x1b\x0e\x17\x2c\x16\xa9\x1b\x32\x16\x52\x8c\x12\x92\xba\x3e\x5b\x0e\xfd\x34\xfd\xeb\x02\x2d\x09\x5d\x4f\x7f\x49\x70\xfc\x97\x4f\x28\x4e\xbd\xeb\xcb\xcb\x8b\x9b\xec\x9b\x08\x44\x89\x7f\x71\x5d\xbe\xbb\xaa\xbe\x3b\x01\x8e\xe9\xf4\x24\x15\x6b\x8a\xd3\x08\x63\x71\x02\x62\x9d\xe0\xe9\x89\xc0\xaf\x42\x62\x9f\xcc\x2a\xec\xc8\xb1\x83\xdd\xd8\x41\xc6\xdf\xa0\xe0\x6f\x89\x5e\xfd\x20\x76\xe7\x8c\x89\x54\x70\x94\xc8\x0f\x92\x45\xc9\xb7\x83\x5e\x70\xca\x96\x78\x78\xed\xde\xb8\x23\x89\x5c\xbb\xec\x2e\x49\xec\xfa\x69\x3a\x28\xc4\x57\x54\x32\xda\xf2\x35\x67\xc1\x1a\x36\xe5\x47\xf9\x52\xd3\x33\xc9\x3d\x38\x91\xb2\x83\x94\xfd\xe4\x02\x52\x14\xa7\x4e\x8a\x39\x59\xdc\xd5\x66\x2c\x11\x0f\x49\xec\xc1\x65\xfd\x72\x82\x82\x80\xc4\xa1\x76\x7d\x8e\xfc\xa7\x90\xb3\x55\x1c\x78\x10\x46\x2c\x15\x2f\x11\x11\xd8\x38\xd7\x99\x33\x21\xd8\xd2\x83\x31\x5e\xee\x06\x6c\xcb\x77\xc3\x73\xf8\x6d\x9d\xb0\x90\xa3\x24\x5a\xc3\xf9\xb0\xbc\x51\xbe\x71\x29\x46\x81\x49\xc2\x94\xfc\x07\x7b\x30\x72\xaf\x5b\x91\x43\xca\xe6\x48\x2e\xc1\x35\x5b\x09\x33\xb8\xcf\x62\x81\x48\x8c\x79\x83\xc2\x4e\xf4\xaa\xf0\x5b\x13\x02\xc5\x88\x2f\xc8\xab\x87\x16\x42\x83\x91\xf0\x38\x16\x1e\x0c\x06\x75\xf5\x04\x24\x4d\x28\x5a\x7b\x30\xa7\xcc\x7f\xaa\xdf\x53\x88\x1e\xcc\x99\x88\xba\x49\x73\xf6\xd2\xa0\x57\x1a\x12\x9c\xf1\x65\xf2\xda\x3b\xdd\xc8\xf4\x5b\xf0\xf6\x87\x4f\x51\x9a\x9e\x4f\x07\x3e\xa3\xce\xe0\x73\xbb\x76\xeb\x6c\xca\xd7\x82\x32\x24\x3c\xa0\x78\x21\x1a\xcb\x8e\xbd\x4a\xa3\xab\x89\x73\xc6\x03\xcc\x9d\x39\xeb\x91\x51\x92\xbf\x69\x50\x7f\x21\x81\x88\x3c\x78\x7f\xf9\x6d\xf7\xdc\x04\xc5\x98\x36\xe6\x0e\xcf\x77\xcc\x8f\xf0\xf2\xae\xba\xaa\x32\x1e\x15\x5f\x1c\x05\x64\x95\x7a\x70\xdd\x14\xae\xea\x3a\x06\xaf\x51\x22\x46\x28\x60\x2f\x1e\x5c\x26\xaf\x30\x4a\x5e\x61\x9c\xbc\xc2\x31\x1e\xcb\x2f\x1b\x7e\xa3\xb1\x71\x4d\x38\x82\x25\x6d\x0e\x5e\x3a\xa9\x5b\xf3\xd2\x8a\x3c\xc5\x80\x94\x51\x12\x28\xae\x42\x44\xe2\x74\xce\x38\xb3\xe0\xc9\xa3\x28\x15\x8e\x1f\x11\xda\x74\xe4\x9c\xb7\x02\xde\xe8\x6a\xc3\x73\xf8\x11\xa3\x00\x73\xb3\x07\xa3\x24\x71\xa2\xec\xfe\xa6\x55\xd5\xc7\xe3\xd1\xf7\x37\x0f\x57\x8d\xb5\xcc\x28\xe3\x46\x33\xec\xdc\x48\x1a\x21\xff\x6e\x89\x8d\xa3\x1b\x69\x23\xed\x7e\xcd\x92\xe0\x48\x23\xaa\x91\xc7\xe3\x9b\xab\xf1\xd5\x6d\x03\x8c\xa5\x44\x10\x16\x7b\x32\x91\x20\x41\x9e\x71\xb7\x56\x2b\x32\x47\xa3\xb6\x10\x70\xa7\x87\xcc\x17\x4c\xc2\x48\xfa\x96\xfc\x85\x79\x7d\x84\x4c\x6c\x8e\xe0\x28\x4e\x17\x8c\x2f\x3d\x58\x25\x09\xe6\x3e\x4a\x2d\x78\x21\x32\x7f\xb7\xf9\xb8\xae\x1b\x45\x09\x51\x12\xc6\x1e\xf8\x38\xd6\x38\xc9\x0d\x63\x54\x55\xd5\xaa\x3c\x9c\xa3\xd3\xf1\xfb\xf7\x17\xb0\xfb\xe1\xde\x9e\x59\xae\xe0\xe3\x8f\xf7\x1f\x1f\x1e\x46\xb6\xd2\x69\x7e\x55\x49\x42\x9a\xdb\x74\x6b\xdb\x44\x28\x5d\x2d\x97\x88\x37\x53\x39\x25\x31\x76\xa2\x1c\xc7\xbd\x69\xcd\x74\x79\xe5\x64\x74\x90\xfc\x5e\x7b\x50\xe8\xcf\x14\x2a\x3d\xb7\xa6\x1a\xb4\x12\xac\xe9\x3f\xaf\x4e\x1e\x65\xbf\xbf\xd1\xcc\x5f\xa5\x7d\xd5\x42\x7b\x78\x0e\x9f\x64\xa1\x03\x6c\x01\x22\xc2\x90\x71\x60\x94\x2f\xbb\xe5\xc6\x2c\xc0\xcd\x25\xd8\xe1\x56\x50\x4d\x74\x24\x56\x8a\x36\xe4\xbb\x42\xf7\x63\x2d\x92\xe7\xf2\xe9\x37\x0a\xc5\x8c\xed\xb4\x9a\x31\xee\x06\xac\x69\xa2\x3e\xee\x0f\x60\xad\x91\x9d\xf4\x01\xed\xb9\xb7\x5a\xf2\x71\xb4\xbe\x6b\x66\xc5\x7a\xbc\x2b\xf2\x96\xf2\xd0\xcb\x0b\xc8\xfe\xb9\xe3\x8a\x6f\x66\xbc\x54\x1d\xd2\xe8\xce\xe3\xb3\x5a\x8e\xed\xd1\xa3\xfb\x01\x2f\x7c\x16\x8f\x4c\xfa\xac\x8a\x30\xa7\xa8\x6a\xe9\x3e\xd4\x7b\x4e\x04\xf1\x11\xed\x83\x3d\xc6\xb7\x23\x3c\xea\x09\x98\x55\xe0\x1f\x49\x18\xf5\x82\x7e\xfc\x7e\xf4\xf1\xe6\xca\x1e\xf4\x27\x1c\x90\xd5\xb2\x17\xf6\xe1\xe1\x87\xef\xc6\x37\xf6\xb0\x8f\xec\xa5\x17\xf3\xf6\x6f\xf7\x57\xd7\x3f\xd8\x63\xfe\x8c\x43\x4a\x42\x32\xa7\x46\x0f\xa8\x41\x5f\x7d\x77\xfd\xdd\xf5\x83\x3d\xf4\x3f\xe3\xa7\x98\xbd\xc4\x6f\x8d\x0b\x6e\xc2\x92\x55\xd2\x56\x32\xc7\x2c\xc6\x46\x67\xbc\xba\xd4\x4b\x88\xd2\xc1\xd1\x3c\x65\x74\xa5\x17\x83\x59\xca\x1a\x5d\x56\x2b\x55\xd0\xcb\x26\x3d\xc3\xe6\x03\xa4\x1f\x7b\xe0\x8c\xde\x6b\x03\xb2\x3b\xe3\x37\x2a\x4f\x3b\xdc\x1c\x6a\xb5\x92\xa1\x4c\xea\xae\x97\x55\x64\x69\x06\x8a\xa2\x1a\x6e\x16\xdf\xaa\xae\xd8\x05\x21\xf9\xa5\xea\x99\x04\x71\x1c\x8b\x7d\x4d\xdc\xdd\xd0\x9d\x9c\xec\x67\x4c\x95\xed\x74\x4b\x66\x76\x78\xdf\x62\xe0\xc2\x7e\x9a\xda\xba\x5a\xb4\x7c\xc1\x5d\x9a\xb3\x45\xb3\x95\xcf\xf4\x9f\xcf\x19\x29\xa3\xaa\x7a\xb7\xdd\x56\xd5\xb2\xb9\xaa\xdf\xb6\xf7\x46\x10\xb5\x8b\x91\xdb\x74\x6f\xc3\xcc\xf1\x82\xf1\x66\xb2\xff\x3f\xb7\x8c\x5c\xf7\xea\x47\xbf\x65\x34\x6f\xfc\x9f\x58\x09\x66\x10\x90\xe7\x86\x91\xd8\x33\xe6\x0b\x2a\xdd\x33\x22\x41\x80\x63\x43\x5f\xb0\x1b\x82\x29\x25\x49\x4a\x52\x6b\xca\x5e\x24\x27\x9b\x22\x3e\x4b\x90\x4f\xc4\xda\x03\xf7\x76\x6f\xb4\xce\x78\x6f\xb0\xb2\x2c\x7e\x0b\x9b\x8e\x6e\x35\x5b\xe5\x46\x0a\xc8\xb2\x5e\x47\xd5\x6a\xdf\x39\xe2\xa9\xb9\xda\x9d\x23\xee\xcc\xc3\x36\x76\xda\x0b\xd9\xb2\x26\xd4\x18\x2a\x78\xbd\x69\xcf\x4f\x76\x5d\xa9\x62\x0d\xed\xb3\x9b\x64\xe5\x9a\x9a\x5f\x68\x7b\x05\x50\xba\x6b\xc7\x46\xa5\x56\x46\x81\xa1\x20\xee\x2b\xd9\x73\x11\xcb\x32\xf3\x8b\x2a\xcc\x02\xac\xac\x2e\xbf\xb0\xb0\x2c\xf0\x54\x51\xf9\x85\xf5\x64\x81\x95\xd7\x92\x5f\x58\x46\x16\x68\x8f\xda\x86\xe5\xde\xd5\x63\x01\x55\xa9\x1c\xbf\xb0\xb8\x2b\x10\x8b\x82\xf1\x20\xb8\xe1\x39\x3c\xaf\x68\x8c\x39\x9a\x13\x4a\x04\xc1\x2d\xee\xdb\xd1\x8c\xcb\x68\x91\xbc\xee\xd5\x4d\x1b\x77\x75\x33\x0a\x6e\x83\x9d\x0b\x7d\xc4\x02\x23\xb1\xe2\x38\x85\x19\xac\x9a\x8b\x6f\x8f\x27\x02\x5d\x4c\x74\x91\xa0\x24\x15\x45\x86\xa9\x97\xcd\x26\xc4\x1c\xa9\x69\x9d\x2f\xda\xdc\xc9\x31\xdb\xb7\x27\xeb\x79\xf5\xf8\xfe\xe1\xc3\xed\x87\x7b\x2b\xcc\x3f\xff\x54\xcf\xac\xda\xf6\xc4\x46\xe6\x9d\x9c\x1d\x4c\x73\x31\x19\x71\xf2\x2a\x64\xec\xde\xe0\x25\x7c\x43\x96\x52\xe7\xa8\xaf\xd0\xad\x22\x37\xb7\x9b\xb4\xfd\xe0\xdb\xe6\xc6\x56\x31\x82\xe7\x7b\x0f\x6e\xcb\x73\x9d\x76\x49\x76\x2d\x7a\x8c\x96\x7a\x11\xa7\x34\x6d\x11\x3c\x35\xd8\x5d\x8f\xde\x8e\x6b\x15\x46\x35\xe4\xac\x49\xef\x40\xb5\x09\xa8\x1a\x6a\xd1\xa5\x77\xe0\xda\x84\x56\x0d\x57\xb5\xe9\x1d\xa0\x36\x41\x56\x03\xad\xf6\xe9\x1d\xd8\x36\xe1\x56\xc3\x2e\x1b\xf5\x83\x80\x87\xe7\x40\xd1\x1a\xb7\x95\x4b\xea\x5e\xfe\xcb\xec\x93\xfe\x8a\xa7\x92\x46\xc2\x88\xbe\x13\x6d\x76\x58\xe8\x8c\x3c\xcd\x28\xd1\x25\x07\x98\x43\xad\x49\x6f\x06\x49\xf2\x0a\xb5\x2b\xd5\x5b\x44\x41\x85\xe8\xfa\x94\xa5\x38\xa8\x44\x6c\xab\x0d\x8e\x8e\x9d\x6c\x47\x56\xf3\x6d\x28\x0b\x8a\xb5\x1d\x8b\x5d\x92\xd3\x6b\xd4\xee\xcd\xe6\xda\x66\x88\x66\xa9\xda\x76\x73\x5f\xb4\xaa\x31\x6f\xda\x53\xee\x7b\x76\x21\x25\xf3\xa0\x47\xe5\x3a\x95\x3d\x7b\xda\x3d\x76\xae\xf5\x6e\xb1\xd2\xec\x77\x56\xc3\x5a\xef\xdb\xfb\x44\x2d\x57\x35\x37\x51\xb6\xd3\x43\x91\x1d\xcc\xfa\xd8\xb3\xc0\x36\xc0\x17\x59\xa2\x1f\xdf\x2a\x57\x18\x28\xc8\x6c\x61\x81\x6e\x93\x33\x0c\xe8\x59\xd6\xb0\xc0\xb7\xc9\x1d\x06\xfc\x47\xf6\x62\x01\x6e\x93\x43\x0c\xe0\xbb\x2c\x62\x41\xc3\x26\x97\x18\x68\xe4\xd9\xe4\x8d\x08\x14\x8d\xaf\x63\x7a\x02\xf6\x55\xa2\x59\xaf\x8f\x49\x62\x4e\x40\x38\xf6\xf3\xde\x9c\xbd\x38\x1c\x3f\x63\x9e\x6a\x9b\xb5\x35\x6f\xbe\x6a\x72\xd1\xbe\x17\x63\xa1\x89\x43\x42\x63\x11\x91\x6e\xbb\x76\x9d\x35\x99\x6d\x99\x79\x9b\x9e\xdc\x88\xfc\x56\x0d\xba\x11\xfc\x2d\xba\x75\x23\xf0\xdb\xb4\xee\x46\xe8\x37\xe8\xe3\x8d\xb8\x6f\xd7\xd4\x1b\xe1\x0f\xee\xf0\x27\xc3\xfc\xd4\xdd\x64\x98\x1d\x61\x3c\x9a\xcc\x59\xb0\xce\x4f\x00\x06\xe4\x19\xd4\x71\xa7\xe9\xa0\x3c\x4e\x36\xd8\x9d\xd0\x9b\xe4\x27\x35\xf2\x21\xbb\xb3\x1b\x95\x31\xd9\xb8\x91\xf1\x3c\xe4\x64\x18\x8d\x2a\x68\xc3\x6c\x72\x7e\x1c\xb0\xc9\xc0\xee\xbc\x42\x71\x2e\x4d\xa3\x32\x9e\xa9\xb3\x95\xfa\x29\xcb\x68\x5c\x1f\xba\xd9\xc0\x3b\x82\xc0\x9b\x82\xbb\xdd\x1e\xd5\x51\xd2\x2c\xfc\x14\x64\xf3\x40\xdc\xa0\x55\x30\xa7\x5f\xcd\xe0\x5f\x88\x88\xe0\x5d\xa3\x21\xb8\x67\xab\x58\x48\xa2\x88\xd2\xdf\x1b\xbd\xc2\x3b\x82\xb6\x5b\x23\xd8\x24\x99\x4d\xd2\x04\x95\x0c\x51\x8c\x82\xc1\x6c\x92\x0a\xce\xe2\x70\xf6\x1b\x13\x88\xaa\x73\xa5\x46\x6a\xae\xba\xbf\xdd\x36\xb7\x70\xa4\xd9\xd5\xfc\xc9\x50\x62\xcf\x26\xc3\xc4\x2c\x4a\xfb\x8d\x8a\x69\xaa\xc9\xca\xa0\xa8\x9d\x5a\xc8\x02\x42\x01\xa7\x66\x56\x33\xf5\x0c\xf2\xa5\x3c\x38\x83\xcb\x16\x8d\x34\xc9\xab\x80\x5d\x4c\x9b\x15\xae\xe0\x41\xa1\xa3\x36\xdd\x34\x08\xca\x85\x52\x6a\xa5\xd5\xb6\x99\x20\x38\x0e\xb6\x5b\x4b\x81\x76\xae\xbf\xaf\x4c\x95\x99\xb3\x4a\x00\xb1\x96\xac\x32\xff\x6b\x09\xf7\xc8\x5e\xf6\x95\x4a\x4e\x99\xc9\x38\x6b\x2d\x87\x9c\xf1\xb5\x04\xc8\x72\xc9\xbe\x32\xe4\xb3\x66\x79\x26\xb2\x96\x24\x9f\xf7\xb5\x84\x91\x19\x77\x5f\x51\xd4\x9c\x99\xca\xd5\xd6\x62\xa8\x39\x5f\x4b\x88\xa2\x26\xd9\x57\x90\x72\xde\xac\xac\x6a\xac\x05\x2a\xe7\x7e\x2d\xa1\xf2\x12\x6e\x5f\x99\x8a\x69\xb3\xa2\x04\xb4\x96\xa8\x98\xb9\xb7\x40\x2d\x69\xa0\x75\x6a\x95\xe7\x7a\x6d\xd2\x91\x0a\x5a\x05\x05\x55\x8e\x4c\x07\x79\x8b\xd1\x26\xe4\xaf\x39\x21\x4d\xd8\x6f\x07\x7d\x62\xb6\xaf\x9b\xc3\x68\x57\x96\xce\x01\xc4\x95\x27\x1d\x46\x38\x77\xc2\x03\x88\xe6\x51\xe8\x30\xb2\x65\x08\x3b\x80\xb0\x0c\xe4\x87\x51\xcd\x52\xc0\x01\x24\x2b\x39\xf0\x30\xca\xb5\x24\x7a\x00\x03\x45\x79\x71\x18\xf5\x5d\x71\xd2\x4d\xba\xe3\x56\x9b\x5f\x1b\xa6\x4c\x86\x79\xf5\x5b\xbf\x7c\xd4\x26\x60\x9b\x9b\x6f\x36\x1c\xc5\x21\x86\x77\x4f\x17\xef\x9e\x65\xc5\xdb\xdc\x1a\x37\x97\xbb\xe5\xb4\x67\x13\xbb\xa8\xa6\xd6\xcd\xc6\xfd\x84\x9f\x31\x27\x62\xbd\xdd\x16\x7f\x95\x74\xbc\xd9\x80\x2b\x8b\x7e\xd8\x6e\x5b\xa2\x4f\x95\xfd\x80\x89\x6e\xa5\x56\xc6\xaa\xb3\x12\x3d\x11\x6d\xb6\x8b\xcd\x6e\xd1\x7a\xd8\xc5\x5e\x35\xbb\x26\x92\xe5\x84\x87\x6c\x53\xbb\x73\x7c\xcb\xad\xc9\x10\x99\x2c\x67\x5e\x2c\xa6\xeb\x0d\xdc\xfc\xe3\xee\x73\xa3\x93\xca\x7a\xbe\x66\xd3\x66\x66\xad\xaa\x78\x14\x63\x6a\x61\xcc\xec\x59\x49\x5f\xce\x21\xc1\x74\xb0\xd9\xb8\x3f\xb1\x54\xfc\x8a\x7d\x1c\x8b\x47\xf5\x8c\x20\xfb\x99\xd9\x6c\x50\x43\xec\x00\x54\xa0\xd1\x55\x6d\x78\xfe\xfc\x62\x00\x01\x12\xc8\x11\x2c\x0c\x29\x76\xd4\x9d\x7e\xba\xb3\x9e\x01\x93\x61\x74\xd5\xc3\x4d\x45\x1f\xc5\xe3\x8e\x1e\x01\xd4\xb4\x15\xed\x1f\x04\x59\xdf\x9c\x79\x68\xca\xb8\xc0\x81\x7d\x13\xab\x51\xa4\xa4\xc1\xa7\x05\x9b\xe5\x64\x5d\xca\x42\xed\xf6\x20\xd9\x1f\x1c\xe6\xde\x5a\x06\x8d\xd2\x5f\x41\xf5\xdc\xea\xd6\xef\x98\xa7\x72\x19\xab\xbb\xf2\x22\x38\x50\xeb\xc8\x17\x08\x16\xc8\xc1\xaf\x3e\x45\x4b\x24\x57\xbc\x23\x38\x41\x71\x28\x97\x01\xe2\x04\x39\xd9\x2e\xe4\x74\x20\xf8\x0a\x0f\x8a\x96\xdb\x5e\xdc\xee\x48\xd0\x7c\x95\x36\x72\x1b\xe6\xb1\x34\x0d\x64\x0b\xa2\x10\xae\x11\xb9\xf7\xd3\x70\xd5\xce\x8d\xa7\xf3\x1b\x28\x83\x5d\x7b\xb0\xee\x84\x46\x46\xe4\xd2\x01\x63\xb4\xc4\xd2\xe7\x2a\x09\xc1\x18\xf1\x7a\xe9\x64\x2b\xa2\x4c\x3e\x68\x89\x07\xc6\x25\xb3\x3f\x72\x16\xc3\xc1\xfd\x80\x53\x9f\x93\x44\x14\xab\x6c\x2f\x7b\x97\x68\x28\x4f\x82\x12\xf1\x91\xc4\x4f\x52\x64\x10\x88\x87\x58\x4c\x07\x73\x8a\xe2\xa7\xc1\x4c\x5e\xdf\x5b\x0b\x93\x21\x25\xfb\x2c\x57\xdb\x68\x62\x8f\xdb\xd5\x9a\x1c\x3a\xb6\x9f\xcf\xbe\x14\x7c\x40\xc6\x35\x5f\x9e\x73\x18\x36\x0b\xb1\x72\x64\xed\x6d\xad\x3c\xab\x26\xdc\x49\xb6\x82\x76\x28\xa7\x8b\x55\xac\xc6\x9e\x9e\xe9\x8f\x5e\x53\x01\x59\x6a\xe2\x29\x4c\x21\x60\xfe\x6a\x89\x63\xe1\xfe\x7b\x85\xf9\xfa\x13\xa6\xd8\x17\x8c\xff\x40\xe9\xe9\xc9\x1f\x5a\x1a\xfb\x7c\x72\xd6\x3c\x00\x10\xa7\x8c\x62\x97\xb2\xf0\xb4\x00\x3d\xbb\xab\x6f\xaa\x2e\x18\x87\xd3\x67\xc4\x81\xc0\xb4\xa4\xec\x52\x1c\x87\x22\x02\x07\x46\x77\x40\x60\x36\x85\xcb\x3b\x20\x8e\xd3\xe4\x17\xd4\x71\xc9\x6c\xce\x1f\xe4\xb3\xcb\x62\x9f\x12\xff\x09\xa6\x50\x8a\x88\x4d\x73\xe4\x0b\xbb\x99\x03\xb8\xd9\x51\xe0\x9f\x59\x80\x5d\xe5\xc8\x8f\x24\x15\x6e\x86\x7a\x7a\x92\x9d\x10\x68\x0a\x26\x5f\xdb\xfa\xa5\xdd\xaa\xda\x9e\x9d\xe6\xc3\x27\xc3\x42\xf5\x93\x61\xb6\x75\x7e\x34\x19\xaa\xff\x27\xe0\xbf\x01\x00\x00\xff\xff\xbc\x30\xfb\x2c\x37\x40\x00\x00")

func templatesAnalysisTemplateHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesAnalysisTemplateHtml,
		"templates/analysis-template.html",
	)
}

func templatesAnalysisTemplateHtml() (*asset, error) {
	bytes, err := templatesAnalysisTemplateHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/analysis-template.html", size: 16439, mode: os.FileMode(420), modTime: time.Unix(1494508002, 0)}
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
	"templates/analysis-template.html": templatesAnalysisTemplateHtml,
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
		"analysis-template.html": &bintree{templatesAnalysisTemplateHtml, map[string]*bintree{}},
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

