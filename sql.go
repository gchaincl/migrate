package migrate

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/gchaincl/migrate/dsl"
)

func TableSQL(t *dsl.Table, d Dialect) string {
	buf := bytes.NewBuffer(nil)
	lines := []string{}
	fmt.Fprintf(buf, "CREATE TABLE %s (\n", t.Name)
	for _, ft := range t.Fields {
		line := bytes.NewBuffer(nil)
		field, name := ft.Field, ft.Name

		// retrieve the datatype name from the dialect or fallback to .String()
		dt := d.DataType(ft.Field.Type)
		if dt == "" {
			dt = field.Type.String()
		}
		fmt.Fprintf(line, "\t%s %s%s", name, dt, fmtArgs(field.Args))

		if field.NotNull {
			fmt.Fprintf(line, " NOT NULL")
		}

		if field.PrimaryKey {
			fmt.Fprint(line, " PRIMARY KEY")
		}

		if def := field.Default; def != nil {
			fmt.Fprintf(line, " DEFAULT %v", def)
		}

		lines = append(lines, line.String())
	}
	fmt.Fprintf(buf, "%s\n", strings.Join(lines, ",\n"))
	fmt.Fprintf(buf, ")\n")

	return buf.String()
}

func fmtArgs(args []interface{}) string {
	nArgs := len(args)
	if nArgs == 0 {
		return ""
	}

	var strs = make([]string, nArgs)
	for i, arg := range args {
		strs[i] = fmt.Sprintf("%v", arg)
	}

	return fmt.Sprintf("(%s)", strings.Join(strs, ","))
}
