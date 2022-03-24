package model

import "fmt"

type MysqlSetting struct {
	sqlDriver string
	user      string
	password  string
	protocol  string
	address   string
	dataBase  string
}

func NewMySqlSetting(
	sqlDriver string,
	user string,
	password string,
	protocol string,
	address string,
	dataBase string) MysqlSetting {
	s := MysqlSetting{
		sqlDriver: sqlDriver,
		user:      user,
		password:  password,
		protocol:  protocol,
		address:   address,
		dataBase:  dataBase,
	}
	return s
}

func (s *MysqlSetting) DriverName() string {
	return s.sqlDriver
}

func (s *MysqlSetting) DataSourceName() string {
	return fmt.Sprintf("%s:%s@%s(%s)/%s",
		s.user,
		s.password,
		s.protocol,
		s.address,
		s.dataBase)
}
