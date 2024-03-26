/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add current day as assisted",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("add called")
		db, _ := sql.Open("sqlite3", "./oat.db")
		query, _ := db.Prepare("CREATE TABLE IF NOT EXISTS Days (id INTEGER PRIMARY KEY AUTOINCREMENT, date DATE UNIQUE)")
		query.Exec()

		currentTime := time.Now()
		fmt.Println(currentTime)
		query, _ = db.Prepare("INSERT INTO Days (date) VALUES (?)")
		query.Exec(currentTime.Format(time.DateOnly))

		rows, _ := db.Query("SELECT * FROM Days")
		var id int
		var date time.Time
		for rows.Next() {
			rows.Scan(&id, &date)
			// if e != nil {
			// 	fmt.Println("Parsing error: ", e)
			// 	return
			// }
			fmt.Println(strconv.Itoa(id))
			fmt.Println(date.Format(time.DateOnly))
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
