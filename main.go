package main

import (
	"log"
	"os"

	"github.com/cyberark/conjur-api-go/conjurapi"
)

//ENV Vars for def in K8s:
// CONJUR_PASS_OBJECT, CONJUR_USER_OBJECT, CONJUR_AUTHN_TOKEN_FILE, CONJUR_APPLIANCE
// CONJUR_VERSION, CONJUR_ACCOUNT

func main() {

	//Defining Username & Password objects to retrieve, as per 12 factor
	//this is being accomplished via env variables.
	variableIdentifier := os.Getenv("CONJUR_PASS_OBJECT")
	variableuserIdentifier := os.Getenv("CONJUR_USER_OBJECT")

	//Loading configuration via defined Env vars:
	//CONJUR_APPLIANCE, CONJUR_VERSION, CONJUR_ACCOUNT
	config, err := conjurapi.LoadConfig()
	if err != nil {
		panic(err)
	}

	//Get Authorization token from shared store from sidecar
	conjur, err := conjurapi.NewClientFromTokenFile(config, os.Getenv("CONJUR_AUTHN_TOKEN_FILE"))

	if err != nil {
		panic(err)
	}

	//Grab Password from Conjur
	secretValue, err := conjur.RetrieveSecret(variableIdentifier)
	if err != nil {
		panic(err)
	}

	//Grab Username from Conjur
	secretValueUser, err := conjur.RetrieveSecret(variableuserIdentifier)
	if err != nil {
		panic(err)
	}

	//Display Username & Password in log.
	log.Printf("%s:%s", "The Username Used: ", secretValueUser)
	log.Printf("%s:%s", "The Password Used: ", secretValue)
}
