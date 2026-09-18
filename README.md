# Go Error Workshop Library

> ⚠️ Are you looking for the workshop code to participate? Then this is the
> wrong Repository. Please check out [the workshop repository](https://github.com/survivorbat/go-workshop-errors).
> This repository is used as a library there.

This library is used in the [error workshop](https://github.com/survivorbat/go-workshop-errors).
Its purpose is to demonstrate a library that goes through several refactors
that improve its error handling.

## Content

This library contains multiple versions of
the _Survivorbat™ WiFi-enabled Dishwasher_ SDK.
Its errors are generated based on the input of the function.
If all values are correct, it sleeps for 20 milliseconds and then returns nil.

### Version 1

Version 1 includes inline use of `errors.New` and `fmt.Errorf`, forcing a caller
to inspect the `.Error()` strings to assert the error.

### Version 2

Version 2 turns the errors into sentinel errors, allowing a caller to more
easily assert the error.
These errors are wrapped to add additional information to the full error string.

### Version 3

Version 3 turns the sentinel errors into full-fledged error structs, allowing
a caller direct access to the data.
