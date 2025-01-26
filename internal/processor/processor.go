package processor

import (
	"fmt"

	"github.com/ChickenWhisky/fmtDB/internal/utils"
)

// ProcessCSVFiles reads multiple CSV files, computes frequency counts, and writes the result to an output file.
func ProcessCSVFiles(filePaths []string, outputPath string) error {
	// Map to store frequency counts
	frequencyMap := make(map[string]int)

	// Iterate over each input CSV file
	for _, filePath := range filePaths {
		fmt.Printf("Processing file: %s\n", filePath)

		// Read the CSV file
		records, err := ReadCSV(filePath)
		if err != nil {
			return fmt.Errorf("error reading file %s: %w", filePath, err)
		}

		// Process each record and update the frequency map
		for _, record := range records {
			if len(record) == 0 {
				continue // Skip empty rows
			}

			// Assume the column of interest is the "Company" column (index 3)
			// Update this index if your column is different
			if len(record) > 3 {
				normalized := utils.NormalizeString(record[3]) // Normalize the company name
				frequencyMap[normalized]++
			}
		}
	}

	// Convert frequency map to a 2D slice for CSV writing
	outputData := [][]string{{"Company", "Count"}} // Header row
	for company, count := range frequencyMap {
		outputData = append(outputData, []string{company, fmt.Sprintf("%d", count)})
	}

	// Write the frequency data to the output CSV file
	if err := WriteCSV(outputPath, outputData); err != nil {
		return fmt.Errorf("error writing to output file %s: %w", outputPath, err)
	}

	fmt.Printf("Frequency counts written to: %s\n", outputPath)
	return nil
}