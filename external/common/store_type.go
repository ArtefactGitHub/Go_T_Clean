package common

type StoreType string

const (
	Memory StoreType = "memory"
	MySql  StoreType = "mysql"
)

func (t StoreType) IsMemory() bool {
	return t == Memory
}

func (t StoreType) IsMySql() bool {
	return t == MySql
}
