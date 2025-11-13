package cacheConfig
type Config struct {
	Driver string   `json:"driver" yaml:"driver"` // redis | memcache | memory
	Addr   string   `json:"addr" yaml:"addr"`
	Password string `json:"password" yaml:"password"`
	DB       int    `json:"db" yaml:"db"`
	Servers  []string `json:"servers" yaml:"servers"` // for Memcache
}
