package cache

import (
	"testing"
	"github.com/stretchr/testify/suite"
	"github.com/kyokan/chaind/pkg/config"
)

func TestRedisCacher(t *testing.T) {
	suite.Run(t, &CacherSuite{
		cacher: NewRedisCacher(&config.RedisConfig{
			URL: "localhost:6379",
		}),
	})
}