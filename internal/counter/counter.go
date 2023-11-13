package counter

import (
	"expvar"
	"strconv"
	"sync"
)

var inRequests *counter
var outRequests *counter
var successRequests *counter
var failedRequests *counter
var hitRedis *counter
var missRedis *counter

type counter struct {
	cnt int
	m   *sync.RWMutex
}

func (c *counter) Add() {
	c.m.Lock()
	defer c.m.Unlock()
	c.cnt++
}

func (c *counter) String() string {
	c.m.RLock()
	defer c.m.RUnlock()
	return strconv.FormatInt(int64(c.cnt), 10)
}

func init() {
	inRequests = &counter{m: &sync.RWMutex{}}
	expvar.Publish("InRequests", inRequests)

	outRequests = &counter{m: &sync.RWMutex{}}
	expvar.Publish("OutRequests", outRequests)

	successRequests = &counter{m: &sync.RWMutex{}}
	expvar.Publish("SuccessRequests", successRequests)

	failedRequests = &counter{m: &sync.RWMutex{}}
	expvar.Publish("FailedRequests", failedRequests)

	hitRedis = &counter{m: &sync.RWMutex{}}
	expvar.Publish("HitRedis", hitRedis)

	missRedis = &counter{m: &sync.RWMutex{}}
	expvar.Publish("MissRedis", missRedis)
}
