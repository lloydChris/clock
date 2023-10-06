# Clock

## Use it
Run the application with `go run clock.go`

edit the output messages with:
`second:<new message>`
`minute:<mes message>`
`hour:<new message>`

This can be done any time the program is running

## Contributing
1 hour is a long time to wait when testing if the hour chime works.
3 flags are provided to help with testing.
`-verbose` adds detailed output
`-longChime <seconds>` can be used to change the interval for the longest chime.  The default is 1 hour
`-shortChime <seconds>` can be used to change the inverval for the middle chime.  The default is 1 minute

A good way to get quick feedback when testing might be:
`go run clock.go -verbose -longChime 10 -shortChime 5`

## Testing
`go test` will run the tests
There is not much to test and this was done mostly to demonstrate that I am capable of writting tests.
