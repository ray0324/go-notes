package main

import (
	"flag"
	"fmt"
	"go-notes/10-protobuf/todo"
	"os"
	"strings"

	"github.com/golang/protobuf/proto"
)

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "miss subcommand:list or add")
		os.Exit(1)
	}

	var err error
	switch cmd := flag.Arg(0); cmd {
	case "list":
		err = list()
	case "add":
		err = add(strings.Join(flag.Args()[1:], " "))
	default:
		err = fmt.Errorf("Unkonw subcommand: %s", cmd)
	}

	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println("TODO")

}

const db = "db.pb"

func add(text string) error {
	task := &todo.Task{
		Text: text,
		Done: false,
	}

	b, err := proto.Marshal(task)

	if err != nil {
		return fmt.Errorf("could not encode task: %v", err)
	}

	f, err := os.OpenFile(db, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		return fmt.Errorf("could not open file: %v", err)
	}

	_, err = f.Write(b)

	if err != nil {
		return fmt.Errorf("could not write task to file: %v", err)
	}

	if err := f.Close(); err != nil {
		return fmt.Errorf("could not close file %s: %v", db, err)
	}

	fmt.Println(proto.MarshalTextString(task))
	return nil
}

func list() error {
	return nil
}
