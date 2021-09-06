package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"user-service/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewDB)

// Data .
type Data struct {
	// TODO wrapped database client
	db  *sqlx.DB
	log *log.Helper
}

// NewData .
func NewData(db *sqlx.DB, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "user-service/data"))

	return &Data{
			db:  db,
			log: log,
		}, func() {

		}, nil
}

func NewDB(conf *conf.Data, logger log.Logger) *sqlx.DB {
	log := log.NewHelper(log.With(logger, "module", "auth-service/data/sqlx"))
	db, err := sqlx.Connect("mysql", conf.Database.Source)
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	return db
}
