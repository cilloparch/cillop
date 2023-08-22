package list

import "time"

type Config struct {
	Offset int64
	Limit  int64
}

type DateConfig struct {
	Config
	StartDate time.Time
	EndDate   time.Time
}

func (c *Config) SetOffset(offset int64) *Config {
	c.Offset = offset
	return c
}

func (c *Config) SetLimit(limit int64) *Config {
	c.Limit = limit
	return c
}

func (c *DateConfig) SetStartDate(startDate time.Time) *DateConfig {
	c.StartDate = startDate
	return c
}

func (c *DateConfig) SetEndDate(endDate time.Time) *DateConfig {
	c.EndDate = endDate
	return c
}