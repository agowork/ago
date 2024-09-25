package config

import "github.com/agowork/utilities/env"

var ServerPort string = env.GetEnv(
	"SERVER_PORT",
	"8080",
)
