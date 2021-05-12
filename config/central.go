package config

import (
	"errors"
	"path/filepath"
	"strings"
)

func (c *Config) CentralConfigReady() bool {
	return (c.Central != nil && c.Central.Enable)
}

func (c *Config) BuildCentralConfig() error {
	if c.Repository == "" {
		return errors.New("repository: not set (or env GITHUB_REPOSITORY is not set)")
	}
	if c.Central == nil {
		return errors.New("central: not set")
	}
	if c.Central.Root == "" {
		c.Central.Root = "."
	}
	if !strings.HasPrefix(c.Central.Root, "/") {
		c.Central.Root = filepath.Clean(filepath.Join(c.Root(), c.Central.Root))
	}
	if c.Central.Reports == "" {
		c.Central.Reports = defaultReportsDir
	}
	if !strings.HasPrefix(c.Central.Reports, "/") {
		c.Central.Reports = filepath.Clean(filepath.Join(c.Root(), c.Central.Reports))
	}
	if c.Central.Badges == "" {
		c.Central.Badges = defaultBadgesDir
	}
	if !strings.HasPrefix(c.Central.Badges, "/") {
		c.Central.Badges = filepath.Clean(filepath.Join(c.Root(), c.Central.Badges))
	}

	return nil
}