/** @file discordJobs.go
 * @brief This file contain all the functions to handle the actions and reactions of the Discord API
 * @author Juliette Destang
 * 
 */

// @cond

package jobs

import (
	"fmt"
	"os"

	"github.com/DisgoOrg/disgohook"

	"AREA/pkg/models"
	"AREA/pkg/utils"
)

// @endcond

/** @brief this function take a user id and a message, and send a discord message thanks to webhook id
 * @param userID uint, params string
 */
func SendMessage(userID uint, params string) {
	paramsArr := utils.GetParams(params)
	if len(paramsArr) != 2 {
		fmt.Fprintln(os.Stderr, paramsArr, "params passed are not correct")
		return
	}

	userToken := *models.FindUserByDiscordWebhook(paramsArr[1])
	messageUrl := fmt.Sprintf("%s/%s", userToken.WebhookID, userToken.WebhookToken)
	webhook, _ := disgohook.NewWebhookClientByToken(nil, nil, messageUrl)

	webhook.SendContent(paramsArr[0])
}
