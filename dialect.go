package migrate

import "github.com/gchaincl/migrate/dsl"

type Dialect interface {
	DataType(dsl.DataType) string
}

type noopDialect struct {
}

func (*noopDialect) DataType(t dsl.DataType) string { return "" }

type PostgresDialect struct {
}

func (*PostgresDialect) DataType(t dsl.DataType) string {
	switch t {
	case dsl.ID:
		return "SERIAL"
	}
	return ""
}

type MySqlDialect struct {
}

func (*MySqlDialect) DataType(t dsl.DataType) string {
	switch t {
	case dsl.ID:
		return "INT AUTO_INCREMENT"
	}
	return ""
}

type SqliteDialect struct {
}

func (*SqliteDialect) DataType(t dsl.DataType) string {
	switch t {
	case dsl.ID:
		return "INT AUTO_INCREMENT"
	}
	return ""
}
