//+build wireinject

package config

import "gen/internal/gen_config"

type Manger struct {
	GenSystemConfig gen_config.GenSystemConfig `yaml:"GenSystemConfig"`
}

//@ConfigEntity
func RedisConfigEntity() RedisConfig  {
	return RedisConfig{}
}

//@ConfigEntity
func CacheConfigEntity() CacheConfig  {
	return CacheConfig{}
}


func Cache2ConfigEntity() CacheConfig  {
	return CacheConfig{}
}

