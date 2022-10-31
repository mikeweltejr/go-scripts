package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

func readCsv(fileName string, streamingService string) []byte {
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	var titles []Title

	records, csvErr := csv.NewReader(file).ReadAll()

	if csvErr != nil {
		fmt.Println(csvErr)
		os.Exit(1)
	}

	for i, line := range records {
		if i > 0 {
			var title Title
			title.Name = line[0]
			title.Type = line[1]
			title.StreamingService = streamingService
			titles = append(titles, title)
		}
	}

	json, jsonErr := json.Marshal(titles)

	if jsonErr != nil {
		fmt.Println(jsonErr)
		os.Exit(1)
	}

	return json
}

func writeJsonToFile(fileName string, data []byte) {
	err := os.WriteFile(fileName, data, 0644)

	if err != nil {
		fmt.Println(err)
	}
}
