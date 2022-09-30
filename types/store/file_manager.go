package store

import (
	"encoding/json"
	"os"

	"github.com/pkg/errors"
)

type FileManager struct {
	Path string
}

func NewFileManager(path string) FileManager {
	return FileManager{path}
}

// Write will use user password to encrypt data and save to file, the file name is user name
func (f FileManager) Write(name string, info KeyInfo) error {
	bytes, err := json.Marshal(info)
	if err != nil {
		return err
	}

	return os.WriteFile(f.Path+"/"+name, bytes, 0600)
}

// Read will read encrypted data from file and decrypt with user password
func (f FileManager) Read(name string) (KeyInfo, error) {
	bytes, err := os.ReadFile(f.Path + "/" + name)
	if err != nil {
		return KeyInfo{}, errors.New("key not found")
	}

	var keyInfo KeyInfo
	err = json.Unmarshal(bytes, &keyInfo)
	return keyInfo, err
}

// Delete will delete user data and use user password to verify permissions
func (f FileManager) Delete(name, password string) error {
	return os.Remove(f.Path + "/" + name)
}

// Has returns whether the specified user name exists
func (f FileManager) Has(name string) bool {
	_, err := os.Stat(f.Path + "/" + name)
	return err == nil || os.IsExist(err)
}

// 判断root文件夹是否存在，不存在则创建
func (f FileManager) CreateRootDir() error {
	_, err := os.Stat(f.Path)
	if err == nil || os.IsExist(err) {
		return nil
	}

	return os.Mkdir(f.Path, os.ModePerm)
}
