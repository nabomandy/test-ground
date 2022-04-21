package dbData

import (
	"database/sql"

	// "strconv"

	"dev.azure.com/innodepcloud/vurix-dms/_git/dms-demo-service/config"
	dao "dev.azure.com/innodepcloud/vurix-dms/_git/dms-demo-service/dao"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

//01.사용자 목록
func getUserList12345(c *gin.Context) (dao.Result, int, Global_Prop) {
	var result dao.Result
	var resultResponse dao.Result
	var count int = 0
	var selectData Global_Prop
	var dmsDBConn *gorm.DB
	var err error = nil
	var rows *sql.Rows = nil
	var resultBody []interface{}

	//result query
	query := FindGlobalProperty

	// db connection 객체
	dmsDBConn, err = dao.ConnectDB(config.GetInstance().VurixDmsDB)
	if err != nil {
		zap.S().Errorf("Fail to DB Connection : %v", err)
		return dao.ResponseFunc2(500, "Internal Server Error", resultResponse), count, selectData
	}
	defer dmsDBConn.Close()

	zap.S().Debugf("Query : %v", query) // 쿼리 체킹

	rows, err = dmsDBConn.Raw(query).Rows()
	if err != nil {
		return dao.ResponseFunc2(500, "Internal Server Error", resultResponse), count, selectData
	} else {

		for rows.Next() { // 존재하는 데이터를 스캐닝

			// dmsDBConn.ScanRows(rows, &selectData)
			selectData = Global_Prop{}
			dmsDBConn.ScanRows(rows, &selectData)
			resultBody = append(resultBody, selectData)
		}

	}
	resultResponse = dao.Result{
		// ResultsAll: ResultsAll,
	}
	result = dao.ResponseFunc2(200, "OK", resultResponse)
	return result, count, selectData
}
