package migrate

import (
	"fmt"
	"log"
	"reflect"

	"github.com/gchaincl/migrate/dsl"
)

func Migrate(m interface{}, dialect Dialect) error {
	value := reflect.ValueOf(m)

	// Dereference if m is a pointer by calling .Elem()
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	// We need to get a struct as argument
	if value.Kind() != reflect.Struct {
		panic("Migrate() can only handle struct{} type")
	}

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		tag := value.Type().Field(i).Tag
		log.Printf("tag = %+v\n", tag)
		switch t := field.Interface().(type) {
		case dsl.CreateTable:
			table, err := t()
			if err != nil {
				return err
			}
			fmt.Printf("%s", TableSQL(table, dialect))
		case dsl.ChangeTable:
			change, err := t()
			if err != nil {
				return err
			}
			log.Printf("change = %+v\n", change)
		case dsl.RemoveTable:
			remove, err := t()
			if err != nil {
				return err
			}
			log.Printf("remove = %+v\n", remove)
		}
	}
	return nil
}
