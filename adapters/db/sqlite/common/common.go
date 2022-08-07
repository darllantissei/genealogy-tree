package common

import (
	"database/sql"
	"errors"
	"strings"
	"sync"

	log "github.com/sirupsen/logrus"
	gouuid "github.com/satori/go.uuid"
)

type CommonDB struct {
	db       *sql.DB
	sourceDB string
	lockFlow *sync.Mutex
}

func NewCommon(fileDataBase string) *CommonDB {
	return &CommonDB{
		sourceDB: fileDataBase,
	}
}

func (c *CommonDB) GetUUID() string {

	uuid := gouuid.NewV4()

	return uuid.String()
}

func (c *CommonDB) TreatError(err error, messageNewError string) error {

	log.WithFields(log.Fields{
		"error_db": err.Error(),
	},
	).Error("err: error in db")

	return errors.New(messageNewError)
}

func (c *CommonDB) connect() {

	var (
		err error
	)

	if c.lockFlow == nil {

		c.lockFlow = &sync.Mutex{}

		c.lockFlow.Lock()

		defer func() {

			c.lockFlow.Unlock()

			c.lockFlow = nil
		}()

	}

	if c.db == nil {

		if strings.EqualFold(c.sourceDB, "") {
			c.db, err = sql.Open("sqlite3", ":memory")
		} else {
			c.db, err = sql.Open("sqlite3", c.sourceDB)
		}

		if err != nil {
			panic("impossible establish connection with DB sqlite")
		}
	}

}

func (c *CommonDB) Query(dql string, args ...interface{}) (*sql.Rows, error) {

	c.connect()

	return c.db.Query(dql, args...)

}

func (c *CommonDB) ExecStatement(dml string, args ...interface{}) (sql.Result, error) {

	c.connect()

	stm, err := c.db.Prepare(dml)

	if err != nil {
		return nil, err
	}

	defer stm.Close()

	return stm.Exec(args...)

}
