package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Hardware struct {
		MotorDegreePerStep float64 `yaml:"motorDegreePerStep" json:"motorDegreePerStep,string"`
		WaitForStep        int     `yaml:"waitForStep" json:"waitForStep,string"`
		GearRatio          float64 `yaml:"gearRatio" json:"gearRatio,string"`
		Camera             struct {
			Height     int `yaml:"height" json:"height,string"`
			Width      int `yaml:"width" json:"width,string"`
			Brightness int `yaml:"brightness" json:"brightness,string"`
			Contrast   int `yaml:"contrast" json:"contrast,string"`
			Sharpness  int `yaml:"sharpness" json:"sharpness,string"`
		} `yaml:"camera" json:"camera"`
	} `yaml:"hardware" json:"hardware"`
	Server struct {
		PhotoDirectory        string `yaml:"photoDirectory" json:"photoDirectory"`
		DistributionDirectory string `yaml:"distributionDirectory" json:"distributionDirectory"`
		Port                  string `yaml:"port" json:"port"`
	} `yaml:"server" json:"server"`
}

func (c *Config) ReadFromFile(filename string) (err error) {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(c)
	return
}

func (c *Config) SaveToFile(filename string) (err error) {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	encoder := yaml.NewEncoder(f)
	err = encoder.Encode(c)
	return
}
