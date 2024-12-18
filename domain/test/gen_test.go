package test

import (
	"fmt"
	"goTest/domain/factory"
	"gorm.io/gen"
	"log"
	"testing"
)

func TestGen1(t *testing.T) {

	db, _ := factory.GetDb().DB()

	tableName := "tbl_order"
	rows, err := db.Query(fmt.Sprintf("SELECT COLUMN_NAME, DATA_TYPE FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME = '%s'", tableName))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Printf("type %s struct {\n", capitalize(tableName))
	for rows.Next() {
		var columnName, dataType string
		if err := rows.Scan(&columnName, &dataType); err != nil {
			log.Fatal(err)
		}
		goType := map[string]string{
			"int":       "int",
			"tinyint":   "int8",
			"smallint":  "int16",
			"mediumint": "int32",
			"bigint":    "int64",
			"float":     "float32",
			"double":    "float64",
			"varchar":   "string",
			"char":      "string",
			"text":      "string",
			"datetime":  "time.Time",
			"timestamp": "time.Time",
			"date":      "time.Time",
			"time":      "time.Time",
			"binary":    "[]byte",
			"varbinary": "[]byte",
			"blob":      "[]byte",
			"json":      "json.RawMessage",
		}[dataType]
		if goType == "" {
			goType = "interface{}"
		}
		fmt.Printf("\t%s %s `json:\"%s\"`\n", capitalize(columnName), goType, columnName)
	}
	fmt.Println("}")
}

func capitalize(s string) string {
	if s == "" {
		return ""
	}
	return string(s[0]-'a'+'A') + s[1:]
}

func TestGen2(t *testing.T) {
	g := gen.NewGenerator(gen.Config{
		OutPath: "../query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	gormdb := factory.GetDb()
	// gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(gormdb) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	allModels := g.GenerateAllTable()

	g.ApplyBasic(
		// Generate struct `User` based on table `users`
		allModels...,
	)

	// Generate the code
	g.Execute()

}

func TestGen3(t *testing.T) {
	//query.Tbl.Table("tbl_users")
}
