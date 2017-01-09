package webtest

import (
	"io"
	"testing"

	"gotrix/lib/logger"

	"gotrix/web"

	"gotrix/app/db/dbtest"

	"gotrix/lib/web/webtest/testclient"

	"github.com/jgroeneveld/schema"
	"github.com/jgroeneveld/trial/assert"
	"github.com/jgroeneveld/trial/th"
)

func TestAPIExpenses(t *testing.T) {
	txMFac := dbtest.NewTxManagerFactory()
	defer txMFac.Close()

	router := web.NewRouter(logger.Discard, txMFac)
	s := testclient.New(router)

	{
		status, body, err := s.PostForm("/api/v1/expenses", "description=asd&amount=12")
		assert.MustBeNil(t, err)
		defer body.Close()

		assert.Equal(t, 201, status)
		AssertJSONSchema(t,
			schema.Map{
				"amount":      12,
				"description": "asd",
			},
			body,
		)
	}

	{
		status, body, err := s.Get("/api/v1/expenses")
		assert.MustBeNil(t, err)
		defer body.Close()

		assert.Equal(t, 200, status)
		AssertJSONSchema(t,
			schema.Array(
				schema.Map{
					"amount":      12,
					"description": "asd",
				}),
			body,
		)
	}
}

func AssertJSONSchema(t *testing.T, matcher schema.Matcher, r io.Reader) {
	err := schema.MatchJSON(matcher, r)
	if err != nil {
		th.Error(t, 1, err.Error())
	}
}
