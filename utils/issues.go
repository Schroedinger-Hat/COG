package utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type Issue struct {
	Name        string
	Description string
	Labels      []string
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

	for _, record := range records[1:] {
		issue := Issue{
			Name:        record[0],
			Description: record[1],
		}

		issue.Labels = strings.Split(record[2], ";")
		if issue.Labels == nil {
			fmt.Println("Error reading labels on csv")
			return nil
		}

		issues = append(issues, issue)
	}

	return issues
}
