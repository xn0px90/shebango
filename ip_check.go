//usr/bin/env go run $0 $@ ; exit
package main
import (
    "os"
    "sync"
    "os/exec"
    "fmt"
)

var wg *sync.WaitGroup

func main() {
    urls := os.Args[1:]

    wg = new(sync.WaitGroup)
    wg.Add(len(urls))

    for _, url := range urls {
        go download(url)
    }

    wg.Wait()
}

func download(url string) {
    defer wg.Done()
    cmd := exec.Command("curl", "-0", url)
    cmd.Run()
    fmt.Printf("%v",cmd)
}
