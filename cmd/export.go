package cmd

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export transactions to a CSV file",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Error: CSV file path is required")
			return
		}
		filePath := args[0]
		exportTransactions(filePath)
	},
}

func exportTransactions(filePath string) {
	home, _ := os.UserHomeDir()
	dataFile := filepath.Join(home, ".finance-cli", "transactions.json")

	var transactions []Transaction
	if _, err := os.Stat(dataFile); err == nil {
		data, _ := os.ReadFile(dataFile)
		json.Unmarshal(data, &transactions)
	}

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Type", "Category", "Amount", "Date"})
	for _, t := range transactions {
		record := []string{t.Type, t.Category, fmt.Sprintf("%.2f", t.Amount), t.Date.Format("2006-01-02")}
		writer.Write(record)
	}

	fmt.Printf("Exported %d transactions to %s\n", len(transactions), filePath)
}

func init() {
	rootCmd.AddCommand(exportCmd)
}
