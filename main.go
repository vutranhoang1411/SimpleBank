package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/vutranhoang1411/SimpleBank/api"
	db "github.com/vutranhoang1411/SimpleBank/db/sqlc"
	"github.com/vutranhoang1411/SimpleBank/util"
	_ "github.com/golang/mock/mockgen/model"
)
const(

)
func main(){
	config,err:=util.LoadConfig("./");
	if (err!=nil){
		log.Fatal(err);
	}
	conn,err:=sql.Open(config.DBDriver,config.DBSource);
	if err!=nil{
		log.Fatal(err);
	}
	server,err:=api.NewServer(config,db.NewStore(conn))
	if err!=nil{
		log.Fatal(err);
	}
	err=server.Start(config.ServerAddress);
	if err!=nil{
		log.Fatal(err)
	}
	log.Print("Server running on port: ",config.ServerAddress);

}