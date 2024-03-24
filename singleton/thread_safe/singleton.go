package main

import "sync"

type Singleton struct {
}

var (
	uniqueInstance *Singleton
	once           sync.Once
)

func GetInstance() *Singleton {
	once.Do(func() {
		uniqueInstance = &Singleton{}
	})
	return uniqueInstance
}

func (s *Singleton) GetDescription() string {
	return "I'm a classic Singleton"
}
