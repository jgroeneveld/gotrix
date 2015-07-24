# gotrix concept

**gotrix** is a blueprint for a go web project, solving the most common problems.
Its not meant to be used as a library but rather to be copied and used as starting point.
This frees from having to 

# Assumptions
A small to middle-sized web application with cli, api and frontend using postgres as db.

# Dependencies

- `benbjohnson/ego` for view rendering
- `julienschmidt/httprouter` for routing

## Recommended further libs

- `jgroeneveld/trial/assert` small and clean assertions

# Application Structure

- `app` application level logic
- `cmd` binaries
- `config` configuration
- `lib` reusable library components

# Errors

- `lib/errors` is used to wrap all unknown error sources to add stacktrace information.
- `app/apperr` provides application level errors (`apperr.Validation`, `apperr.RecordNotFound`)
- `lib/web/httperr` converts application level errors into http errors with status codes. They can be rendered by api or html middlewares into error responses.