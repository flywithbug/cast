package cast

import (
	"sync"
)

var con *Config

//自增Id
var increaseId = 100

type Config struct {
	Lock     sync.RWMutex
	Template Template
}

func Conf() *Config {
	if con == nil {
		con = confInit()
	}
	return con
}

func confInit() *Config {
	return &Config{
		Lock:     sync.RWMutex{},
		Template: NewTemplate(),
	}
}

func (c *Config) increaseId() int {
	c.Lock.Lock()
	increaseId++
	c.Lock.Unlock()
	return increaseId
}
