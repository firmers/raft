// Copyright 2017-2021 Lei Ni (nilei81@gmail.com) and other contributors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	  http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vfs

import (
	"os"
	"syscall"
	"testing"

	"github.com/cockroachdb/errors"
	"github.com/cockroachdb/pebble/vfs"
)

var Default = &fs{FS: vfs.Default}

type File vfs.File

type FS interface {
	vfs.FS
	OpenForAppend(name string) (vfs.File, error)
}

type fs struct {
	vfs.FS
}

func (that *fs) OpenForAppend(name string) (vfs.File, error) {
	f, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND|syscall.O_CLOEXEC, 0)
	if nil != err {
		return nil, err
	}
	return &fileCompat{File: f}, errors.WithStack(err)
}

type fileCompat struct {
	*os.File
}

func (f *fileCompat) Preallocate(offset, length int64) error {
	return nil
}

func (f *fileCompat) SyncTo(length int64) (fullSync bool, err error) {
	return true, f.Sync()
}

func (f *fileCompat) SyncData() error {
	return f.Sync()
}

func (f *fileCompat) Prefetch(offset int64, length int64) error {
	return nil
}

func (f *fileCompat) Fd() uintptr {
	return f.File.Fd()
}

func NewStrictMem() FS {
	return &fs{FS: vfs.NewStrictMem()}
}

func NewMem() FS {
	return &fs{FS: vfs.NewMem()}
}

func ReportLeakedFD(fs vfs.FS, t *testing.T) {
	//mf, ok := fs.(*vfs.MemFS)
	//if !ok {
	//	return
	//}
	//ff := func(path string, isDir bool, refs int32) error {
	//	if refs != 0 {
	//		t.Fatalf("%s (isDir %t) is not closed", path, isDir)
	//	}
	//	return nil
	//}
	//if err := mf.Iterate(ff); err != nil {
	//	t.Fatalf("fs.Iterate failed %v", err)
	//}
}
