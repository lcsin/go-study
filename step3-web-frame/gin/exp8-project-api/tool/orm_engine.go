package tool

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/lcsin/go-study/step3/gin/exp8-project-api/model"
)

var DbEngine *Orm

type Orm struct {
	*xorm.Engine
}

func OrmEngine(conf *Config) (*Orm, error) {
	database := conf.Database
	conn := database.User + ":" + database.Password +
		"@tcp(" + database.Host + ":" + database.Port + ")/" +
		database.DbName + "?charset=" + database.Charset

	engine, err := xorm.NewEngine(database.Driver, conn)
	if err != nil {
		return nil, err
	}

	engine.ShowSQL(database.ShowSql)
	engine.Sync(new(model.Member))

	orm := new(Orm)
	orm.Engine = engine
	DbEngine = orm
	return orm, nil
}
