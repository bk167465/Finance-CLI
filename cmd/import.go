package cmd

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import transactions from a CSV file",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args)<1 {
			fmt.Println("Error: CSV file path is required")
			return
		}
		filePath := args[0]
		fmt.Println(filePath)
		importTransactions(filePath)
	},
}

func importTransactions(filePath string){
	file, err := os.Open(filePath)
    if err!=nil {
		fmt.Println("Error opening file1")
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err!=nil {
		fmt.Println("Error opening file")
		return
	}

	fmt.Println(records)
	fmt.Println(records[0])
	fmt.Println(records[1])

	var transactions []Transaction
	for _, record := range records[1:] {
        amount, _ := strconv.ParseFloat(record[2], 64)
        date, _ := time.Parse("2006-01-02", record[3])
        newTransaction := Transaction{record[0], record[1], amount, date}
		transactions = append(transactions, newTransaction)
	}
	saveTransactions(transactions)
	fmt.Printf("Imported %d transactions from %s\n", len(transactions), filePath)
} 

func saveTransactions(transactions []Transaction){
    home, _ := os.UserHomeDir()
	dataFile := filepath.Join(home, "finance-cli", "data", "transactions.json")

	var existingTransactions []Transaction
	if _, err := os.Stat(dataFile); err!=nil {
		panic(err)
	}
	data, _ := os.ReadFile(dataFile)
	json.Unmarshal(data, &existingTransactions)

	transactions = append(existingTransactions, transactions...)
	
	data, _ = json.MarshalIndent(transactions, "", "")
    os.MkdirAll(filepath.Dir(dataFile), os.ModePerm)
	os.WriteFile(dataFile, data, 0644)
}

func init() {
	rootCmd.AddCommand(importCmd)
}
