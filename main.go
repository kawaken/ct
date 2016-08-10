package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/alecthomas/kingpin"
)

const defaultCountFile = ".config/ct/count"

var eLog = log.New(os.Stderr, "", 0)

var (
	mod           = kingpin.Flag("mod", "Number of modulo.").Short('m').Default("0").Int()
	countFilePath = kingpin.Flag("file", "File to save number.").Short('f').String()
	up            = kingpin.Command("up", "Count up the number.")
	upStep        = setStep(up)
	down          = kingpin.Command("down", "Count down the number.")
	downStep      = setStep(down)
	reset         = kingpin.Command("reset", "Reset the stored number.")
	resetNum      = reset.Arg("number", "Reset number.").Default("0").Int()
)

func setStep(c *kingpin.CmdClause) *int {
	return c.Arg("step", "Counting step.").Default("1").Int()
}

func openCountFile(path string) (*os.File, bool, error) {
	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); err != nil {
		err = os.MkdirAll(dir, 0777)
		if err != nil {
			return nil, false, err
		}
	}

	var f *os.File
	var err error
	var isNew bool

	if _, err = os.Stat(path); err == nil {
		f, err = os.OpenFile(path, os.O_RDWR, 0666)
	} else {
		f, err = os.Create(path)
		isNew = true
	}
	return f, isNew, err
}

func cmdMain() int {
	cmd := kingpin.Parse()

	path := filepath.Join(os.Getenv("HOME"), defaultCountFile)
	if *countFilePath != "" {
		path = *countFilePath
	}

	f, new, err := openCountFile(path)
	if err != nil {
		eLog.Printf("error open file '%s': %s", path, err)
		return 1
	}
	defer f.Close()

	var num int
	if new {
		num = -1
	} else {
		_, err = fmt.Fscanf(f, "%d", &num)
		if err != nil {
			eLog.Printf("error read file '%s': %s", path, err)
			return 1
		}
	}

	switch cmd {
	case "up":
		num += *upStep
	case "down":
		num -= *downStep
	case "reset":
		num = *resetNum
	}

	if *mod > 0 {
		num %= *mod
	}

	_, err = f.Seek(0, 0)
	if err != nil {
		eLog.Printf("error write file '%s': %s", path, err)
		return 1
	}

	_, err = fmt.Fprintf(f, "%d\n", num)
	if err != nil {
		eLog.Printf("error write file '%s': %s", path, err)
		return 1
	}

	fmt.Println(num)

	return 0
}

func main() {
	os.Exit(cmdMain())
}
