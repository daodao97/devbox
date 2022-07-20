package app

import "github.com/daodao97/fly"

type App struct {
	apiList   fly.Model
	apiGroup  fly.Model
	workspace fly.Model
}

func NewApp() *App {
	return &App{
		apiList:   fly.New("api_list"),
		apiGroup:  fly.New("api_group"),
		workspace: fly.New("workspace"),
	}
}
