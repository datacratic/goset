# goset #

Simple set library for golang which puts an emphasis on usability over
performance.


## Installation ##

You can download the code via the usual go utilities:

```
go get github.com/datacratic/goset
```

To build the code and run the test suite along with several static analysis
tools, use the provided Makefile:

```
make test
```

Note that the usual go utilities will work just fine but we require that all
commits pass the full suite of tests and static analysis tools.


## Examples ##

Usage examples are available in the following [test suite](example_test.go).


## Performance ##

Set is currently fairly slow and lots of work has to be done to tune it
performance-wise. The main bottleneck that was observed on live code was the
underlying go map data-structure which forces lots of memory allocation.


## License ##

The source code is available under the Apache License. See the LICENSE file for
more details.
