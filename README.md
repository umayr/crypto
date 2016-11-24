A proof of concept to replace golang's standard crypto with openssl wrapped library.

### Usage
Consider the following snippet:

```go
package main

import (
	"crypto/aes"
	"fmt"
)

func main() {
	key := make([]byte, 32)
	c, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	plaintext := make([]byte, 32)
	out := make([]byte, len(plaintext))
	c.Encrypt(out, plaintext)

	fmt.Println(out)
	fmt.Println(plaintext)
}
```

Run it with `go`:

```
λ go run ./examples/aes.go

[220 149 192 120 162 64 137 137 173 72 162 20 146 132 32 135 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
```

Now run with the shell script in the `build` directory:

```
λ ./build/go.sh run ./examples/aes.go

2016/11/24 07:34:58 crypto/aes: encrypt with openssl
[220 149 192 120 162 64 137 137 173 72 162 20 146 132 32 135 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]

```



