package migrate

import (
	"fmt"
	"testing"

	. "github.com/gchaincl/migrate/dsl"
	"github.com/stretchr/testify/assert"
)

func TestSQL(t *testing.T) {
	table := NewTable("test", func(t *Table) {
		t.Field("id", ID, PrimaryKey, NotNull)
		t.Field("name", VARCHAR, Args(128), NotNull)
		t.Field("created_at", DATETIME, Default("NOW"))
	})
	expected := `CREATE TABLE test (
	id ID NOT NULL PRIMARY KEY,
	name VARCHAR(128) NOT NULL,
	created_at DATETIME DEFAULT NOW
)
`
	assert.Equal(t, expected, TableSQL(table, &noopDialect{}))

}

func TestDialects(t *testing.T) {
	table := NewTable("test", func(t *Table) {
		t.Field("id", ID)
	})

	tests := []struct {
		name     string
		dialect  Dialect
		expected string
	}{
		{"noop", &noopDialect{}, "id ID"},
		{"mysql", &MySqlDialect{}, "id INT AUTO_INCREMENT"},
		{"sqlite", &SqliteDialect{}, "id INT AUTO_INCREMENT"},
		{"psql", &PostgresDialect{}, "id SERIAL"},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%v", test.dialect)
		t.Run(name, func(t *testing.T) {
			sql := TableSQL(table, test.dialect)
			assert.Contains(t, sql, test.expected)
		})
	}
}
