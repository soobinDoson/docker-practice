package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"runtime/debug"

	"github.com/soobinDoson/docker-practice.git/db"
	"github.com/urfave/cli/v2"
)

func main() {
	d := &db.DB{}
	if err := d.ConnectDb("host=127.0.0.1 user=postgres password=123 dbname=docker_practice port=5432 sslmode=disable"); err != nil {
		debug.PrintStack()
		log.Panicln(err)
	}
	log.Println("connect db successful!")
	err := HTTPServe()
	if err != nil {
		log.Panicln(err)
	}
}

func CliTool() error {
	app := &cli.App{
		Name:  "boom",
		Usage: "make an explosive entrance",
		Action: func(*cli.Context) error {
			fmt.Println("boom! I say!")
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		return err
	}
	return nil
}

func HTTPServe() error {
	r := http.DefaultServeMux
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Pong!, %q", html.EscapeString(r.URL.Path))
	})
	err := http.ListenAndServe(":3001", nil)
	if err != nil {
		return err
	}
	return nil
}
