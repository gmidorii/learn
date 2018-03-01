package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("vim-go")
	fmt.Fprintf(os.Stdout, "Now: %v\n", time.Now())

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "\t")
	encoder.Encode(map[string]string{
		"hello": "world",
	})

	csvWriter := csv.NewWriter(os.Stdout)
	csvWriter.Write([]string{"a", "b", "c"})
	csvWriter.Flush()
	csvWriter.WriteAll([][]string{
		[]string{"d", "e", "f"},
		[]string{"g", "h", "i"},
	})
	csvWriter.Flush()
}
