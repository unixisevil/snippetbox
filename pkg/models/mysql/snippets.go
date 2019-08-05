package mysql

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/unixisevil/snippetbox/pkg/models"
)

type SnippetModel struct {
	DB *sqlx.DB
}

func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	stmt := `insert into snippets (title, content, created, expires)
values(?, ?, utc_timestamp(), date_add(utc_timestamp(), interval ? day))`

	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	stmt := `select id, title, content, created, expires from snippets
where expires > utc_timestamp() and id = ?`

	//row := m.DB.QueryRow(stmt, id)
	s := &models.Snippet{}
	//err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	err := m.DB.Get(s, stmt, id)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}
	return s, nil
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	stmt := `select id, title, content, created, expires from snippets
where expires > utc_timestamp() order by created desc limit 10`

	/*rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	*/
	snippets := []*models.Snippet{}
	if err := m.DB.Select(&snippets, stmt); err != nil {
		return nil, err
	}
	return snippets, nil
}
