/** @file discordController.go
 * @brief This file contain all the functions to communicate with the discord API
 * @author Juliette Destang
 * 
 */

// @cond

package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
	"encoding/json"

	"github.com/tidwall/gjson"
	
	"AREA/pkg/utils"
	"AREA/pkg/models"
)

// @endcond

/** @brief callback of the discord API. This function retrieve a token for the connected user and store it in the database.
 * @param w http.ResponseWriter, r *http.Request
 */
func AuthDiscord(w http.ResponseWriter, r *http.Request){

	authUrl := "https://discordapp.com/api/v6/oauth2/token";

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	
	data := url.Values{}
	data.Set("client_id", utils.GetEnv("DISCORD_CLIENT_ID"))
	data.Set("client_secret", utils.GetEnv("DISCORD_CLIENT_SECRET"))
	data.Set("grant_type", "authorization_code")
	data.Set("redirect_uri", "http://localhost:8080/discord/auth")
	data.Set("scope", "webhook.incoming")
	data.Set("code", r.FormValue("code"))
	encodedData := data.Encode()

	req, err := http.NewRequest("POST", authUrl, strings.NewReader(encodedData))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res, _ := json.Marshal("bad request")
		w.Write(res)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	response, _ := client.Do(req)
	body, _ := ioutil.ReadAll(response.Body)

	message := gjson.GetBytes(body, "message")

	if (message.String() == "Maximum number of webhooks reached (10)") {
		http.Redirect(w, r, "http://localhost:8081/user/services", http.StatusSeeOther)
		return
	}

	webhookId := gjson.GetBytes(body, "webhook.id").String()
	webhookToken := gjson.GetBytes(body, "webhook.token").String()
	requestUser, _ := GetUser(w, r)
	models.UpdateDiscordWebhook(requestUser.ID, webhookId, webhookToken)

	http.Redirect(w, r, "http://localhost:8081/user/services", http.StatusSeeOther)
}

/** @brief on a request, retrieve the discord redirect url
 * @param w http.ResponseWriter, r *http.Request
 */
func GetDiscordUrl(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(&w)
	discordID := utils.GetEnv("DISCORD_CLIENT_ID");
	res, _ := json.Marshal(fmt.Sprintf("https://discord.com/api/oauth2/authorize?client_id=%s&redirect_uri=http://localhost:8080/discord/auth&response_type=code&scope=webhook.incoming&permissions=536870912", discordID))
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}