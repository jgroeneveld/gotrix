# gotrix concept

**gotrix** is a blueprint for a go web project, solving the most common problems.
Its not meant to be used as a library but rather to be copied and used as starting point.
This frees from having to fulfill every need at every point in the lib.

# Assumptions
A small to middle-sized web application with cli, api and frontend using postgres as db.

# Dependencies

- `benbjohnson/ego` for view rendering
- `julienschmidt/httprouter` for routing

## Recommended further libs

- `jgroeneveld/trial/assert` small and clean assertions

# Application Structure

- `app` application level logic
- `cfg` configuration
- `cmd` binaries
- `lib` reusable library components
- `scripts` contains `go run`able scripts
- `web` web handlers and web-related logic (e.g. views)

## app

- `apperrors` see [Errors](#errors)
- `db` database persistence layer and connection management
- `db/migrations` see [Migrations](#migrations)
- `model` provides data structures that are managed by `db`.
- `service` service layer containing use cases for the application logic.

## cfg

- `config.go` contains the definition and loading of the config from file, env and defaults
- `defaults.go` contains the default configuration for the different environments.


## cmd

- `gtmigrate` runs the migrations
- `gtserver` starts the webserver


## lib

.......

## scripts

- `ego` contains a wrapped version of `benbjohnson/ego` so that every developer uses the same version
- `goassets` a script to bundle assets into the binary

## web

- `api` related handlers, serializers etc.
- `frontend` related handlers, views, assets etc.
- `webtest` global tests for the web layer like *end to end* tests
- `router.go` main entry point for the web layer

## Flow

- The access/presentation layer (`cli` / `web/api` / `web/frontend`) 
    - provides handlers that translate and validate user input into calls for the
`app/service` layer. 
    -  translate output to display for the user
    
- The service layer (`app/service`)
    - provides use cases, validates input and executes any calls to the persistence layer 
    - or any other data sources
     
- The persistence layer (`app/db`) provides access to the database in a structured way. 

# Errors

- `lib/errors` is used to wrap all unknown error sources to add stacktrace information.
- `app/apperr` provides application level errors (`apperr.Validation`, `apperr.RecordNotFound`)
- `lib/web/httperr` converts application level errors into http errors with status codes. They can be rendered by api or html middlewares into error responses.

# Testing

To allow for test helpers that are importable from other packages and to 
prevent cyclic dependencies, all tests will be contained in special `*test` packages.
All tests for the `db` package are contained in `db/dbtext` for example. 
`dbtest` -> `db`
`dbtest` -> `fabricate`
`fabricate` -> `db`
This way `fabricate` can be used by `dbtest` which both depend on `db`.

Furthermore, testing helpers are not directly mixed with the production code.

Sometimes this `*test` packages need to contain an `empty.go` file so that `go get` does not break for this package.