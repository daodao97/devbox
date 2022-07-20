package app

import (
	"github.com/davecgh/go-spew/spew"
	"time"

	"github.com/daodao97/fly"
)

type Workspace struct {
	ID     int64      `db:"id"    json:"id"`
	Name   string     `db:"name"     json:"name"`
	Config string     `db:"config"   json:"config"`
	Ctime  time.Time  `db:"ctime"    json:"ctime"`
	Mtime  time.Time  `db:"mtime"    json:"mtime"`
	Group  []ApiGroup `json:"group"`
}

func (a *App) CreateWorkspace(g *Workspace) (bool, error) {
	lastId, err := a.workspace.Insert(g)
	if err != nil {
		return false, err
	}
	return lastId > int64(0), nil
}

func (a *App) UpdateWorkspace(g *Workspace) (bool, error) {
	g.Mtime = time.Now()
	return a.workspace.Update(g)
}

func (a *App) WorkspaceList() (list []Workspace, err error) {
	err = a.workspace.Select().Binding(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (a *App) DeleteWorkspace(id int64) (bool, error) {
	return a.workspace.Delete(fly.WhereEq("id", id))
}

func (a *App) LoadWorkspace(id int64) (*Workspace, error) {
	var w Workspace
	err := a.workspace.SelectOne(fly.WhereEq("id", id)).Binding(&w)
	if err != nil {
		return nil, err
	}
	var g []ApiGroup
	err = a.apiGroup.Select(fly.WhereEq("workspace_id", id), fly.OrderByDesc("id")).Binding(&g)
	if err != nil {
		return nil, err
	}
	var gid []interface{}
	for _, v := range g {
		gid = append(gid, v.ID)
	}

	var l []ApiList
	err = a.apiList.Select(fly.WhereIn("gid", gid)).Binding(&l)
	if err != nil {
		return nil, err
	}

	spew.Dump(l)

	for i, _g := range g {
		var tmp []ApiList
		for _, _l := range l {
			if _g.ID == _l.Gid {
				tmp = append(tmp, _l)
			}
		}
		g[i].List = tmp
	}

	w.Group = g

	return &w, nil
}
