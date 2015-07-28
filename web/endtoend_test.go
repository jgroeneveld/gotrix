package web

import (
	"database/sql"
	"testing"

	"net/http/httptest"

	"net/http"

	"gotrix/lib/db"
	"gotrix/lib/db/dbtest"
	"gotrix/lib/logger"
	"github.com/jgroeneveld/trial/assert"
	"gotrix/cfg"
)

func TestAPIExpenses(t *testing.T) {
	txManager := dbtest.NewTestTxManager(newTestCon())
	defer txManager.Rollback()

	router := NewRouter(logger.Discard, txManager)
	s := httptest.NewServer(router)

	resp, err := http.Get(s.URL + "/api/v1/expenses")
	assert.MustBeNil(t, err)
	defer resp.Body.Close()
	assert.Equal(t, 200, resp.StatusCode)
}

var cachedTestCon *sql.DB

func newTestCon() *sql.DB {
	if cachedTestCon == nil {
		con, err := db.Connect(cfg.Config.DatabaseURL, cfg.Config.ApplicationName)
		if err != nil {
			panic(err)
		}
		cachedTestCon = con
	}
	return cachedTestCon
}
