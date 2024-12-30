package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Record struct {
	name, surname, tel string
}

func (r Record) String() string {
	return fmt.Sprintf("%20s %20s %20s\n", r.name, r.surname, r.tel)
}

type phonebook struct {
	records []Record
}

func (p *phonebook) add(name, surname, telephone string) {
	p.records = append(p.records, Record{
		name,
		surname,
		telephone,
	})
}

func (p *phonebook) search(name string) []Record {
	var output []Record
	for _, record := range p.records {
		if record.name == name {
			output = append(output, record)
		}
	}
	return output
}

func (p *phonebook) list() []Record {
	return p.records
}

func main() {
	log.SetFlags(0)
	var pBook phonebook
	pBook.records = []Record{
		{
			"raul",
			"gale",
			"96474747474",
		},
		{
			"hem",
			"kaps",
			"98804833838",
		},
	}
	if len(os.Args) < 2 {
		usage()
	}
	command := os.Args[1]
	commandArgs := os.Args[2:]
	switch command {
	case "add":
		assertRequiredArgs(command, commandArgs)
		pBook.add(commandArgs[0], commandArgs[1], commandArgs[2])
	case "search":
		assertRequiredArgs(command, commandArgs)
		records := pBook.search(commandArgs[0])
		print(records)
	case "list":
		assertRequiredArgs(command, commandArgs)
		print(pBook.list())
	default:
		usage()
	}
}

func usage() {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintln("Usage: phonebook <command>"))
	sb.WriteString(fmt.Sprintln("\tcommand:"))
	sb.WriteString(fmt.Sprintln("\t\t1. add <name> <surname> <telephone>:"))
	sb.WriteString(fmt.Sprintln("\t\t2. search <name>"))
	sb.WriteString(fmt.Sprintln("\t\t3. list"))
	log.Println(sb.String())
	os.Exit(1)
}

func print(records []Record) {
	if len(records) == 0 {
		log.Printf("No records found.")
		return
	}
	fmt.Printf("%20s %20s %20s\n", "name", "surname", "telephone")
	for _, r := range records {
		fmt.Printf("%s", r)
	}
}

func assertRequiredArgs(command string, args []string) {
	length := len(args)
	switch {
	case command == "add" && length != 3:
		usage()
	case command == "search" && length != 1:
		usage()
	case command == "list" && length != 0:
		usage()
	}
}
