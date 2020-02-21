package config

type Selection struct {
	StaticConfigurationFilePath string
	DynamicConfigurationFilePath string
	Profile string
}

type Static struct {
	Server struct {
		Name string
		Host string
		Port int
	}
	Swagger struct {
		DocumentationHost     string
		DocumentationBasePath string
	}
	Auth struct {
		Secret string
	}
	IsRelease bool
}

type Dynamic struct {
	Threshold struct {
		Min float64
		Max float64
	}
}
