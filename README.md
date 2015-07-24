# gotrix concept

## Errors

- `lib/errors` is used to wrap all unknown error sources to add stacktrace information.
- `app/apperr` provides application level errors (`apperr.Validation`, `apperr.RecordNotFound`)
- `lib/web/httperr` converts application level errors into http errors with status codes. They can be rendered by api or html middlewares into error responses.