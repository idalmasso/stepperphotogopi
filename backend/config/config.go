package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Hardware struct {
		MotorDegreePerStep float64 `yaml:"motorDegreePerStep"`
		WaitForStep        int     `yaml:"waitForStep"`
		GearRatio          float64 `yaml:"gearRatio"`
	} `yaml:"hardware"`
	PhotoDirectory        string `yaml:"photoDirectory"`
	DistributionDirectory string `yaml:"distributionDirectory"`
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
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	encoder := yaml.NewEncoder(f)
	err = encoder.Encode(c)
	return
}