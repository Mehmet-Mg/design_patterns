package main

type Singleton struct {
}

var uniqueInstance *Singleton

func GetInstance() *Singleton {
	if uniqueInstance == nil {
		uniqueInstance = &Singleton{}
	}
	return uniqueInstance
}

func (s *Singleton) GetDescription() string {
	return "I'm a classic Singleton"
}
