package queries

import (
	"auto-garden/argparsing"
	"context"
)

func InsertIntoHumidityLevelDB(plantId int64, currentlyOn bool, currentHumidity int64) error {
	_, err := argparsing.ParseArgs().DbConn.Query(
		context.Background(),
		plantHumidityLevelQuery,
		currentlyOn,
		currentHumidity,
		plantId)
	return err
}
