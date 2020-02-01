package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	ServerListenAddress   string `json:"server_listen_address"`
	ClientSocketIOAddress string `json:"client_socket_address"`
}

//NewConfig - create streuct with default values
func NewConfig() *Config {
	mycfg := Config{
		ServerListenAddress:   ":8080",
		ClientSocketIOAddress: ":8080",
	}
	return &mycfg
}

//LoadConfigFile - load the damn file!
func LoadConfigFile(c *Config) {
	jsonFile, err := os.Open("chaos-monkey.json")
	if err != nil {
		log.Println(err)
	}
	log.Println("Successfully Opened chaos-monkey.json")

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	log.Printf("got %+v from file", string(byteValue))
	jerr := json.Unmarshal(byteValue, c)
	if jerr != nil {
		log.Printf("Could not unmarshall config file: %s", jerr)
	}
}
