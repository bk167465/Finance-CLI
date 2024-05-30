# Finance CLI
Finance CLI is a command-line tool to track and manage personal finances. It allows you to add and categorize expenses and income, generate monthly reports, and import/export data in different formats (CSV, JSON). The tool leverages Cobra for the CLI framework and Viper for configuration management.

> Build the project using go build .

## Usage
### Add Transaction
./finance-cli add --type <type> --category <category> --amount <amount>

### Generate Report
./finance-cli report --month <month> --year <year>

### Import Transactions
./finance-cli import <local_path_to_csv_file>

### Export Transactions
./finance-cli export <local_path_to_csv_file>

