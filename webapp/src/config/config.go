package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// APIURL is the base URL for the API communication
	APIURL = ""
	// Port is the port on which the web application will run
	Port = 0
	// Hashkey is used to authenticate the cookie
	HashKey []byte
	// BlockKey is used to encrypt the cookie data
	BlockKey []byte
)

func LoadConfig() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		log.Fatal(err)
	}

	APIURL = os.Getenv("API_URL")
	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))
}
