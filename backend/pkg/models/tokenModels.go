/** @file tokenModels.go
 * @brief This file contains all the functions to handle the tokens in the database
 * @author Juliette Destang
 * 
 */

 // @cond
package models

import (
	"github.com/jinzhu/gorm"
)

type GithubWebhook struct {
	gorm.Model
	UserId              uint     `json:"user_id"`
	WebhookID     		string   `json:"webhook_id"`
}

type DiscordWebhook struct {
	gorm.Model
	UserId              uint     `json:"user_id"`
	JobId               uint     `json:"job_id"`
	WebhookID     		string   `json:"webhook_id"`
	WebhookToken     	string   `json:"webhook_token"`
}

type Token struct {
	gorm.Model
	UserId              uint     `json:"user_id"`
	CurrentDiscordWebhookId		string   `json:"current_discord_webhook_id"`
	CurrentDiscordWebhookToken	string   `json:"current_discord_webhook_token"`
	SpotifyToken        string   `json:"spotify_token"`
	DeezerToken        	string   `json:"deezer_token"`
	SpotifyRefreshToken string   `json:"spotify_refresh_token"`
	Email               string   `json:"email"`
	EmailPassword       string   `json:"email_password"`
	GithubToken         string   `json:"github_token"`
}

// @endcond

/** @brief Creates a new token user
 * @param newToken *Token
 * @return *Token
 */
func (newToken *Token) CreateTokenUser() *Token {
	db.NewRecord(newToken)
	db.Create(&newToken)
	return newToken
}

/** @brief Find the user token by user ID
 * @param id uint
 * @return *Token
 */
func FindUserToken(id uint) *Token {
	var getToken Token
	db.Where("user_id = ?", id).Find(&getToken)
	return &getToken
}

/** @brief Find the user discord webhook token by user ID
 * @param id uint
 * @return *DiscordWebhook
 */
func FindUserByDiscordWebhook(id string) *DiscordWebhook {
	var getToken DiscordWebhook
	db.Where("job_id = ?", id).Find(&getToken)
	return &getToken
}

/** @brief This function check if the user is connected to a given service
 * @param token Token, service string
 * @return bool
 */ 
func CheckIfConnectedToService(token Token, service string) bool {
	returnValue := false
	switch service {
	case "email":
		if token.Email != "" {
			returnValue = true
			break
		}
	case "discord":
		if token.CurrentDiscordWebhookToken != "" {
			returnValue = true
			break
		}
	case "spotify":
		if token.SpotifyToken != "" {
			returnValue = true
			break
		}
	case "github":
		if token.GithubToken != "" {
			returnValue = true
			break
		}
	case "deezer":
		if token.DeezerToken != "" {
			returnValue = true
			break
		}
	}
	return returnValue
}

/** @brief This function set a given token to a user in the database
 * @param cookie string, column string, token string
 */ 
func SetUserToken(userId uint, column string, token string) {
	db.Model(&Token{}).Where("user_id = ?", userId).Update(column, token)
}

/** @brief This function find a user thanks to a given github webhook token
 * @param token string
 * @return *GithubWebhook
 */ 
func FindUserByWebhookToken(token string) *GithubWebhook {
	var getToken GithubWebhook
	db.Where("webhook_id = ?", token).Find(&getToken)
	return &getToken
}

/** @brief This function find a user thanks to a given github webhook token
 * @param token string
 * @return *GithubWebhook
 */ 
func SetGithubWebhook(userId uint, newWebhook string) {
	var newGithubWebhook GithubWebhook
	newGithubWebhook.UserId = userId
	newGithubWebhook.WebhookID = newWebhook
	db.Create(&newGithubWebhook)
	// db.Model(&Token{}).Where("user_id = ?", userId).Update("webhook_id", newWebhook)
}

/** @brief This function create a new raw with a user ID and a new webhook ID and webhook token
 * @param userId uint, newWebhookID string, newWebhookToken string
 */ 
func SetDiscordWebhook(userId uint, jobId uint, newWebhookId string, newWebhookToken string) {
	var newDiscordWebhook DiscordWebhook
	newDiscordWebhook.UserId = userId
	newDiscordWebhook.WebhookID = newWebhookId
	newDiscordWebhook.WebhookToken = newWebhookToken
	newDiscordWebhook.JobId = jobId
	db.Create(&newDiscordWebhook)
}

func UpdateDiscordWebhook(userId uint, newWebhookID string, newWebhookToken string) {
	SetUserToken(userId, "current_discord_webhook_id", newWebhookID)
	SetUserToken(userId, "current_discord_webhook_token", newWebhookToken)
}
