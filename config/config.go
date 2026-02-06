package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Lnt Lenta      `yaml:"Lenta"`
	Pr  Proxy      `yaml:"Proxy"`
	Ctg Categories `yaml:"Categories"`
	Str Store      `yaml:"Store"`
	Exp Export     `yaml:"Export"`
}

type Lenta struct {
	URL        string `yaml:"URL"`
	APIURL     string `yaml:"APIURL"`
	Timeout    int    `yaml:"Timeout"`
	MaxRetries int    `yaml:"MaxRetries"`
	RetryDelay int    `yaml:"RetryDelay"`
}

type Proxy struct {
	URL     string `yaml:"URL"`
	Type    string `yaml:"Type"`
	Enabled bool   `yaml:"Enabled"`
}

type Store struct {
	StoreID      string  `yaml:"StoreID"`
	StoreName    string  `yaml:"StoreName"`
	StoreAddress string  `yaml:"StoreAddress"`
	Latitude     float64 `yaml:"Latitude"`
	Longitude    float64 `yaml:"Longitude"`
}

type Categories struct {
	ID   string `yaml:"ID"`
	Name string `yaml:"Name"`
	Path string `yaml:"Path"`
}

type Export struct {
	JSONEnabled bool   `yaml:"JSONEnabled"`
	CSVEnabled  bool   `yaml:"CSVEnabled"`
	OutputDir   string `yaml:"OutputDir"`
}

func LoadConfig(configPath string) (*Config, error) {
	data, err := os.ReadFile(configPath)

	if err != nil {
		return nil, fmt.Errorf("Error - %w", err)
	}

	var cnf Config

	if err = yaml.Unmarshal(data, &cnf); err != nil {
		return nil, fmt.Errorf("Error - %w", err)
	}

	fmt.Println(cnf)

	return &cnf, nil
}
