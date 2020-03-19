package packrfix

import (
	"net/http"
	"os"
	"time"

	"github.com/gobuffalo/packr/v2"
)

//BoxExtended ...
type BoxExtended struct {
	packr.Box
	createTime time.Time
}

//New create new BoxExtended
func New(originalBox *packr.Box) *BoxExtended {
	box := new(BoxExtended)
	box.Box = *originalBox
	box.createTime = time.Now()

	return box
}

// Open returns a File using the http.File interface
func (be BoxExtended) Open(name string) (http.File, error) {
	theFile, err := be.Box.Open(name)
	if err != nil {
		panic(err)
	}
	theInfo, err := theFile.Stat()

	ffixed := new(VariableFileFixed)
	ffixed.File = theFile
	ffixed.info = *new(fileInfo)
	ffixed.info.Path = theInfo.Name()
	ffixed.info.size = theInfo.Size()
	ffixed.info.isDir = theInfo.IsDir()
	ffixed.info.modTime = be.createTime

	return ffixed, nil
}

type fileInfo struct {
	Path    string
	size    int64
	modTime time.Time
	isDir   bool
}

func (f fileInfo) Name() string {
	return f.Path
}

func (f fileInfo) Size() int64 {
	return f.size
}

func (f fileInfo) Mode() os.FileMode {
	return 0444
}

func (f fileInfo) ModTime() time.Time {
	return f.modTime
}

func (f fileInfo) IsDir() bool {
	return f.isDir
}

func (f fileInfo) Sys() interface{} {
	return nil
}

//VariableFileFixed ...
type VariableFileFixed struct {
	http.File
	info fileInfo
}

//Seek ...
func (f *VariableFileFixed) Seek(offset int64, whence int) (int64, error) {
	return f.Seek(offset, whence)
}

//Close ...
func (f *VariableFileFixed) Close() error {
	return nil
}

//Readdir ...
func (f VariableFileFixed) Readdir(count int) ([]os.FileInfo, error) {
	return f.Readdir(count)
}

//Stat ...
func (f VariableFileFixed) Stat() (os.FileInfo, error) {
	return f.info, nil
}

func (f VariableFileFixed) String() string {
	return string(f.String())
}
