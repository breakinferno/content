package content_storage

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"sync"
)

// ContentStorage 获取content的存储仓库
type ContentStorage struct {
	db *mongo.Client
}

var _cs *ContentStorage
var _csOnce sync.Once

// ContentStorageInstance 返回 ContentStorage 单例对象
func ContentStorageInstance() *ContentStorage {
	_csOnce.Do(func() {
		db, _ := initDB()
		_cs = &ContentStorage{
			db: db,
			// cache??
		}
	})
	return _cs
}

