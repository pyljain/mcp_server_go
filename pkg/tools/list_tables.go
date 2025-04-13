package tools

import (
	"database/sql"
	"strings"
)

type ListTables struct {
	conn *sql.DB
}

func NewListTables(conn *sql.DB) *ListTables {
	return &ListTables{conn: conn}
}

func (lt *ListTables) Execute(arguments map[string]interface{}) (string, error) {
	query := `
		SELECT name
		FROM sqlite_master
		WHERE type='table' AND name NOT LIKE 'sqlite_%'
		ORDER BY name;
	`

	rows, err := lt.conn.Query(query)
	if err != nil {
		return "", err
	}

	defer rows.Close()

	var tables []string
	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			return "", err
		}

		tables = append(tables, table)
	}

	if err := rows.Err(); err != nil {
		return "", err
	}

	return strings.Join(tables, "\n"), nil
}
