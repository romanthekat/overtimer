package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const configFilename = "overtimer.json"

func getApp() (*App, error) {
	jsonFile, err := os.Open(configFilename)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("%s does not exist, creating new file\n", configFilename)

			jsonFile, err = os.Create(configFilename)
			if err != nil {
				return nil, err
			}
			_, err := jsonFile.WriteString("{}")
			if err != nil {
				return nil, err
			}

			return &App{}, nil
		} else {
			return nil, err
		}
	}
	defer jsonFile.Close()

	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var app App
	err = json.Unmarshal(bytes, &app)

	if err != nil {
		return nil, err
	}

	return &app, nil
}
