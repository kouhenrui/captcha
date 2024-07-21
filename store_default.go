package main

import (
	"github.com/patrickmn/go-cache"
	"time"
)

type StoreDefalut struct {
	t     time.Duration
	catch *cache.Cache
}

func NewStoreDefalut(selftime, defaulttime, exptime time.Duration) *StoreDefalut {
	return &StoreDefalut{t: selftime, catch: cache.New(defaulttime, exptime)}
}
func (sd *StoreDefalut) Get(key string, clear bool) string {
	value, found := sd.catch.Get(key)
	if found {
		return value.(string)
	}
	return ""
}

func (sd *StoreDefalut) Set(key, value string) error {
	sd.catch.Set(key, value, sd.t)
	return nil
}

func (sd *StoreDefalut) Verify(id, answer string, clear bool) bool {
	if id == "" || answer == "" {
		return false
	}
	v := sd.Get(id, clear)
	return v != "" && v == answer
}
