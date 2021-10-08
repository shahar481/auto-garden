package queries

import (
	"auto-garden/argparsing"
	"context"
)

// Inserts a user into the users table
func InsertIntoUsersTable(user string, password []byte, salt string, writePermission bool, readPermission bool) error {
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

// Checks if user exists in db
func DoesUserExist(userName string) (bool, error) {
	values, err := argparsing.ParseArgs().DbConn.Query(context.Background(),
		doesUserExistQuery, userName)
	if err != nil {
		return false, err
	}
	return values.Next(), nil
}

// Deletes user from users table
func DeleteUserFromUsersTable(userName string) error {
	_, err := argparsing.ParseArgs().DbConn.Query(
		context.Background(),
		removeUserQuery,
		userName)
	return err
}

func GetUserSalt(userName string) (string, error) {
	var salt string
	err := argparsing.ParseArgs().DbConn.QueryRow(context.Background(),
		getUserSalt, userName).Scan(&salt)
	return salt, err
}

func IsValidUserCredentials(userName string, hashedPassword [32]byte) (bool, error) {
	values, err := argparsing.ParseArgs().DbConn.Query(context.Background(),
		verifyCredentialsQuery, userName, hashedPassword)
	if err != nil {
		return false, err
	}
	return values.Next(), nil
}