package main

import (
	"SpeedLoggerz/server/models"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func LoadJsonFile(path string) ([]models.SpeedyLogFile, error) {
	collectedLogs := make([]models.SpeedyLogFile, 0)
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("Error walking the path: %v\n", err)
			return err
		}
		fmt.Printf("Processing file: %s\n", path)
		fmt.Printf("Name: %s\n", info)

		if !info.IsDir() && filepath.Ext(path) == ".json" {
			// Open and read the file
			content, err := os.ReadFile(path)
			if err != nil {
				log.Printf("Error reading file: %v\n", err)
				return err
			}
			thing := string(content)
			thingy := []string{thing}
			output, err := exec.Command(`C:\Users\User\AppData\Local\Programs\Python\Python312\Scripts\json-repair.py`, thingy...).Output()
			repaired := output
			if err != nil {
				log.Printf("Error repairing json: %v\n", err)
			}

			var logs models.SpeedyLogFile
			jsonBytes, err := json.Marshal(repaired)
			if err != nil {
				log.Print("Error re-marshalling data:", err)
			}
			err = json.Unmarshal(jsonBytes, &logs)
			// For SyntaxError or a ParseError:
			if jsonErr, ok := err.(*json.UnmarshalTypeError); ok {
				problemPart := content[jsonErr.Offset-10 : jsonErr.Offset]
				fmt.Printf("SyntaxError ~ error near '%s' (offset %d)\n", problemPart, jsonErr.Offset)
				fmt.Printf("UnmarshalTypeError Value: %v || Type: %v - %v\n", jsonErr.Value, jsonErr.Type, jsonErr.Offset)
				fmt.Printf("Logs: %v\n", logs)
			} else if terr, ok := err.(*time.ParseError); ok {
				// For time-related errors:
				fmt.Printf("ParseError ~ time is not in RFC 3339 format.\n %s", terr)

			} else {
				// For other errors:
				fmt.Printf("General error unmarshalling JSON:: %v\n", err)
			}
			if err != nil {
				log.Printf(" %v\n", err)
				return err
			}

			// Process the JSON data
			fmt.Printf("Successfully read from file: %s\n", path)

			collectedLogs = append([]models.SpeedyLogFile{}, logs)
		}
		return nil
	})
	if err != nil {
		fmt.Print("oops")
		return collectedLogs, err
	}
	return collectedLogs, err
}

func SyncLogFiles(files []models.SpeedyLogFile) (int64, error) {
	bf, err := json.Marshal(files)
	body := bytes.NewBuffer(bf)
	resp, err := http.Post(models.URL, "application/json", body)
	if err != nil {
		fmt.Printf("HTTP Error: %s", resp.StatusCode)
	}
	return 0, nil
}

func main() {
	// url value currently set to mine. Allow users to set theirs.
	logs, err := LoadJsonFile(models.PATH_TO_LOG_FILES)
	if err != nil {
		log.Print("Unable to load JSON file.")
	}
	id, err := SyncLogFiles(logs)
	if err != nil {
		log.Printf("Unable to sync logs. User ID: %v.", id)
	}
}
