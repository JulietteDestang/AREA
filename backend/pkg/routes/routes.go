/** @file routes.go
 * @brief This file contains the AreaRouter, used to handle all our API endpoints
 * @author Juliette Destang
 * 
 */

 // @cond
package routes

import (
	"github.com/gorilla/mux"

	"AREA/pkg/controllers"
)

var AreaRouter = func(router *mux.Router) {
	router.HandleFunc("/about.json", controllers.CORS(controllers.GetAboutJson)).Methods("GET")

	router.HandleFunc("/register/", controllers.CORS(controllers.CreateUser)).Methods("POST")
	router.HandleFunc("/login/", controllers.CORS(controllers.LoginUser)).Methods("POST")
	router.HandleFunc("/login/", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/logout/", controllers.CORS(controllers.Logout)).Methods("GET")
	router.HandleFunc("/login/{userID}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/login/{userID}", controllers.DeleteUser).Methods("DELETE")

	router.HandleFunc("/webhook/", controllers.Webhook).Methods("POST")

	router.HandleFunc("/discord/auth", controllers.AuthDiscord).Methods("GET")
	router.HandleFunc("/discord/auth/url", controllers.GetDiscordUrl).Methods("GET")

	router.HandleFunc("/github/auth/url", controllers.GetGithubUrl).Methods("GET")
	router.HandleFunc("/github/auth", controllers.AuthGithub).Methods("GET")

	router.HandleFunc("/spotify/auth/url", controllers.CORS(controllers.GetSpotifyUrl)).Methods("GET")
	router.HandleFunc("/spotify/auth", controllers.AuthSpotify).Methods("GET")
	
	router.HandleFunc("/deezer/auth/url", controllers.CORS(controllers.GetDeezerUrl)).Methods("GET")
	router.HandleFunc("/deezer/auth", controllers.AuthDeezer).Methods("GET")
	
	router.HandleFunc("/email/login", controllers.CORS(controllers.AuthEmail)).Methods("POST")

	router.HandleFunc("/area/user/areas", controllers.CORS(controllers.GetUserJobs)).Methods("GET")
	router.HandleFunc("/area/user/propositions", controllers.CORS(controllers.GetUserPropositions)).Methods("GET")

	router.HandleFunc("/area/get", controllers.CORS(controllers.GetUserJobs)).Methods("GET")
	router.HandleFunc("/area/create", controllers.CORS(controllers.AddJobToUser)).Methods("POST")
	router.HandleFunc("/area/delete/{ID}", controllers.CORS(controllers.RemoveJob)).Methods("GET")
}

// @endcond