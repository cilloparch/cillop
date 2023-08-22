package env

import "go.deanishe.net/env"

func Load(v interface{}) {
	err := env.Bind(v)
	if err != nil {
		panic(err)
	}
}
