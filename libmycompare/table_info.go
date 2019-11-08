package libmycompare

const TableInfoTypeUnknow = 0
const TableInfoTypeBaseTable = 1
const TableInfoTypeView = 2

//TableInfo table infomation
type TableInfo struct {
	TableName string
	TableType int
}

//SetTableType set table type
func (t *TableInfo) SetTableType(tableType string) {
	t.TableType = TableInfoTableType(tableType)
}

//TableInfoTableType convert table string to int
func TableInfoTableType(tableType string) int {
	if tableType == "BASE TABLE" {
		return TableInfoTypeBaseTable
	} else if tableType == "VIEW" {
		return TableInfoTypeView
	}
	return TableInfoTypeUnknow
}
