/** @file deezerController.go
 * @brief Oauth and reactions for Deezer
 * @author Erwan
 */

// @cond

package controllers

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/tidwall/gjson"

	"AREA/pkg/models"
	"AREA/pkg/utils"
)

// @endcond

/** @brief Returns the current user's ID
 *
 * @param[string] accessToken
 *
 * @return int64 userId
 */
func GetUserId(accessToken string) int64 {
	Url := "https://api.deezer.com/user/me?access_token=" + accessToken
	resp, err := http.Get(Url)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return -1
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, "client: could not read response body:", err)
		return -1
	}
	id := gjson.GetBytes(body, "id")
	idConverted := id.Int()
	return (idConverted)
}

/** @brief Connects the user with OAuth2
 *
 * Connexion via OAuth2 and then putting the connexion token inside the database,
 * the user is then redirected to the services page
 */
func AuthDeezer(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["code"]

	if !ok || len(keys[0]) < 1 {
		fmt.Fprintln(os.Stderr, "Url Param 'code' is missing")
	}
	code := keys[0]

	authUrl := "https://connect.deezer.com/oauth/access_token.php"
	data := url.Values{}
	data.Set("app_id", "564742")
	data.Set("secret", "7fd8db961fb009848532ca1c901bbc76")
	data.Set("code", string(code))
	encodedData := data.Encode()
	newUrl := authUrl + "?" + encodedData

	resp, err := http.Get(newUrl)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, "client: could not read response body:", err)
		return
	}
	accessToken := strings.Split(string(body), "=")[1]
	accessToken = strings.Split(accessToken, "&")[0]

	requestUser, _ := GetUser(w, r)
	models.SetUserToken(requestUser.ID, "deezer_token", accessToken)
	http.Redirect(w, r, "http://localhost:8081/user/services", http.StatusSeeOther)
}

/** @brief Returns the URL for OAuth2
 *
 * It has the full URL to ask for permissions and the redirection URL asked by the frontend
 */
func GetDeezerUrl(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(&w)
	deezerID := utils.GetEnv("DEEZER_APP_ID")
	json.Marshal(html.EscapeString("https://connect.deezer.com/oauth/auth.php?app_id=" + deezerID + "&redirect_uri=http://localhost:8080/deezer/auth&perms=basic_access,email,offline_access,manage_library,listening_history"))
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "https://connect.deezer.com/oauth/auth.php?app_id=%s&redirect_uri=http://localhost:8080/deezer/auth&perms=basic_access,email,offline_access,manage_library,listening_history", deezerID)
}
