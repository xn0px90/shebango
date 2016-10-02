//usr/bin/env go run $0 $@ ; exit
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if !Exists("redis.tar.gz") {
		Download("http://download.redis.io/redis-stable.tar.gz", "redis.tar.gz")
	}
	if Exists("redis") {
		V(os.RemoveAll("redis"))
	}
	V(os.MkdirAll("redis", 0755))
	Run("tar zxvf redis.tar.gz -C redis --strip-components=1")
	V(os.Chdir("redis"))
	Run("make")
}

func V(err error) {
	if err != nil {
		panic(err)
	}
}

func Exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
		V(err)
	}
	return true
}

func Download(fileURL string, filename string) {
	resp, err := http.Get(fileURL)
	V(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	V(err)
	V(ioutil.WriteFile(filename, body, 0644))
}

func Run(command string) {
	fmt.Printf("run: %s\n", command)
	args := strings.Split(command, " ")
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("command failed: %s\n", command)
		V(err)
	}
}
