package webtest

import (
	"testing"

	"gotrix/lib/logger"

	"gotrix/web"

	"gotrix/app/db/dbtest"

	"gotrix/web/webtest/testclient"

	"github.com/jgroeneveld/trial/assert"
)

func TestAPIExpenses(t *testing.T) {
	txManager := dbtest.NewTxManager()
	defer txManager.Rollback()

	router := web.NewRouter(logger.Discard, txManager)
	s := testclient.New(router)

	status, body, err := s.PostForm("/api/v1/expenses", "description=asd&amount=12")
	assert.MustBeNil(t, err)
	assert.Equal(t, 201, status)
	assert.MustBeEqual(t, "{\"amount\":12,\"description\":\"asd\"}\n", string(body))

	status, body, err = s.Get("/api/v1/expenses")
	assert.MustBeNil(t, err)
	assert.Equal(t, 200, status)
	assert.MustBeEqual(t, "[{\"amount\":12,\"description\":\"asd\"}]\n", string(body))
}
