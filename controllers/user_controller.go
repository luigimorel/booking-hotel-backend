package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/morelmiles/booking-backend/config"
	"github.com/morelmiles/booking-backend/middleware"
	"github.com/morelmiles/booking-backend/models"
	"golang.org/x/crypto/bcrypt"
)

func ComparePassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func GetUsers(w http.ResponseWriter, r *http.Request) {

	var users []models.User

	config.DB.Find(&users)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&users)
}

// GetUserById - Fetches a list of all users.
func GetUserById(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["id"]
	if !checkIfUserExists(userId) {
		json.NewEncoder(w).Encode("user not found!")
		return
	}

	user := models.User{}

	config.DB.First(&user, userId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func checkIfUserExists(userId string) bool {
	var user models.User
	config.DB.First(&user, userId)

	return user.ID != 0
}

// CreateUser - Creates a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		middleware.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		middleware.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	user.BeforeSave()
	err = user.Validate("")
	if err != nil {
		middleware.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	json.NewDecoder(r.Body).Decode(&user)

	// Fix this middleware
	if _, err := middleware.HashPassword(user.Password); err != nil {
		log.Printf("Error: %d", err)
		return
	}

	newUser := config.DB.Create(&user)
	err = newUser.Error

	if err != nil {
		log.Panic(err)
	} else {
		json.NewEncoder(w).Encode(&user)
	}
}

// UpdateUserById -  Updates a single user by the ID specified
func UpdateUserById(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["id"]
	if !checkIfUserExists(userId) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("user not found!")
		return
	}

	var user models.User

	config.DB.First(&user, userId)
	json.NewDecoder(r.Body).Decode(&user)
	config.DB.Save(&user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// DeleteUserById - Updates a single user by the ID specified.
func DeleteUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userId := mux.Vars(r)["id"]
	if !checkIfUserExists(userId) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("user not found!")
		return
	}

	var user models.User
	config.DB.Delete(&user, userId)
	json.NewEncoder(w).Encode(user)
}

// Login - Allows the user to login
func Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		middleware.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		middleware.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()

	err = user.Validate("login")
	if err != nil {
		middleware.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := SignIn(user.Email, user.Password)
	if err != nil {
		formattedError := middleware.FormatError(err.Error())
		middleware.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	middleware.JSON(w, http.StatusOK, token)

}

// Sign in - enables the user to sign in into their account
func SignIn(email, password string) (string, error) {
	var err error

	user := models.SignInInput{}
	err = config.DB.Debug().Model(models.User{}).Where("email=?", email).Take(&user).Error

	if err != nil {
		return " ", err
	}

	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return middleware.CreateToken(user.ID)
}

// GetAllPropertiesByUser fetches all properties owned by user.
// Figure own how to write this on Gorm - select * from users join properties on users.id=properties.owner_id;
func GetAllPropertiesByUser(w http.ResponseWriter, r *http.Request) {
	var users models.User
	config.DB.Raw("SELECT * FROM users JOIN properties ON users.id = properties.owner_id")

	json.NewEncoder(w).Encode(users)
}
