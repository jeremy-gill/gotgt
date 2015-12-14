/*
Copyright 2015 The GoStor Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package backingstore

import "github.com/gostor/gotgt/pkg/scsi"

func init() {
	scsi.RegisterBackingStore("file", new)
}

type FileBackingStore struct {
	BaseBackingStore
}

func new() (scsi.BackingStore, error) {
	return NullBackingStore{
		Name:            "file",
		DataSize:        0,
		OflagsSupported: 0,
	}, nil
}

func (bs *FileBackingStore) Open(dev *SCSILu, path string, fd *int, size *uint64) error {
	return nil
}

func (bs *FileBackingStore) Close(dev *SCSILu) error {
	return nil
}

func (bs *FileBackingStore) Init(dev *SCSILu, Opts string) error {
	return nil
}

func (bs *FileBackingStore) Exit(dev *SCSILu) error {
	return nil
}

func (bs *FileBackingStore) CommandSubmit(cmd *SCSICommand) error {
	return nil
}