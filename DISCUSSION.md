# Discussion

While puting together the learnings from differents projects we did into something that could be used as starting point for future ones some clever solutions should up, but also some inherent problems in the go language showed up.


## The error Dilemma

Go has an internal interface for errors. This is used for almost all error reporting and has only one requirements for its types: they must return a string with the `Error` method. This is plain simple and leads to the following code pattern found everywhere:

    err := someMethod()
    if err != nil {
      return err
    }

This might not be the nicest thing to read, but gets familiar soon. But there is an inherent problem: Errors don't have location information, i.e. which call produced the error? This can be resolved with the wrapping of errors in a type that has a stack trace added:

    err := someMethod()
    if err != nil {
      return errors.Wrap(err)
    }

Now the location can be printed in `Error` method. But on the other hand the original error is hidden behind the wrapper, i.e. more action is required to retrieve the original error, if one wants to discriminate the error:

    err := db.SelectEntity(id)
    e := errors.Unwrap(err)
    if e == sql.ErrNoRows {
        return errors.New("no entry with id %q exists", id)
    } else if err != nil {
        return err
    }

With multiple layers of error wrapping (think a domain layer's error types that all contain the wrapped error) the actual error handling becomes a burden, as many layers of reflection are required to determine which error is actually to handle.


## The Package Catastrophe

With modularization comes the intent to have many small modules, for which intent can be seen easily. But with packages of same intent in different domains there is the inherent problem, that go's package names don't give any respect to the ancestry, i.e. a package `some.domain.with.errors` will loose all information the `domain` aspect for this `errors` package, i.e. for `domain1` and `domain2` one can't say which `errors` package was meant (if both have one and both are required). The workaround is an named import, but the question remains how to chose those names. Something like `domain1errors` doesn't read very nicely.

This problem striked mostly with us dividing the gotrix file hierarchy in two mayor branches `lib` (with everything general for the different domains) and `app` and `web` (specific content).

The biggest problem was testing though. There is no way to use test helpers from another package, without making those helpers available in production code. Especially for fabrications this is a problem, as they should be used for testing the db layer itself, but should also be available when testing other packages.