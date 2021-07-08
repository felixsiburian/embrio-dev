// All Init func here

package config

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type Config struct {
	echo echo.Echo
}

func (c *Config) InitEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}

func CatchEror(err error) {
	if err != nil {
		panic(err)
	}
}
