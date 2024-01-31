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
			Name:        removeIfSpaceIsFirst(string(record[0])),
			Description: removeIfSpaceIsFirst(string(record[1])),
		}

		rawLabels := strings.Split(record[2], ";")
		if rawLabels == nil {
			fmt.Println("Error reading labels on csv")
			return nil
		}
		for _, label := range rawLabels {
			issue.Labels = append(issue.Labels, removeIfSpaceIsFirst(label))
		}

		issues = append(issues, issue)
	}

	return issues
}

func removeIfSpaceIsFirst(s string) string {
	numberOfSpace := 0
	for index := range s {
		if !strings.HasPrefix(s[index:], " ") {
			break
		}
		numberOfSpace += 1
	}
	return s[numberOfSpace:]
}
