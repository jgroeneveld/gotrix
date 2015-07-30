# Discussion

We recently did some projects in go with a web serving component having a HTML UI. While condensing the lessons learned into something reusable, some inherent problems in the go language showed up. This documents discusses those.


## The error Dilemma

Go has an internal interface for errors. This is used for almost all error reporting and has only one requirements for its types: they must return a string with the `Error` method. This is plain simple and leads to the following code pattern found everywhere:

    err := someMethod()
    if err != nil {
      return err
    }

This might not be beautiful code (at least when you need to check many errors in a single function), but gets familiar soon. There is an inherent problem though: errors don't have location information in form of a stack trace, i.e. which call produced the error? One possible solution is to use a dedicated error type, that is used to create new one via something like `errors.New` and wrap all incoming errors:

    v, err := someMethod()
    if err != nil {
      return errors.Wrap(err)
    }
    
    if v == 0 {
      errors.New("returned value is 0")
    }

The custom error type can have a stack trace appended and use it while building up the message in the `Error` method. But on the other hand the original error is now hidden behind the wrapper type, i.e. more action is required to retrieve the original error, if one wants to handle specific errors differently, like in the following example:

    err := db.SelectEntity(id)
    e := errors.Unwrap(err)
    if e == sql.ErrNoRows {
        return errors.New("no entry with id %q exists", id)
    } else if err != nil {
        return err
    }

When working with domain specific errors this gets even worse, as it requires a lot of type casting to determine what kind of error was actually received. This gets pretty tedious and error prone soon. But is by far the best solution we've found. Others we considered were:

* Using the actual error types of the different domains as return values would be way more explicit, but is produces so much boiler plate, as it requires the package prefix most of the time (like in `apperror.Error`) and all errors must be returned as pointers.
* Using panics is not something that is used a lot in go, but is a perfectly valid strategy inside a library, i.e. the panic should leave the boundaries of the library. But with the different domain errors this would require a lot of panic recovery, type switching, selecting the panics that should be handled and panicking again on everything else. Basically poor-mans-exceptions without the helpers provided by languages using exceptions as error transport.

The current solution with wrapped errors is not optimal for the given reason, but best the language permits.


## The Package And Testing Catastrophe

The primary entity of modularization within go is the package. While trying to structure the gotrix application we had the following intent:

* Have everything application specific in dedicated packages following a domain approach.
* Extract generic stuff to a separate structure that could be published as library in the future.

This led to three major package trees: `app`, `web` and `lib`. The prior two are application related and the last one contains the generic stuff. Those trees share a lot of common structure, i.e. there are many packages that share the same purpose and therefore name. This is a problem as go's package names don't give any respect to the ancestry, i.e. a package `some.domain.with.errors` will loose all information on the `domain` aspect. The workaround is an named import, but the question remains how to chose those names. Something like `domain1errors` doesn't read very nicely, requires manual effort while building import statements and a lot of additional typing.

Those things are inconvenient, but testing is a real problem. There are lots of testing helpers, that should be shared between the different packages, like for example the fabrication* functions. This isn't possible in go without making those helpers available on the public production interface, as entities in test files (having the `_test` suffix) are not exported, not even for other packages' tests. Generally the testing code is put side by side with production code, but we decided to add separate "test" packages marked by their name for everything but simple unit tests. This makes the intent of the functions pretty clear (like in `dbtest.FabricateSomething`) and prevents many of the import loop problems.

*fabrication: Functions that insert data in the database with predefined defaults so the tests can concentrate on the test subject
