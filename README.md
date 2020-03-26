# Migrate
Migrate is a DSL for the Go programming lanuage that helps to migrate your SQL database.
Migrate is vendor independent and currently SQLite, MySQL and Postgres dialects are supported.

This is experimental software and the API will change.

# Usage

```go
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
```
