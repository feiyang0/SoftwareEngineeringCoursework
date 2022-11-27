package mysql

import (
	"fmt"
	"log"
	"sync"

	"github.com/spf13/viper"
	"gorm.io/gorm"

	"SoftwareEngine/internal/pkg/logger"
	v1 "SoftwareEngine/internal/pkg/model/server/v1"
	"SoftwareEngine/internal/server/store"
	"SoftwareEngine/pkg/database"
)

type datastore struct {
	db *gorm.DB
}

func (ds *datastore) Users() store.UserStore {
	return newUsers(ds)
}

func (ds *datastore) Problems() store.ProblemStore {
	return newProblems(ds)
}

func (ds *datastore) Students() store.StudentStore {
	return newStudents(ds)
}

//func (ds *datastore) Teachers() store.TeacherStore {
//	return newTeachers(ds)
//}
//
//func (ds *datastore) Admins() store.AdminStore {
//	return newAdmin(ds)
//}

func (ds *datastore) Close() error {
	if ds.db == nil {
		return nil
	}

	db, err := ds.db.DB()
	if err != nil {
		return err
	}

	return db.Close()
}

var (
	mysqlFactory store.Factory
	once         sync.Once
)

// GetMySQLFactoryOr create mysql factory with the given config.
func GetMySQLFactoryOr() (store.Factory, error) {
	var err error
	var dbIns *gorm.DB
	once.Do(func() {
		options := &database.Options{
			Host:                  viper.GetString("db.host"),
			Username:              viper.GetString("db.username"),
			Password:              viper.GetString("db.password"),
			Database:              viper.GetString("db.database"),
			MaxIdleConnections:    viper.GetInt("db.max-idle-connections"),
			MaxOpenConnections:    viper.GetInt("db.max-open-connections"),
			MaxConnectionLifeTime: viper.GetDuration("db.max-connection-life-time"),
			LogLevel:              viper.GetInt("db.log-level"),
			Logger:                logger.New(viper.GetInt("db.log-level")),
		}

		dbIns, err = database.New(options)

		// uncomment the following line if you need auto migration the given models
		// not suggested in production environment.
		if err = migrateDatabase(dbIns); err != nil {
			log.Fatalln("migrateDatabase:", err)
		}

		mysqlFactory = &datastore{dbIns}
	})

	if mysqlFactory == nil || err != nil {
		return nil, fmt.Errorf("failed to get mysql store fatory, mysqlFactory: %+v, error: %w", mysqlFactory, err)
	}
	return mysqlFactory, nil
}

// cleanDatabase tear downs the database tables.
// nolint:unused // may be reused in the feature, or just show a migrate usage.
func cleanDatabase(db *gorm.DB) error {
	if err := db.Migrator().DropTable(&v1.User{}); err != nil {
		return err
	}

	//if err := db.Migrator().DropTable(); err != nil {
	//	return err
	//}

	return nil
}

// migrateDatabase run auto migration for given models, will only add missing fields,
// won't delete/change current data.
// nolint:unused // may be reused in the feature, or just show a migrate usage.
func migrateDatabase(db *gorm.DB) error {
	// userTable
	if err := db.AutoMigrate(&v1.User{}); err != nil {
		return err
	}

	// 问题
	if err := db.AutoMigrate(&v1.Problem{}, &v1.Tag{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&v1.StudentProblem{}); err != nil {
		return err
	}
	//if err := db.SetupJoinTable(&v1.User{}, "Problems", &v1.StudentProblem{}); err != nil {
	//	return err
	//}

	return nil
}

// resetDatabase resets the database tables.
// nolint:unused,deadcode // may be reused in the feature, or just show a migrate usage.
func resetDatabase(db *gorm.DB) error {
	if err := cleanDatabase(db); err != nil {
		return err
	}
	if err := migrateDatabase(db); err != nil {
		return err
	}

	return nil
}
