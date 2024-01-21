package connection

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)
  
var (
	DB  *sql.DB
	err error
)
func ConnectPostgresDB()(*sql.DB, error) {
	url := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSL"),
	)

	DB, err := sql.Open("postgres", url)
	  if err != nil {
		  panic(err)
	  }
	  _,err = DB.Exec(`CREATE TABLE IF NOT EXISTS "product" (
		"id" UUID NOT NULL,
		"name" VARCHAR(255) NOT NULL,
		"description" VARCHAR(255) NOT NULL,
		"priceInCents" INTEGER NOT NULL,
		"stock" INTEGER NOT NULL,
		"createdAt" DATE NOT NULL,
		"updatedAt" DATE NULL,
		"isActive" BOOLEAN NOT NULL,
		PRIMARY KEY("id")
	);
	
	CREATE TABLE IF NOT EXISTS "service" (
		"id" UUID NOT NULL,
		"name" VARCHAR(255) NOT NULL,
		"description" VARCHAR(255) NOT NULL,
		"priceInCents" BIGINT NOT NULL,
		"hasDate" BOOLEAN NOT NULL,
		"date" DATE NOT NULL,
		"isActive" BOOLEAN NOT NULL,
		"createdAt" DATE NOT NULL,
		"updatedAt" DATE NULL,
		PRIMARY KEY("id")
	);
	
	CREATE TABLE IF NOT EXISTS "user" (
		"id" UUID NOT NULL,
		"name" VARCHAR(255) NOT NULL,
		"email" VARCHAR(255) NOT NULL,
		"password_hash" VARCHAR(255) NOT NULL,
		"role" VARCHAR(255) NOT NULL,
		PRIMARY KEY("id")
	);
	
	CREATE TABLE IF NOT EXISTS "reservation" (
		"id" UUID NOT NULL,
		"createdAt" DATE NOT NULL,
		"DateOfReservation" DATE NOT NULL,
		"ReservedId" UUID NOT NULL,
		"tableNumber" INTEGER NOT NULL,
		"createdId" UUID NOT NULL,
		"isFinished" BOOLEAN NOT NULL,
		"isActiveNow" BOOLEAN NOT NULL,
		PRIMARY KEY("id"),
		FOREIGN KEY("ReservedId") REFERENCES "service"("id"),
		FOREIGN KEY("createdId") REFERENCES "user"("id")
	);
	
	CREATE TABLE IF NOT EXISTS "client" (
		"id" UUID NOT NULL,
		"name" VARCHAR(255) NOT NULL,
		"fone" INTEGER NULL,
		"creditInCents" INTEGER NOT NULL,
		"cpf" INTEGER NULL,
		"isActive" BOOLEAN NOT NULL,
		"createdAt" DATE NOT NULL,
		PRIMARY KEY("id")
	);
	
	ALTER TABLE "reservation" ADD CONSTRAINT "reservation_clientid_foreign" FOREIGN KEY("id") REFERENCES "client"("id");
	`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Tables created successfully!")
	  err = DB.Ping()
	  return DB, err
	  
}
