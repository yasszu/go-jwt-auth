package persistence

import (
	"fmt"
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
	tables, err := getTables(t)
	if err != nil {
		t.Error(err)
	}

	for _, table := range tables {
		if err := truncate(t, table); err != nil {
			t.Error(err)
		}
	}

	fn()
}

func getTables(t *testing.T) ([]string, error) {
	var tables []string
	q := `
SELECT tablename
FROM pg_catalog.pg_tables
WHERE schemaname != 'pg_catalog'
  AND schemaname != 'information_schema'
`
	if err := db.Raw(q).Scan(&tables).Error; err != nil {
		t.Error(err)
		return nil, err
	}
	t.Logf("TABLES: %v", tables)

	return tables, nil
}

func truncate(t *testing.T, table string) error {
	t.Logf("TRUNCATE TABLE: %s", table)

	q := fmt.Sprintf("TRUNCATE %s RESTART IDENTITY CASCADE", table)
	if err := db.Exec(q).Error; err != nil {
		return err
	}

	return nil
}
