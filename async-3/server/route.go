package server

import (
	"async-3/server/controllers"
	"database/sql"
	"fmt"
	"net/http"
)

func StartServer(router *http.ServeMux, port string, db *sql.DB) {
	buildRoute(router, db)

	fmt.Println("Server running at", port)
	http.ListenAndServe(port, router)
}

func buildRoute(router *http.ServeMux, db *sql.DB) {
	homeRoute(router)
	employeeRoute(router, db)
}

func employeeRoute(router *http.ServeMux, db *sql.DB) {

	employeeController := controllers.NewEmployeeController(db)

	router.HandleFunc("/employees", employeeController.Index)
	router.HandleFunc("/employees/update", employeeController.UpdateByID)
	router.HandleFunc("/employees/add", employeeController.Add)
	router.HandleFunc("/employees/delete", employeeController.DeleteByID)
}

func homeRoute(router *http.ServeMux) {
	homeController := controllers.NewHomeController()

	router.HandleFunc("/", homeController.Index)

}
