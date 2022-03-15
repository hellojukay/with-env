package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/joho/godotenv"
)

// first load env from ~/.env
// second load env from .env
func init() {
	home, _ := os.UserHomeDir()
	loadEnv(filepath.Join(home, ".env"), ".env")
}

func loadEnv(files ...string) {
	for _, envfile := range files {
		if _, err := os.Stat(envfile); !os.IsNotExist(err) {
			_ = godotenv.Overload(envfile)
		}
	}

}
func main() {
	c := exec.Command(os.Args[1], os.Args[2:]...)
	c.Stderr = os.Stderr
	c.Stdout = os.Stdout
	if err := c.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "%q\n", err)
		os.Exit(1)
	}
}
