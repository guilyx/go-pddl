package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Env     string `envconfig:"env" default:"dev"`
	Version string `envconfig:"project_version"`
	Test    bool   `envconfig:"test" default:"false"`
	Domain  string `envconfig:"domain" default:"./data/domain.pddl"`
	Problem string `envconfig:"problem" default:"./data/problem.pddl"`
	MaxPeek int    `envconfig:"max_peek" default:"2"`
}

func NewConfig() (*Config, error) {
	config := Config{}
	err := envconfig.Process("", &config)
	if err != nil {
		return nil, fmt.Errorf("Failed to process env: %v", err)
	}
	if len(config.Problem) < 1 {
		return nil, fmt.Errorf("Problem file isn't parse-able")
	}
	if len(config.Domain) < 1 {
		return nil, fmt.Errorf("Domain file isn't parse-able")
	}
	return &config, nil
}
