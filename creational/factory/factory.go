package factory

import (
	"fmt"
)

type (
	mongoDB struct {
		database map[string]interface{}
	}

	sqlite struct {
		database map[string]interface{}
	}

	file struct {
		name    string
		content string
	}

	ntfs struct {
		files map[string]file
	}

	ext4 struct {
		files map[string]file
	}

	FileSystem interface {
		CreateFile(string)
		FindFile(string) file
	}

	Database interface {
		GetData(string) interface{}
		PutData(string, interface{})
	}
)

func (mdb mongoDB) GetData(query string) interface{} {
	if _, ok := mdb.database[query]; !ok {
		return nil
	}

	fmt.Println("MongoDB")
	return mdb.database[query]
}

func (sql sqlite) GetData(query string) interface{} {
	if _, ok := sql.database[query]; !ok {
		return nil
	}

	fmt.Println("Sqlite")
	return sql.database[query]
}

func (mdb mongoDB) PutData(query string, data interface{}) {
	fmt.Println("Writing in mongodb")
	mdb.database[query] = data
}

func (sql sqlite) PutData(query string, data interface{}) {
	fmt.Println("Writing in sqlite")
	sql.database[query] = data
}

func (ntfs ntfs) CreateFile(path string) {
	file := file{"NTFS file", path}
	ntfs.files[path] = file
	fmt.Println("NTFS")
}

func (ext ext4) CreateFile(path string) {
	file := file{"Ext file", path}
	ext.files[path] = file
	fmt.Println("Ext4")
}

func (ntfs ntfs) FindFile(path string) file {
	if _, ok := ntfs.files[path]; !ok {
		return file{}
	}
	return ntfs.files[path]
}

func (ext ext4) FindFile(path string) file {
	if _, ok := ext.files[path]; !ok {
		return file{}
	}
	return ext.files[path]
}

func DatabaseFactory(env string) interface{} {
	switch env {
	case "production":
		return mongoDB{
			database: make(map[string]interface{}),
		}
	case "development":
		return sqlite{
			database: make(map[string]interface{}),
		}
	default:
		return nil
	}
}

func FilesystemFactory(env string) interface{} {
	switch env {
	case "production":
		return ntfs{
			files: make(map[string]file),
		}
	case "development":
		return ext4{
			files: make(map[string]file),
		}
	default:
		return nil
	}
}

type Factory func(string) interface{}

func AbstractFactory(fact string) Factory {
	switch fact {
	case "database":
		return DatabaseFactory
	case "filesystem":
		return FilesystemFactory
	default:
		return nil
	}
}
