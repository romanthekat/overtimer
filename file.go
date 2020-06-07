package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

const configFilename = "overtimer.json"

func getApp() (*App, error) {
	jsonFile, err := openConfigFile()
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var app = NewAppDefault()
	err = json.Unmarshal(bytes, &app)

	if err != nil {
		return nil, err
	}

	return app, nil
}

func openConfigFile() (*os.File, error) {
	jsonFile, err := os.OpenFile(configFilename, os.O_RDWR, os.ModeExclusive)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("%s does not exist, creating new file\n", configFilename)

			jsonFile, err = createNewConfigFile()
			if err != nil {
				return nil, err
			}

			return jsonFile, nil
		} else {
			return nil, err
		}
	}

	return jsonFile, nil
}

func createNewConfigFile() (*os.File, error) {
	jsonFile, err := os.Create(configFilename)
	if err != nil {
		return nil, err
	}

	_, err = jsonFile.WriteString("{}")
	jsonFile.Seek(0, io.SeekStart)

	if err != nil {
		return nil, err
	}

	return jsonFile, nil
}

func (app *App) save() error {
	jsonFile, err := openConfigFile()
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	bytes, err := json.MarshalIndent(app, "", "  ")
	if err != nil {
		return err
	}

	_, err = jsonFile.Write(bytes)
	return err
}
