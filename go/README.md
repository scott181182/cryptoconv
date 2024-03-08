# Cryptocurrency Converter - Go

## Running

In order to run this program, you must have [Go](https://go.dev/doc/install) installed (I did not want to Dockerize the build process).

With Go installed, the following can be used to build and run the program:

```sh
# Build the program.
go build

# Run the code
./cryptoconv 100 BTC ETH
```

## Testing

There are some limited tests in this library. To run, simply execute `go test`
