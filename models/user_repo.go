package models

import (
	"log"

	"apidev.fatihmert.dev/states"
)

func (mdl *User) All() []*User {
	m := make([]*User, 0)

	rows, err := states.DB.Query(`SELECT * FROM users`)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		mItem := new(User)
		if err := rows.Scan(
			&mItem.ID, &mItem.Mail, &mItem.Password, &mItem.Token, 
		); err != nil {
			panic(err)
		}
		m = append(m, mItem)
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}

	return m
}

func (mdl *User) FindFromId(id int) *User {
	m := new(User)
	rows, err := states.DB.Query(`SELECT * FROM users WHERE id=?`, id)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&m.ID, &m.Mail, &m.Password, &m.Token, 
		); err != nil {
			panic(err)
		}
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}

	return m
}