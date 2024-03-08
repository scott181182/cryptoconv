# Cryptocurrency Converter - Rust

## Running

In order to run this program, you must have [Rust](https://www.rust-lang.org/tools/install) installed (I did not want to Dockerize the build process).

With Rust installed, there are two approaches to build and run the program:

```sh
# Build and run the program in a single command
cargo run -- 100 BTC ETH

# Build and _then_ run the program as a separate binary executable
# Note the executable path changes if you do a release build
cargo build
./target/debug/cryptoconv 100 BTC ETH
```

## Testing

There are some limited tests in this library. To run, simply execute `cargo test`
