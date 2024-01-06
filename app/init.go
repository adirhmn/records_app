package app

import (
	"belajar-golang-restful-api/helper"
	"database/sql"
	"fmt"
)

func tableExists(db *sql.DB, tableName string) bool {
	query := `
		SELECT EXISTS (
			SELECT FROM information_schema.tables 
			WHERE table_schema = 'public' 
			AND table_name = $1
		);
	`

	var exists bool
	err := db.QueryRow(query, tableName).Scan(&exists)
	if err != nil {
		helper.PanicifError(err)
		return false
	}
	return exists
}

func createTable( db *sql.DB) {
	createTableSQL := `
		CREATE TABLE IF NOT EXISTS records (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		marks INT[] NOT NULL,
		createdAt TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
	);;`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		helper.PanicifError(err)
	}
	fmt.Println("Table created successfully.")
}

func insertSampleData( db *sql.DB) {
	insertSQL := `
		INSERT INTO records (name, marks)
		VALUES
			('Student1', '{90, 50, 40, 20}'),
			('Student2', '{80, 55, 50, 15}'),
			('Student3', '{70, 60, 55, 10}'),
			('Student4', '{60, 65, 60, 5}'),
			('Student5', '{50, 70, 65, 5}'),
			('Student6', '{40, 75, 70, 5}'),
			('Student7', '{30, 80, 75, 5}'),
			('Student8', '{20, 85, 80, 5}'),
			('Student9', '{10, 90, 85, 5}'),
			('Student10', '{5, 95, 90, 5}'),
			('Student11', '{70, 50, 80}'),
			('Student12', '{65, 55, 85}'),
			('Student13', '{60, 60, 90}'),
			('Student14', '{55, 65, 95}'),
			('Student15', '{50, 70, 100}'),
			('Student16', '{90, 30, 80}'),
			('Student17', '{85, 35, 85}'),
			('Student18', '{80, 40, 90}'),
			('Student19', '{75, 45, 95}'),
			('Student20', '{70, 50, 100}'),
			('Student21', '{50, 40, 80, 30}'),
			('Student22', '{45, 35, 85, 35}'),
			('Student23', '{40, 30, 90, 40}'),
			('Student24', '{35, 25, 95, 45}'),
			('Student25', '{30, 20, 100, 50}'),
			('Student26', '{85, 15, 80, 20}'),
			('Student27', '{80, 10, 85, 25}'),
			('Student28', '{75, 5, 90, 30}'),
			('Student29', '{70, 5, 95, 35}'),
			('Student30', '{65, 5, 100, 40}');
		`

	_, err := db.Exec(insertSQL)
	helper.PanicifError(err)
	fmt.Println("Data records inserted successfully.")
}

func InitDB(db *sql.DB)  {
	if !tableExists(db, "records"){
		createTable(db);
		insertSampleData(db);
	}
	
}
