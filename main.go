package main

import (
	"errors"

	"fmt"
	"os"

	"github.com/ernesto27/firestore-backup/backup"
	"github.com/ernesto27/firestore-backup/firestore"
)

func main() {
	args := os.Args[1:]

	errParams := validateParams(args)
	if errParams != nil {
		panic(errParams)
	}

	firestore, errF := firestore.New(args[0])
	if errF != nil {
		panic(errF)
	}

	collections := args[1:]

	for _, collection := range collections {
		data, errData := firestore.GetData(collection)
		if errData != nil {
			panic(errData)
		}
		jsonData, err := backup.GetJSONFromData(data)
		if err != nil {
			panic(err)
		}
		errSave := backup.SaveJSONToFile(collection+".json", jsonData)
		if errSave != nil {
			fmt.Println(errSave)
		}
	}
}

func validateParams(params []string) error {
	if len(params) < 2 {
		return errors.New("Invalid params")
	}
	return nil
}
