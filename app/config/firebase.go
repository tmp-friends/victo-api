package config

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
)

func NewFirebaseApp() *firebase.App {
	// 環境変数にサービスアカウントファイルパスを設定することで暗黙的にadmin sdkが使える
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("err initializing firebase app: %v\n", err)
	}

	return app
}
