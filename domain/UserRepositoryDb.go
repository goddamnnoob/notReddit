package domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type UserRepositoryDb struct {
	client *sql.DB
}

func (d UserRepositoryDb) GetAllUsers() ([]User, error) {
	getAll := "select customer_id,name,city,zipcode,date_of_birth, status from customers"
	rows, err := d.client.Query(getAll)
	if err != nil {
		log.Println("Error Querying Customer Table " + err.Error())
		return nil, err
	}
	users := make([]User, 0)
	for rows.Next() {

		var u User
		err := rows.Scan(&u.Id, &u.Name, &u.City, &u.Zipcode, &u.DateOfBirth, &u.Status)
		if err != nil {
			log.Println("Error while scanningUsers" + err.Error())
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func NewUserRepositoryDb() UserRepositoryDb {
	client, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	client.SetConnMaxIdleTime(time.Minute * 3)
	client.SetMaxIdleConns(10)
	client.SetMaxOpenConns(10)
	return UserRepositoryDb{client}
}

func (d UserRepositoryDb) ById(id string) (*User, error) {
	byId := "select customer_id,name,city,zipcode,date_of_birth, status from customers where customer_id=?"
	rows := d.client.QueryRow(byId, id)
	var u User
	err := rows.Scan(&u.Id, &u.Name, &u.City, &u.Zipcode, &u.DateOfBirth, &u.Status)
	if err != nil {
		log.Println("Error while scan user: " + err.Error())
		return nil, err
	}
	return &u, nil
}