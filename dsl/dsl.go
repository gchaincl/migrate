package dsl

// DataType indicates the possible fields data types
type DataType int

const (
	// ID is an abstract type, use this if you want a numerical auto-increment primary key field
	// Each dialect interpret this DataType differently
	ID DataType = iota

	CHAR
	VARCHAR
	BINARY
	VARBINARY
	TEXT

	BOOL
	INT
	SERIAL
	BIGINT
	FLOAT

	DATE
	DATETIME
)

// Field represent a database column
type Field struct {
	Type       DataType
	Args       []interface{}
	NotNull    bool
	PrimaryKey bool
	Default    interface{}
}

// NewField returns a new Field given a DataType and zero or more options
func NewField(_type DataType, opts ...FieldOpt) *Field {
	f := &Field{Type: _type}
	for _, opt := range opts {
		opt(f)
	}
	return f
}

type FieldOpt func(f *Field)

// Args are DataType args. Args are passed to the DataType when creating a new Field
func Args(args ...interface{}) FieldOpt {
	return func(f *Field) { f.Args = args }
}

// NotNull indicates that the field can't be null
func NotNull(f *Field) { f.NotNull = true }

// PrimaryKey sets the field as a primary key
func PrimaryKey(f *Field) { f.PrimaryKey = true }

// Default specifies the default field value
func Default(v interface{}) FieldOpt {
	return func(f *Field) { f.Default = v }
}

type FieldTuple struct {
	Name  string
	Field *Field
}

// Table defines a table creation behavior an its fields
type Table struct {
	Name   string
	Fields []FieldTuple
}

// NewTable returns a new table given one or more fields
// fn gets the table reference.
func NewTable(name string, fn func(*Table)) *Table {
	t := &Table{
		Name: name,
	}

	fn(t)
	return t
}

// Field creates a new Field definition for a given table
func (t *Table) Field(col string, _type DataType, opts ...FieldOpt) *Table {
	t.Fields = append(t.Fields, FieldTuple{
		Name:  col,
		Field: NewField(_type, opts...),
	})
	return t
}

type CreateTable func() (*Table, error)

type Change struct {
	name      string
	removes   []string
	newFields map[string]*Field
}

func NewChange(name string, fn func(*Change)) *Change {
	c := &Change{
		name:      name,
		newFields: make(map[string]*Field),
	}
	fn(c)
	return c
}

func (c *Change) Remove(col string) *Change {
	c.removes = append(c.removes, col)
	return c
}
func (c *Change) Add(col string, _type DataType, opts ...FieldOpt) *Change {
	c.newFields[col] = NewField(_type, opts...)
	return c
}

type ChangeTable func() (*Change, error)

type RemoveTable func() (string, error)
