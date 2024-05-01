package cypher

import (
	"fmt"
	"reflect"

	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// Parse parses a cypher query result record and store the result in target interface
//   - record: neo4j result record
//   - alias: the alias used in cypher query (e.g. m.title)
//   - target: target interface (e.g. movie.Entity)
//     Target object should a "neo4j" tab (e.g. `neo4j:"title"`)
func Parse(record *neo4j.Record, alias string, target interface{}) error {
	elem := reflect.ValueOf(target).Elem()
	for i := 0; i < elem.Type().NumField(); i++ {
		structField := elem.Type().Field(i)
		tag := structField.Tag.Get("neo4j")
		fieldType := structField.Type
		fieldName := structField.Name
		if val, ok := record.Get(fmt.Sprintf("%s.%s", alias, tag)); ok {
			if val == nil {
				continue
			}
			field := elem.FieldByName(fieldName)
			if field.IsValid() {
				t := fieldType.String()
				switch t {
				case "string":
					field.SetString(val.(string))
				case "int64":
					field.SetInt(val.(int64))
				case "uuid.UUID":
					field.Set(reflect.ValueOf(uuid.MustParse(val.(string))))
				default:
					return fmt.Errorf("invalid type: %s", t)
				}
			}
		}

	}

	return nil
}

// ParseStruct parses a cypher query result record and store the result in target interface
//   - record: neo4j result record
//   - alias: the alias used in cypher query (e.g. m.title)
//   - target: target interface (e.g. movie.Entity)
//     Target object should a "neo4j" tab (e.g. `neo4j:"title"`)
func ParseStruct(record *neo4j.Record, alias string, target interface{}) error {
	val, ok := record.Get(alias)
	if ok {
		valDb := val.(neo4j.Node)
		mapstructure.Decode(valDb.GetProperties(), &target)
	}
	return nil
}
