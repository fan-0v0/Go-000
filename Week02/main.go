package main

import (
	"database/sql"
	"errors"
)

type user struct {
	Id   int
	Name string
}

type Dao struct {
	db *Db
}

func (d *Dao) Query() ([]user, error) {
	data, err := d.db.query("select * from users where ...")
	if err != nil {
		if errors.Is(sql.ErrNoRows) {
			return nil, errors.New("query error: " + sql.ErrNoRows.Error())
		}
		return nil, errors.wrap(err, "query error")
	}
	return data, err
}

type Service struct {
	dao *Dao
}

func (s *Service) Query() ([]user, error) {
	return s.dao.Query()
}

func main() {
	svc := Service{}
	//省略query前逻辑
	data, err := svc.Query()
	if err != nil {
		log.Errorf("error info: %v:", err)
		return
	}
	log.Infof("query result: %v", data)
	// .....
}
