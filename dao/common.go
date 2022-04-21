package dao

import (
	"bytes"
	"database/sql"
	"io"

	"dev.azure.com/innodepcloud/vurix-dms/_git/dms-demo-service/config"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

type TableColumnInfo struct {
	TableName     string `json:"table_name" gorm:"column:table_name"`
	TableOID      string `json:"table_oid" gorm:"column:table_oid"`
	ColumnName    string `json:"column_name" gorm:"column:column_name"`
	ColumnOID     string `json:"column_oid" gorm:"column:column_oid"`
	ColumnComment string `json:"column_comment" gorm:"column:column_comment"`
}

func GetColName(tableName string) ([]interface{}, error) {
	var dmsDBConn *gorm.DB
	var tableColumn TableColumnInfo
	var commentList []interface{} = []interface{}{}
	var err error = nil
	var rows *sql.Rows = nil

	// db connect
	dmsDBConn, err = ConnectDB(config.GetInstance().VurixDmsDB)
	if err != nil {
		zap.S().Errorf("Fail to DB Connection : %v", err)
		return commentList, err
	}
	defer dmsDBConn.Close()

	// make select data query
	query := ` select c.relname as "table_name"
					, a.attrelid as "table_oid"
					, a.attname as "column_name"
					, a.attnum as "column_oid"
					, (select col_description(a.attrelid, a.attnum)) as "column_comment"
				from pg_catalog.pg_class c
					inner join pg_catalog.pg_attribute a on a.attrelid = c.oid
				where c.relname = ?
					and a.attnum > 0
					and a.attisdropped is false
					and pg_catalog.pg_table_is_visible(c.oid)
				order by a.attrelid, a.attnum `

	zap.S().Infof("%v : %v", query, tableName)
	rows, err = dmsDBConn.Raw(query, tableName).Rows()

	if err != nil {
		return commentList, err
	}

	for rows.Next() {
		tableColumn = TableColumnInfo{}
		rows.Scan(&tableColumn.TableName, &tableColumn.TableOID, &tableColumn.ColumnName, &tableColumn.ColumnOID, &tableColumn.ColumnComment)
		commentList = append(commentList, tableColumn.ColumnComment)
	}

	return commentList, nil
}

func StreamToString(stream io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.String()
}
