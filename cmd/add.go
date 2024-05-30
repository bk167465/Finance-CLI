package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"time"
)

type Transaction struct {
	Type     string    `json:"type"`
	Category string    `json:"category"`
	Amount   float64   `json:"amount"`
	Date     time.Time `json:"date"`
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new transaction",
	Run: func(cmd *cobra.Command, args []string) {
	    tType, _ := cmd.Flags().GetString("type")
		category, _ := cmd.Flags().GetString("category")
		amount, _ := cmd.Flags().GetFloat64("amount")
		date := time.Now()

		transaction := Transaction{tType, category, amount, date}
		saveTransaction(transaction)

		fmt.Printf("Added %s: %s of %.2f on %s\n", tType, category, amount, date.Format("2006-01-02"))
	},
}

func saveTransaction(newTransaction Transaction){
	home, _ := os.UserHomeDir()
	dataFile := filepath.Join(home, "finance-cli", "data", "transactions.json")

	fmt.Println(dataFile)

	var transactions []Transaction
	if _, err := os.Stat(dataFile); err!=nil {
		panic(err)
	}
	data, _ := os.ReadFile(dataFile)
	json.Unmarshal(data, &transactions)

	transactions = append(transactions, newTransaction)

	data, _ = json.MarshalIndent(transactions, "", "")
    os.MkdirAll(filepath.Dir(dataFile), os.ModePerm)
	os.WriteFile(dataFile, data, 0644)
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringP("type", "t", "expense", "Type of transaction (expense or income)")
	addCmd.Flags().StringP("category", "c", "", "Category of the transaction")
	addCmd.Flags().Float64P("amount", "a", 0, "Amount of the transaction")

	addCmd.MarkFlagRequired("type")
	addCmd.MarkFlagRequired("category")
	addCmd.MarkFlagRequired("amount")
}