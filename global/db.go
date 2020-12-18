package global

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var (
	DBLink *gorm.DB
)

func SetupDBLink() (error) {
	var err error
	DBLink, err = gorm.Open(DatabaseSetting.DBType,
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
			DatabaseSetting.UserName,
			DatabaseSetting.Password,
			DatabaseSetting.Host,
			DatabaseSetting.DBName,
			DatabaseSetting.Charset,
			DatabaseSetting.ParseTime,
	))
	if err != nil {
		return err
	}

	if ServerSetting.RunMode == "debug" {
		DBLink.LogMode(true)
	}
	DBLink.SingularTable(true)
	DBLink.DB().SetMaxIdleConns(DatabaseSetting.MaxIdleConns)
	DBLink.DB().SetMaxOpenConns(DatabaseSetting.MaxOpenConns)
	return nil
}

