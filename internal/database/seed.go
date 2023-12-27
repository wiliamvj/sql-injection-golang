package database

import (
  "log"
)

func SeedUsers() error {
  // drop table users
  _, err := DBConnection.Exec(`DROP TABLE IF EXISTS users`)
  if err != nil {
    log.Fatal(err)
  }

  // create table users
  createTableQuery := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(256) NOT NULL,
			email VARCHAR(256) NOT NULL UNIQUE,
      password VARCHAR(256) NOT NULL
		)
	`
  _, err = DBConnection.Exec(createTableQuery)
  if err != nil {
    log.Fatal(err)
  }
  log.Println("Tabela de usuários criada com sucesso.")

  insertUserQuery := `
		INSERT INTO users (name, email, password) VALUES
		('John Doe', 'john.doe@example.com', 123456),
		('Bob', 'bob@example.com', 123456),
		('Charlie', 'charlie@example.com', 123456),
    ('Slash', 'slash@example.com', 098765),
    ('Gilmour', 'gilmour@example.com', 1255657),
    ('Steve Vai', 'steve_vai@example.com', 1255657)
	`

  _, err = DBConnection.Exec(insertUserQuery)
  if err != nil {
    log.Fatal(err)
  }
  log.Println("Usuários inseridos com sucesso.")
  return nil
}
