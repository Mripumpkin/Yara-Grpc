package common

import (
	"github.com/bluele/gcache"
)

type DBconfig struct {
	gc gcache.Cache
}

func newDBConfig() *DBconfig {
	gConfig := new(DBconfig)
	return gConfig
}

func GetGCacheHandle() gcache.Cache {
	if GDB.gc == nil {
		panic("GDB.gc nil")
	}
	return GDB.gc
}

func InitCache() {
	gc := gcache.New(300).LRU().Build()
	GDB.gc = gc
}
