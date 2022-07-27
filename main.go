package main

import (
	"context"

	"cookbook/api"
	"cookbook/db"

	"github.com/sirupsen/logrus"
)

func main() {
	client, err := db.InitDB()
	if err != nil {
		logrus.WithError(err).Fatal("unable to connect to db")
	}

	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			logrus.WithError(err).Fatal("disconnect failed")
		}
	}()

	api := api.InitAPI()
	api.Run()
}
