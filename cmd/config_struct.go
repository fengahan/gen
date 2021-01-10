package cmd

//@ConfigStruct()
type MysqlConfig struct {
	Host string `yaml:"Host"`
}

//新建一个struct 接收yaml 的值 key的值和配置文件的名字一致

//@ConfigStruct(key="Redis2")
type RedisConfig struct {
	Host string `yaml:"Host"`
	Port string `yaml:"Port"`
}

//@ConfigStruct(key="Cache")
type CacheConfig struct {
	Host string `yaml:"Host"`
	Port string `yaml:"Port"`
}
