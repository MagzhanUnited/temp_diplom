package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/umangraval/Go-Mongodb-REST-boilerplate/controllers"
	middlewares "github.com/umangraval/Go-Mongodb-REST-boilerplate/handlers"
)

// Routes -> define endpoints
func Routes() *mux.Router {
	router := mux.NewRouter()

	// Create a subrouter for /api routes
	apiRouter := router.PathPrefix("/api").Subrouter()

	apiRouter.HandleFunc("/person", controllers.CreatePersonEndpoint).Methods("POST")
	apiRouter.HandleFunc("/residentialas/{category}", controllers.GetResidentialas).Methods("GET")
	apiRouter.HandleFunc("/residential", controllers.CreateResidential).Methods("POST")
	apiRouter.HandleFunc("/order", controllers.CreateOrder).Methods("POST")
	apiRouter.HandleFunc("/residential/{id}", controllers.DeleteResidentEndpoint).Methods("DELETE")
	apiRouter.HandleFunc("/residential/{id}", controllers.GetResidentEndpoint).Methods("GET")
	apiRouter.HandleFunc("/auth", controllers.Auths).Methods("POST")
	apiRouter.HandleFunc("/people", middlewares.IsAuthorized(controllers.GetPeopleEndpoint)).Methods("GET")
	apiRouter.HandleFunc("/person/{id}", controllers.GetPersonEndpoint).Methods("GET")
	apiRouter.HandleFunc("/person/{id}", controllers.DeletePersonEndpoint).Methods("DELETE")
	apiRouter.HandleFunc("/person/{id}", controllers.UpdatePersonEndpoint).Methods("PUT")
	apiRouter.HandleFunc("/upload", controllers.UploadFileEndpoint).Methods("POST")
	apiRouter.PathPrefix("/static/").Handler(http.StripPrefix("/api/static/", http.FileServer(http.Dir("./uploaded/"))))

	return router
}
