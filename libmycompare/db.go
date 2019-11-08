package libmycompare

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type database struct {
	dbconn *sql.DB
	Config ConnConfig
}

func (d *database) Open() error {
	conn, err := sql.Open("mysql", d.Config.ConnString)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	d.dbconn = conn
	return nil
}

func (d *database) Close() error {
	if d.dbconn == nil {
		return fmt.Errorf("dbconn is nil")
	}
	return d.dbconn.Close()
}

func (d database) ListAllTables() ([]TableInfo, error) {
	rows, err := d.dbconn.Query("show full tables")
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	defer rows.Close()

	var tables []TableInfo
	for rows.Next() {
		tableName := ""
		tableType := ""
		err := rows.Scan(&tableName, &tableType)
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}
		table := TableInfo{
			TableName: tableName,
		}
		tables = append(tables, table)
	}
	return tables, nil
}

func (d database) Schema(tbName string) ([]SchemaInfo, error) {
	rows, err := d.dbconn.Query(fmt.Sprintf("desc `%s`", tbName))
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	defer rows.Close()

	var sinfos []SchemaInfo
	for rows.Next() {
		fieldName := ""
		fieldType := ""
		fieldNull := ""
		fieldKey := ""
		fieldDefult := sql.NullString{}
		fieldExtra := ""
		err := rows.Scan(&fieldName, &fieldType, &fieldNull, &fieldKey, &fieldDefult, &fieldExtra)
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}
		sinfo := SchemaInfo{
			FieldName:   fieldName,
			FieldType:   fieldType,
			FieldNull:   fieldNull,
			FieldKey:    fieldKey,
			FieldDefult: fieldDefult.String,
			FieldExtra:  fieldExtra,
		}
		sinfos = append(sinfos, sinfo)
	}

	return sinfos, nil
}
