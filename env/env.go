package env

import "go.deanishe.net/env"

// Load loads the environment variables into the struct.
func Load(v interface{}) {
	err := env.Bind(v)
	if err != nil {
		panic(err)
	}
}
