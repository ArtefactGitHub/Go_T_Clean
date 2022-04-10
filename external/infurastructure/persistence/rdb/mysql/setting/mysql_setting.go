package common

import "fmt"

type MySqlSetting struct {
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
	dataBase string) MySqlSetting {
	s := MySqlSetting{
		sqlDriver: sqlDriver,
		user:      user,
		password:  password,
		protocol:  protocol,
		address:   address,
		dataBase:  dataBase,
	}
	return s
}

func (s *MySqlSetting) DriverName() string {
	return s.sqlDriver
}

func (s *MySqlSetting) DataSourceName() string {
	return fmt.Sprintf("%s:%s@%s(%s)/%s",
		s.user,
		s.password,
		s.protocol,
		s.address,
		s.dataBase)
}
