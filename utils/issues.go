package utils

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Issue struct {
	Name        string
	Description string
}

func GetIssue(csvFilePath string) []Issue {
	file, err := os.Open(csvFilePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return nil
	}

	var issues []Issue

	for _, record := range records {
		issue := Issue{
			Name:        record[0],
			Description: record[1],
		}

		issues = append(issues, issue)
	}

	return issues
}
