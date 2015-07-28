// DO NOT EDIT

package views

import (
	"fmt"
	"html"
	"io"
)

var _ = fmt.Sprint("") // just so that we can keep the fmt import for now
//line web/frontend/views/_footer.ego:1
func Footer(w io.Writer) error {
//line web/frontend/views/_footer.ego:2
	_, _ = io.WriteString(w, "\n\n</div>\n</body>\n</html>\n")
	return nil
}

//line web/frontend/views/_header.ego:1
func Header(w io.Writer) error {
//line web/frontend/views/_header.ego:2
	_, _ = io.WriteString(w, "\n\n<!DOCTYPE html>\n<html>\n<head>\n    <meta charset=\"utf-8\">\n    <meta name=\"viewport\" content=\"width=device-width\">\n    <title>gotrix</title>\n    <link rel=\"stylesheet\" href=\"/assets/css/lib/bootstrap.min.css\" type=\"text/css\" media=\"screen\" title=\"no title\" charset=\"utf-8\">\n    <link rel=\"stylesheet\" href=\"/assets/css/app.css\" type=\"text/css\" media=\"screen\" title=\"no title\" charset=\"utf-8\">\n\n    <script src=\"/assets/js/lib/jquery-2.0.3.min.js\"></script>\n    <script src=\"/assets/js/app.js\"></script>\n</head>\n<body>\n<div class=\"container\">\n\n    <nav class=\"navbar navbar-default\">\n        <div class=\"container-fluid\">\n            <div class=\"navbar-header\">\n                <button type=\"button\" class=\"navbar-toggle collapsed\" data-toggle=\"collapse\" data-target=\"#navbar\" aria-expanded=\"false\" aria-controls=\"navbar\">\n                    <span class=\"sr-only\">Toggle navigation</span>\n                    <span class=\"icon-bar\"></span>\n                    <span class=\"icon-bar\"></span>\n                    <span class=\"icon-bar\"></span>\n                </button>\n                <a class=\"navbar-brand\" href=\"/\">gotrix</a>\n            </div>\n            <div id=\"navbar\" class=\"navbar-collapse collapse\">\n                <ul class=\"nav navbar-nav\">\n                    <li><a href=\"/\">Expenses</a></li>\n                </ul>\n            </div>\n        </div>\n    </nav>\n\n")
	return nil
}

//line web/frontend/views/expenses_list.ego:1
func writeExpensesList(w io.Writer, v *ExpensesList) error {
//line web/frontend/views/expenses_list.ego:2
	_, _ = io.WriteString(w, "\n\n<h1>Expenses</h1>\n\n")
//line web/frontend/views/expenses_list.ego:5
	for _, e := range v.Expenses {
//line web/frontend/views/expenses_list.ego:6
		_, _ = io.WriteString(w, "\n<div>\n    ")
//line web/frontend/views/expenses_list.ego:7
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", e.Description)))
//line web/frontend/views/expenses_list.ego:7
		_, _ = io.WriteString(w, " - ")
//line web/frontend/views/expenses_list.ego:7
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprintf("%v", e.Amount)))
//line web/frontend/views/expenses_list.ego:8
		_, _ = io.WriteString(w, "\n</div>\n")
//line web/frontend/views/expenses_list.ego:9
	}
	return nil
}
