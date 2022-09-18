package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Connection struct {
	Conn *gorm.DB
}

func createConnection(cnfg *Config) *Connection {
	dsn := Resolve(cnfg)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return &Connection{
		Conn: db,
	}
}

func (c *Connection) Load() *Connection {
	cnfgInstance := Config{}
	cnfg := *cnfgInstance.bootCnfg()

	return createConnection(&cnfg)
}