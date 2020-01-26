package migrate

import (
	"database/sql"
	"testing"

	. "github.com/gchaincl/migrate/dsl"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

type migration struct {
	CreateT1 CreateTable `version:"1"`
	ChangeT1 ChangeTable `version:"2"`
	RemoveT1 RemoveTable `version:"3"`
}

var Migration = &migration{
	CreateT1: func() (*Table, error) {
		return NewTable("t1", func(t *Table) {
			t.Field("id", SERIAL, PrimaryKey)
			t.Field("email", VARCHAR, NotNull)
			t.Field("username", VARCHAR, NotNull)
			t.Field("password", VARCHAR, NotNull)
			t.Field("pin", VARCHAR, Args(4)) // This produces a VARCHAR(4)
		}), nil
	},

	ChangeT1: func() (*Change, error) {
		return NewChange("t1", func(c *Change) {
			c.Remove("pin")
			c.Add("address", TEXT)
		}), nil
	},

	RemoveT1: func() (string, error) {
		return "t1", nil
	},
}

func TestMigration(t *testing.T) {
	t.Run("SQLite3", func(t *testing.T) {
		db, err := sql.Open("sqlite3", ":memory:")
		require.NoError(t, err)
		_ = db
	})
	Migrate(Migration, &SqliteDialect{})
}
