/** @file githubController.go
 * @brief This file contain all the functions to communicate with the discord API
 * @author Juliette Destang
 * 
 */

// @cond
package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	"os"

	"github.com/tidwall/gjson"

	"AREA/pkg/jobs"
	"AREA/pkg/models"
	"AREA/pkg/utils"
	// "log"
)

// @endcond

/** @brief on a request, retrieve the github redirect url
 * @param w http.ResponseWriter, r *http.Request
 */
func GetGithubUrl(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(&w)
	githubID := utils.GetEnv("GITHUB_ID")
	res, _ := json.Marshal(fmt.Sprintf("https://github.com/login/oauth/authorize?client_id=%s&scope=admin:repo_hook repo&state=random", githubID))
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

/** @brief on a request, retrieve a token for the connected user
 * @param w http.ResponseWriter, r *http.Request
 */
func AuthGithub(w http.ResponseWriter, r *http.Request) {
	fmt.Println("redirect")
	url := fmt.Sprintf("https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s", utils.GetEnv("GITHUB_ID"), utils.GetEnv("GITHUB_SECRET"), r.FormValue("code"))

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res, _ := json.Marshal("bad request")
		w.Write(res)
	}
	req.Header.Add("Accept", "application/json")

	response, _ := client.Do(req)
	body, _ := ioutil.ReadAll(response.Body)
	accessToken := gjson.GetBytes(body, "access_token")
	requestUser, _ := GetUser(w, r)
	models.SetUserToken(requestUser.ID, "github_token", accessToken.String())
	http.Redirect(w, r, "http://localhost:8081/user/services", http.StatusSeeOther)
}

/** @brief on a request, when the user create an area with github, create a webhook that can handle the action of the user.
 * store the webhook id in the database
 * @param w http.ResponseWriter, r *http.Request
 */
func CreateWebhook(userID uint, action string, params string) {
	split := strings.Split(params, "@@@")
	username := split[0]
	repository := split[1]

	if (username == "") || repository == "" {
		return
	}
	if models.CheckExistingGitAction(userID, action) {
		fmt.Fprintln(os.Stderr, "webhook already exist")
		return
	}

	userToken := *models.FindUserToken(userID)

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/hooks", username, repository)

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	body := []byte(`{"name":"web","active":true,"events":["` + action + `"],"config":{"url":"` + utils.GetEnv("NGROK_REDIRECT") + `/webhook/` + `","content_type":"json","insecure_ssl":"0"}}`)

	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		fmt.Fprintln(os.Stderr, "erreur when trying to create webhook")
		return
	}
	req.Header.Add("Authorization", "token "+userToken.GithubToken)
	req.Header.Add("Accept", "application/vnd.github+json")

	response, _ := client.Do(req)
	newbody, _ := ioutil.ReadAll(response.Body)

	if (gjson.GetBytes(newbody, "message")).String() == "Bad credentials" {
		fmt.Fprintln(os.Stderr, "please re log to github")
		return
	}

	if (gjson.GetBytes(newbody, "message")).String() == "Validation Failed" {
		fmt.Fprintln(os.Stderr, "webhook already exist")
		return
	}

	webhookID := gjson.GetBytes(newbody, "id")
	models.SetGithubWebhook(userID, webhookID.String())
}

/** @brief this function is called when a ping webhook is send to the api
 * @param w http.ResponseWriter, r *http.Request
 */
func Webhook(w http.ResponseWriter, r *http.Request) {
	webhookID := r.Header.Get("X-Github-Hook-Id")
	webhookEvent := r.Header.Get("X-Github-Event")
	userToken := *models.FindUserByWebhookToken(webhookID)
	fmt.Println(userToken.UserId, webhookEvent)
	jobs.ExecGithJob(userToken.UserId, webhookEvent)
}
