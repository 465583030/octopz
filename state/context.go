package state

import (
	"gopkg.in/redis.v5"
)

type Context struct {
	Redis *redis.Client
}
