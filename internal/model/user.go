package model

import (
	"time"
)

var UserStore []User = make([]User, 0)

type User struct {
	Id            uint64        `json:"id,omitempty"`
	Name          string        `json:"name"`
	SpeakDuration time.Duration `json:"speakDuration,omitempty"`
	Speaking      bool          `json:"speaking,omitempty"`
	StartDate     time.Time     `json:"startDate,omitempty"`
}

func FindMaxId() uint64 {
	var highest uint64 = 0
	for _, user := range UserStore {
		if user.Id > highest {
			highest = user.Id
		}
	}

	return highest + 1
}
