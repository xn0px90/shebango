# SheBanGo

# Example of how to execute a Go program as a shell script.

```
//usr/bin/env go run $0 $@ ; exit

package main
import "fmt"
func main() {
    fmt.Printf("hello, world\n")
}
```

# Make it an executable

```
chmod +x hello.go
./hello.go
hello, world
```
# There you go !
