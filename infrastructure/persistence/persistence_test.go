package persistence

import (
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/yasszu/go-jwt-auth/infrastructure/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func TestMain(m *testing.M) {
	var err error

	db, err = postgres.NewTestConn()
	if err != nil {
		panic(err)
	}

	code := m.Run()

	os.Exit(code)
}

func prepare(t *testing.T, fn func()) {
	tables, err := getTables()
	if err != nil {
		t.Error(err)
	}

	for _, table := range tables {
		if err := truncate(table); err != nil {
			t.Log(err)
		}
	}

	fn()
}

func getTables() ([]string, error) {
	var tables []string
	q := `
SELECT tablename
FROM pg_catalog.pg_tables
WHERE schemaname != 'pg_catalog'
  AND schemaname != 'information_schema'
`
	if err := db.Raw(q).Scan(&tables).Error; err != nil {
		return nil, err
	}

	return tables, nil
}

func truncate(table string) error {
	if err := db.Exec("TRUNCATE TABLE ? RESTART IDENTITY CASCADE", table).Error; err != nil {
		return err
	}
	return nil
}
