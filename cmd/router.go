package main

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/naveeshkumar24/internals/handlers"
	"github.com/naveeshkumar24/internals/middleware"
	"github.com/naveeshkumar24/repository"
)

func registerRouter(db *sql.DB) *mux.Router {
	router := mux.NewRouter()
	router.Use(middleware.CorsMiddleware)

	Studentrepo := repository.NewStudentLoginRepository(db)
	StudentPoHandler := handlers.NewStudentLoginHandler(*Studentrepo)
	Faculturepo := repository.NewFacultyLoginRepository(db)
	FacultyPoHandler := handlers.NewFacultyLoginHandler(*Faculturepo)
	Wardenrepo := repository.NewWardenLoginRepository(db)
	WardenPoHandler := handlers.NewWardenLoginHandler(*Wardenrepo)
	Leaveformrepo := repository.NewLeaveFormRepository(db)
	LeaveformPoHandler := handlers.NewLeaveFormHandler(*Leaveformrepo)
	router.HandleFunc("/leaveformsubmit", LeaveformPoHandler.SubmitLeaveForm).Methods("POST")
	router.HandleFunc("/leaveformgetpendingleave", LeaveformPoHandler.GetPendingLeaves).Methods("GET")
	router.HandleFunc("/leaveformupdateleaveform", LeaveformPoHandler.UpdateLeaveStatus).Methods("PUT")
	router.HandleFunc("/leaveformsanctionedleaveform", LeaveformPoHandler.GetSanctionedLeaves).Methods("GET")
	router.HandleFunc("/wardenregister", WardenPoHandler.SubmitWardenRegisterForm).Methods("POST")
	router.HandleFunc("/wardenlogin", WardenPoHandler.SubmitWardenLoginForm).Methods("POST")
	router.HandleFunc("/facultyregister", FacultyPoHandler.SubmitFacultyRegisterForm).Methods("POST")
	router.HandleFunc("/facultylogin", FacultyPoHandler.SubmitFacultyLoginForm).Methods("POST")
	router.HandleFunc("/studentregister", StudentPoHandler.SubmitStudentRegisterForm).Methods("POST")
	router.HandleFunc("/studentlogin", StudentPoHandler.SubmitStudentLogin).Methods("POST")
	return router
}
