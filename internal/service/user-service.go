package service

import (
	"time"

	"github.com/SubhaNITS150/dob-calculator/internal/models"
	"github.com/google/uuid"
)

func CalculateAge(dob time.Time) int {
	now := time.Now()
	age := now.Year() - dob.Year()

	if now.Month() < dob.Month() ||
		(now.Month() == dob.Month() && now.Day() < dob.Day()) {
		age--
	}

	return age
}

func MapUserFromSQLC(
	id uuid.UUID,
	name string,
	dob time.Time,
) models.User {

	return models.User{
		ID:   id,
		Name: name,
		DOB:  dob.Format("2006-01-02"),
		// Age:  CalculateAge(dob),
	}
}