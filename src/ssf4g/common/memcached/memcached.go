package memcached

import (
	"ssf4g/libs/compress"

	"github.com/Terry-Mao/gomemcache/memcache"
)

type Memcached struct {
	_pool *memcache.Pool
}

func NewMemcached(server string, maxidle int) (*Memcached, error) {
	memcached := &Memcached{}

	// 测试Server
	_, err := memcache.Dial("tcp", server)

	if err != nil {
		return nil, err
	}

	memcached._pool = memcache.NewPool(func() (memcache.Conn, error) {
		conn, err := memcache.Dial("tcp", server)
		if err != nil {
			return nil, err
		}
		return conn, nil
	}, maxidle)

	return memcached, nil
}

func (memcached *Memcached) Set(key string, value []byte, timeout int32) error {
	conn := memcached._pool.Get()
	defer conn.Close()

	if value != nil {
		memValue, err := compress.DataToSnappy(value)

		if err != nil {
			return err
		}

		if err := conn.Store("set", key, memValue, 0, timeout, 0); err != nil {
			return err
		}
	}

	return nil
}

func (memcached *Memcached) Get(key string) ([]byte, error) {
	conn := memcached._pool.Get()
	defer conn.Close()

	var memValue []byte = nil

	if err := conn.Get("get", func(r *memcache.Reply) {
		if r.Key == key && r.Flags == 0 {
			memValue = r.Value
		}
	}, key); err != nil {
		return nil, err
	}

	if memValue != nil {
		value, err := compress.SnappyToData(memValue)

		if err != nil {
			return nil, err
		}

		return value, nil
	}

	return nil, nil
}

func (memcached *Memcached) GetMulti(keys []string) (map[string][]byte, error) {
	conn := memcached._pool.Get()
	defer conn.Close()

	memValues := make(map[string][]byte)
	values := make(map[string][]byte)

	if err := conn.Get("get", func(r *memcache.Reply) {
		if r.Flags == 0 {
			memValues[r.Key] = r.Value
		}
	}, keys...); err != nil {
		return nil, err
	}

	for key, memValue := range memValues {
		if memValue != nil {
			value, err := compress.SnappyToData(memValue)

			if err != nil {
				continue
			}

			values[key] = value
		}
	}

	return values, nil
}

func (memcached *Memcached) Delete(key string) error {
	conn := memcached._pool.Get()
	defer conn.Close()

	if err := conn.Delete(key); err != nil {
		if err != memcache.ErrNotFound {
			return err
		}

		return nil
	}

	return nil
}
