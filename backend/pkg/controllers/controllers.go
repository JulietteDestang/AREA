/** @file controllers.go
 * @brief This file contain all the functions that log, log out or modify a User
 * @author Juliette Destang
 */

// @cond

package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"

	"AREA/pkg/config"
	"AREA/pkg/jobs"
	"AREA/pkg/models"
	"AREA/pkg/utils"
)

var SecretKey = utils.GetEnv("RAPID_API_KEY")

//@endcond

/** @brief Gets all the users from the database
 * @param w http.ResponseWriter, r *http.Request
 */
func GetAllUsers(w http.ResponseWriter, r *http.Request){
	newUsers:=models.GetAllUsers()
	res, _ :=json.Marshal(newUsers)
	utils.EnableCors(&w)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

/** @brief on a request, get the users by id from the database with the vars "userID"
 * @param w http.ResponseWriter, r *http.Request
 */
func GetUserById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	userId := vars["userID"]
	ID, err:= strconv.ParseInt(userId,0,0)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error while parsing", err)
	}
	userDetails, _:= models.GetUserById(ID)
	res, _ := json.Marshal(userDetails)
	utils.EnableCors(&w)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

/** @brief on a request, create new user in the database with an encrypted password
 * @param w http.ResponseWriter, r *http.Request
 */
func CreateUser(w http.ResponseWriter, r *http.Request) {
	NewUser := &models.User{}
	utils.ParseBody(r, NewUser)

	sameUser := models.FindUser(NewUser.Email)
	if (sameUser.Email != "") {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(NewUser.Password), 14)

	NewUser.Password = password
	b := NewUser.CreateUser()
	NewUserToken := &models.Token{}
	NewUserToken.UserId = b.ID
	NewUserToken.CreateTokenUser()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

/** @brief on a request, log the user retrieved in the request.
 * this function check if the password match the password of the user in the database,
 * then create a new json web token to identify the user trought his use of AREA.
 * A cookie is set with the json web token
 * @param w http.ResponseWriter, r *http.Request
 */
func LoginUser(w http.ResponseWriter, r *http.Request) {
	LoginUser := &models.User{}
	utils.ParseBody(r, LoginUser)

	user := *models.FindUser(LoginUser.Email)

	if (user.Email == "") {
		fmt.Println("bad email")
		res, _ := json.Marshal("bad email")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return
	}
	err := bcrypt.CompareHashAndPassword(user.Password, []byte(LoginUser.Password))
	if (err != nil) {
		fmt.Println("not hash")
		w.WriteHeader(http.StatusBadRequest)
		res, _ := json.Marshal("bad password")
		w.Write(res)
		return
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		Path:     "/",
		HttpOnly: false,
	}

	jobs.AddUserJobsOnLogin(user.ID)
	http.SetCookie(w, cookie)
	res, _ := json.Marshal("sucess")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

/** @brief delete a user in the database with a given ID
 * @param w http.ResponseWriter, r *http.Request
 */
func DeleteUser(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	user := models.DeleteUser(ID)
	res, _ := json.Marshal(user)
	utils.EnableCors(&w)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

/** @brief update the user with a given ID. This fuction can change the firstname, lastname and email of the user
 * @param w http.ResponseWriter, r *http.Request
 */
func UpdateUser(w http.ResponseWriter, r *http.Request){
	var updateUser = &models.User{}
	utils.ParseBody(r, updateUser)
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0,0)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error while parsing")
		return
	}
	userDetails, db:=models.GetUserById(ID)
	if updateUser.Firstname != ""{
		userDetails.Firstname = updateUser.Firstname
	}
	if updateUser.Lastname != ""{
		userDetails.Lastname = updateUser.Lastname
	}
	if updateUser.Email != ""{
		userDetails.Email = updateUser.Email
	}
	db.Save(&userDetails)
	res, _ := json.Marshal(userDetails)
	utils.EnableCors(&w)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

/** @brief log out a user with a given ID. The jwt token set in the cookie is removed.
 * @param w http.ResponseWriter, r *http.Request
 */
func Logout(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		Path:     "/",
		HttpOnly: false,
	}
	requestUser, _ := GetUser(w, r)
	http.SetCookie(w , cookie)
	jobs.SuprUserJobsOnLogout(requestUser.ID)
	res, _ := json.Marshal("sucess")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

/** @brief get the user data from the jwt cookie. This function parse the jwt of the user and retrieve 
 * a model of this user.
 * @param w http.ResponseWriter, r *http.Request
 */
func GetUser(w http.ResponseWriter, r *http.Request) (models.User, error) {
	cookie, cookieErr := r.Cookie("jwt")
	var user models.User
	if (cookieErr != nil) {
		w.WriteHeader(http.StatusBadRequest)
		return user, nil
	}
	token, err := jwt.ParseWithClaims(cookie.Value, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return user, err
	}

	claims := token.Claims.(*jwt.StandardClaims)
	config.GetDb().Where("id = ?", claims.Issuer).First(&user)
	return user, nil
}

/** @brief Sets the CORS for the front
 * @param next http.HandlerFunc
 * @return http.HandlerFunc
 */
func CORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	  w.Header().Add("Access-Control-Allow-Origin", "http://localhost:8081")
	  w.Header().Add("Access-Control-Allow-Credentials", "true")
	  next(w, r)
	}
}