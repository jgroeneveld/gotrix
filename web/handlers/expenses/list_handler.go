package expenses

import (
	"github.com/jgroeneveld/bookie2/app/model"
	"github.com/jgroeneveld/bookie2/web/ctx"
	"github.com/jgroeneveld/bookie2/web/views"
	"net/http"
)

func ListHandler(rw http.ResponseWriter, r *http.Request, c *ctx.Context) error {
	view := &views.ExpensesList{
		Expenses: []*model.Expense{
			&model.Expense{Description: "Fahrrad", Amount: 109900},
			&model.Expense{Description: "iPhone", Amount: 14999},
		},
	}

	return views.RenderWithLayout(rw, view)
}
