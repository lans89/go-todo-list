package db

import (
	"api-todolist/conf"
	"database/sql" //libreria interfaz generica para base de datos
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq" //libreria implementacion de postgres
)

var connStr string = "postgres://%s:%s@%s:%d/%s?sslmode=disable"

var SqlDB *sql.DB = nil

func InitDB() {
	var err error = nil
	dbParams := conf.GetConfig().DB
	dbConnection := fmt.Sprintf(connStr, dbParams.Username, dbParams.Password, dbParams.Server, dbParams.Port, dbParams.DataBaseName)
	SqlDB, err = sql.Open("postgres", dbConnection)
	if err != nil {
		log.Fatal(err)
	}
	// Configurar el pool de conexiones
	SqlDB.SetMaxOpenConns(dbParams.MaxIdleConnections)                                 // Número máximo de conexiones abiertas a la base de datos
	SqlDB.SetMaxIdleConns(dbParams.MaxIdleConnections)                                 // Número máximo de conexiones inactivas en el pool
	SqlDB.SetConnMaxLifetime(time.Duration(dbParams.MaxLifetimeMin) * time.Minute)     // Tiempo máximo que una conexión puede ser reutilizada
	SqlDB.SetConnMaxIdleTime(time.Duration(dbParams.MaxIdleConnections) * time.Minute) // Tiempo máximo que una conexión puede estar inactiva

	// Verificar conexión
	err = SqlDB.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("¡Conectado exitosamente a la base de datos!")
}
