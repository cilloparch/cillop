package list

type list struct{}

func New() *list {
	return &list{}
}

func (l *list) Config() *Config {
	return &Config{}
}

func (l *list) DateConfig() *DateConfig {
	return &DateConfig{}
}