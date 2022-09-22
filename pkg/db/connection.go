package db

import (
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var lock = &sync.Mutex{}

type Connection struct {
	Conn *gorm.DB
}

var connectionInstance *Connection


func GetInstance() *gorm.DB {
    if connectionInstance == nil {
        lock.Lock()
        defer lock.Unlock()
        if connectionInstance == nil {
			// create database instance
			connection := Connection{}
            connectionInstance = connection.load()
        }
    }

    return connectionInstance.Conn
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

func (c *Connection) load() *Connection {
	cnfgInstance := Config{}
	cnfg := *cnfgInstance.bootCnfg()

	return createConnection(&cnfg)
}