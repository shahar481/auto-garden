package db

import (
	"auto-garden/consts"
	"auto-garden/db/queries"
	"auto-garden/crypto"
	"errors"
	"github.com/thanhpk/randstr"
	"time"
)

func ShouldWater(plantId int64, currentlyOn bool, currentHumidity int64) (bool, error) {
	wantedHumidity, err := queries.GetPlantWantedHumidity(plantId)
	if err != nil {
		return false, err
	}
	err = queries.InsertIntoHumidityLevelDB(plantId, currentlyOn, currentHumidity)
	return currentHumidity < wantedHumidity, err
}

func AddUser(user string, password string, writePermission bool, readPermission bool) error {
	err := errDoesUserExist(user)
	if err != nil {
		return err
	}
	salt := randstr.Hex(16)
	hashedPassword := crypto.HashPassword(password, salt)
	return queries.InsertIntoUsersTable(user, hashedPassword[:], salt, writePermission, readPermission)
}

func DeleteUser(user string) error {
	exists, err := queries.DoesUserExist(user)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("user does not exist")
	}
	return queries.DeleteUserFromUsersTable(user)
}

func errDoesUserExist(userName string) error {
	exists, err := queries.DoesUserExist(userName)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("user exists")
	}
	return nil
}

func IsSessionValid(sessionId int64) (bool, error) {
	date, err := queries.GetSessionDate(sessionId)
	if err != nil {
		return false, err
	}
	return isValidDate(date), nil
}

func isValidDate(date time.Time) bool {
	minutesTime := time.Since(date).Minutes()
	return minutesTime > consts.SessionMinutesTimeout
}

