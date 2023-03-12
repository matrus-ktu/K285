package config

type Config struct {
	Server struct {
		Address string `yaml:"address"`
		Port    string `yaml:"port"`
		Secret  string `yaml:"session_secret"`
	}
	Database struct {
		Address  string `yaml:"address"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
	}
}
