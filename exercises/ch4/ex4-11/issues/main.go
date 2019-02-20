// CRUD commandline app to access GitHub Issues API

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var (
		userEVK  = "GHISSUES_USER"
		tokenEVK = "GHISSUES_TOKEN"
		ownerEVK = "GHISSUES_OWNER"
		repoEVK  = "GHISSUES_REPO"
	)
	user, err := lookupEnv(userEVK)
	trip(err)
	token, err := lookupEnv(tokenEVK)
	trip(err)
	owner, _ := lookupEnv(ownerEVK)
	repo, _ := lookupEnv(repoEVK)

}

func trip(err error) {
	if err != nil {
		panic(err)
	}
}
func lookupEnv(key string) (string, error) {
	val, ok = os.LookupEnv(key)
	if !ok {
		return "", fmt.Errorf("cannot find envvar %s", key)
	}
	return val, nil
}
