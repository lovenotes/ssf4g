package accountmemcached

import (
	"ssf4g/common/memcached"
)

type AccountMemcached struct {
	_memcached *memcached.Memcached
}

func NewAccountMemcached(memcached *memcached.Memcached) *AccountMemcached {
	return &AccountMemcached{
		_memcached: memcached,
	}
}
