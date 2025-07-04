package config

type Config struct {
	Server struct {
		Port int `yml:"port"`
	} `yml:"port"`

	Database struct {
		Port     int    `yml:"port"`
		Host     string `yml:"host"`
		User     string `yml:"user"`
		DBName   string `yml:"dbname"`
		SSlModel string `yml:"sslmode"`
	} `yml:"database"`
}
