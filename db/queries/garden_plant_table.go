package queries

import (
	"auto-garden/argparsing"
	"context"
)

func GetPlantWantedHumidity(plantId int64) (int64, error) {
	var wantedHumidity int64
	err := argparsing.ParseArgs().DbConn.QueryRow(context.Background(),
		doesPlantExistQuery, plantId).Scan(&wantedHumidity)
	return wantedHumidity, err
}