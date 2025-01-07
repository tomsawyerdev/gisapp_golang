package database

import (
	"context"
	"fmt"
	//"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

// var DB *pgx.Conn
var DB *pgxpool.Pool

//var CTX context.Context

// https://stackoverflow.com/questions/31218008/sharing-a-globally-defined-db-conn-with-multiple-packages

func ConfigureDb() {
	var err error

	/*
		user:=   os.Getenv("MYSQL_USER")
		password:= os.Getenv("MYSQL_PASSWORD")
		host:=   os.Getenv("MYSQL_HOST")
		port:=   os.Getenv("MYSQL_PORT")
		database:= os.Getenv("MYSQL_DATABASE")

	*/

	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Africa/Lagos", host, username, password, databaseName, port)
	// fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	const (
		user     = "gisuser"
		password = "gisuser"
		host     = "localhost"
		port     = 5432
		database = "gisapp"
	)

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, database)

	fmt.Println("Database DSN:", dsn)
	//CTX = context.Background()
	//DB, err = pgx.Connect(context.Background(), dsn)
	DB, err = pgxpool.New(context.Background(), dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		//panic(err)
		os.Exit(1)
	}

	fmt.Println("ConfigureDb Ping:", DB.Ping(context.Background()) == nil)

}

// Avoid Shadow https://stackoverflow.com/questions/26125143/invalid-memory-address-error-when-running-postgres-queries
