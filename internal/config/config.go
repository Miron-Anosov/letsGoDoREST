package config


type Config struct {
	Env string `yaml:"env" env-default:"ENV envDefault:"development"`
}