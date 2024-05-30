package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"time"
)

// reportCmd represents the report command
var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Generate a report of transactions",
	Run: func(cmd *cobra.Command, args []string) {
		month, _ := cmd.Flags().GetInt("month")
		year, _ := cmd.Flags().GetInt("year")
		generateReport(month, year)
	},
}

func generateReport(month, year int) {
    home, _ := os.UserHomeDir()
	dataFile := filepath.Join(home, "finance-cli", "data", "transactions.json")

	var transactions []Transaction
	if _, err := os.Stat(dataFile); err!=nil {
		panic(err)
	}
	data, _ := os.ReadFile(dataFile)
	json.Unmarshal(data, &transactions)

	var totalIncome, totalExpense float64
	for _, t := range transactions {
		if t.Date.Month() == time.Month(month) && t.Date.Year() == year {
            if t.Type == "expense" {
				totalExpense += t.Amount
			} else {
				totalIncome += t.Amount
			}
		}
	}

	fmt.Printf("Report for %d-%02d\n", year, month)
	fmt.Printf("Total income: $%.2f\n", totalIncome)
	fmt.Printf("Total expenses: $%.2f\n", totalExpense)
	fmt.Printf("Net Change $%.2f\n", totalIncome - totalExpense)
} 

func init() {
	rootCmd.AddCommand(reportCmd)

	reportCmd.Flags().IntP("month", "m", int(time.Now().Month()), "Month for the report")
    reportCmd.Flags().IntP("year", "y", int(time.Now().Year()), "Year for the report")
}
