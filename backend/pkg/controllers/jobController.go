/** @file jobController.go
 * @brief This file contain all the functions to handle the job
 * @author Timothee de Boynes
 */

// @cond
package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"

	"AREA/pkg/jobs"
	"AREA/pkg/models"
	"AREA/pkg/utils"
)

// @endcond

/** @brief on a request, add a new job to a given user
 *
 * @param w http.ResponseWriter, r *http.Request
 */
func AddJobToUser(w http.ResponseWriter, r *http.Request) {
	newJob := &models.Job{}
	utils.ParseBody(r, newJob)

	requestUser, _ := GetUser(w, r)
	newJob.UserId = requestUser.ID
	newJob.ActionExecuted = false

	gitActions := utils.GetKeyFromMap(jobs.GitHubActions)
	if utils.ArrayContainsString(gitActions, newJob.ActionFunc) {
		CreateWebhook(requestUser.ID, jobs.GitHubActions[newJob.ActionFunc], newJob.ActionFuncParams)
	}

	jobId := newJob.CreateJob()
	userToken := models.FindUserToken(newJob.UserId)

	if newJob.ReactionService == "discord" {
		newJob.ReactionFuncParams = newJob.ReactionFuncParams + "@@@" + strconv.FormatUint(uint64(jobId), 10)
		models.SetDiscordWebhook(newJob.UserId, jobId, userToken.CurrentDiscordWebhookId, userToken.CurrentDiscordWebhookToken)
		models.UpdateJobField(jobId, "reaction_func_params", newJob.ReactionFuncParams)
	}

	jobs.AddJob(*newJob)
	res, _ := json.Marshal(jobId)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

/** @brief on a request, remove a job to a given user
 *
 * @param w http.ResponseWriter, r *http.Request
 */
func RemoveJob(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	jobId, err := strconv.ParseUint(vars["ID"], 10, 32)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	jobs.RemoveJobByID(uint(jobId))
	models.DeleteUserJob(uint(jobId))
}

/** @brief on a request, retrieve all a user's job
 *
 * @param w http.ResponseWriter, r *http.Request
 */
func GetUserJobs(w http.ResponseWriter, r *http.Request) {
	requestUser, _ := GetUser(w, r)
	jobs := models.GetJobsByUserId(requestUser.ID)

	res, _ := json.Marshal(jobs)
	w.Write(res)
}

/** @brief on a request, retrieve the actions and reaction if the user is connected to the service
 *
 * @param w http.ResponseWriter, r *http.Request
 */
func GetUserPropositions(w http.ResponseWriter, r *http.Request) {
	requestUser, _ := GetUser(w, r)
	FillServices()
	services := Services
	var servicesOptions []Service

	tokens := *models.FindUserToken(requestUser.ID)
	for _, service := range services {
		if service.Name == "email" && !models.CheckIfConnectedToService(tokens, "email") {
			continue
		}
		if service.Name == "discord" && !models.CheckIfConnectedToService(tokens, "discord") {
			continue
		}
		if service.Name == "spotify" && !models.CheckIfConnectedToService(tokens, "spotify") {
			continue
		}
		if service.Name == "github" && !models.CheckIfConnectedToService(tokens, "github") {
			continue
		}
		if service.Name == "deezer" && !models.CheckIfConnectedToService(tokens, "deezer") {
			continue
		}
		servicesOptions = append(servicesOptions, service)
	}
	res, _ := json.Marshal(servicesOptions)
	w.Write(res)
}
