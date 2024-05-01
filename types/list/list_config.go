package list

import "time"

type Config struct {
	Offset int
	Limit  int
}

type DateConfig struct {
	Config
	StartDate time.Time
	EndDate   time.Time
}

func (c *Config) SetOffset(offset int) *Config {
	c.Offset = offset
	return c
}

func (c *Config) SetLimit(limit int) *Config {
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
