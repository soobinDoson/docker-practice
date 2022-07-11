package db

import (
	"log"
	"time"

	pb "github.com/soobinDoson/docker-practice.git/proto"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	engine *gorm.DB
}

// ConnectDb open connection to db
// ConnectDb expose ...
func (d *DB) ConnectDb(sqlDSN string) error {
	db, err := gorm.Open(postgres.New(
		postgres.Config{
			DSN:                  sqlDSN,
			PreferSimpleProtocol: true,
		}),
		&gorm.Config{
			// NamingStrategy: schema.NamingStrategy{
			// 	TablePrefix:   "partner.", // schema name
			// 	SingularTable: false,
			// },
			Logger: logger.Default.LogMode(logger.Info),
		})
	if err != nil {
		return err
	}
	sqlDb, err := db.DB()
	if err != nil {
		return err
	}
	ticker := time.NewTicker(5 * time.Minute)
	go func() {
		for {
			<-ticker.C
			if err := sqlDb.Ping(); err != nil {
				log.Print(err)
			}
		}
	}()
	// // SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	// sqlDb.SetMaxIdleConns(10)

	// // SetMaxOpenConns sets the maximum number of open connections to the database.
	// sqlDb.SetMaxOpenConns(100)
	d.engine = db
	return nil
}

func (d *DB) listUserPartnerQuery(req *pb.UserPartnerRequest) *gorm.DB {
	ss := d.engine.Table(tblUserPartner)
	if req.GetUserId() != "" {
		ss.Where(tblUserPartner+".user_id = ?", req.GetUserId())
	}
	if req.GetPhone() != "" {
		ss.Where(tblUserPartner+".phone = ?", req.GetPhone())
	}
	return ss
}

func (d *DB) ListUserPartner(rq *pb.UserPartnerRequest) ([]*pb.UserPartner, error) {
	log.Println("req: ", rq)
	var userParters []*pb.UserPartner
	ss := d.listUserPartnerQuery(rq)
	if rq.GetLimit() != 0 {
		ss.Limit(int(rq.GetLimit()))
	} else {
		if rq.GetLimit() != 0 {
			ss.Limit(int(rq.GetLimit()))
		}
	}
	err := ss.Order("created desc").Find(&userParters).Error
	if err != nil {
		return nil, err
	}
	return userParters, nil
}
