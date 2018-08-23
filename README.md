## Gocsv

Service to parse CSV file with PRN expressions and Refs (A1, B1, etc)

### Building, running, testing

```
$ go get github.com/dneprix/gocsv
$ cd ~/go/src/github.com/dneprix/gocsv
$ go run main.go files/test.csv

$ go test ./...
```

### Comments

* Reference LETTER NUMBER should have only one letter (A1, B1, etc)
* Circle refs returns #ERR
* Script handles /0 operations and +-Inf values
* I don't use Interfaces because we no need them in this task. But I can add Interfaces if task will be complicated: several file parsers, several calc algorithms, etc.  
* I wrote tests only for important functions because of time. It is not a problem to cover all code if it is needed.
* I didn't test very big CSV files. Maybe I will need to update parsing CSV code for them.
