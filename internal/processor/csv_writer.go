package processor

import (
    "encoding/csv"
    "os"
)

func WriteCSV(filePath string, data [][]string) error {
    file, err := os.Create(filePath)
    if err != nil {
        return err
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    return writer.WriteAll(data)
}
