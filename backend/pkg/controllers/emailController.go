/** @file emailController.go
 * @brief This file contain a functions for the email API
 * @author Juliette Destang
 */

// @cond

package controllers

import (
	"net/http"
	"io"

	"AREA/pkg/models"
	"github.com/tidwall/gjson"
)

// @endcond

/** @brief on a request, store the given email and password to the database.
 * @param w http.ResponseWriter, r *http.Request
 */
func AuthEmail(w http.ResponseWriter, r *http.Request) {
	b, _ := io.ReadAll(r.Body)
    email := gjson.GetBytes(b, "email")
    password := gjson.GetBytes(b, "password")
	requestUser, _ := GetUser(w, r)
	models.SetUserToken(requestUser.ID, "email", email.String())
	models.SetUserToken(requestUser.ID, "email_password", password.String())
	w.WriteHeader(http.StatusOK)
}