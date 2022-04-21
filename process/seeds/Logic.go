package seeds

import (
	"database/sql"
	"io/fs"
	"io/ioutil"

	"dev.azure.com/innodepcloud/vurix-dms/_git/dms-demo-service/config"
	"dev.azure.com/innodepcloud/vurix-dms/_git/dms-demo-service/dao"
	"go.uber.org/zap"
)

func (s Seed) CreateDemoProgramSeeder() {
	var err error
	if s.dmsDBConn, err = dao.ConnectDB(s.dmsDBInfo); err != nil {
		zap.S().Errorf("Fail to connect DB: %v : %v", s.dmsDBInfo, err)
		return
	}
	defer s.dmsDBConn.Close()

	query := ` create table if not exists public."DemoProgramSeeder" (
		file_name varchar(255) not null,
		constraint demoprogram_seeder_pk primary key (file_name)); `

	zap.S().Infof(query)
	s.dmsDBConn.Exec(query)
}

func (s Seed) CheckExists() {
	var err error
	var cnt int
	var rows *sql.Rows = &sql.Rows{}
	var seederFile string
	var seederFiles []string = []string{}
	var sqlFiles []fs.FileInfo

	if s.dmsDBConn, err = dao.ConnectDB(s.dmsDBInfo); err != nil {
		zap.S().Errorf("Fail to connect DB: %v : %v", s.dmsDBInfo, err)
		return
	}
	defer s.dmsDBConn.Close()

	cntQuery := ` select count(file_name) as cnt from public."DemoProgramSeeder"; `
	zap.S().Infof("%v", cntQuery)
	cntRow := s.dmsDBConn.Raw(cntQuery).Row()
	cntRow.Scan(&cnt)

	zap.S().Infof("Number of 'DemoProgramSeeder' table rows: %v", cnt)

	if cnt > 0 {
		query := ` select file_name from public."DemoProgramSeeder"; `
		zap.S().Infof("%v", query)

		if rows, err = s.dmsDBConn.Raw(query).Rows(); err != nil {
			zap.S().Errorf("DB Error: %v\n", err)
			return
		}

		for rows.Next() {
			seederFile = ""
			rows.Scan(&seederFile)
			seederFiles = append(seederFiles, seederFile)
		}
	}

	if sqlFiles, err = ioutil.ReadDir(config.GetInstance().SeederFilePath); err != nil {
		zap.S().Errorf("Get Seed Sql Folder Error: %v\n", err)
		return
	}

	for _, f := range sqlFiles {
		if !dao.Contains(seederFiles, f.Name()) {
			s.ExecuteQuery(f.Name())
		}
	}
}

func (s Seed) ExecuteQuery(fileName string) {
	var err error

	byte, err := ioutil.ReadFile(config.GetInstance().SeederFilePath + fileName)

	if err != nil {
		zap.S().Errorf("Fail to get seeder file: %v\n", err)
		return
	}

	zap.S().Infof(" ===> Excute Seeder Query: %v", fileName)
	execQuery := string(byte)
	zap.S().Infof("\n%v\n", execQuery)
	if dmsDB := s.dmsDBConn.Exec(execQuery); dmsDB.Error != nil {
		zap.S().Errorf("Fail to exec seeder file: [%v]\n%v\n", fileName, err)
		return
	}

	insertQuery := ` insert into public."DemoProgramSeeder" (file_name) values(?); `
	zap.S().Infof("%v [%v]", insertQuery, fileName)
	s.dmsDBConn.Exec(insertQuery, fileName)
}

/** GetSourcePath() 실행 시 /go/app 리턴
func GetSourcePath() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}
*/
