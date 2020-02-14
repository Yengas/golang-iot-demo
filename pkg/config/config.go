package config

type Config struct {
	ServerConfig struct {
		Name string `yaml:"name"`
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"server"`
	SwaggerConfig struct {
		DocumentationHost     string `yaml:"documentation_host"`
		DocumentationBasePath string `yaml:"documentation_base_path"`
	} `yaml:"swagger"`
	Auth struct {
		Secret string `yaml:"secret"`
	} `yaml:"auth"`
}
