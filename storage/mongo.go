package storage

import "github.com/hozzer/apitemplate/types"

type MongoStorage struct{}

func (s *MongoStorage) Get(key string) (*types.User, error) {
	return &types.User{
		ID:   1,
		Name: "Foo",
	}, nil
}
