# goset #

Simple set library for golang which puts an emphasis on usability over
performance. Examples can be found in the [example_test.go](example_test.go)
test suite.

## Performance ##

Set is currently fairly slow and lots of work has to be done to tune it
performance-wise. The main bottleneck that was observed on live code was the
underlying go map data-structure which forces lots of memory allocation.
