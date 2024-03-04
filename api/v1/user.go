package v1

import (
	"encoding/json"
	"github.com/ChristianHamm/stopwatch/internal/model"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

func HttpOk(writer http.ResponseWriter, payload []byte) {
	writer.Header().Add("Content-Type", "application/json")
	//writer.Header().Add("Access-Control-Allow-Origin", "*")
	_, _ = writer.Write(payload)
}

func ListUsers(writer http.ResponseWriter, request *http.Request) {
	res, _ := json.Marshal(&model.UserStore)
	HttpOk(writer, res)
}

func AddUser(writer http.ResponseWriter, request *http.Request) {
	newUser := model.User{}

	if err := json.NewDecoder(request.Body).Decode(&newUser); err != nil {
		http.Error(writer, "Could not unmarshal JSON user", http.StatusBadRequest)
		return
	}

	newUser.Id = model.FindMaxId()
	newUser.Speaking = false
	newUser.SpeakDuration = 0
	model.UserStore = append(model.UserStore, newUser)

	res, _ := json.Marshal(&model.UserStore)
	HttpOk(writer, res)
}

func ToggleUser(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	idString := vars["id"]

	var reset bool
	if request.URL.Query().Has("reset") && request.URL.Query().Get("reset") == "true" {
		reset = true
	}

	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// Handle Reset
	if reset {
		for i, user := range model.UserStore {
			if user.Id == id {
				model.UserStore[i].Speaking = false
				model.UserStore[i].SpeakDuration = 0
				model.UserStore[i].StartDate = time.Now()
			}
		}
	}

	// Toggle the speaking flag and calculate time, stop all other users
	for i, user := range model.UserStore {
		if user.Id == id && !user.Speaking {
			model.UserStore[i].StartDate = time.Now()
			model.UserStore[i].Speaking = true
		} else {
			model.UserStore[i].Speaking = false
		}
	}

	res, _ := json.Marshal(&model.UserStore)
	HttpOk(writer, res)
}

func DeleteUser(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	idString := vars["id"]

	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	for i, user := range model.UserStore {
		if user.Id == id {
			model.UserStore = append(model.UserStore[:i], model.UserStore[i+1:]...)
		}
	}

	res, _ := json.Marshal(&model.UserStore)
	HttpOk(writer, res)
}
