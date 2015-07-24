package db

import (
	"database/sql"
	"strconv"
	"strings"
)

type Opt struct {
	WhereQuery string
	WhereArgs  []interface{}
	Order      string
	Limit      *int
	ForUpdate  bool
}

func OrderBy(field string) func(o *Opt) {
	return func(o *Opt) {
		o.Order = field
	}
}

func Limit(l int) func(o *Opt) {
	return func(o *Opt) {
		o.Limit = &l
	}
}

func Where(q string, args ...interface{}) func(o *Opt) {
	return func(o *Opt) {
		o.WhereQuery = q
		o.WhereArgs = args
	}
}

func queryWithOpts(con Con, query string, opts ...func(*Opt)) (*sql.Rows, error) {
	query, args := queryFromOpts(query, opts...)
	return con.Query(query, args...)
}

func queryFromOpts(query string, opts ...func(*Opt)) (string, []interface{}) {
	o := &Opt{}
	for _, f := range opts {
		f(o)
	}

	q := []string{query}

	args := []interface{}{}
	if o.WhereQuery != "" {
		q = append(q, "WHERE "+replaceQuestionMarksInQuery(o.WhereQuery))
		args = append(args, o.WhereArgs...)
	}
	if o.Order != "" {
		q = append(q, "ORDER BY "+o.Order)
	}
	if o.Limit != nil {
		q = append(q, "LIMIT "+strconv.Itoa(*o.Limit))
	}

	return strings.Join(q, " "), args
}

func replaceQuestionMarksInQuery(query string) string {
	i := 1
	return replaceOutsideQuoted(query, '?', func() string {
		s := "$" + strconv.Itoa(i)
		i++
		return s
	})
}

func replaceOutsideQuoted(s string, old rune, rep func() string) string {
	res := ""
	quoteCount := 0
	for _, r := range s {
		if r == '\'' {
			quoteCount = (quoteCount + 1) % 2
		}
		switch {
		case r == old && quoteCount == 0:
			res += rep()
		default:
			res += string(r)
		}
	}
	return res
}
