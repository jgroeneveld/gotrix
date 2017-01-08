package middleware

import (
	"github.com/jgroeneveld/trial/assert"
	"net/http"
	"testing"
	"gotrix/app/db/dbtest"
	"gotrix/lib/web/ctx"
	"gotrix/lib/logger"
	"sync"
	"gotrix/lib/web"
	"gotrix/lib/db"
)

// This tests for a race condition in the tx middleware. The same tx manager is
// used over and over again. For sequential requests this is fine, but for parallel
// ones this is a race condition on the tx. In this test two requests r1 and r2 are
// constructed both facilitating a tx. The first request will be blocked until the
// second one finished. This creates the situation with the race condition.
func TestTransaction(t *testing.T) {
	txManager := dbtest.NewTxManager()

	mw := TxMiddleware(txManager)

	mx := new(sync.Mutex)
	mx.Lock()

	mwStack1 := mw.Bind(waitHandler(mx))
	mwStack2 := mw.Bind(successHandler())

	wg := new(sync.WaitGroup)
	wg.Add(2)
	var err1, err2 error

	go func() {
		defer wg.Done()
		err1 = mwStack1(nil, nil, ctx.NewContext(logger.Discard, nil))
	}()

	go func() {
		defer wg.Done()
		err2 = mwStack2(nil, nil, ctx.NewContext(logger.Discard, nil))
		assert.True(t, txManager.CloseSuccessCalled)
		mx.Unlock()
	}()

	wg.Wait()

	assert.Nil(t, err1)
	assert.Nil(t, err2)
}

func waitHandler(mx *sync.Mutex) web.HTTPHandle {
	return func(w http.ResponseWriter, r *http.Request, c *ctx.Context) error {
		defer mx.Unlock()
		mx.Lock()

		con, err := c.TxManager.Begin()
		if err != nil {
			return err
		}

		return testQuery(con)
	}
}

func successHandler() web.HTTPHandle {
	return func(w http.ResponseWriter, r *http.Request, c *ctx.Context) error {
		con, err := c.TxManager.Begin()
		if err != nil {
			return err
		}
		return testQuery(con)

	}
}

func testQuery(con db.Con) error {
	_, err := con.Exec("select * from pg_catalog.pg_database")
	return err
}

