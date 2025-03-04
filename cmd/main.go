package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/naveeshkumar24/pkg/database"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("unable to load env: %v", err)
	}
	conn := NewConnection()
	defer conn.DB.Close()
	server := &http.Server{
		Addr:    os.Getenv("PORT"),
		Handler: registerRouter(conn.DB),
	}

	query := database.NewStudentQuery(conn.DB)
	err := query.CreateStudentTables()
	if err != nil {
		log.Fatal("Unable to create database %v", err)
	}

	facultyquery := database.NewFacultyQuery(conn.DB)
	err = facultyquery.CreateFacultyTable()
	if err != nil {
		log.Fatal("Unable to create database %v", err)
	}
	wardenQuery := database.NewWardenQuery(conn.DB)
	err = wardenQuery.CreateWardenTable()
	if err != nil {
		log.Fatal("Unable to create database %v", err)
	}
	leaveformquery := database.NewLeaveQuery(conn.DB)
	err = leaveformquery.CreateLeaveFormTable()
	if err != nil {
		log.Fatal("Unable to create database %v", err)
	}
	log.Printf("server is running at port %s", os.Getenv("PORT"))
	err = server.ListenAndServe()
	if err != nil {
		log.Fatalf("unable to start the server: %v", err)
	}
}
