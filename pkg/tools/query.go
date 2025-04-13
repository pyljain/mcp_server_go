package tools

import (
	"database/sql"
	"strings"
)

type Query struct {
	conn *sql.DB
}

func NewQuery(conn *sql.DB) *Query {
	return &Query{
		conn: conn,
	}
}

func (q *Query) Execute(arguments map[string]interface{}) (string, error) {
	rows, err := q.conn.Query(arguments["query"].(string))
	if err != nil {
		return "", err
	}

	defer rows.Close()

	// Get names of the columns
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}

	var sb strings.Builder
	// Add each column to the top row separated by  tabs
	sb.WriteString(strings.Join(columns, "\t") + "\n")

	// Create an []byte for each column as with scan values of each row will be read into rawResult. Can't use string as strings don't hold nils but values could be nil
	rawResult := make([][]byte, len(columns))

	// dest is meant to be a pointer to the raw result (row) but it conforms to the format scan expects i.e. interface
	dest := make([]interface{}, len(columns))
	for i := range rawResult {
		dest[i] = &rawResult[i]
	}

	for rows.Next() {
		err := rows.Scan(dest...)
		if err != nil {
			return "", err
		}

		var row []string
		for _, raw := range rawResult {
			if raw == nil {
				row = append(row, "NULL")
			} else {
				row = append(row, string(raw))
			}
		}
		sb.WriteString(strings.Join(row, "\t") + "\n")
	}

	if err := rows.Err(); err != nil {
		return "", err
	}

	return sb.String(), nil
}
