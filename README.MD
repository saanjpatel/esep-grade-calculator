# ESEP Grade Calculator

This library is a simple utility to calculate your grades.

Setup Instructions:
Set up with Visual Studio Code:
1. Install Visual Studio Code if not already installed.
2. Install Go at this link https://go.dev/dl/ according to system.
3. Install Go extension with Visual Studio Code.
Note: Here is a link to what other editor plugins and IDEs can be used with Go https://go.dev/doc/editors.

Respository Summary:
Besides the README.MD, there are three files in the grade-calculator folder.

Runtime requirements:
Golang 1.24+

Run unit tests:
```
cd grade_calculator/
go test ./...
```

Run unit tests with code coverage:
```
cd grade_calculator/
go test -coverprofile=coverage.txt ./...
```