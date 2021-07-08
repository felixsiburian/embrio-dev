// Connection to DB here without gorm

package db

import (
	"context"
	"database/sql"
	"embrio-dev/service/model/db"
	"embrio-dev/service/model/econst"
	"embrio-dev/service/tools"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"strconv"
	"time"
)

var availableDrivers = []string{"postgres"}

type DB struct {
	*sqlx.DB
}

func InitSQLServer(conn db.DBConn) (err error) {
	port, err := strconv.Atoi(conn.DBPort)
	if err != nil {
		log.Fatal(err.Error())
	}
	connStr := fmt.Sprintf("server=%s;user id=%s;port=%d;database=%s", conn.DBEngine, conn.DBPassword, port, conn.DBName)
	dbx, err := sql.Open("sqlserver", connStr)
	if err != nil {
		log.Fatal(err.Error())
	}

	ctx := context.Background()
	err = dbx.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("Connected \n")

	return nil
}

func InitPostgres() (err error) {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		err = errors.New(fmt.Sprintf("Can't convert to int : %+v", err))
		return err
	}

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), port, os.Getenv("DB_USER"), os.Getenv("DB_PWD"), os.Getenv("DB_NAME"))
	dbs, err := sql.Open("postgres", connStr)
	if err != nil {
		err = errors.New(fmt.Sprintf("Can't connect to DB : %+v", err))
		return err
	}
	defer dbs.Close()

	err = dbs.Ping()
	if err != nil {
		err = errors.New(fmt.Sprintf("Can't PING DB : %+v", err))
		return err
	}
	log.Println("We are now connected")
	return nil
}

func NewConnection(driver string, connectionString string, connLifeTime int64) (*DB, error) {
	if tools.CheckArr(availableDrivers, driver) == false {
		return nil, errors.New("driver not available")
	}

	db, err := sqlx.Connect(driver, connectionString)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("db : failed to connect to database %+v", err))
	}

	durationLifetime, err := strconv.Atoi(tools.Env(econst.DBConnLifetime, "15"))
	if err != nil {
		err = errors.New("while convert to int")
		return nil, err
	}

	durationIdleCons, err := strconv.Atoi(tools.Env(econst.DBConnMaxIdle, "5"))
	if err != nil {
		err = errors.New("while convert to int")
		return nil, err
	}

	durationMaxCons, err := strconv.Atoi(tools.Env(econst.DBConnMaxOpen, "0"))
	if err != nil {
		err = errors.New("while convert to int")
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * time.Duration(int64(durationLifetime)))
	db.SetMaxIdleConns(durationIdleCons)
	db.SetMaxOpenConns(durationMaxCons)

	return &DB{db}, nil
}

func NewPostgreConnection(connectionString string, connLifeTime int64) (*DB, error) {
	return NewConnection("postgres", connectionString, connLifeTime)
}
