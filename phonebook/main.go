package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Record struct {
	name, surname string
	tel           int
}

type phonebook []Record

func (p *phonebook) add()    {}
func (p *phonebook) search() {}
func (p *phonebook) list() []Record {
	return make([]Record, 0)
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}
}

func usage() {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintln("Usage: phonebook <command>"))
	sb.WriteString(fmt.Sprintln("\tcommand:"))
	sb.WriteString(fmt.Sprintln("\t\t1. add:"))
	sb.WriteString(fmt.Sprintln("\t\t2. search:"))
	sb.WriteString(fmt.Sprintln("\t\t3. list:"))
	log.Println(sb.String())
}
