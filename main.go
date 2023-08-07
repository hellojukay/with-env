package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime/debug"

	"github.com/joho/godotenv"
)

var (
	showVersion = flag.Bool("version", false, "pring progream version")
)

// first load env from ~/.env
// second load env from .env
func init() {
	flag.Parse()
	if *showVersion {
		PrintVersion()
		os.Exit(0)
	}
	if len(os.Args) < 2 {
		fmt.Printf("Usage: with-env {command}\n")
		os.Exit(1)
	}
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

func PrintVersion() {
	info, ok := debug.ReadBuildInfo()
	if ok {
		println(info.Main.Version, info.Main.Sum)
	}
}
