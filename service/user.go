package service

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"sync"

	"lumosh_klinik/model"
)

var (
	users      []model.User
	userMutex  sync.Mutex
	userFile   = "users.json"
	sessionMap = make(map[string]bool)
)

func LoadUsers() error {
	data, err := ioutil.ReadFile(userFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	return json.Unmarshal(data, &users)
}

func SaveUsers() error {
	data, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(userFile, data, 0644)
}

func Login(username, password string) (model.User, error) {
	for _, user := range users {
		if user.Username == username && user.Password == password {
			sessionMap[username] = true
			return user, nil
		}
	}
	return model.User{}, errors.New("username atau password salah")
}

func Logout(username string) {
	delete(sessionMap, username)
}

func IsSessionValid(username string) bool {
	_, exists := sessionMap[username]
	return exists
}

func AddUser(user model.User) error {
	userMutex.Lock()
	defer userMutex.Unlock()
	users = append(users, user)
	return SaveUsers()
}

func GetUsers() []model.User {
	return users
}
