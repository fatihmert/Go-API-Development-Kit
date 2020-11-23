package models

import (
	"log"

	"apidev.fatihmert.dev/states"
)

func (mdl *{{ .Name }}) All() []*{{ .Name }} {
	m := make([]*{{ .Name }}, 0)

	rows, err := states.DB.Query(`SELECT * FROM {{ .Table }}`)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		mItem := new({{ .Name }})
		if err := rows.Scan(
			{{ range $i, $e := .Columns}}&mItem.{{ $e }}, {{end}}
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

func (mdl *{{ .Name }}) FindFromId(id int) *{{ .Name }} {
	m := new({{ .Name }})
	rows, err := states.DB.Query(`SELECT * FROM {{ .Table }} WHERE id=?`, id)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			{{ range $i, $e := .Columns}}&m.{{ $e }}, {{end}}
		); err != nil {
			panic(err)
		}
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}

	return m
}