package gcp

import (
	"context"
	"encoding/json"
	"os"
	"path"

	logging "cloud.google.com/go/logging/apiv2"
	"google.golang.org/api/option"
)

var IsGCloud = false

type Auth struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RefreshToken string `json:"refresh_token"`
	Type         string `json:"type"`
}

func LoggingClient(ctx context.Context) (*logging.Client, error) {
	if !IsGCloud {
		return logging.NewClient(ctx, option.WithCredentialsFile(authFile()))
	} else {
		return logging.NewClient(ctx)
	}
}

func authDir() string {
	hd, _ := os.UserHomeDir()
	dir := path.Join(hd, ".loggo", "auth")
	return dir
}

func authFile() string {
	return path.Join(authDir(), "gcp.json")
}

func Delete() {
	_ = os.Remove(authFile())
}

func (a *Auth) Save() error {
	if err := os.MkdirAll(authDir(), os.ModePerm); err != nil {
		return err
	}
	b, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return err
	}

	file, err := os.Create(authFile())
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(b)
	return err
}
