// Code generated by vfsgen; DO NOT EDIT.

package migrations

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

// SqlFiles statically implements the virtual filesystem provided to vfsgen.
var SqlFiles = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Time{},
		},
		"/1_base_schema.down.sql": &vfsgen۰CompressedFileInfo{
			name:             "1_base_schema.down.sql",
			modTime:          time.Time{},
			uncompressedSize: 88,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x76\xf6\x70\xf5\x75\x54\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x86\x0a\xc4\x3b\x3b\x86\x38\xfa\xf8\xbb\x2b\x38\x3b\x06\x3b\x3b\xba\xb8\x5a\x73\xe1\x55\x1d\x10\xe4\xef\x0b\x57\x0a\x08\x00\x00\xff\xff\xac\xa9\x98\xf1\x58\x00\x00\x00"),
		},
		"/1_base_schema.up.sql": &vfsgen۰CompressedFileInfo{
			name:             "1_base_schema.up.sql",
			modTime:          time.Time{},
			uncompressedSize: 24403,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x3c\x7f\x73\xe3\xb6\xb1\xff\xeb\x53\xec\x74\x9c\x4a\xbc\x88\x3a\xbb\x69\x3b\xaf\x76\x7c\x33\x3a\x99\xf6\xf1\x45\xa6\x5c\x49\xce\x5d\x5e\x26\xc3\x81\x48\x48\x42\x4c\x91\x0a\x41\xd9\xd6\x9b\x7e\xf8\x37\x58\x80\x24\x40\x52\xb2\xe4\xbb\x4b\xdb\x79\xd5\x1f\xc9\x99\x04\x81\xc5\xfe\xde\xc5\x2e\x6c\xdb\x1b\x4d\x9d\x49\xcb\xb6\xa7\x4b\xc6\x21\x48\x42\x0a\x84\xf3\xcd\x8a\x72\xc8\x96\x24\x83\x8c\xcc\x22\x0a\x31\x11\x0f\x02\x12\x43\x12\x47\x5b\x98\x51\xf8\xeb\x77\x10\x2c\x49\xca\x21\x4a\xe2\x45\xab\x35\x18\x3b\xfd\xa9\x03\x93\xc1\x07\xe7\xb6\x0f\xee\x35\x78\xa3\x29\x38\x9f\xdc\xc9\x74\xa2\x1e\xfa\x83\xfe\xb4\x3f\x1c\xdd\x5c\x80\x6d\x43\x40\x32\x12\x25\x0b\x39\x3b\x87\x6f\x81\xc5\x19\x4d\x63\x12\xc1\x7c\x13\x07\x19\x4b\x62\x7e\xc8\x94\x77\xe3\xd1\x2d\xce\x17\x92\x8c\x94\x93\xad\x37\xb3\x88\x05\xda\x54\xf9\x5c\xce\xa7\xa9\xe3\x4d\xdc\x91\x57\x99\x2e\x63\x2b\xca\x03\x12\xd1\x70\x06\x1f\xdd\xe9\x87\x7c\x51\x39\xd1\x45\xcb\x6e\xfe\xb5\x6c\x1b\xa6\x88\x9f\x90\xce\x59\xcc\x70\x31\xc0\xe7\xcd\xe3\x73\x38\xa6\xfd\xf7\x43\xa7\x82\x97\x1e\xa7\x29\xa3\x1c\x3a\x2d\x00\x00\x16\xc2\x8c\x2d\xc4\x23\x12\xc1\xdd\xd8\xbd\xed\x8f\x7f\x82\x1f\x9c\x9f\xba\xf8\x76\x45\xb3\x94\x05\x3e\x0b\x05\xde\xe4\xa3\x88\xcc\x68\xc4\xc5\xdf\x3f\xff\x22\x9f\xdc\x7b\xee\xdf\xef\x9d\x8e\x7c\x61\x81\xeb\x0d\x86\xf7\x57\x0e\x74\x58\x68\xb5\xac\x8b\x1c\x14\xd7\xbb\x72\x3e\x81\x5c\xdb\x97\x63\xc5\xbc\x23\x6f\x07\x78\xf7\x13\xd7\xbb\x81\x1b\xd7\x83\x7c\xe6\x8b\xfd\xdb\xc2\x51\xe5\xae\xe4\x96\x24\x84\x0f\x74\x0b\x53\xe7\xd3\x54\xfe\xf5\x48\xa2\x0d\x85\x8c\x3e\xab\x1d\x69\xbb\x46\xa0\xcb\x1d\x3c\xd0\x6d\x57\x0e\xb7\xf4\xad\x1a\x2f\x6a\xfb\x7d\x19\x48\xff\x81\x6e\xfd\x75\xc2\x91\x8e\x0a\x62\x89\x69\x0d\x2a\x13\xe6\x75\xc2\x4b\x12\xe4\x50\xc8\x6f\xba\x62\xa8\x06\xc6\x3a\xe1\x75\xbc\xd7\xd1\x5c\x87\x64\xcf\x7c\x2f\xec\x4a\x01\x5f\xe0\x7e\xe2\x8c\xdd\xfe\x70\x27\x3b\x09\x21\xd7\x76\x8a\xe2\x24\x1f\x8a\xff\xc8\x87\x21\x9d\x93\x4d\x94\xf9\xc1\x72\x13\x3f\xf8\x28\xb5\x8f\x24\x82\xf7\xa3\xd1\xd0\xe9\x7b\x70\xe5\x5c\xf7\xef\x87\x53\xc8\xd2\x8d\xfa\x20\xa5\x19\x8d\xc5\x36\xfc\x35\x4d\x59\x12\x82\xeb\x4d\x9d\xf1\x8f\xfd\x61\x31\xd6\xbb\x1f\x0e\xbb\x60\xdb\xe2\xff\x90\x25\xb0\xe1\x14\xb2\x25\xcd\x97\xaa\xcd\xd0\x80\x6b\x84\x52\xc3\x4d\x09\xba\xc1\x20\xfa\xf3\x17\x59\x22\x5f\xbe\x63\x90\xbd\x8e\x3c\xc9\xb6\xe2\x1d\x4e\xe9\x7a\x13\x67\x3c\x15\xbb\x1c\xed\x98\x11\xb9\x54\x31\xe9\x8f\xfd\xe1\xbd\x33\x69\x75\xda\x26\x3e\xdb\x5d\xe8\x14\x78\x6a\xff\x17\x2c\x93\x4d\xca\xdb\xd6\xf9\xb9\xa0\x8e\xd5\x6d\x75\xda\x55\xa4\x88\x2f\xfe\x76\x0a\x6f\x4a\xf4\xb6\xff\xf4\x67\xfc\xae\xf8\xec\xa2\x55\x6c\x78\x34\x86\xb1\x73\x37\xec\x0f\x1c\xb8\xbe\xf7\x06\x53\xb7\xce\x87\x0b\x9a\xf9\xcd\xc4\xee\x58\xb8\xef\xb1\x33\xbd\x1f\x7b\x93\x62\xc1\x56\x7f\x02\x27\x42\xe5\x9e\xe0\xeb\x89\x33\x74\x06\x53\x89\x9d\xf3\xf3\x02\xaa\xeb\xf1\xe8\x76\x17\xa6\x3f\x7e\x70\xc6\x8e\xc0\xf4\x65\x15\x1d\x17\x2d\x35\xf3\xb0\xef\xdd\xdc\xf7\x6f\x1c\xe0\xbf\x45\x30\x91\x74\xbb\xeb\x8f\xfb\xc3\xa1\x33\x84\x49\xff\xda\xb9\x78\xed\x1e\xab\x08\xfd\x3d\x76\x59\x23\xe2\x51\xfb\x3c\x62\xa3\x2b\xf2\x40\x7d\x25\x2a\x28\x03\x95\xdd\x65\x29\x5b\x2c\x68\x8a\xcf\xca\x0d\x5e\x39\x83\x61\x7f\xec\xb4\xde\x3b\x37\xae\x27\xde\x39\x9f\x9c\xc1\xfd\xd4\x81\x79\x92\xae\x48\xd6\x69\x37\x89\x8f\x30\xc9\xbd\x6f\xdc\x8e\x30\xa9\x30\x75\x6f\x9d\xc9\xb4\x7f\x7b\x37\xfd\x1f\xa5\x97\xe1\x6a\x74\x8f\x9b\x19\x3b\x03\x57\x58\xe2\x6e\x6e\x7b\x18\xaa\x06\xab\x2d\xa5\xaa\xfa\xf3\x9c\x8f\x3d\x4d\x7a\x2f\xf6\x80\x53\xd5\xac\x0a\x22\xe8\x14\x0b\x75\xd1\xe0\x6b\xfa\x42\x4a\xe3\x11\x4b\xdf\x39\xe3\xeb\xd1\xf8\x16\x82\x94\x92\x8c\xfa\xcb\xed\x9a\xa6\x12\xb1\x39\x30\xe6\xe2\xed\x6e\x75\x9a\x2e\xb4\x05\x10\x3b\xd6\x2c\x7e\x52\x10\xc4\xc8\x42\x1a\x2e\xdf\x1d\x21\xaa\x8d\x98\x3a\x91\xcc\x0b\xfd\xe1\xd4\x19\x37\x92\x0f\x26\xce\x54\x69\x3e\x34\x05\xa5\x7f\xd4\x0b\x92\xd5\x3a\xa5\x9c\x77\xf7\xbe\xf5\x39\x5d\xac\x68\x9c\xcd\xb6\x70\x09\xed\x02\xf3\xed\x17\xbe\x4a\xd2\x90\xa6\xf2\x1b\xc4\x0e\x8e\xb6\x2e\xe0\xe4\xa4\x86\xc0\x8b\x96\x78\x69\xdb\xb8\x63\x0e\x4f\x4b\x9a\x4a\xbb\x41\xe3\x10\x27\x07\xc6\x61\x46\xe7\x49\x4a\x21\x4e\x9e\x3a\x96\x7d\x76\x0a\x2b\x16\x6f\x32\xca\xe1\x89\x45\x91\xf0\x62\xf3\x85\x69\xa8\x53\x95\x84\xa1\x5f\x80\x24\xe7\xf7\xd7\x49\xc4\x82\xed\x11\xe4\x2d\x15\x71\xb9\x6e\x5b\x92\x43\x8a\x9d\xf8\xe4\xa2\xe5\x78\x57\x35\xa1\x5f\x47\xeb\x05\xff\x2d\xd2\x0c\xd4\xd8\xbd\xb9\x71\xc6\x50\x93\x63\xdf\x90\xdc\x6b\x41\x4e\x65\x80\x1a\x94\x00\x7e\x87\x23\xaf\x47\x63\x70\xfa\x83\x0f\x30\x1e\x7d\xc4\x07\x39\x83\xdc\x8d\x47\x03\xe7\xea\x7e\x5c\x77\x25\xea\x1a\x44\xa8\xa0\xd6\x2e\x57\x17\x7d\x63\xb7\xe6\xd4\xef\x71\x8e\x6d\x31\x17\x8c\x69\xb6\x49\x63\x20\x5a\xdc\x01\xb3\x0d\x8b\x32\x98\xa7\xc9\x0a\x88\xe1\xab\x90\x38\x04\x02\x7c\x33\x9f\xb3\xe7\x1e\x3a\xe3\x4b\x9a\x7b\x6c\x38\x80\x71\xe1\x89\xc4\x01\xc9\x68\x08\x3c\x51\x11\xcd\x92\xaa\x6f\x20\x48\x36\x51\x08\x73\x96\x01\x8b\x61\xbe\x89\xa2\xde\x31\x4a\x55\xa7\x83\x58\xce\x7f\x62\xd9\xd2\x97\x53\x97\xb2\x53\x73\xae\xf2\xc5\xd1\x28\x1b\x6a\x58\x8c\x69\x36\x30\x1d\xbe\x99\xf1\x2c\x65\xf1\x42\xf7\x78\x84\x3c\xc3\x5f\xbf\xb3\x3b\x22\x12\xf3\x23\x1a\x2f\xb2\x65\x47\xce\x6e\x7d\x7b\x66\x59\xf0\x8f\x7f\x40\xdb\x6f\x8b\xff\xa9\xa7\xe7\xe7\xb8\x46\x93\x95\x71\x6f\x6f\xef\x9b\x0d\x8d\x4e\x96\x98\x3e\x41\xb1\x34\x29\xdc\x63\x81\x03\x45\x01\xc4\xb9\x88\x14\x34\x57\x4e\x1f\x56\xa0\x1a\xde\x6f\x32\x60\x73\x31\x40\x7c\xa9\xd3\x2d\x4c\x28\x8f\xdb\x99\xa0\x4c\x17\x16\x34\xa6\x29\x11\x52\x2b\x97\xdf\xc4\xec\xb7\x8d\x64\x0d\x5c\xd2\x4b\x32\x9a\x53\x96\xc9\x20\x55\xac\xbb\x59\xe3\xd2\x31\x7d\xce\x84\xe9\x81\x64\xbe\x83\x80\x48\x3a\x81\x9f\x67\x31\x1b\x4f\x80\x65\xc0\x97\xc8\x19\x42\x43\x90\x28\xa2\xa1\x0c\x7b\xd9\xbc\x60\x4c\x01\x21\xc4\x49\x06\x5b\x9a\x01\x7d\x66\x3c\x3b\x86\x75\x62\xfa\xe4\xd7\xd8\xa7\x91\x65\x7c\x92\x2e\x14\xdb\x18\x41\xdf\xa1\x8c\x33\xe8\x4f\x9c\x62\xde\x8f\x1f\x1c\x0f\x74\x5e\xa9\x2c\x64\xc1\xf7\x22\xae\x9f\x7e\x70\x3c\xc3\x2a\x55\x86\x29\x26\xca\xdf\x3a\x43\x6d\x09\x5c\xfa\x55\x72\xb2\x63\x31\x6d\xdf\xd2\x8f\x35\x3e\xb0\x4a\x28\x1a\xf4\xa8\x60\xeb\x1f\x47\xc3\xfe\xd4\x6d\xe4\xea\x01\x1a\xef\x9c\xad\x24\x5d\x25\x5b\x2f\xd8\x23\x8d\x75\x8e\xec\xe5\x09\x92\x0d\xa7\x5c\xb0\x16\x4f\x56\x14\x38\xfd\x6d\x43\xe3\x80\x72\xc1\x35\x8a\x65\xf2\xfc\x88\xe4\x9b\x96\x6d\xbb\xc8\xe1\x5f\x84\x6d\x94\xb7\x61\xa8\xe1\x17\x98\x66\x74\x3f\x05\x95\x26\xc0\x7f\x57\x82\x3a\xab\x55\x77\xf7\x40\xa0\x43\xb1\xd9\x85\x72\xfe\xd4\x93\x4b\x14\x27\xe1\x58\xac\x17\xbe\x70\x3b\x64\x34\xef\xe7\x88\x28\xec\xa2\x49\xfb\x76\xb7\xcd\xc2\xb6\x65\x9d\x9f\xe3\x94\xc3\xd1\xe8\x0e\xc1\xde\x13\x2b\xe5\xa1\xab\x70\xd6\xb4\x9d\x75\x41\x0f\xe1\x0a\x86\x93\xbc\xae\xe0\xae\x7b\x54\x55\x96\xaa\x0d\x38\x48\x3c\x6b\x8c\x29\x97\x93\x60\x8c\x3c\x18\x8c\xbc\xeb\xa1\x3b\x98\xc2\xd5\x08\xbc\xd1\xf4\x83\xeb\xdd\x68\x42\xea\x7a\x37\xcd\x7b\xec\x89\x2d\x36\xbf\x29\x17\x57\xe8\x9a\x8e\x00\xbd\xd7\xe2\x39\xba\x13\x60\xdb\xb0\x89\x43\x9a\xc2\x92\x2d\x96\x10\x24\x71\xb0\x49\x53\x1a\x07\x5b\xe4\x3c\x16\x73\x9a\x66\xb0\x22\x5b\xe4\xbc\x54\xa9\xf2\x78\x9b\x2d\x59\xbc\xe8\xa2\x5d\x4c\xb7\xc2\x90\xd2\x88\x06\x19\x5a\xd5\x28\x49\xd6\xf9\xd4\xcb\x2c\x5b\xf3\xf3\xb7\x6f\x79\x46\x82\x87\xe4\x91\xa6\xf3\x28\x79\x12\x2e\xdb\x5b\xf2\xf6\xec\x2f\x7f\xfb\xcb\xe9\x77\x7f\xfa\xb3\x72\x22\xdc\xa9\xd4\x31\xd7\xa3\x7b\xef\x4a\x3a\x68\x39\x71\x56\xb8\xcf\xd5\x01\x7b\x92\x1e\x4a\x43\x04\xa5\x78\x62\xd5\x52\xba\x6c\xec\x18\xc6\xf5\xb2\x4a\x67\x05\x40\x0d\x2c\xc7\xbb\x02\xc1\x83\xcd\xce\xd7\xdd\xf0\xee\x66\xf2\xf7\xe1\x3e\xc5\x01\x37\x34\x53\x5a\x43\x66\xb7\x48\x9a\x92\x2d\x14\xa9\x23\xa9\x44\xe4\xab\x07\xba\xed\xc1\xb5\x78\x10\x6f\x15\x80\x5d\x31\xc5\x13\x85\x27\x12\x4b\x6f\x24\xff\x10\x6d\xe6\x8c\x02\xe1\xe8\x98\x12\x41\x0c\x2e\xde\x72\x66\xd8\x57\x54\x42\xa8\x81\xd6\x29\xcd\xb2\x2d\x2c\x29\x79\xdc\x42\x94\x04\x0f\xa8\x8a\x84\xf9\xe3\x6b\x22\xdc\x85\x68\x7b\x8c\x82\x11\x22\x2d\xb8\x7a\x9d\x70\x7f\x9e\xa4\xfe\x03\xdd\xee\x73\x64\x1e\xe8\xb6\xfc\xd3\xb4\x48\x2c\xce\x1a\x55\x0b\x94\x58\x42\x5d\x20\x9e\x08\x9d\xe2\xd7\x1f\xeb\x12\x08\x5e\xff\xd6\xb9\x28\x23\x51\xb0\x6d\xb1\xc9\x30\xd9\x88\x97\xc1\x92\x06\x0f\xb8\x7d\x16\x2f\x40\x04\x02\x6a\xcc\x9c\xf1\x0c\x92\x75\xc6\x56\x8c\x67\x2c\x90\x03\xcf\x35\xb6\x2c\x36\xb7\x4e\x78\xc1\x77\xad\x1d\x6a\xa1\x9e\x9a\x2b\x39\xb1\x82\x25\x93\x19\x8b\x97\x7d\xef\x0a\xd3\x99\x97\x05\xea\x4a\x21\xc8\xe7\x54\x5c\xeb\x5e\x4b\x76\x35\xad\xb1\x8a\x1c\xca\xb1\xca\xf4\x81\x7b\x6d\x8a\xdb\x1e\x71\xc2\x98\x45\x50\x3a\x49\xfd\x06\x8b\x52\x53\x76\x56\x09\xa3\x3e\x2e\x57\x3d\x02\xed\x82\x4d\x05\xd3\xeb\x01\xe5\x8c\x06\x44\xd0\xe8\x89\x02\x49\xd1\xeb\xa3\xf3\xb9\xd0\x2f\xc1\x92\xc4\x0b\x41\x28\x74\xc4\x83\x25\x5d\x11\x9d\x66\x24\xe2\x09\x06\x3a\x1c\xf8\x46\x85\x74\x26\x87\xcc\x68\x94\x3c\x01\x9e\x56\xa4\xa9\x98\x91\xc5\x90\xd1\x74\xc5\x85\x8f\xa7\x69\x3f\x23\xb2\xc9\x03\xb7\xe1\x68\xf0\x43\x73\xc8\xeb\x7a\x30\xf9\xd0\x1f\x3b\x70\x7f\x77\x25\x0f\x0b\x06\xc3\xfb\x89\xfb\xa3\x03\xb7\xa3\x2b\xa7\xdd\x35\x76\x6f\xe5\xdb\xe7\x34\x48\xe2\x50\xb1\x20\x99\x67\x34\x45\x46\xfc\x57\xe2\xb1\xaf\xc0\x5f\x25\x28\xe4\x19\x33\xd1\xf0\x2d\x9c\x7d\xbd\xcd\xc9\x1d\x18\x5c\x50\x6e\xc3\x64\x0e\x77\x82\x09\x65\x73\x5b\xe6\x90\xf3\x4b\x38\xc3\x53\xa3\x33\x9b\xc5\x21\x7d\xa6\xa1\x54\xe0\xbc\xbe\xdb\x3d\x1e\xca\x8e\xcd\x88\x9f\xcc\xea\x1a\xc9\xe9\x52\x59\x76\x4d\x68\x6a\xee\x43\x31\xcb\x2e\x37\x42\xe7\xa7\x5d\xa4\xf5\x46\xd3\x46\xf2\xf6\xdd\x89\x03\xed\x01\x7a\xaa\xc2\x19\x98\x33\x0c\xa1\x85\x25\xcb\x27\x69\x9b\x44\x6f\xe4\x8a\xfd\x49\x0b\xdb\x6e\xf4\x85\x95\x74\x10\x25\x2b\x2a\x62\x53\x01\xb9\x34\x96\x79\xdc\x87\x9e\xf1\x2b\xec\x96\xa4\x09\x0b\x3b\x86\x65\x52\xa9\x47\xfd\x81\x72\x8b\x5d\x6f\xaa\x7b\xc0\xd2\xbc\x34\xb9\xa7\x7b\x39\x5a\x3f\x76\x6a\x95\xf4\x2f\xbe\x29\xa0\xe9\x96\x70\xbc\xde\x6b\xc4\x25\x7b\x2c\xd4\xdd\xa7\x5d\x9e\x4e\x93\xc4\x56\xbf\xdd\xe9\x6e\x49\x92\x44\x0d\x22\xfa\x40\xb7\xba\x7a\xe9\x7b\x57\xc5\x2b\x99\xe4\xbd\xd4\x30\xfe\x59\x5e\x18\x72\xd3\x53\x4a\xd6\x6b\xc1\x39\x69\xb2\x89\x43\xf8\x95\x27\xf1\xcc\xa7\x24\x58\xfa\x82\x98\xc2\x67\x12\xf1\x1a\x10\x98\xd1\x4c\x70\x58\x9a\x3c\xf9\x94\x67\x6c\x45\x32\xda\xb2\x6d\x61\x98\xd4\xe9\x68\xe7\xec\x14\xd9\xfe\xec\xf4\xd4\x3a\x82\xbd\x24\x5b\x55\xd6\xed\xfc\xca\x25\x28\x5d\x40\x76\x12\x48\x29\x99\xab\x3c\xcc\xb4\x5a\x85\x53\x34\x71\xa6\xa3\x6b\x48\x69\x90\xa4\x61\x0b\x8a\xbd\xe6\x67\xdf\xad\x5d\x39\x18\x98\x4c\xc7\x82\x45\xc6\xa3\x8f\x13\x38\x3b\x2d\x38\x56\x08\xe3\x49\x05\xac\xf2\x45\x13\xee\x36\x71\x4c\x79\x89\xb2\x12\x61\x90\x23\xec\xf3\x70\x24\xe7\x97\x27\xc3\xbe\xf4\x8a\x49\xbc\xc5\x7f\xd4\xf0\x40\xe2\x2d\x8d\xe8\x8a\xc6\xd9\x17\xc3\x05\x2e\xa4\x80\x30\x10\xb1\x33\xe9\xb8\xe7\xd7\xf4\x0d\xdc\xc9\xd2\x82\xfe\x9d\xcb\xe1\xc0\x6f\x5e\x5c\xe7\x00\x2c\x17\x3e\x5b\x3d\x99\xc2\xe6\x3e\xaa\x4b\xbe\xdb\x53\x37\x5d\x73\x89\xd4\x4e\x9e\x15\xd8\x93\x11\x28\xdd\x49\x33\x42\x2b\x13\x40\x2f\x85\x69\x2a\x4a\xeb\x99\x71\xda\x0b\x1b\xd1\x47\x1f\x9b\xd3\xc9\x09\x94\x13\x1f\x59\x9d\x4a\x1b\x83\xab\x27\x7a\x3a\xa6\x9e\xed\x29\x62\x2c\x3c\x21\x90\xee\x71\x43\x0a\x73\x0e\x2c\x7b\x65\x32\xe7\x20\xf7\xfb\xdf\x8a\x94\x7b\x36\xd2\xab\x78\xa8\xf7\x9e\xc0\x44\x7f\x38\xd4\x00\x7a\xb3\x6b\xf5\xa6\x74\xd7\x11\xeb\x21\xaa\x86\xee\xad\x3b\x85\xb3\x63\xd9\x68\xbd\x9b\x8b\xf6\xc7\xfa\xc7\x30\x80\xe1\x41\xbe\x32\xce\x76\xbd\x29\x52\xf9\x44\x85\x23\xe8\x72\xd1\x67\x1a\xe0\x31\x17\x32\x6e\x92\x52\xa0\xcf\x6b\x1a\x73\xa1\xf2\xf3\xf4\x46\xb1\x35\x99\x81\x6f\x74\xc0\x1a\xbc\x87\x7f\x6a\x9c\x6c\x70\x4f\x15\xb2\x03\x52\x19\x8d\xfe\xb8\xc4\x67\xc1\x27\x1a\x8f\xe4\xfe\xc7\x0d\xd5\x7c\x54\x9f\x85\x8a\xe6\xa5\xd7\x07\x6b\xc2\xd2\xe3\x29\xcf\xc2\x8e\xee\x52\xed\x71\x58\xf7\xd3\x7c\xce\x52\x9e\xe5\xd9\xbb\x2c\x81\x75\x4a\x1f\x69\x9c\x15\xe9\x69\x79\x74\x36\xa3\x22\xe6\xde\x70\x1a\xc2\x26\xcf\xed\x09\x4b\x19\x50\xce\x49\xca\xa2\x6d\x13\x52\x5f\x72\x0f\x3f\xd7\x39\x7c\x35\x59\x6b\x9e\xbe\x8e\xb3\x97\x49\x8a\x2a\x5e\x3f\x58\x2a\x93\x0b\x84\xe7\xf1\xa0\xc4\x9b\xa0\x3c\xfa\x58\x2d\xdb\x3e\xe5\x90\xd2\x75\x4a\xb9\x40\xef\x03\xdd\xaa\x1a\x49\x82\x47\xcc\x02\xe1\x19\x74\x9e\x28\x84\x89\x90\xa1\x0d\xa7\x18\x8a\x8a\x78\x88\x09\x32\xb0\x38\x93\xf3\x16\x86\x83\x6f\xd6\xeb\x24\xcd\x80\x65\x56\x71\xc2\xc0\x8a\x57\x34\x85\x35\x4d\x31\x71\x21\x3e\x0f\x52\x96\xb1\x80\x44\x20\x67\xcb\x96\x8c\xb7\x6c\x9b\x71\x19\x66\x21\x61\x85\xa6\x2a\x33\xbe\x41\xc4\x04\x9c\x24\x0e\x31\x59\x42\x82\x25\x0d\xc5\xfb\x94\x1e\x6c\xa7\xa4\x73\x99\x25\xbe\xe6\xd1\x15\x8e\xaf\xd5\xd2\x38\xf2\xe7\x5f\xa0\xe4\x49\xac\xa1\x64\xe1\xb3\xff\x48\x22\xf1\xb8\x53\x49\xd5\x1b\x09\x78\xdb\x96\x3b\x10\x51\x22\x82\x5f\x9e\x16\x67\x49\x6e\x82\x45\x84\x2a\x38\xab\x48\x95\x56\xa7\xc0\xac\x35\xea\x30\x16\x72\xa5\xd4\xb6\x8a\x12\xa8\xcd\xa0\x23\xc4\x56\x47\x68\x4a\x09\x4f\x62\x6e\x19\x53\x05\x09\x89\x28\x0f\x68\x27\x7a\x58\xf7\xd6\x09\xaf\x1e\x16\xec\x57\xe2\xbf\x72\xfb\xdd\xbb\xb6\x2f\x1d\x1a\xbf\xdd\x05\xda\x7b\xa0\x5b\xcb\x12\xc8\xe8\xee\x58\xa7\x57\x3f\xb2\xd8\xa9\x2f\x70\x3a\x31\xab\x8c\x36\x2d\xc1\xf4\xc5\xb7\x3b\xa5\xb4\x21\x74\xb1\x80\x9a\x6b\x0e\x9d\xeb\x29\xfc\xf7\xc8\x6d\x76\xec\x21\xaa\x40\x28\x42\xd7\x4e\xd4\x93\xc2\x8e\x50\xa1\xd2\x8e\x7a\xb9\x8c\xe7\x20\xb6\x0e\x5f\xc4\xac\xbf\x8c\x1e\xd6\xf5\x35\xab\x4f\xea\xc7\x87\x20\x3e\xec\x15\xd6\xa5\x42\x10\x43\x1d\x99\x9f\x68\x5b\xa9\x8e\x28\x37\x61\xdb\x31\xa5\x21\x32\x26\x96\xaa\xc0\x6c\x2b\x03\xbf\x52\xeb\x86\x94\x84\x32\x13\xcf\xe6\xa0\x13\x0f\x85\x50\x70\xb3\xd0\xc3\x32\x20\x2d\xe6\x1d\x8d\xaf\x9c\x31\xbc\xff\x09\xa2\x62\x7d\x4b\x4f\xe7\xf6\xc7\xe3\xfe\x4f\x55\x29\x2a\x79\x48\x89\x9a\x40\x79\x17\x4e\x2d\x83\x23\x8c\xcd\xe4\x2a\xcf\x97\x15\x39\x4d\xe8\x03\x38\x6b\xae\x49\xea\xe4\x27\x39\xe4\x59\x2c\x68\x49\x7e\x53\x4b\x9b\x74\xb6\x60\xb1\x83\xee\xb9\x52\x10\xec\x93\x43\xcd\xc2\x67\xe1\x49\x5a\x6a\xdb\x15\x7d\xbd\xdb\x4d\x3b\x50\x87\x09\xae\x92\xb6\x41\x86\x87\x15\x6d\xa6\xbb\x5b\x58\xd8\x0b\x05\x33\x72\x34\xbe\x3f\xff\x92\x3f\xc2\x59\xf2\x87\xff\xd1\x7e\x35\xed\x67\xfa\x57\x8f\x5f\x56\xf5\xc9\xf9\x70\xde\x9d\xca\x0f\x93\x02\xe2\x5f\x1d\x23\x11\x21\x48\x69\x75\xe1\xde\xf3\x9c\xc9\xb4\xa3\xd3\xd2\xb2\x04\x81\x1e\x1e\x6b\x29\xbc\x3a\xe7\x1e\xaf\x16\x25\xc4\x15\xbd\x58\x80\xff\x4f\x56\x8c\x3a\xdb\xbf\xa4\x14\xe5\x46\x76\x6b\xc5\x42\x7b\x69\x03\xff\xa3\xbe\x5e\x52\x5f\xb6\x2d\xab\xaa\x78\xe9\x62\xaa\xc8\x42\xb5\x5a\x60\x53\x09\x0d\x85\x8d\x51\x11\xa6\x54\x59\x2d\xdb\xae\xd4\x3c\x15\xc1\x5c\x59\xb7\x24\x4f\x3d\xff\x97\xca\x68\x45\x13\xfc\x43\xd5\xa6\xb6\xa0\x50\x99\x0f\x8f\x1d\xa3\xc5\x24\xcf\x79\x96\x3a\x52\xa5\x3d\x4b\xfd\x58\x6a\xc4\x8a\xde\x93\x93\x92\xc5\x42\x8a\x8a\xd5\x35\x9e\x68\xe2\xa1\x31\xc0\x8b\xe9\x46\x6e\x15\x01\x9a\xfa\xc4\xf5\x3c\x67\xbc\x4f\x78\x95\xb4\x62\x61\x4b\xfe\x6d\x9d\x80\x3b\xca\xaf\x6d\xfb\x2a\x41\x4f\x1d\x35\xb7\x2a\x34\xc5\xa3\x0d\x79\xb6\x99\x57\x1a\x2a\x5a\xd6\x72\x45\xc7\xd7\xfd\x54\x99\xbe\xa1\x0b\x08\xf2\x4e\x20\x95\x7c\xd5\xda\x81\x50\x50\xef\xa7\x5a\xb9\xf5\x7b\xf7\xe6\xe0\x03\x90\xe6\x6e\xa0\x4e\x01\x82\x32\x91\xdc\x10\xfb\xea\x5b\xc5\xbe\x70\xe8\xc9\x87\x7e\x50\x51\x80\x7d\xc0\x59\x47\xf3\x87\x3b\x23\x59\x39\xa2\x0c\x65\xf3\xd3\x82\x4b\x1d\xea\xaf\x54\x4c\x52\x67\x82\xdd\x59\x83\x62\x27\x98\xd0\x90\x47\x4f\x92\x93\x2b\xe1\x98\x24\x6c\xe9\x91\xa0\x43\x32\x98\x3a\x86\x33\xa2\x50\xf5\x62\xb4\x87\xff\x46\xb2\x5a\xad\x2a\x86\xf7\x63\xb3\x82\xcc\x5c\x01\xbf\x91\x9f\x05\x99\x14\xf2\xa6\x94\xe0\x5e\xfe\xef\x1c\x93\x48\xc5\xd5\x8d\x10\xc0\xb2\xb0\x12\xa9\x01\x1a\x33\x61\xf8\x7a\x2f\xf4\x50\x0a\x56\x1c\xd3\x2f\xe0\x89\xbe\x96\xf0\x87\xbb\xc8\x3a\x4c\x26\x30\xff\xb6\x2c\xa2\xa7\x8c\xbf\x30\x6f\xc8\x32\xf8\x3b\x92\x92\x15\xcd\x68\x0a\x2b\x12\xb3\xf5\x26\x22\x32\x73\x5c\xb4\xa7\x1e\x76\x06\x54\x62\xaf\xda\xe6\xe1\x27\xb1\x99\x25\xaf\x73\x12\xd6\x2a\xe6\xdd\x7a\x79\x6b\x42\xc9\x38\x8f\x09\x0b\x2b\x05\xc4\xb6\xcd\x69\x06\xc5\x37\x4f\x4b\x16\x51\x20\x61\x88\x69\xc4\xb3\x6f\x20\x99\x43\x4a\xe2\x30\x59\xc5\x94\x63\xac\x22\x33\x4f\x6a\x78\x5e\x6c\xaf\xba\x33\xf2\x0c\x19\x89\xd8\x22\x2e\x6b\xf1\xd5\x42\xda\x20\x9e\x91\xc5\x82\xa6\xca\x6d\xc9\x9b\x30\x04\xba\x7e\x4d\x66\xbc\xa7\x2b\xf9\x12\x0f\x46\x7b\x4c\x69\x1f\x77\xf5\x6a\x74\x6a\x35\x52\x9f\x57\x1f\x65\x59\xe7\xe7\x29\x5d\x04\x11\xd1\xbb\x62\x0c\x8c\xbf\x81\xce\x59\xef\xf4\xdb\x4e\x47\xa2\xac\x63\xbd\x39\xed\x9d\x9e\x59\xf6\x69\xef\xf4\xf4\x2f\x96\x65\xd5\x1b\xbf\x74\xbe\x3a\x3c\xd6\xe5\xbb\x5b\x81\x2a\x1d\x9b\x75\x1e\x50\x3d\x9c\x9a\xdb\x76\x60\x2f\x23\x18\xcd\x8c\x0d\xbd\x8c\xe6\x03\xd5\x8e\x58\x73\x03\x3a\xd8\xdf\x7a\x35\xca\xeb\xaf\x26\xce\xb4\xc8\x18\x63\x2d\xd6\x95\x73\x25\x9d\x43\xd3\xd8\x7f\x96\x78\x54\x81\xb3\x76\x7a\x09\x5a\xfb\x8c\xd4\x5a\xcd\x78\xae\x94\xe1\xa5\x02\xda\x93\x1d\x07\x4f\xc7\xd1\x55\x41\x5d\x21\x6b\x5d\xd4\x8f\x22\xb4\x2e\xef\x1b\x4e\x5f\x3a\xcf\xd3\xab\xfa\x64\xf9\xaa\x2c\x58\x9d\x45\x54\xfc\x53\x68\x8c\xb7\xd2\x15\x7e\x8b\x05\xaf\xd8\x97\xcf\x44\xcc\xb0\xa0\x3c\x93\x5d\x56\x7a\x12\x2d\xdd\xc4\xb9\xeb\xbc\x59\x87\x24\xa3\x42\x33\x60\x39\x02\x9e\x04\x9b\xef\x7a\x0d\x54\x3f\x48\x58\x77\x62\xcf\x38\x2e\x54\x94\xcb\xb9\xaf\xf1\x24\x54\xf0\xe4\x8e\x16\xe8\x4b\x98\x93\x88\x53\x8d\x41\xb0\xc0\xa8\x30\x26\x2c\x6c\xd6\x32\xfb\x4e\xf4\x0f\x03\xdc\xfa\x9a\xf2\xd0\xc8\xcf\x5f\x44\x57\xa5\xf4\x70\xae\x7e\x89\x79\x5f\x4f\x34\xb1\x23\x93\x66\x97\x9f\x45\xb2\xaf\x46\x98\x7d\x27\x62\x7b\xba\x3f\x8f\xa6\xde\x2b\xaa\x51\x6a\xbd\xd3\xbb\xc9\x67\xf4\x51\x9b\x9d\x4a\xa3\xfe\xd0\x99\x0c\x9c\xce\xaa\x57\x9d\xaf\xd6\x1f\xb1\xbf\x71\xfb\x25\x1d\x6e\x74\x10\x7c\xbe\x90\xee\x41\x84\x29\xa6\x3b\x0f\x3b\x5f\xd1\x97\xde\xe4\x8a\x96\xbd\xe2\xaf\x73\x18\x6a\x4b\xed\xbc\xbb\xe1\x2b\x38\x0d\x0d\x17\x1a\x54\x1f\x7d\x09\xc7\xe1\xeb\xda\xe6\x17\x05\x41\x3a\xe2\x47\xe2\xf5\xff\x9b\x8d\xde\x2b\x45\x87\x5a\xe9\x1a\x8e\x2f\x1b\x51\xff\xb5\xcc\xf5\x7e\x4d\xf0\x3b\x19\xd5\x23\xf4\xf2\x2b\xcd\x6a\x03\x96\xb1\xe4\xe1\xeb\x19\xd4\xa3\xcd\x99\x6d\x87\x69\xb2\xce\x03\x51\x2c\xe9\x90\x53\xf2\xfc\xc2\x28\x12\x87\x10\xd2\x88\xaa\x4a\x3f\xb2\x5e\xa7\xc9\x3a\x65\xc8\xe9\x98\x4f\x38\xa6\x85\x4a\x2c\x66\x38\x35\xbc\x41\x09\x24\x51\x48\x53\x3f\x5b\x92\x58\xbf\x22\xc3\xac\xf2\xc9\x29\x02\x8d\x77\x72\x40\x73\x83\x14\xe0\x95\x11\x34\x90\x31\xb1\x3e\xb9\x7c\x57\x2e\x2c\x81\xab\x8f\xc0\x58\x3a\x64\x2b\x1a\x8b\xa8\x5b\xdd\xca\x21\x5f\xe5\x29\x72\x75\x04\xa0\xb7\x63\x35\x77\x1d\x49\x1b\x20\xeb\x75\x75\x60\x6b\xf6\xf9\xe8\x90\xdb\x14\x20\x0d\x9d\xdf\xea\xf7\x2e\xc8\xfb\x6f\x4a\x50\x4a\xcc\xa8\xef\xcb\xda\x2e\x44\x57\xb1\x6d\x50\x45\x5e\xf5\x37\xfa\xb2\xa1\x51\x49\xaf\xf6\x59\xc3\x5f\xb9\x59\x5f\xbb\xf2\xc2\x57\x37\xa0\xf5\xca\xab\x43\x60\xd9\xaa\x9c\x5a\x34\x7e\x50\x02\x89\xb7\x74\x75\x42\x6d\x0a\xd9\xa3\xbb\xec\xe5\x5d\xa9\x52\x04\x97\x3d\xd9\x65\x95\xd7\x6c\xea\x09\x11\xac\x88\x80\xa5\xd6\x96\x59\x9e\x0d\x96\xb4\x2a\x0e\xf8\xc4\x96\xa1\x3f\x19\xe8\x3e\x88\x81\x4b\x02\x19\x5b\x2c\x33\x9d\x24\x9d\xa2\x91\xca\x6a\x32\x4d\x0f\x71\xf2\x84\x77\x00\xc8\x49\xe8\x33\x09\x32\x08\x36\x99\x9d\xcc\xe7\xc5\xad\x1e\x2c\x5e\x94\x97\x76\x08\x11\x5b\x2b\x3b\xa5\x48\x61\x60\x2a\xaf\x21\xef\x65\x89\x7c\x9e\x91\xd5\xba\x93\x92\x78\x41\x7d\x1a\x87\x5a\x3f\x5b\x09\xe5\x0b\x54\x92\xc2\x12\x1c\x44\x20\xe9\x84\x07\x49\xcc\xb3\x94\xb0\x38\x83\x20\x40\x42\x05\xf2\xa0\x29\x08\xd4\x08\x56\x40\x72\x20\xc1\x7d\x1e\xb1\x80\x42\xc8\x25\xdd\x79\x31\x5f\x65\x44\x31\xb3\x6d\x17\x9b\x16\x06\x9e\x3e\x07\xd1\x06\x8b\x40\x31\xfd\x26\x6b\xcd\xe8\x23\x4d\x65\x27\x30\x7c\x6f\x50\xed\x69\xc9\x82\xa5\x18\x81\x0d\x79\xc5\xb7\x3a\x63\x85\xbc\x67\x68\x8a\xcb\x06\xed\x21\xd8\x2b\xe4\xbd\x12\x90\xef\x2f\x77\x53\x6b\x13\xb3\x67\x7f\xc5\x82\x34\x91\x6d\x75\xbc\x53\x42\x64\x99\x9c\x58\x4e\x78\xe5\x34\xf2\xa3\x7b\xad\x6f\xa7\xb1\x2d\x4c\xf5\x35\x61\xdc\xde\xd0\xea\x66\xdb\xc1\x92\x60\x7b\x3f\x49\x69\x79\xa2\x27\x94\x8a\xea\x65\xc2\x6b\xca\x84\x75\x59\x27\x82\xd0\xc8\xa0\x4b\xf2\xa8\x2a\xc6\x13\x9e\x01\x67\x2b\x16\x91\xb4\xc8\xa8\xe6\x37\x63\x3c\x89\xd9\x18\xcf\x79\x19\xbb\xb0\x65\x1d\xe7\x9c\x45\x99\x2c\x2c\x22\x51\x94\x1f\x1f\xe2\xe2\x38\xf3\x8c\xd2\xd8\x90\x00\xdb\x9e\x6d\xb2\xa2\x0e\x31\x6e\xcb\x76\x48\xf1\xa7\x9c\x4f\x82\x1b\x63\xe5\x5f\x8c\x8d\x95\x45\x5f\xe5\xd6\xf8\x82\xa2\x27\xc8\x69\xa6\x0e\xba\x8c\xbe\x49\x7c\x76\xf2\xdb\x86\xa6\xdb\x93\x02\x7f\x78\xbc\xb0\x4e\xd0\x05\x20\x51\xb4\xf5\xd1\xf8\x29\x90\x8d\xc2\x17\x5d\x6b\x32\x9e\xb1\x38\xc8\x2a\x07\x73\xf9\xaf\x66\x16\xbe\x39\x3b\x71\x8d\x11\x92\xf7\x50\x2d\x7f\x0f\xdf\xfc\xe9\x64\x68\xbc\x75\x3e\x0d\x9c\xbb\xe9\xd7\x5e\xf8\xdd\x25\xae\x8c\xdc\x9d\x43\xf2\x9d\x06\x89\xd5\x85\x20\x89\xe7\x2c\x5d\xd1\xf0\x20\xac\xec\x81\x69\x07\x82\x1b\x40\xd3\x6e\xb4\x6c\x28\x8d\x50\x2b\x9d\xd5\xdf\xe0\x32\xb5\xbd\x03\xf2\x83\x72\xc3\xea\x1f\x29\x15\x50\x0e\xe9\x95\x27\xcb\x97\xbb\x80\xd6\xc6\x14\xa8\x13\xb8\xfc\xae\x42\x45\xfc\xc9\x68\x47\xaa\x5e\xb2\x5e\x0b\x59\xff\x56\x96\xb3\x47\xec\x81\x46\x58\x75\x87\x0d\x8c\x3c\x59\x51\xa9\xc2\x78\x46\x52\xac\xaf\x23\x19\x50\x92\x46\x0c\xbb\x9d\xd8\x8a\xd6\x67\x2f\x34\x09\x02\x91\xdb\x34\xe3\x97\x07\xd9\xfa\x33\x4b\xa7\xb1\xf4\x1a\xc3\x1d\xc4\xbd\x72\x86\x8e\x90\x20\xe1\x72\xee\x3e\x72\xd6\xb1\x69\x86\x20\x25\xae\xe4\x31\x54\x13\x43\xe9\x85\x2b\xfa\xc1\x79\xb7\x5a\x61\x58\xbb\x2d\x43\xd6\xe4\xa8\x3f\xae\xdc\xc9\xd4\xf5\x06\x53\xa8\x54\x54\x10\x5e\x2d\xaa\x50\xdc\x62\xee\xdc\xd2\xd5\x83\xd9\x28\xad\x3b\xbb\x5d\xcd\x03\xb3\xa4\x05\x2e\x7c\xca\x42\xe7\x16\x71\xa7\x88\x59\x81\xd3\x35\x49\x85\x27\x8e\x73\xa3\x1e\xc3\x73\x24\x3c\xd5\x2c\x0b\xa1\xa1\xf8\xea\x0f\x9c\xd2\x3f\xa8\xa9\xb4\xfa\x8b\x34\x79\xe2\x39\xd0\x40\x66\xc9\x23\xde\xb2\xa0\x1e\xf4\x0c\x8d\xa7\x6b\x39\xd4\x70\x15\xc4\xab\xd3\xca\x5d\x92\x5c\xc3\x57\x81\x33\x85\xdb\x93\xb3\x12\xaf\xbc\x53\xd6\x9d\xd4\x85\xeb\x8b\xc8\x73\xe5\xfe\x56\xc5\x54\xfb\xa5\xda\x18\xd4\x53\x1b\xfe\xe3\x1f\x25\xcf\xfc\x2c\xff\xee\xe5\x90\xff\x72\xb4\xe0\x14\xff\x52\x12\xb2\xbf\xbd\x00\x76\x88\xc7\x9b\x46\xb1\x28\x2e\xa5\xd5\x38\x52\x5d\x5a\x5b\xe1\xb5\xfc\xb6\x37\xfc\x4c\x85\x6a\xa5\x1b\x7c\xf9\xce\xe4\x62\xcd\x85\xbe\x7c\x67\xba\xd0\x3a\x8b\x5f\xbe\xd3\x3c\x16\xfd\xb2\x37\x19\xac\x1e\x74\xdb\xdb\x81\x0d\xcf\x2a\x8a\x15\x4b\x65\xbe\xe0\x7f\xbf\xdc\x49\xa7\x8c\xec\x65\xb3\x65\xf3\x61\x56\x99\xb5\x03\xfd\xa6\x96\x37\xa6\xac\xef\x4b\xa2\x96\xc4\x69\x64\xd4\x9c\x47\xeb\xa5\x7c\x7c\x99\x3c\xe5\x58\x2f\x03\x98\xcb\x77\xf9\xa1\xed\x37\xae\xbc\x57\xaf\x82\xe9\x55\xaf\x76\xb7\x6b\xe3\x4f\xa7\x88\x37\xfa\xd8\xb1\xc0\x3e\x3c\x7b\x6d\x26\x6b\xf4\xea\x6f\x79\x7c\xab\x6a\xbf\xd1\x31\xd6\xfb\x6d\x32\x92\x3e\x12\xe3\x26\x00\xdd\x5d\xc5\x73\xdf\xdd\x09\xdb\xb2\xeb\x6c\x9d\x26\x01\x0d\xd1\x47\x4b\xb4\xf6\xf9\xd9\x16\x82\x34\x89\x73\x2e\xa9\x5d\xd5\x87\xfb\xd2\x99\xd9\xaa\xf1\x98\x22\xb8\x9e\x44\x48\x61\xec\x0c\x46\xe3\x2b\xf3\x72\x95\x30\x81\x24\xa6\x78\x0d\x10\x3c\xb1\x6c\x09\xfc\x81\xad\xf1\x62\x0b\xa1\x3e\x73\x67\x52\x0c\x41\x4f\x73\x26\x0b\xde\x77\xe3\xe2\x7a\x34\x86\x14\x5c\xaf\xca\x6b\xfb\x39\xed\x00\x2e\x47\x6d\x93\xd7\xab\x41\x79\x21\x89\x82\x83\x97\xf7\x81\x64\xba\xa0\x43\x12\x03\xc7\xee\xaa\x22\xc6\x05\x4d\x25\xbc\x52\x04\x56\x32\x22\x4b\x7b\xba\xb5\x1c\x8d\xc1\x1b\x61\x37\x5b\x9e\x4d\xfb\xc1\xbd\x83\xe1\x68\xf0\x83\x93\x57\xad\x89\xdf\x60\xe4\x4d\x5d\xef\xde\x91\xc5\x65\xc5\xfd\x0d\xda\x88\x1c\xb8\x97\x13\x4e\x69\xcf\x38\xd1\x3b\x96\xff\xd3\x6a\xb2\xb2\x84\xf1\xf6\xd6\x9d\x96\x91\x92\xac\x7b\xfb\x9d\x09\xfc\xaf\x8a\x06\x47\x10\xeb\xe4\xe4\xa2\xf5\x7f\x01\x00\x00\xff\xff\xe3\x9e\xc6\x09\x53\x5f\x00\x00"),
		},
		"/2_post-0.1.0-alpha.1.up.sql": &vfsgen۰CompressedFileInfo{
			name:             "2_post-0.1.0-alpha.1.up.sql",
			modTime:          time.Time{},
			uncompressedSize: 403,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\xcf\xb1\x4e\xc3\x30\x14\x85\xe1\xdd\x4f\x71\x86\x4a\xa5\x0b\x02\x16\x2a\x21\x86\x5b\xeb\xb6\x54\xba\xb1\x83\xed\x6c\x48\x91\x69\x1d\x29\xc2\x75\x50\xd2\x20\x1e\x9f\x01\x86\x8e\x30\x70\xf6\xf3\x4b\x9f\x76\x4c\x81\x61\x1d\x1c\xd7\x42\x9a\xb1\x6d\x8c\x0e\x7b\x6b\xe0\xf5\x13\x57\xd4\xd6\xce\x56\xd7\xfd\xd4\x4e\xe7\x98\x53\x7b\x8a\xe3\x5b\x1a\xaf\x3e\x62\x9e\x13\x8e\xc3\xfc\x9a\x13\xde\xc7\x74\xe8\xa7\x7e\x28\x2b\xe5\x38\x34\xce\x78\x6c\xac\x15\x26\xa3\xc8\x63\xd1\xcd\xe5\xb0\x50\x00\xe0\x59\x58\x07\x74\x79\x88\xe7\xf5\x94\xca\xf1\xbb\xb3\xc2\x23\x96\x2f\x9f\xf7\x5d\x77\x73\xb1\xbb\xa5\xfa\xb9\x0a\x99\x5d\x43\x3b\x86\x7f\x16\xec\xab\xaa\x09\xb4\x11\x46\x4d\x8e\x44\x58\xe0\x69\xcb\x0f\x4a\xfd\x9e\x52\x86\xf1\x14\x73\x5b\x62\xf9\x2f\xc8\xfa\x02\x72\xfb\x47\xc8\x57\x00\x00\x00\xff\xff\x13\x5a\xdf\xb2\x93\x01\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/1_base_schema.down.sql"].(os.FileInfo),
		fs["/1_base_schema.up.sql"].(os.FileInfo),
		fs["/2_post-0.1.0-alpha.1.up.sql"].(os.FileInfo),
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
