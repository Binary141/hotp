# HMAC-Based One-Time Password (HOTP) Algorigthm

## What is this?
A Golang implementation of [RFC 4226](https://datatracker.ietf.org/doc/html/rfc4226) which serves as the basis for Time-Based One-Time Passwords ([RFC 6238](https://datatracker.ietf.org/doc/html/rfc6238)). This implementation relies only on packages within the standard library, so no external imports (outside of testing) needed. 

## Install
Requires >=1.24.2
```sh
go get github.com/binary141/hotp-go-rfc@v1.0.0
```
## Usage
```golang
package main

import (
	"fmt"

	"github.com/binary141/hotp-go-rfc"
)

func main() {
	secret := "12345678901234567890"
	var counter uint64 = 1
	digits := 6

	fmt.Println(hotp.Hotp(string(secret), counter+2, digits))
}
```

## Testing
```sh
go test
```

