package libmycompare

import "fmt"

//MyCompare main of lib
type MyCompare struct {
	SrcConfig  ConnConfig
	DestConfig ConnConfig
}

//Run start compare
func (m MyCompare) Run() error {

	err := m.exec()
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (m MyCompare) exec() error {
	srcDb := database{
		Config: m.SrcConfig,
	}

	destDb := database{
		Config: m.DestConfig,
	}

	//open
	err := srcDb.Open()
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	defer srcDb.Close()

	err = destDb.Open()
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	defer destDb.Close()

	tableSrcs, err := srcDb.ListAllTables()
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	tableDests, err := destDb.ListAllTables()
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	notExists, err := m.compareTables(tableSrcs, tableDests)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	//จุดนี้ table ตรงกันหมดแล้วใช้ชื่อใน tableSrcs
	for _, tbInfo := range tableSrcs {
		if m.isContain(notExists, tbInfo.TableName) {
			continue
		}
		srcSchs, err := srcDb.Schema(tbInfo.TableName)
		if err != nil {
			return fmt.Errorf("%w", err)
		}

		destSchs, err := destDb.Schema(tbInfo.TableName)
		if err != nil {
			return fmt.Errorf("%w", err)
		}

		err = m.compareSchema(tbInfo.TableName, srcSchs, destSchs)
		if err != nil {
			return fmt.Errorf("%w", err)
		}

	}

	return nil
}

func (m MyCompare) isContain(notExists []string, tbname string) bool {
	for _, n := range notExists {
		if n == tbname {
			return true
		}
	}
	return false
}

func (m MyCompare) compareSchema(tbName string, srcSchs, destSchs []SchemaInfo) error {

	for _, srcSch := range srcSchs {
		found := false
		for _, destSch := range destSchs {
			if srcSch.FieldName == destSch.FieldName {
				found = true

				var fieldNotMatchs []string
				if srcSch.FieldType != destSch.FieldType {
					fieldNotMatchs = append(fieldNotMatchs, fmt.Sprintf("Type src is '%s' , dest is '%s'", srcSch.FieldType, destSch.FieldType))
				}
				if srcSch.FieldNull != destSch.FieldNull {
					fieldNotMatchs = append(fieldNotMatchs, fmt.Sprintf("Null src is '%s' , dest is '%s'", srcSch.FieldNull, destSch.FieldNull))
				}
				if srcSch.FieldKey != destSch.FieldKey {
					fieldNotMatchs = append(fieldNotMatchs, fmt.Sprintf("Key src is '%s' , dest is '%s'", srcSch.FieldKey, destSch.FieldKey))
				}
				if srcSch.FieldDefult != destSch.FieldDefult {
					fieldNotMatchs = append(fieldNotMatchs, fmt.Sprintf("Default src is '%s' , dest is '%s'", srcSch.FieldDefult, destSch.FieldDefult))
				}
				if srcSch.FieldExtra != destSch.FieldExtra {
					fieldNotMatchs = append(fieldNotMatchs, fmt.Sprintf("Extra src is '%s' , dest is '%s'", srcSch.FieldExtra, destSch.FieldExtra))
				}

				for _, fieldNotMatch := range fieldNotMatchs {
					fmt.Printf("field not match: %s.%s \t\t\t %s\n", tbName, srcSch.FieldName, fieldNotMatch)
				}

				break
			}
		}
		if !found {
			fmt.Printf("field not found: %s.%s \n", tbName, srcSch.FieldName)
		}
	}

	return nil
}

func (m MyCompare) compareTables(tableSrcs, tableDests []TableInfo) ([]string, error) {
	var notExists []string
	for _, tableSrc := range tableSrcs {
		found := false
		for _, tableDest := range tableDests {
			if tableSrc.TableName == tableDest.TableName && tableSrc.TableType == tableDest.TableType {
				found = true
				break
			}
		}
		if !found {
			fmt.Printf("not found %s\n", tableSrc.TableName)
			notExists = append(notExists, tableSrc.TableName)
		}
	}

	return notExists, nil
}
