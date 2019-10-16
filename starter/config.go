package starter

import (
	"github.com/IMQS/nf/nfdb"
	serviceconfig "github.com/IMQS/serviceconfigsgo"
)

type Config struct {
	HttpPort int
	DB       nfdb.DBConfig
}

func (c *Config) Load() error {
	return serviceconfig.GetConfig("", "starter", 0, "starter.json", c)
}
