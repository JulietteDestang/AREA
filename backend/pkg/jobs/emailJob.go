/** @file emailJob.go
 * @brief This file contain all the functions to handle the actions and reactions of the Email API
 * @author Juliette Destang
 * 
 */

// @cond

package jobs

import (
	"fmt"
	"net/smtp"
	"os"

	"AREA/pkg/models"
	"AREA/pkg/utils"
)

// @endcond

/** @brief this function take a user id and a message, and send an email to a receiver
 * @param userID uint, params string
 */
func SendEmail(userID uint, params string) {
	paramsArr := utils.GetParams(params)
	if len(paramsArr) != 2 {
		fmt.Fprintln(os.Stderr, paramsArr, "params passed are not correct")
		return
	}
	receiver := paramsArr[1]
	message := paramsArr[0]

	requestUser := *models.FindUserToken(userID)
	from := requestUser.Email
	password := requestUser.EmailPassword

	to := []string{
		receiver,
	}
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	messagePayload := []byte(message)

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, messagePayload)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}
