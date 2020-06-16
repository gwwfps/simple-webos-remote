package config

import (
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"syscall"
)

type Config struct {
	TvAddr    string
	TvMac     string
	ClientKey string
	TvInputs  []TvInput
}

type TvInput struct {
	Id    string
	Alias string
}

var configPath string
var inputs = []string{"HDMI_1", "HDMI_2", "HDMI_3", "HDMI_4"}

func init() {
	configPath = filepath.Join(os.Getenv("HOME"), ".config/simple-webos-remote.json")
	viper.SetConfigFile(configPath)

	viper.SetDefault("TvAddr", "LGsmartTV.lan")
}

func Read() (*Config, error) {
	err := viper.ReadInConfig()

	if e, ok := err.(*os.PathError); ok && e.Err == syscall.ENOENT {
		log.Info().Msg("creating config file")
		err = ioutil.WriteFile(configPath, []byte("{}"), 0644)
		if err != nil {
			return nil, errors.WithMessage(err, "cannot create config file")
		}
		err = viper.ReadInConfig()
	}

	if err != nil {
		return nil, errors.WithMessage(err, "cannot read config file")
	}

	c := &Config{}
	err = viper.Unmarshal(c)
	if err != nil {
		return nil, errors.WithMessage(err, "cannot unmarshal config file")
	}

	c.tidy()

	return c, nil
}

func (c *Config) Save() {
	typ := reflect.TypeOf(*c)
	val := reflect.ValueOf(*c)

	for i := 0; i < typ.NumField(); i++ {
		viper.Set(typ.Field(i).Name, val.Field(i).Interface())
	}

	err := viper.WriteConfig()
	if err != nil {
		log.Error().Err(err).Msg("cannot save config file")
	}
}

func (c *Config) tidy() {
	var validInputs []TvInput
	for i, inputId := range inputs {
		for _, input := range c.TvInputs {
			if input.Id == inputId {
				validInputs = append(validInputs, input)
				break
			}
		}
		if len(validInputs) < i+1 {
			validInputs = append(validInputs, TvInput{
				Id:    inputId,
				Alias: inputId,
			})
		}
	}
	c.TvInputs = validInputs
}
