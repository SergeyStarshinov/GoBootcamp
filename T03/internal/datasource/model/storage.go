package datamodel

import "sync"

type Storage struct {
	Data sync.Map
}

func NewStorage() *Storage {
	return &Storage{}
}
