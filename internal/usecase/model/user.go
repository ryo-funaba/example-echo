package model

import (
	"time"

	"github.com/ryo-funaba/example_echo/internal/domain/model"
	"github.com/ryo-funaba/example_echo/internal/utils/timeutil"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func UserFromDomainModel(m *model.User) *User {
	return &User{
		ID:   int(m.ID),
		Name: m.Name,
		Age:  calcAge(m.CreatedAt),
	}
}

// calcAge calculates age from birthday.
func calcAge(birthday time.Time) int {
	now := timeutil.GetTruncatedNowJSTTime()

	thisYear, thisMonth, thisDay := now.Date()
	age := thisYear - birthday.Year()

	// when birthday is yet to come
	if thisMonth < birthday.Month() || thisMonth == birthday.Month() && thisDay < birthday.Day() {
		age--
	}

	return age
}
