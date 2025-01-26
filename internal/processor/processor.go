package processor

import (
	"fmt"
	"strings"

	"github.com/ChickenWhisky/fmtDB/internal/utils"
)

// ProcessCSVFiles reads multiple CSV files, computes frequency counts, and writes the result to an output file.
func ProcessCSVFiles(filePaths []string, outputPath string) error {
	// Map to store frequency counts
	frequencyMap := make(map[string]int)
	companyMapToCollege := make(map[string][]string)
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
			if record[3] ==""||record[3] ==" "{
				continue
			}

			// Assume the column of interest is the "Company" column (index 3)
			// Update this index if your column is different
			if len(record) > 3 {
				normalized := utils.NormalizeString(record[3]) // Normalize the company name
				frequencyMap[normalized]++
				companyMapToCollege[normalized] = append(companyMapToCollege[normalized], record[0])
			}
		}
	}

	// Convert frequency map to a 2D slice for CSV writing
	outputData := [][]string{{"Company", "Count","Colleges"}} // Header row
	for company, colleges := range companyMapToCollege {
		collegeList := strings.Join(colleges, ", ") // Join college names with a comma
		outputData = append(outputData, []string{company, fmt.Sprintf("%d", frequencyMap[company]), collegeList})
	}



	// Write the frequency data to the output CSV file
	if err := WriteCSV(outputPath, outputData); err != nil {
		return fmt.Errorf("error writing to output file %s: %w", outputPath, err)
	}

	fmt.Printf("Frequency counts written to: %s\n", outputPath)
	return nil
}
