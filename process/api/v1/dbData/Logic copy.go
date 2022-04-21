package dbData

import (

	// "strconv"

	"database/sql"
	"fmt"

	// "encoding/json"

	"dev.azure.com/innodepcloud/vurix-dms/_git/dms-demo-service/config"
	dao "dev.azure.com/innodepcloud/vurix-dms/_git/dms-demo-service/dao"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

func GetDataList2() []Global_Property {
	var resultBody []Global_Property = []Global_Property{}
	var selectData Global_Property
	var dmsDBConn *gorm.DB
	var err error = nil
	var rows *sql.Rows = nil

	query := FindGlobalProperty // 쿼리

	// db connection 객체
	dmsDBConn, err = dao.ConnectDB(config.GetInstance().VurixDmsDB)
	if err != nil {
		zap.S().Errorf("Fail to DB Connection : %v", err)
	}
	defer dmsDBConn.Close()

	zap.S().Debugf("Query : %v", query) // 쿼리 체킹
	rows, err = dmsDBConn.Raw(query).Rows()
	if err != nil {
	} else {
		for rows.Next() { // 존재하는 데이터를 스캐닝
			selectData = Global_Property{}
			dmsDBConn.ScanRows(rows, &selectData)

			resultBody = append(resultBody, selectData)
		}
		// json.Unmarshal([]byte(selectData), &Global_Property)
		fmt.Println(selectData.Prop_value)
		DICONTAINER := selectData.Prop_value
		fmt.Println(DICONTAINER)
	}
	return resultBody
}
