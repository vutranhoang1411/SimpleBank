package db

import (
	"database/sql"
	"log"
	"os"
	"testing"
)
const(
	dbDriver="postgres"
	dbSource="postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

func TestMain(m *testing.M){
	conn,err:=sql.Open(dbDriver,dbSource)
	if err!=nil{
		log.Fatal("Cannot connect to db:",err)
	}
	_=New(conn)

	os.Exit(m.Run());
}