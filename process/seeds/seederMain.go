package seeds

import (
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

type Seed struct {
	dmsDBInfo []string
	dmsDBConn *gorm.DB
}

func NewSeeder(dmsDBInfo []string) *Seed {
	zap.S().Info("### New DB Seeder ... ")
	return &Seed{
		dmsDBInfo: dmsDBInfo,
		dmsDBConn: nil,
	}
}
