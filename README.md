# check-go

This is a fork of Go at [golang/go@358d354](https://github.com/golang/go/commit/358d35455d06b1ebee948efff123842490dcb797) to add
a check keyword to the language for error-handling.

To try it out, execute
```
$ go run cmd/check-go file.go
```
where `file.go` contains Go source with code that uses the check keyword. Running this command
will place `generated.file.go` into the same directory, containing the equivalent
Go source without the check keyword.
