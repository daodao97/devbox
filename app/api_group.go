package app

import (
	"github.com/daodao97/fly"
	"time"
)

type ApiGroup struct {
	ID          int64     `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	WorkspaceId int64     `db:"workspace_id" json:"workspace_id"`
	Config      string    `db:"config" json:"config"`
	Ctime       time.Time `db:"ctime" json:"ctime"`
	Mtime       time.Time `db:"mtime" json:"mtime"`
	List        []ApiList `json:"children"`
}

func (a *App) CreateGroup(g *ApiGroup) (bool, error) {
	lastId, err := a.apiGroup.Insert(g)
	if err != nil {
		return false, err
	}
	return lastId > int64(0), nil
}

func (a *App) UpdateGroup(g *ApiGroup) (bool, error) {
	g.Mtime = time.Now()
	return a.apiGroup.Update(g)
}

func (a *App) DeleteGroup(id int64) (bool, error) {
	return a.apiGroup.Delete(fly.WhereEq("id", id))
}
