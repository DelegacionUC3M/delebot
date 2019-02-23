package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

// TomlConfig hols the data of all the necesary databases
type TomlConfig struct {
	Bot BotData  `toml:"bot"`
	BD  Database `toml:"database"`
}

// Database data to initialize a connection
type Database struct {
	Name     string
	User     string
	Password string
}

// BotData contains the token of the bot
type BotData struct {
	Token string
}

// ParseConfig parses the toml file to retrieve the data from the toml file
func ParseConfig(file string) (TomlConfig, error) {

	var config TomlConfig
	_, err := toml.DecodeFile(file, &config)

	return config, err
}

// CreateDbInfo formats the data to create a connection to the database
func CreateDbInfo(config Database) string {
	connection := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", config.User, config.Password, config.Name)

	return connection
}
