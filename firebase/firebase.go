package firebase

import (
	"context"
	"time"

	googlefirebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"firebase.google.com/go/db"
	"firebase.google.com/go/messaging"
	"firebase.google.com/go/storage"
	"google.golang.org/api/option"
)

type Firebase struct {
	Auth      *auth.Client
	Database  *db.Client
	Storage   *storage.Client
	Messaging *messaging.Client
}

func NewFirebase(config *Config) (*Firebase, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opt := option.WithCredentialsFile(config.CredentialsFile)
	c := googlefirebase.Config{
		DatabaseURL:   config.DatabaseUrl,
		StorageBucket: config.StorageBucket,
	}
	app, err := googlefirebase.NewApp(ctx, &c, opt)
	if err != nil {
		return nil, err
	}

	a, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}

	d, err := app.Database(ctx)
	if err != nil {
		return nil, err
	}

	s, err := app.Storage(ctx)
	if err != nil {
		return nil, err
	}

	m, err := app.Messaging(ctx)
	if err != nil {
		return nil, err
	}

	f := Firebase{
		Auth:      a,
		Database:  d,
		Storage:   s,
		Messaging: m,
	}

	return &f, nil
}
