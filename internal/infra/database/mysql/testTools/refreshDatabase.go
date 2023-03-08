package testTools

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ShemaInformation struct {
	TableName string
}

type DatabaseOperations struct {
	connection *sqlx.DB
}

func NewTestDatabaseOperations(connection *sqlx.DB) *DatabaseOperations {
	return &DatabaseOperations{
		connection: connection,
	}
}

func (t *DatabaseOperations) RefreshDatabase() {
	t.truncateTables(t.getTables())
}

func (t *DatabaseOperations) getTables() []ShemaInformation {
	schemaInformation := []ShemaInformation{}

	rows, err := t.connection.Query("SELECT TABLE_NAME FROM information_schema.TABLES WHERE TABLE_SCHEMA = 'siacob' AND NOT TABLE_NAME = 'schema_migrations'")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		schema := ShemaInformation{}
		rows.Scan(&schema.TableName)
		schemaInformation = append(schemaInformation, schema)
	}

	return schemaInformation
}

func (t *DatabaseOperations) truncateTables(shemaInformation []ShemaInformation) {
	for _, schema := range shemaInformation {
		_, _ = t.connection.Exec("SET FOREIGN_KEY_CHECKS=0;")
		_, err := t.connection.Query(fmt.Sprintf("TRUNCATE TABLE %s;", schema.TableName))
		if err != nil {
			panic(err)
		}
		_, _ = t.connection.Query("SET FOREIGN_KEY_CHECKS=1;")
	}
}
