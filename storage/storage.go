package storage

import (
	"github.com/hozzer/apitemplate/types"
)

type Storage interface {
	Get(key string) (*types.User, error)
}
