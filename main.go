package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"

	"github.com/nsf/termbox-go"
)

func main() {
	err := run()
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
}

func run() error {
	err := termbox.Init()
	if err != nil {
		return errors.Wrap(err, "initializing termbox")
	}
	defer termbox.Close()

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Ch == 'q' {
				return nil
			}
		}
	}
}
