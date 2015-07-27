package web

import (
	"database/sql"
	"testing"

	"net/http/httptest"

	"net/http"

	"github.com/jgroeneveld/gotrix/lib/db"
	"github.com/jgroeneveld/gotrix/lib/db/dbtest"
	"github.com/jgroeneveld/gotrix/lib/logger"
	"github.com/jgroeneveld/trial/assert"
	"github.com/jgroeneveld/gotrix/cfg"
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
