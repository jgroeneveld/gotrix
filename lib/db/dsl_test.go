package db

import (
	"testing"

	"github.com/jgroeneveld/trial/assert"
)

func TestQueryFromOpts(t *testing.T) {
	query, args := FormatOpts(
		"SELECT name FROM expenses",
		OrderBy("name ASC"),
		Where("ignore = '?' AND name = ? AND age > ? AND created_at < ?", "harald", 12, "2015"),
		Limit(12),
	)

	assert.Equal(t, "SELECT name FROM expenses WHERE ignore = '?' AND name = $1 AND age > $2 AND created_at < $3 ORDER BY name ASC LIMIT 12", query)
	assert.DeepEqual(t, []interface{}{"harald", 12, "2015"}, args)
}
