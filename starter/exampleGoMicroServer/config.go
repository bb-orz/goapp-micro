package exampleGoMicroServer

type Config struct {
	Name       string
	ServerHost string
	ServerPort int64
}

func DefaultConfig() *Config {
	return &Config{
		Name:       "ExampleGoMicro",
		ServerHost: "127.0.0.1",
		ServerPort: 8889,
	}
}
