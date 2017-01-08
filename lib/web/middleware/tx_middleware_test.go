package middleware

import (
	"github.com/jgroeneveld/trial/assert"
	"net/http"
	"testing"
	"gotrix/app/db/dbtest"
	libdbtest "gotrix/lib/db/dbtest"
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
	txMFac := dbtest.NewTxManagerFactory()

	mw := TxMiddleware(txMFac)

	mx := new(sync.Mutex)
	mx.Lock()

	mwStack1 := mw.Bind(waitHandler(mx))
	mwStack2 := mw.Bind(successHandler())

	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		defer wg.Done()
		c := ctx.NewContext(logger.Discard, nil)
		err := mwStack1(nil, nil, c)
		assert.Nil(t, err, "first request should succeed")
		assert.True(t, c.TxManager.(*libdbtest.TxManager).CloseSuccessCalled)
	}()

	go func() {
		defer wg.Done()
		c := ctx.NewContext(logger.Discard, nil)
		err := mwStack2(nil, nil, c)
		assert.Nil(t, err, "second request should succeed")
		assert.True(t, c.TxManager.(*libdbtest.TxManager).CloseSuccessCalled)
		mx.Unlock()
	}()

	wg.Wait()
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

