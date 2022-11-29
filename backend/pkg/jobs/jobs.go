/** @file jobs.go
 * @brief This file contain all the functions to handle the actions and reactions of the Email API
 * @author Juliette Destang
 * 
 */

// @cond

package jobs

import (
	"AREA/pkg/models"
)

var currentJobs []models.Job

var GitHubActions = map[string]string{
	"Push action happened on your repository" : "push",
	"Pull request action happened on your repository": "pull_request",
	"If branch protection changed on your repository": "branch_protection_rule",
	"Checks run action on your repository": "check_run",
	"Checks suite action on your repository": "check_suite",
	"Someone created a git reference on your repository": "create",
	"Someone deleted a git reference on your repository": "delete",
	"Check deploy action on your repository": "deployment",
	"Check deploy status action on your repository": "deployment_status",
	"A discussion started on you repository": "discussion",
	"A comment was added to discussion on you repository": "discussion_comment",
	"Fork action on your repository": "fork",
	"Gollum action on your repository": "gollum",
	"Someone wrote an issue comment": "issue_comment",
	"Someone created an issue": "issues",
	"Someone created a label on your repository": "label",
	"A merge was perfomed on your repository": "merge_group",
	"A milestone was created on your repository" : "milestone",
	"Someone pushed on your publishing branch" : "page_build",
	"Created a project card on your repository" : "project_card",
	"Moved a card to a column in you repository" : "project_column",
	"Set your repoistory to public" : "public",
	"Add a comment to a pull request" : "pull_request_comment",
	"Someone reviewed your pull request" : "pull_request_review",
	"Added a comment to a review on some pull request" : "pull_request_review_comment",
	"Some activity was detected on a pull request" : "pull_request_target",
	"Some activity was detected on your package" : "registry_package",
	"Released action on your repository" : "release",
	"Dispatched action on your repository" : "repository_dispatch",
	"Added a schedule on your repository" : "schedule",
	"Changed the status of your repository" : "status",
	"Someone watched your repository" : "watch",
	"There was some activity on your workflow" : "workflow_call",
	"Dispatched action on your workflow" : "workflow_dispatch",
	"Ran your repository workflow" : "workflow_run",
}

var ActionMap = map[string]func(string) bool {
	"The temperature is over a given value": TemperatureIsOverN,
	"The temperature is under a given value": TemperatureIsUnderrN,
	"Check if the player main Teemo": IsPlayingTeemo,
	"The player winrate is over a given %": WinrateIsOverN,
	"The player KDA is over a given value": KDAIsOverN,
	"The covid cases are over a given number": CovidCaseIsOverN,
	"The covid critical cases are over a given number": CovidCriticalCaseIsOverN,
	"Play heads or tails": HeadsOrTails,
	"A choosen crypto is over a given number": CryptoIsOverN,
	"A choosen crypto is under a given number": CryptoIsUnderN,
}

var ReactionMap = map[string]func(uint, string) {
	"Adds a given song to the user's queue": AddSongToQueue,
	"Sends an email from user to given receiver": SendEmail,
	"Sends a webhook message on selected channel": SendMessage,
	"Adds a given song to the given playlist": AddSongToPlaylist,
}

// @endcond

/** @brief This function take a user id and activate his job on login
 * @param userID uint
 */
func AddUserJobsOnLogin(userId uint) {
	jobs := models.GetJobsByUserId(userId)
	currentJobs = append(currentJobs, jobs...)
}

/** @brief This function take a Job model and append a new job to the currentJob
 * @param newJob models.Job
 */
func AddJob(newJob models.Job) {
	currentJobs = append(currentJobs, newJob)
}

/** @brief Remove a given job to the currentJob
 * @param jobId uint
 */
func RemoveJobByID(jobId uint) {
	var newCurrentJobs []models.Job
	for _, job := range currentJobs {
		if (job.ID == jobId) {
			continue
		}
		newCurrentJobs = append(newCurrentJobs, job)
	}
	currentJobs = newCurrentJobs
}

/** @brief Remove* all job from the currentJob when a user logout
 * @param userID uint
 */
func SuprUserJobsOnLogout(userId uint) {
	var newCurrentJobs []models.Job

	for _, job := range currentJobs {
		if (job.UserId == userId) {
			continue
		}
		newCurrentJobs = append(newCurrentJobs, job)
	}
	
	currentJobs = newCurrentJobs
}

/** @brief Execute all active jobs each X seconds thanks to a crontab
 */
func ExecAllJob() {
	for index := range currentJobs {
		if ActionMap[currentJobs[index].ActionFunc] != nil && ActionMap[currentJobs[index].ActionFunc](currentJobs[index].ActionFuncParams) {
			if (currentJobs[index].ActionExecuted == false) {
				currentJobs[index].ActionExecuted = true
				ReactionMap[currentJobs[index].ReactionFunc](currentJobs[index].UserId, currentJobs[index].ReactionFuncParams)
			}
		} else {
			currentJobs[index].ActionExecuted = false
		}
	}
}

/** @brief On ping from github api, execute the réaction of a given gitAction
 * @param userID uint, githAction string
 */
func ExecGithJob(userID uint, githAction string) {
	for _, job := range currentJobs {
		if (githAction == GitHubActions[job.ActionFunc]) && (job.UserId == userID){
			ReactionMap[job.ReactionFunc](job.UserId, job.ReactionFuncParams)
		}
	}
}