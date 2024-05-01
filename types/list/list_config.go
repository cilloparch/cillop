package list

import "time"

// Config is a struct that defines the configuration of a list
// It has an offset and a limit
// Offset is the number of items to skip
// Limit is the maximum number of items to return
type Config struct {

	// Offset is the number of items to skip
	Offset int

	// Limit is the maximum number of items to return
	Limit int
}

// DateConfig is a struct that defines the configuration of a list with dates
type DateConfig struct {
	Config

	// StartDate is the start date of the list
	StartDate time.Time

	// EndDate is the end date of the list
	EndDate time.Time
}

// SetOffset sets the offset of the list
func (c *Config) SetOffset(offset int) *Config {
	c.Offset = offset
	return c
}

// SetLimit sets the limit of the list
func (c *Config) SetLimit(limit int) *Config {
	c.Limit = limit
	return c
}

// SetStartDate sets the start date of the list
func (c *DateConfig) SetStartDate(startDate time.Time) *DateConfig {
	c.StartDate = startDate
	return c
}

// SetEndDate sets the end date of the list
func (c *DateConfig) SetEndDate(endDate time.Time) *DateConfig {
	c.EndDate = endDate
	return c
}
