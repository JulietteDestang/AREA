/** @file models.go
 * @brief This file contain all the functions to handle the user model in the database
 * @author Juliette Destang
 * 
 */

// @cond
package models

import (
	"github.com/jinzhu/gorm"
	"AREA/pkg/config"
)

var db * gorm.DB


type User struct {
	gorm.Model
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Email string `json:"email"`
	Password []byte `json:"password"`
}
// @endcond

/** @brief Initialize the tables of the databases thanks to model struct
 */
func init() {
	config.Connect()
	db = config.GetDb()
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Token{})
	db.AutoMigrate(&Job{})
	db.AutoMigrate(&GithubWebhook{})
	db.AutoMigrate(&DiscordWebhook{})
}

/** @brief Create a new raw with a new user in the database
 * @param newUser *User
 * @return *User
 */
func (newUser *User) CreateUser() *User{
	db.NewRecord(newUser)
	db.Create(&newUser)
	return newUser
}

/** @brief Find a user in the database thanks to a given email
 * @param Email string
 * @return *User
 */
func FindUser(Email string) *User{
	var getUser User
	db.Where("email = ?", Email).Find(&getUser)
	return &getUser
}

/** @brief Find the id of a user thanks to a given email
 * @param Email string
 * @return *User
 */
func FindUserID(Email string) *uint{
	var getUser User
	db.Where("email = ?", Email).Find(&getUser)
	return &getUser.ID
}

/** @brief Retrieve all the users present in the database
 * @return []User
 */
func GetAllUsers() []User{
	var Users []User
	db.Find(&Users)
	return Users
}

/** @brief Retrieve a user thanks to a given ID
 * @param Id int64
 * @return *User, *gorm.DB
 */
func GetUserById(Id int64) (*User, *gorm.DB){
	var getUser User
	db:=db.Where("ID=?", Id).Find(&getUser)
	return &getUser, db
}

/** @brief Delete a user thanks to a given ID
 * @param ID int64
 * @return User
 */
func DeleteUser(ID int64) User{
	var User User
	db.Where("ID=?", ID).Delete(User)
	return User
}
