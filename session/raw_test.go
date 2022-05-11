package session

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"testing"
)

var TestDB *sql.DB

func TestMain(m *testing.M) {
	TestDB, _ = sql.Open("sqlite3", "gee.db")
	code := m.Run()
	_ = TestDB.Close()
	os.Exit(code)
}

func newSession() *Session {
	s := &Session{db: TestDB}
	_, _ = s.Raw("drop table if exists User;").Exec()
	_, _ = s.Raw("create table User(Name text);").Exec()
	_, _ = s.Raw("insert into User(Name) values (?), (?)", "Jack", "Jam").Exec()
	return s
}

func TestSession_Exec(t *testing.T) {
	s := newSession()
	result, _ := s.Raw("insert into User(Name) values (?), (?);", "Tom", "Sam").Exec()
	if count, err := result.RowsAffected(); err != nil || count != 2 {
		t.Fatal(err.Error(), "expected 2 but get", count)
	}
}

func TestSession_QueryRow(t *testing.T) {
	s := newSession()
	row := s.Raw("select count(*) from user").QueryRow()
	var count int
	if err := row.Scan(&count); err != nil || count != 2 {
		t.Fatal(err.Error(), "expected 2 but get", count)
	}
}

func TestSession_QueryRows(t *testing.T) {
	s := newSession()
	rows, _ := s.Raw("select * from user").QueryRows()
	var names []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			t.Fatal(err.Error())
		}
		names = append(names, name)
	}
	fmt.Println(names)
}

