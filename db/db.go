package db

import (
	"auto-garden/argparsing"
	"context"
	"crypto/sha256"
	"errors"
	"github.com/thanhpk/randstr"
)

func ShouldWater(plantId int64, currentlyOn bool, currentHumidity int64) (bool, error) {
	var wantedHumidity int64
	err := argparsing.ParseArgs().DbConn.QueryRow(context.Background(),
		doesPlantExistQuery, plantId).Scan(&wantedHumidity)
	if err != nil {
		return false, err
	}
	insertIntoHumidityLevelDB(plantId, currentlyOn, currentHumidity)
	return currentHumidity < wantedHumidity, nil
}

func AddUser(user string, password string, writePermission bool, readPermission bool) error {
	err := errDoesUserExist(user)
	if err != nil {
		return err
	}
	salt := randstr.Hex(16)
	hashedPassword := sha256.Sum256([]byte(password + salt))
	return insertIntoUsersDB(user, hashedPassword[:], salt, writePermission, readPermission)
}

func DeleteUser(user string) error {
	exists, err := doesUserExist(user)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("user does not exist")
	}
	return deleteUserFromUsersDB(user)
}

func errDoesUserExist(userName string) error {
	exists, err := doesUserExist(userName)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("user exists")
	}
	return nil
}
