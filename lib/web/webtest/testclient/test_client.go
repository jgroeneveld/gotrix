package testclient

import (
	"bytes"
	"gotrix/lib/errors"
	"io"
	"net/http"
	"net/http/httptest"
)

func New(handler http.Handler) *TestClient {
	return &TestClient{
		testServer: httptest.NewServer(handler),
	}
}

type TestClient struct {
	testServer *httptest.Server
}

func (s *TestClient) PostForm(path string, data string) (status int, responseBody io.ReadCloser, err error) {
	body := bytes.NewBufferString(data)

	req, err := http.NewRequest("POST", s.testServer.URL+path, body)
	if err != nil {
		return 0, nil, errors.Wrap(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, nil, errors.Wrap(err)
	}
	return rsp.StatusCode, rsp.Body, nil
}

func (s *TestClient) PostFormData(path string, contentType string, data io.Reader) (status int, responseBody io.ReadCloser, err error) {
	req, err := http.NewRequest("POST", s.testServer.URL+path, data)
	if err != nil {
		return 0, nil, errors.Wrap(err)
	}
	req.Header.Add("Content-Type", contentType)

	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, nil, errors.Wrap(err)
	}

	return rsp.StatusCode, rsp.Body, nil
}

func (s *TestClient) PatchForm(path string, data string) (status int, responseBody io.ReadCloser, err error) {
	body := bytes.NewBufferString(data)

	req, err := http.NewRequest("PATCH", s.testServer.URL+path, body)
	if err != nil {
		return 0, nil, errors.Wrap(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, nil, errors.Wrap(err)
	}

	return rsp.StatusCode, rsp.Body, nil
}

func (s *TestClient) PostJSON(path, jsonBody string) (status int, responseBody io.ReadCloser, err error) {
	body := bytes.NewBufferString(jsonBody)
	rsp, err := http.Post(s.testServer.URL+path, "application/json", body)
	if err != nil {
		return 0, nil, errors.Wrap(err)
	}

	return rsp.StatusCode, rsp.Body, nil
}

func (s *TestClient) PatchJSON(path, jsonBody string) (status int, responseBody io.ReadCloser, err error) {
	body := bytes.NewBufferString(jsonBody)

	req, err := http.NewRequest("PATCH", s.testServer.URL+path, body)
	if err != nil {
		return 0, nil, errors.Wrap(err)
	}
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, nil, errors.Wrap(err)
	}

	return rsp.StatusCode, rsp.Body, nil
}

func (s *TestClient) Get(path string) (status int, responseBody io.ReadCloser, err error) {
	rsp, err := http.Get(s.testServer.URL + path)
	if err != nil {
		return 0, nil, errors.Wrap(err)
	}

	return rsp.StatusCode, rsp.Body, nil
}

func (s *TestClient) Delete(path string) (status int, responseBody io.ReadCloser, err error) {
	req, err := http.NewRequest("DELETE", s.testServer.URL+path, nil)
	if err != nil {
		return 0, nil, errors.Wrap(err)
	}

	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, nil, errors.Wrap(err)
	}

	return rsp.StatusCode, rsp.Body, nil
}
