//usr/bin/env go run $0 $@ ; exit
//You need to chmod +x r2-example.go
package main

import (
	"fmt"

	"github.com/radare/r2pipe-go"
)

func main() {
 r2p, err := r2pipe.NewPipe("plaid/butterfly_33e86bcc2f0a21d57970dc6907867bed")
	if err != nil {
		panic(err)
	}
	defer r2p.Close()

	_, err = r2p.Cmd("aaaa")
	if err != nil {
		panic(err)
	}
	buf0, err := r2p.Cmd("af@sym._main")
	if err != nil {
		panic(err)
	}
	fmt.Println(buf0)

	buf1, err := r2p.Cmd("pd 100")
	if err != nil {
		panic(err)
	}
	fmt.Println(buf1)
	buf2, err := r2p.Cmd("S*")
	if err != nil {
		panic(err)
	}
	fmt.Println(buf2)
}
