package db

import (
	"auto-garden/argparsing"
	"context"
	"log"
)

func insertIntoUsersDB(user string, password []byte, salt string, writePermission bool, readPermission bool) error {
	_, err := argparsing.ParseArgs().DbConn.Query(
		context.Background(),
		insertUserQuery,
		user,
		password,
		salt,
		writePermission,
		readPermission)
	return err
}

func insertIntoHumidityLevelDB(plantId int64, currentlyOn bool, currentHumidity int64) {
	_, err := argparsing.ParseArgs().DbConn.Query(
		context.Background(),
		plantHumidityLevelQuery,
		currentlyOn,
		currentHumidity,
		plantId)
	if err != nil {
		log.Fatal(err)
	}
}

func deleteUserFromUsersDB(userName string) error {
	_, err := argparsing.ParseArgs().DbConn.Query(
		context.Background(),
		removeUserQuery,
		userName)
	return err
}

func doesUserExist(userName string) (bool, error) {
	var userId int64
	err := argparsing.ParseArgs().DbConn.QueryRow(context.Background(),
		doesUserExistQuery, userName).Scan(&userId)
	if err == nil {
		return true, nil
	}
	if err.Error() == "no rows in result set" {
		return false, nil
	}
	return false, err
}