package config

import (
	"log"
	"os"
	"time"
)

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-default:"local"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	TimeOut     time.Duration `yaml:"timeout" env-default:"4 s "`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

 func MustLoad() {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		 log.Fatal("Config path is not set ")
	}

	//check if file exist
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file is not exist: %s", configPath)
	}

	var cfg Config 
 } 