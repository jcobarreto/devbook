package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

// func init() {
// 	key := make([]byte, 64)

// 	if _, err := rand.Read(key); err != nil {
// 		log.Fatal("Error generating random key: ", err)
// 	}

// 	stringBase64 := base64.StdEncoding.EncodeToString(key)
// 	fmt.Printf("Generated random key: %s\n", stringBase64)
// }

func main() {
	config.Load()
	r := router.GenerateRouter()

	fmt.Printf("Listening on port %d...\n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
