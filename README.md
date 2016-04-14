# SheBanGo

System administrator will write shell scripts to accomplish a certain task.
Many times you will encounter scripting languages have limited capabilities.
For example establishing secure transport layer or upload files to a remote location securely and efficiently.
Using Go to replace shell scripts can be achieved quickly.

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
#There you go !
