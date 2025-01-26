package processor

import (

	"github.com/ChickenWhisky/fmtDB/internal/utils"
)

func ComputeFrequencies(records [][]string, columnIndex int) map[string]int {
    frequencies := make(map[string]int)
    for _, record := range records {
        if len(record) <= columnIndex {
            continue
        }
        normalized := utils.NormalizeString(record[columnIndex])
        frequencies[normalized]++
    }
    return frequencies
}
