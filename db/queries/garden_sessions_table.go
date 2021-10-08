package queries

import (
	"auto-garden/argparsing"
	"context"
	"time"
)

func GetSessionDate(id int64) (time.Time, error) {
	var date time.Time
	err := argparsing.ParseArgs().DbConn.QueryRow(context.Background(),
		getSessionDateQuery,
		id).Scan(&date)
	if err != nil {
		return date, err
	}
	return date, nil
}

func GetSessionId(userName string) (int64, error) {
	var sessionId int64
	err := argparsing.ParseArgs().DbConn.QueryRow(context.Background(),
		getSessionIdQuery,
		userName).Scan(&sessionId)
	return sessionId, err
}