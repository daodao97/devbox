package app

import (
	"github.com/daodao97/fly"
	"github.com/davecgh/go-spew/spew"
	"time"
)

type ApiList struct {
	ID      int64     `db:"id" json:"id"`
	Gid     int64     `db:"gid" json:"gid"`
	Name    string    `db:"name" json:"name"`
	Method  string    `db:"method" json:"method"`
	Options string    `db:"options" json:"options"`
	Ctime   time.Time `db:"ctime,ii" json:"ctime"`
	Mtime   time.Time `db:"mtime" json:"mtime"`
}

func (a *App) CreateApi(g *ApiList) (int64, error) {
	lastId, err := a.apiList.Insert(g)
	if err != nil {
		return 0, err
	}
	return lastId, nil
}

func (a *App) UpdateApi(g *ApiList) (bool, error) {
	g.Mtime = time.Now()
	spew.Dump(g)
	return a.apiList.Update(g)
}

func (a *App) DeleteApi(id int64) (bool, error) {
	return a.apiList.Delete(fly.WhereEq("id", id))
}
