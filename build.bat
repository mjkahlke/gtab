@ECHO OFF
:: A simple batch file to build the implementation and run unit tests. No
:: error checking is done. It is assumed that everything runs successfully
:: which is true if nothing has been changed after the code pull.
CD impl
DEL go.mod
go mod init "gtab/impl"
go mod tidy
go test
go build

:: Build the command line application and run unit tests
CD ..\cmd\cli
DEL go.mod
go mod init "gtab/cli"
go mod edit -replace "gtab/impl"=..\..\impl
go mod tidy
go test
go build

:: Build the gin-based webserver to respond to REST API requests
CD ..\rest
DEL go.mod
go mod init "gtab/rest"
go mod edit -replace "gtab/impl"=..\..\impl
go mod tidy
:: TODO: add some unit tests, this will just complain that there are no tests
go test
go build

:: Be a good citizen and return the caller to the starting directory
CD ..\..

