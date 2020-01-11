package models

import "fmt"

func CreateTest() error {
	sqlStatement := `
	INSERT INTO test (test_id, test_word)
	VALUES ($1, $2)
	RETURNING test_id`
	id := 0
	err := db.QueryRow(sqlStatement, 123, "hello world").Scan(&id)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
