# Running Tests StrOps

Open a command prompt in this directory (*stringopsgo/strops/v3*) and run the
following commands.

## Test Execution WITHOUT Code Coverage
Run this in *strops/v3* directory:

### Windows Command Without Coverage
  `go test -v > xx_tests.txt`

### Linux Command Without Coverage
`go test -v | tee xx_tests.txt`

This will generate test results in the *stringopsgo/strops/v3* 
directory which are stored in the text file, `xx_tests.txt`. 

## Running Tests with code coverage

First pull down and install the `cover` package.
 
`go get golang.org/x/tools/cmd/cover`
  
Next, follow the test execution protocol.  
  
## Test Execution With Code Coverage
Run this in *strops/v3* directory:

### Windows Command With Coverage
`go test -cover -v > xx_tests.txt`

### Linux Command With Coverage
`go test -cover -v | tee xx_tests.txt`


## Cover Profile

Generate the code coverage detail. Run this command
in the *stringopsgo/strops/v3* directory:

`go test -coverprofile=xx_coverage.out`


The following provides for code coverage display in your
browser. Run this on the terminal command line and run it
in the *stringopsgo/strops/v3* directory:

`go tool cover -html=xx_coverage.out`