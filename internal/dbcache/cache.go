package dbcache

import (
	"go.uber.org/zap"
	"time"
)

type ThroughDB struct {
	DBCache DBCache
}

func NewThroughDB(conn Connector, logger *zap.SugaredLogger) *ThroughDB {
	return &ThroughDB{DBCache: DBCache{Repo{conn}, logger}}
}

// Get reads a value through the proxy and set the cache
func (rt *ThroughDB) Get(key string, req func() (interface{}, error), expire time.Duration) (interface{}, error) {
	// Get Get Cache from Proxy
	ok, val, err := rt.DBCache.Get(key)

	// return the cache if found
	if ok {
		return val, err
	}

	// Get from origin
	val, err = req()
	if err != nil {
		return val, err
	}

	// Set the value from origin to the proxy cache
	err = rt.DBCache.Set(key, val, expire)
	if err != nil {
		return nil, err
	}
	// return the value got from origin
	return val, err
}
