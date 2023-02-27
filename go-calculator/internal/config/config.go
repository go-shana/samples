package config

import (
	"context"
	"time"

	"github.com/go-shana/core/config"
	"github.com/go-shana/core/validator/numeric"
)

var (
	// The config for this service.
	Service = config.New[Config]("service")
)

// Config is the config for this service.
type Config struct {
	TTL time.Duration `shana:"ttl"`
}

// Init initializes the config's default.
func (c *Config) Init(ctx context.Context) {
	if c.TTL == 0 {
		c.TTL = 10 * time.Second
	}
}

// Validate validates the config.
func (c *Config) Validate(ctx context.Context) {
	// TTL must be in [0, 1h].
	numeric.InRangInclusive(c.TTL, 0, time.Hour)
}
