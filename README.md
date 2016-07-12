Random string generator
========================

This package aims to provide a function to generate random strings.
It uses the crypto/rand package to generate a random number from 0 up to 2^64-1 and then converts it to a Base62 string.

Although it is possible to generate __unique__ strings with this, it is possible that it returns duplicates.
The chance of duplicates is very small, but when working with a very large number of generations duplicates may occur.

## Usage

See `example/example.go`

## License

MIT
