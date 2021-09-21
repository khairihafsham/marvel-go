package main

import (
	"encoding/json"
	"fmt"
	"log"
	"log/syslog"
	"marvel/service"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// The goal is to only retrieve a new batch of ids IFF the total has changed
// Assumption: Marvel only adds characters and don't remove them
//
// - check if totalPath exists.
// - if no, potentially a sign records were never processed. return true
// - if yes, read and compare total, if the same, return false
// - defaults return true
func shouldRefreshData(totalPath string, total int) bool {
	if _, err := os.Stat(totalPath); os.IsNotExist(err) {
		return true
	}

	totalContent, err := os.ReadFile(totalPath)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	totalInFile, err := strconv.Atoi(strings.TrimSuffix(string(totalContent), "\n"))

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	if totalInFile == total {
		log.Printf("Recorded total is still the same. Stop processing")

		return false
	}

	return true
}

func updateTotal(totalPath string, total int) {
	f, err := os.Create(totalPath)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf("%d", total))

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	log.Printf("Written total to %s", totalPath)
}

// - Get all character ids
// - If there is any error, log and exit
// - Otherwise write to temporary file then switch name with target file
func main() {
	logwriter, err := syslog.New(syslog.LOG_INFO, "updateAllCharacterIds")

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	log.SetOutput(logwriter)

	log.Printf("Starting")

	if len(os.Args) < 3 {
		log.Fatalf("Error: 2 file paths required")
	}

	path := strings.TrimSpace(string(os.Args[1]))
	totalPath := strings.TrimSpace(string(os.Args[2]))
	oneCharacter, err := service.GetAllCharacter(service.GetTs, 0, 1, http.DefaultClient)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	if shouldRefreshData(totalPath, oneCharacter.Data.Total) == false {
		return
	}

	log.Printf("Fetching data")

	result, err := service.GetAllCharacterId(service.GetTs, 100, http.DefaultClient)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	ids, err := json.Marshal(result)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	err = os.WriteFile(path, ids, 0644)

	log.Printf("Written data to %s", path)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	updateTotal(totalPath, oneCharacter.Data.Total)
}
