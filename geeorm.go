package geeorm

import (
	"database/sql"
	"github.com/foolish06/geeorm/logger"
	"github.com/foolish06/geeorm/session"
)

// Engine is the main struct of geeorm, manages all db sessions and transactions.
type Engine struct {
	db *sql.DB
}

// NewEngine create an instance of Engine
// connect database and ping it to test whether it's alive
func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		logger.Error(err)
		return
	}
	// Send a ping to make sure the database connection is alive.
	if err = db.Ping(); err != nil {
		logger.Error(err)
		return
	}

	e = &Engine{db: db}
	logger.Info("Success to connect database")
	return
}

func (engine *Engine) Close() {
	if err := engine.db.Close(); err != nil {
		logger.Error("Fail to close database")
		return
	}
	logger.Info("Success to close database")
}

// NewSession creates a new session for next operations
func (engine *Engine) NewSession() *session.Session{
	return session.New(engine.db)
}