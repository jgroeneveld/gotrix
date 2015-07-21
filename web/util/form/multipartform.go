package form

import (
	"mime/multipart"
	"net/http"
)

func NewMultipart(r *http.Request) (*MultipartForm, error) {
	err := r.ParseMultipartForm(10 * 1024 * 1024)
	if err != nil {
		return nil, err
	}

	return &MultipartForm{
		Form: Form{
			Values: r.MultipartForm.Value,
		},
		files: r.MultipartForm.File,
	}, nil
}

type MultipartForm struct {
	files map[string][]*multipart.FileHeader
	Form
}

func (f *MultipartForm) ReqFile(key string) *multipart.FileHeader {
	fh := f.GetFile(key)
	if fh == nil {
		f.addErrorMissingParameter(key)
	}
	return fh
}

func (f *MultipartForm) GetFile(key string) *multipart.FileHeader {
	if len(f.files) == 0 {
		return nil
	}
	return f.files[key][0]
}
