// Code generated by vfsgen; DO NOT EDIT.

package install

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// templates statically implements the virtual filesystem provided to vfsgen.
var templates = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2019, 8, 15, 15, 27, 5, 612581097, time.UTC),
		},
		"/flux-account.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "flux-account.yaml.tmpl",
			modTime:          time.Date(2019, 8, 12, 11, 23, 23, 262609384, time.UTC),
			uncompressedSize: 836,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x92\x4b\xaf\xd3\x30\x10\x85\xf7\xfe\x15\x47\xba\x8b\x0b\xe8\x26\xa8\x3b\x94\x5d\xdb\x05\x0b\x10\x8b\xf0\xd8\x20\x16\x63\x7b\x42\x4d\x5d\x3b\xf2\x23\x3c\xa2\xfc\x77\x94\xa4\x95\x9a\xb6\x20\x55\xba\x3b\x7b\x7c\xc6\x73\xe6\xe8\x2b\x8a\x42\x3c\xe0\xd3\x8e\x11\x39\x74\x46\x31\x48\x29\x9f\x5d\x7a\x82\xb2\x39\x26\x0e\x08\xde\x72\x7c\x02\x39\xbd\x28\x41\x1a\xa7\x8d\xfb\x0e\x0a\x2c\x1e\xe0\x9d\xfd\x0d\xc7\xac\x59\xa3\xf1\x01\xef\xb2\xe4\xe0\x38\x71\xc4\x4f\x93\x76\x53\x4b\x21\x29\xb2\x1e\x27\x70\x8c\x50\xde\xa5\xe0\x2d\x5e\xd4\x9b\xf5\xf6\x65\x29\xa8\x35\x5f\x38\x44\xe3\x5d\x85\x6e\x25\xf6\xc6\xe9\x0a\x1f\x67\x57\xeb\xd9\x94\x38\x70\x22\x4d\x89\x2a\x01\x58\x92\x6c\xe3\x78\x02\x1c\x1d\xb8\x42\x63\xf3\x2f\x71\x7e\xe9\x7b\x98\x06\xe5\x07\x3a\x70\x6c\x49\x31\x86\xe1\xf8\x3e\x5d\x2b\xf4\xfd\xf2\xb5\xef\xc1\x4e\x0f\x83\x18\x73\x39\x37\x14\x24\xa9\x92\x72\xda\xf9\x60\xfe\x50\x32\xde\x95\xfb\x37\xb1\x34\xfe\x75\xb7\x92\x9c\xe8\xe4\x77\x3b\x27\x54\x7b\xcb\xf7\x9a\x15\x21\x5b\x9e\x24\x05\xa8\x35\x6f\x83\xcf\x6d\xac\xf0\xf5\xf1\xd5\xe3\xb7\xa9\x2f\x70\xf4\x39\x28\x5e\x14\x3b\x0e\xf2\xac\x50\xc0\x79\x57\x1f\x85\x9f\xeb\xf7\xff\xd6\x3e\xc3\x86\x9b\x99\x80\xfb\x17\xf5\x96\x6b\x6e\x46\xd1\x69\xd1\xff\xcc\x17\xc0\x75\xb6\x8b\xff\x62\x96\x3f\x58\xa5\x63\x76\x37\xc1\xb9\xb2\x73\x89\xc1\x25\x27\xb7\xc8\xb0\x71\x3c\x69\x6e\x28\xdb\x34\xa3\x32\x12\xf5\x37\x00\x00\xff\xff\xfd\x7f\x67\x6a\x44\x03\x00\x00"),
		},
		"/flux-deployment.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "flux-deployment.yaml.tmpl",
			modTime:          time.Date(2019, 8, 15, 15, 27, 5, 612581097, time.UTC),
			uncompressedSize: 6211,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x58\x5f\x6f\x1b\x37\x12\x7f\xf7\xa7\x18\x28\x0f\x69\x00\x69\x65\xd5\x6d\x71\xd8\x9e\x0b\xa4\x49\xe3\xcb\xa5\x71\x8c\xf8\x72\x87\x7b\xba\x52\xdc\x91\x96\x10\x97\xdc\xe3\x90\x52\x05\xa3\xdf\xfd\x30\xe4\xfe\xe1\x5a\x72\x52\xe4\xed\xf2\x10\xdb\xdc\xe1\xfc\x25\x7f\xf3\x1b\x2e\x16\x8b\x0b\xd1\xaa\x7f\xa2\x23\x65\x4d\x09\xa2\x6d\x69\xb9\x5f\x5d\xec\x94\xa9\x4a\x78\x8d\xad\xb6\xc7\x06\x8d\xbf\x68\xd0\x8b\x4a\x78\x51\x5e\x00\x18\xd1\x60\x09\x1b\x1d\x7e\x7f\x78\x00\xb5\x81\xe2\x56\x34\x48\xad\x90\x08\x7f\xfc\xd1\x7d\x8f\x7f\x96\xf0\xf0\x30\xfd\xfa\xf0\x00\x68\x2a\x16\xa3\x16\x25\x2b\x73\xd8\x6a\x25\x05\x95\xb0\xba\x00\x20\xd4\x28\xbd\x75\xfc\x05\xa0\x11\x5e\xd6\xbf\x8a\x35\x6a\x4a\x0b\xb9\x6d\x96\xf6\x4e\x78\xdc\x1e\xd3\x47\x7f\x6c\xb1\x84\x8f\x28\x1d\x0a\x8f\x17\x00\x1e\x9b\x56\x0b\x8f\x9d\xb2\x2c\x02\xfe\x27\x8c\xb1\x5e\x78\x65\xcd\xa0\x1c\xa0\x75\xb6\x41\x5f\x63\xa0\x42\xd9\x65\x6b\x9d\x2f\x61\x76\x75\x79\xb5\x9a\xc1\x33\xf0\xa8\x75\x26\x01\xde\x02\x49\x27\x5a\x84\x65\x83\xde\x29\x49\x1c\x5c\x6b\x95\xf1\xcf\x09\x78\x73\xd1\x29\xd6\x93\x18\x1e\x45\x01\xd0\xe7\x22\xfe\x8e\x6e\xaf\x24\xbe\x94\xd2\x06\xe3\x6f\xa7\x82\x00\x7b\xab\x43\x83\x83\xaa\x45\xa7\x6a\xab\xfc\x62\x87\xc7\xc1\x00\x71\x16\xfc\x68\xb0\x5f\x19\xf5\x2d\x78\x4b\x15\x0b\x9c\x49\x55\xb8\x11\x41\xfb\xf7\xb6\xc2\x12\x2e\xbf\xbb\xbc\x84\x67\x70\xa8\xd1\x40\xc3\xde\x60\x05\x0e\x45\xb5\xb0\x46\x1f\xe7\x70\x40\x38\x58\xf3\xdc\xc3\x1a\x41\xac\x35\x72\x3e\x64\xdd\xd8\xea\xa2\x53\xf8\x0c\xfe\x51\x2b\x02\x45\x20\xc0\x37\xed\x86\x20\x10\x56\xb0\xb1\x0e\xb6\x68\xd0\x09\xaf\xcc\x16\xee\xef\xff\x06\x3b\x3c\x52\x01\x6f\x0d\xbc\xfb\x0b\xc1\x4f\xd7\xb0\x2a\x56\x97\xf3\x41\x4b\x6f\x3b\x85\x40\x20\x1c\xe6\x7e\x90\x65\x57\x0c\x62\x05\x02\x08\x5b\xc1\x87\xa2\x4b\x14\x1c\x70\x50\x23\x85\x81\x83\x53\x9e\x1d\x2d\xce\xe7\x6f\x8b\x66\x48\x06\x36\xad\x3f\xbe\x56\x2e\x4f\x62\x83\x95\x0a\x4d\x09\xef\xb1\xb1\xee\x98\xc7\x89\xb0\xb1\x5a\xdb\x03\x47\xd4\x99\x56\x14\x43\x0d\xc4\x6b\x02\x64\x20\x6f\x1b\xc5\x19\xd8\x19\x7b\x30\xff\xa9\x2d\x79\x1a\x54\x6c\x94\xc6\x39\x1c\x6a\x25\x6b\x38\xda\x00\x07\xa5\x75\x0a\xca\x5b\xa8\x2c\xdf\x33\x5e\xe6\x4d\xfc\x8b\x03\x7b\x30\xec\xf6\xa0\xc0\x61\x6b\xc1\x09\x5f\xa3\x03\x5f\x0b\xd3\x19\xde\x2a\x5f\x87\x35\x58\x5e\x44\xd0\x6a\x87\x05\xfc\xdb\x86\xe7\x5a\x83\xd0\x64\x7b\x13\xd3\x64\x83\xf2\xa0\x8c\xb7\x71\x8f\xb4\xc6\x0b\x65\xd0\xcd\x61\x8d\xda\x1e\x0a\xb8\xc7\x31\xab\xb5\xf7\x2d\x95\xcb\x65\x65\x25\x15\x7c\xb0\x64\xc5\x57\x07\xcd\x92\xaf\x1e\xf9\xe5\x36\xa8\x0a\x69\x19\x08\x17\xad\x53\x7b\xe1\x31\x1e\x3d\x0e\xa4\xa8\x7d\xa3\x07\x4d\x7d\x2d\x88\xea\x85\xb4\x66\xa3\xb6\xc3\x27\x80\xb4\xf0\x5e\xb4\x65\xb6\x98\x5f\xa4\x45\xb6\xed\x6b\xeb\x52\xec\xc2\x1a\x97\x49\xc9\x78\xfc\xbe\x58\x93\x83\xa2\x9a\x57\x6a\xb1\x47\x10\x50\xa9\xcd\x06\x1d\x83\x66\xaf\xa1\xbb\x55\x23\x30\xc6\x12\x24\x75\x79\x11\x18\x5c\xf6\xaa\xc2\x3e\xed\x1b\xb5\x6d\x44\x3b\x3a\xa2\x7c\x0d\xc2\x00\x1a\xef\x8e\x31\x86\xdf\x92\xd0\x6f\x73\x10\xa6\x82\x60\xa4\x6d\x18\xad\xe3\xfe\x14\xed\xfb\x58\x4e\x61\xaa\x41\x0b\x9a\x7d\xd4\xa0\x90\xba\x7a\x9e\x54\x80\xd3\xf0\x15\x15\xc8\xb6\x7d\xb1\x02\x11\x09\xbc\x05\xd5\x30\x4e\xc2\xcd\xdd\x4d\x04\x01\xf8\x86\xc3\x22\xb5\x35\xca\x8c\xc6\x39\xb8\x3d\x3a\xb5\x51\x32\x02\x36\xb4\xc1\xb5\x96\x90\x5e\xfc\x89\x44\x0e\x5a\x12\x7c\xa4\x2c\x72\x82\xd8\xde\x9f\x48\x1c\x08\xb7\x1d\xaf\xe9\x13\x19\xdb\xb6\x5b\xc6\x0f\xca\x52\x33\x85\xe0\x67\x4f\x80\xf0\xe9\xbe\x33\x20\xdc\xa7\x73\xb8\x89\x27\xf8\x9f\x75\x88\x2e\xeb\x0e\x23\x4e\x1a\x0b\xb3\x32\xdd\xc4\x19\xa8\x46\x6c\x31\x9d\x7e\xde\x50\xc0\x1b\x65\xaa\x18\x73\xc3\xb0\xe2\x50\x8e\xa7\x36\x41\x8a\x46\x41\xc8\xe0\x11\xb7\x72\x11\x98\x27\x80\xf0\xc3\xbd\xaf\xc3\xba\xa8\xac\xdc\xa1\x2b\xa4\x6d\x96\x6e\x79\x40\xb1\xc7\x83\x75\x3b\x5a\xb2\x91\xa5\x17\x43\xfa\xfa\x5a\x72\xcf\x67\x3e\xc0\x96\xbd\xd8\x02\x7b\x5b\x0c\x32\xd1\x54\x09\x9d\x52\x65\x97\x09\x55\xe2\x8f\x72\x55\xac\xae\x8a\xab\xa9\xec\x5d\xd0\xfa\xce\x6a\x25\x8f\x25\xbc\xdd\xdc\x5a\x7f\xe7\x90\xf2\x48\x1c\x92\x0d\x4e\x22\xe5\x58\xee\xf0\xbf\x01\xc9\x4f\xd6\x00\x64\x1b\x4a\xf8\xfe\xb2\x99\x2c\x36\x11\xee\x4b\xf8\xe1\xbb\xf7\x6a\xa4\x0a\xd6\xe5\x9b\x17\x63\x75\xee\x22\x6d\xb8\xba\xbc\xe2\xee\xa9\xcc\xc6\xba\x26\x1e\x5b\xa1\x07\x69\xad\xf6\x68\x90\xe8\xce\xd9\x35\xe6\x1e\x70\x5a\x6f\xa6\x9d\x3b\x99\x4a\x0a\xa7\xcb\xc2\xd7\x25\x2c\x45\xab\x52\xa6\xf7\x3f\x2c\x55\x85\xc6\x2b\x7f\x2c\xda\xb0\xce\x64\x95\x51\x5e\x09\xfd\x1a\xb5\x38\xde\xf3\x1d\xad\xa8\x84\xef\x33\x01\xaf\x1a\xb4\xc1\x9f\xf9\xc6\x8d\x56\xfd\x7f\xb8\x9a\x5d\xdc\x49\x61\xce\x53\x24\x48\xad\xee\x2e\x79\x86\x5e\x46\xcf\xaa\x25\x51\xcd\x5c\xcf\x26\xf6\x09\xda\x76\x98\xb3\xe5\x92\x81\x32\xe9\xcc\x3d\xa7\xb4\x87\xa8\x5e\x4e\xa0\xb2\xcf\xd9\x07\xa3\x8f\x25\x78\x17\x90\xb5\x31\x0f\x8a\x28\xb5\xee\xc0\x9d\xaf\x55\x8b\x6e\x63\x9d\x44\x56\x9a\x88\x0f\xf3\x9e\xa7\x1c\xcf\xb9\xc9\xd4\xf7\xbd\x70\x9d\xef\x49\xec\xeb\xdc\xcf\xee\xe8\x5b\x23\x75\x88\xe8\xc9\xf4\x2d\x35\xb9\x1e\x59\x13\x3f\xf8\x02\x9d\xe9\x09\xcd\x8f\xbc\xf5\x11\xd5\x18\x10\x16\x2a\x94\x5a\x38\xa6\x6d\x6b\xbb\xcf\x00\xe0\x33\x54\x20\x41\x64\x1e\xbc\xb3\xd6\x2f\x0b\xa2\xfa\xc9\x00\x84\x99\x58\x9d\x8d\x6d\x6a\x96\x2c\xcf\x7b\x91\x4c\x03\x9a\xbd\x72\xd6\xc4\xa6\x90\xfa\xed\xec\xdd\xa7\x9f\x7f\x79\xf5\xe1\xf6\xcd\xdb\x9b\x59\x6a\x03\x73\xce\x87\xdd\xa3\x73\xd3\x9e\x9d\xa9\x89\x6d\x6e\x7d\x4c\x1d\xd5\xeb\x73\x31\x9e\x34\xdb\xd3\x18\xc7\xc3\xc9\xc2\x4f\x06\xca\x7d\x8f\x87\x8f\xde\x1a\xc3\x74\x46\x47\x3a\xef\x62\x4d\x32\x15\x8f\x49\x4d\x5e\xf4\xc8\x68\x7a\xfa\x2d\x0c\x08\xed\xd1\x19\xa6\xd7\x27\x1e\x6f\x9c\x6d\xf8\x58\xf4\xac\x65\x0e\x82\xf8\xb8\x75\x9d\x95\xd3\xa0\xad\xdc\xd1\x69\xb1\xd1\xec\xcb\x33\x79\x19\xd3\x3d\xc9\xcb\x5e\xe8\x80\x27\x39\xf9\xd2\x21\x7e\x7c\x06\xfa\xbe\xfb\x99\x13\xc0\x6d\x7f\xda\xee\x3f\xd3\xf0\x9f\x38\x97\x2c\x95\x18\xce\x44\x6e\x8a\x0f\xa3\xd3\x6c\xb2\x9c\xc4\x90\xca\x90\xc6\x34\xac\xb8\x11\x49\x21\x6b\xac\x38\xb3\x79\x69\x07\x66\xc9\x45\xe4\xb4\xcc\x33\x2d\xd6\x75\xd4\x31\xdb\xd0\x8d\x99\x71\xe3\x3c\x1a\xe1\xf1\x88\x42\xdb\xea\x23\x27\x82\xf2\x54\x8c\x04\xce\x1f\x2c\x7b\x19\xb8\xa4\xf1\xc0\xc5\x99\x38\xd6\x01\x6a\x7b\x88\x23\xa0\x35\x06\xa5\x8f\xe4\xce\x4f\x53\xb7\x58\x0c\x01\x44\xfe\xcf\xc6\xaf\x87\xa5\xa2\xe3\x3d\x05\xed\x65\x21\x75\x20\x8f\xae\x60\xfc\xd2\x79\x4a\x3e\x51\xba\x6a\x63\x2a\x5e\x25\xd1\xb7\x77\x93\xa0\xf8\xd6\x11\xfa\x38\x62\x4e\x0b\x3b\xfa\xd0\xcb\xf3\x20\xef\x1d\x4b\xc6\xa1\x2f\x43\xe0\xdc\xe3\x4e\xfa\xfa\x62\x42\xb4\x14\x41\x13\x28\x0e\xc1\x31\x7b\x0a\xab\x74\x9a\xd6\x11\xd7\x23\xc5\x89\xb3\xef\x37\xfd\x40\xf9\x22\xf7\xa5\xbf\x5b\xe9\x14\x32\x33\xcb\x46\xe0\x89\x23\x8c\x85\x09\xdf\x17\x95\x72\xd7\x27\xa8\x9f\xbb\xf5\x31\x23\x58\x63\xf1\x3e\x7d\xfc\x35\xcd\xe8\xc2\x6c\xd3\xb7\x1b\xe5\xe3\xdc\x48\xca\x5b\x77\x1c\xd0\xea\x0d\x93\xc3\x89\x71\xee\x41\xc1\xe9\xeb\x87\x07\x28\x6e\x94\x67\x4d\xf1\xa9\x67\x2a\xb1\x76\xc2\xc8\xba\x17\xfa\x39\xfe\x95\x1e\x7d\xd4\x26\x2e\xf1\xdd\xa0\x73\x3b\x99\x1f\xf0\xbe\xfb\x58\x06\xfa\xbb\x55\x26\xdb\x30\x9b\xcf\xba\xb7\x23\x4d\x98\x6f\x67\x7a\x75\xda\xaa\x0e\xc2\xc4\xe3\xe7\x90\xab\x2a\x13\xab\x6f\x84\x51\x1b\xe6\x7b\x7c\x40\x49\x55\xe8\x52\xac\x8f\x98\x73\x9c\x79\x2d\x21\x04\x53\xa1\x7b\x94\x40\x87\x5a\x78\xb5\xc7\x48\x67\xa8\x2f\xef\x76\x92\xc4\x47\x07\x7e\x08\x8e\xc2\xba\x52\x6e\x35\x4f\x3f\xbf\x1d\x1e\xc2\xc6\xe4\xc4\x87\xae\x73\xc9\x89\xaf\x47\x7d\x56\x7b\xa9\x33\x0a\x3e\x11\xba\x73\xfb\x03\xa1\x1b\x2a\xc7\x32\x70\x7e\xff\x2f\x8d\x50\x67\x1d\x40\xfe\xd0\x6b\xe8\xa5\xc6\xa7\xbc\xb3\xa0\x8b\x7c\x4f\x0f\x96\x13\x8a\x26\x3e\x0f\x71\x9e\xb8\x1b\x28\xff\x68\xc0\xcb\x73\xd5\xe1\x6a\x87\x9a\xd7\x9f\x81\xd1\x7e\x47\xa7\x8b\x77\x5d\xff\x75\x87\x47\x50\xd5\x4f\x83\xd8\x67\x5a\x65\xe6\x15\xab\x10\x3e\x38\x9c\x4c\x99\x67\x6c\xc5\xcf\xc7\xc5\x20\x4f\x4f\xc7\x6e\xf0\x77\xdf\x27\x20\x83\x44\x61\x60\x16\x5a\xf2\x0e\x45\x33\xeb\x01\x2b\x53\xf2\x0d\x16\xdb\x62\x0e\xff\xe2\x21\x0a\x5e\x69\x1b\xaa\x17\x45\x1c\xa2\xbd\xdd\x31\x7f\x23\x68\x85\xf3\x4a\x06\x2d\x1c\x74\x0f\x34\x9d\x96\xc7\x58\xdb\x59\xbd\x3e\x10\xcf\x69\x92\x75\x15\x71\x38\x2b\xd2\x74\xd6\x93\xf1\x47\xdb\xa2\xa1\x6b\xb1\x96\xab\x6f\xaf\x4e\xff\xcf\x03\xbe\x47\xb7\x3f\xf3\xf6\xc9\xb4\x63\x6c\x34\x5c\xb5\x1f\x73\xc4\x13\x3b\x46\x4a\x97\x50\x17\x7d\xf6\xa0\xfa\x3c\x7b\x93\xcd\x1e\x57\x39\xc4\xf8\x48\x10\x5b\xff\x14\x97\xb4\x22\x8f\x66\xd1\xb9\x70\x5d\x5e\x5d\x5e\xad\x2e\xba\x13\xfd\xb2\xaa\x54\x1a\xbb\x18\xcf\x5e\x72\x3b\x9f\x40\xc7\xf8\x7d\x6c\x69\x0f\x0f\xe0\x22\x3a\x7e\x61\xf7\x22\xbe\x6c\x4f\x6e\xc1\xf8\x5b\x6f\xe0\x43\xdb\xa9\x7f\x7d\x7b\xdf\xf7\x22\x9a\x77\x14\x29\xb8\xae\x33\x81\xa9\xac\x27\xb0\x51\x18\x1a\x71\x8c\xe3\xaa\xde\x8f\x0f\x17\x86\xb4\xb5\xbb\xd0\x82\x22\x0a\x48\x60\x0d\x90\x6d\x10\xde\x85\x35\x3a\x83\x1e\x89\xb5\x87\x96\xc6\x77\x89\xca\x50\x3f\x11\xcf\x6e\xad\xc1\x59\xfe\xe5\x55\x74\x20\x7f\x99\x48\xc6\x69\xfa\x58\xd1\x53\x9d\xe8\xdf\xe4\xcb\xc0\xc2\x66\xab\xd9\xc5\xff\x02\x00\x00\xff\xff\x4e\x9e\x09\xda\x43\x18\x00\x00"),
		},
		"/flux-secret.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "flux-secret.yaml.tmpl",
			modTime:          time.Date(2019, 8, 5, 10, 14, 20, 433647241, time.UTC),
			uncompressedSize: 137,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\xca\x31\x0a\xc2\x40\x10\x85\xe1\x7e\x4f\xf1\x2e\xb0\x82\xed\x1c\x42\x0b\xc1\x7e\xc8\xbe\xc8\x62\xb2\x19\x93\x89\x18\x86\xdc\x5d\x14\x1b\xcb\x9f\xff\xcb\x39\x27\xb5\x7a\xe5\xbc\xd4\xa9\x09\x9e\xc7\x74\xaf\xad\x08\x2e\xec\x66\x7a\x1a\xe9\x5a\xd4\x55\x12\xd0\x74\xa4\xa0\x1f\xd6\x57\xbe\x55\xcf\x85\x36\x4c\x5b\x04\x6a\x8f\xc3\x49\x47\x2e\xa6\x1d\xb1\xef\x3f\xfa\x4d\x41\xc4\xff\x8d\x00\x5b\xf9\x30\xdf\x8c\x82\xb3\xe9\x63\x65\x7a\x07\x00\x00\xff\xff\x40\x21\xa1\xbb\x89\x00\x00\x00"),
		},
		"/memcache-dep.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "memcache-dep.yaml.tmpl",
			modTime:          time.Date(2019, 8, 5, 10, 14, 20, 433647241, time.UTC),
			uncompressedSize: 874,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x93\xcd\x6e\x9c\x40\x10\x84\xef\x3c\x45\x49\x7b\x0d\x1b\x61\x69\x2f\xdc\xa2\x38\x89\x2c\x25\xd6\x5e\x9c\x7b\x7b\x68\xf0\x28\xf3\x97\xe9\x66\xb3\x04\xf9\xdd\xa3\xd9\x5f\x36\xf6\x9c\x80\xaa\xaf\xa7\xa6\x80\xba\xae\xab\x15\x3c\x7b\x43\xe6\x85\x3b\x74\x9c\x5c\x9c\x3c\x07\xc5\x28\xdc\xe1\x79\xc2\x57\x37\xee\xa1\x11\x07\x47\xb5\x82\x89\x41\xc9\x06\xce\xb0\x9e\x06\x86\x67\xa5\x8e\x94\xd6\x15\x25\xfb\x93\xb3\xd8\x18\x5a\x50\x4a\xf2\x71\xd7\x54\xbf\x6c\xe8\x5a\xdc\x5f\xc6\x56\x67\x7b\x5b\x01\x81\x3c\xb7\xd7\xdd\xe7\x19\xb6\xc7\xfa\x91\x3c\x4b\x22\xc3\x78\x7d\x3d\x99\x0e\xb7\x2d\xe6\xf9\x56\x9d\x67\x70\xe8\x8a\x4d\x12\x9b\x32\x31\x73\x72\xd6\x90\xb4\x68\x2a\x40\xd8\xb1\xd1\x98\x8b\x02\x78\x52\xf3\xf2\x9d\x9e\xd9\xc9\xf1\xc1\x9b\x00\x15\xa0\xec\x93\x23\xe5\x13\xb2\x08\x5b\x96\xbb\xa1\xdf\xe3\x81\x73\x94\xb2\x2e\x5d\x5d\x98\xfa\x5d\xa6\xac\x43\x9b\x0b\xa1\x6d\xd6\x9b\x75\xb3\xb9\xd5\xb7\xa3\x73\xdb\xe8\xac\x99\x5a\x3c\xf4\x8f\x51\xb7\x99\xa5\xd4\x7a\x76\x51\x1e\x16\xf9\x6a\xd4\x1e\x9b\xe6\x0e\xc0\x0a\x3f\x68\x6f\xfd\xe8\xcb\x0e\x31\x4f\xe5\x95\x8e\xc2\x1f\x60\x03\x3c\x0f\xf4\x3c\x29\xcb\x12\x7c\xc0\xc6\xe3\x06\x14\xfb\x97\xd1\xc7\x8c\x18\x18\x56\xd9\x2f\xed\x09\x4d\x73\xd7\x34\x58\xe1\x9e\x7b\x1a\x9d\x22\xc5\x7c\xcd\xb5\x2a\x9e\xdd\xee\x78\xf9\x14\x4c\xf4\x87\x8f\x4c\x23\x06\x56\xb8\x38\x08\x62\x0f\x26\xf3\x82\xcc\xbf\x47\x16\x05\x85\x0e\x99\x25\xc5\x20\xbc\xbe\x0c\x2a\x53\x6f\x4e\x78\xec\xd3\x38\xcb\x41\xaf\x07\x58\x74\xbf\x8d\x59\xdb\x63\xba\x8b\x2c\x6c\xc6\x6c\x75\xfa\x1c\x83\xf2\x5e\xdb\x05\x97\xc7\xf0\x49\x9e\x84\xf3\xff\xcc\x49\xfa\x96\xe3\x98\xde\x6a\xe4\x5c\xfc\xb3\xcd\x76\x67\x1d\x0f\xfc\x45\x0c\x39\xd2\xc3\xaf\xd0\x93\x13\xae\xfe\x05\x00\x00\xff\xff\x5d\x9a\x63\xab\x6a\x03\x00\x00"),
		},
		"/memcache-svc.yaml.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "memcache-svc.yaml.tmpl",
			modTime:          time.Date(2019, 8, 5, 10, 14, 20, 433647241, time.UTC),
			uncompressedSize: 206,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x8c\x3d\x0e\x02\x21\x10\x46\x7b\x4e\xf1\x5d\x00\x13\x2c\x39\x84\x8d\x89\xfd\x04\x3e\x23\x51\x58\x02\x64\x9b\xc9\xde\xdd\xb0\x6b\xe3\x76\xf3\xf3\xde\xb3\xd6\x1a\xa9\xe9\xc1\xd6\xd3\x52\x3c\x56\x67\xde\xa9\x44\x8f\x3b\xdb\x9a\x02\x4d\xe6\x90\x28\x43\xbc\x01\x8a\x64\x7a\x64\xe6\x20\xe1\xc5\xa8\x8a\xf4\xc4\xe5\x26\x99\xbd\x4a\x20\xb6\xed\x07\xed\xab\x87\xea\xff\x57\x15\x2c\x71\x62\xbd\x32\xcc\x62\x5d\xda\xe8\x73\x00\xec\x39\xbf\x5f\x0f\xc4\xc3\xb9\xab\x73\x06\xe8\xfc\x30\x8c\xa5\x1d\xce\xd9\xf8\x06\x00\x00\xff\xff\x20\x2f\xef\xba\xce\x00\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/flux-account.yaml.tmpl"].(os.FileInfo),
		fs["/flux-deployment.yaml.tmpl"].(os.FileInfo),
		fs["/flux-secret.yaml.tmpl"].(os.FileInfo),
		fs["/memcache-dep.yaml.tmpl"].(os.FileInfo),
		fs["/memcache-svc.yaml.tmpl"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}