package dbData

import (
	parameter "dbAPI/Parameter/28gang"
	"fmt"

	"gorm.io/gorm"
)

type SearchQuery struct {
	Table       string   `json:"tableName"`
	ColumnIndex []Column `json:"columnIndex"`
}

type Column struct {
	ColumnName string      `json:"columnName"`
	Judge      string      `json:"compare"`
	Value      interface{} `json:"value"`
}

func FindTable(db *gorm.DB, tablename string) bool {
	return db.Migrator().HasTable(tablename)
}

func DelTable(db *gorm.DB, tablename string) {
	db.Migrator().DropTable(tablename)
	fmt.Printf("Delete =Table:%s \n", tablename)

}

func (input *SearchQuery) SetQuery(db *gorm.DB, tablename string) *gorm.DB {

	query := db.Table(tablename)
	var tmp []parameter.GameData
	for i := 0; i < len(input.ColumnIndex); i++ {
		condition := fmt.Sprintf("%s%s?", input.ColumnIndex[i].ColumnName, input.ColumnIndex[i].Judge)
		query = query.Where(condition, input.ColumnIndex[i].Value)
	}

	query = query.Find(&tmp)
	if query.Error != nil {
		panic(query.Error)

	}

	return query

}

func FindData(db *gorm.DB, input SearchQuery) []parameter.GameData {

	db = input.SetQuery(db, input.Table)

	var tmp []parameter.GameData
	db.Find(&tmp)

	return tmp

}

func DeleteData(db *gorm.DB, input SearchQuery) {

	db = input.SetQuery(db, input.Table)

	var tmp parameter.GameData
	db.Delete(&tmp)

}
