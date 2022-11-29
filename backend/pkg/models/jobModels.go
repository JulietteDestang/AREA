/** @file jobModels.go
 * @brief This file contains all the functions to maipulate the jobs in the db
 * @author Juliette Destang
 * 
 */

// @cond
package models

import (
	"github.com/jinzhu/gorm"
)

type Job struct {
	gorm.Model
	UserId uint `json:"user_id"`
	ActionExecuted bool `json:"action_executed"`
	ActionService string `json:"action_service"`
	ActionFunc string `json:"action_func"`
	ActionFuncParams string `json:"action_func_params"`
	ReactionService string `json:"reaction_service"`
	ReactionFunc string `json:"reaction_func"`
	ReactionFuncParams string `json:"reaction_func_params"`
}

// @endcond

/** @brief Creates a new job and returns its id
 * @param job *Job
 * @return uint
 */
func (job *Job) CreateJob() uint {
	db.NewRecord(job)
	db.Create(&job)
	return job.ID
}

/** @brief Creates a new job and returns its id
 * @param job *Job
 * @return uint
 */
func GetJobById(Id uint) ([]Job){
	var jobs []Job
	db.Where("ID=?", Id).Find(&jobs)
	return jobs
}

/** @brief Returns all the jobs related to a User based on his ID
 * @param userId uint
 * @return []Job
 */
func GetJobsByUserId(userId uint) ([]Job){
	var jobs []Job
	db.Where("user_id=?", userId).Find(&jobs)
	return jobs
}

/** @brief Update a job column based on the given params
 * @param jobId uint, column string, value string
 */
func UpdateJobField(jobId uint, column string, value string) {
	db.Model(&Job{}).Where("ID = ?", jobId).Update(column, value)
}

/** @brief Deletes a job based on its ID
 * @param ID uint
 * @return Job
 */
func DeleteUserJob(ID uint) Job{
	var job Job
	db.Unscoped().Where("ID = ?", ID).Delete(&job)
	return job
}

/** @brief Deletes all jobs related to a User based on his ID 
 * @param userId uint
 * @return []Job
 */
func DeleteAllUserJob(userId uint) []Job{
	var job []Job
	db.Unscoped().Where("user_id = ?", userId).Delete(&job)
	return job
}

/** @brief Checks if a User already has a job on a Github action
 * @param id uint, action string
 * @return bool
 */
func CheckExistingGitAction(userId uint, action string) bool{
	var job Job
	db.Where("user_id = ?", userId).Where("action_func = ?", action).Find(&job)
	if (job.ActionFunc == "") {
		return false
	} else {
		return true
	}
}