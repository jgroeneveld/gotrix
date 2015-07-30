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
* Using panics is not something that is used a lot in go, but is a perfectly valid strategy inside a library, i.e. the panic should leave the boundaries of the library. But with the different domain errors this would require a lot of panic recovery, type switching, selecting the panics that should be handled and panicking again on everything else.

The current solution with wrapped errors is not optimal for the given reason, but best the language permits.


## The Package Catastrophe

With modularization comes the intent to have many small modules, for which intent can be seen easily. But with packages of same intent in different domains there is the inherent problem, that go's package names don't give any respect to the ancestry, i.e. a package `some.domain.with.errors` will loose all information the `domain` aspect for this `errors` package, i.e. for `domain1` and `domain2` one can't say which `errors` package was meant (if both have one and both are required). The workaround is an named import, but the question remains how to chose those names. Something like `domain1errors` doesn't read very nicely.

This problem striked mostly with us dividing the gotrix file hierarchy in two mayor branches `lib` (with everything general for the different domains) and `app` and `web` (specific content).

The biggest problem was testing though. There is no way to use test helpers from another package, without making those helpers available in production code. Especially for fabrications this is a problem, as they should be used for testing the db layer itself, but should also be available when testing other packages.