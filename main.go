package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	c := exec.Command(os.Args[1], os.Args[2:]...)
	c.Env = append(c.Env, os.Environ()...)
	c.Env = append(c.Env, "https_proxy=http://127.0.0.1:7890")
	c.Env = append(c.Env, "http_proxy=http://127.0.0.1:7890")
	c.Env = append(c.Env, "HTTPS_PROXY=http://127.0.0.1:7890")
	c.Env = append(c.Env, "HTTP_PROXY=http://127.0.0.1:7890")
	c.Stderr = os.Stderr
	c.Stdout = os.Stdout
	if err := c.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "%q\n", err)
		os.Exit(1)
	}
}
