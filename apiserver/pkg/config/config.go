package config

import (
	"github.com/containers-ai/alameda/apiserver/pkg/config/datahub"
	"github.com/containers-ai/alameda/internal/pkg/database/ldap"
	"github.com/containers-ai/alameda/pkg/utils/log"
)

const (
	defaultBindAddress = ":50055"
)

type Config struct {
	BindAddress string          `mapstructure:"bindAddress"`
	Datahub     *datahub.Config `mapstructure:"datahub"`
	Ldap        *ldap.Config    `mapstructure:"ldap"`
	Log         *log.Config     `mapstructure:"log"`
}

func NewDefaultConfig() Config {
	var (
		defaultLogConfig       = log.NewDefaultConfig()
		defaultAPIServerConfig = datahub.NewDefaultConfig()
		defaultLDAPConfig      = ldap.NewDefaultConfig()
		config                 = Config{
			BindAddress: defaultBindAddress,
			Datahub:     defaultAPIServerConfig,
			Ldap:        defaultLDAPConfig,
			Log:         &defaultLogConfig,
		}
	)

	return config
}

func (c *Config) Validate() error {
	return nil
}