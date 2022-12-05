package storage

import "github.com/hozzer/apitemplate/types"

type MemoryStorage struct{}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{}
}

func (s *MemoryStorage) Get(key string) (*types.User, error) {
	return &types.User{
		ID:   1,
		Name: "Foo",
	}, nil
}
