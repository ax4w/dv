package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func MustReadString(r *bufio.Reader, delim byte) string {
	s, err := r.ReadString(delim)
	if err != nil {
		panic(err.Error())
	}
	return strings.TrimSpace(s)
}

func main() {
	var (
		iter0 = make(iteration)
		rows  []string
		cols  []string
	)
	r := bufio.NewReader(os.Stdin)
	println("Distance-Vector")
	println("Enter the tables you want (f.e A,B,C,D):")
	tables := strings.Split(MustReadString(r, '\n'), ",")
	println("setup rows for each table (f.E a,b,c,d). The rows represent the DESTINATION")
	rows = strings.Split(MustReadString(r, '\n'), ",")
	for _, k := range tables {
		println("setting up", k)
		currTable := make(table)
		for _, v := range rows {
			currTable[v] = make(map[string]int)
			fmt.Printf("Added row %s to table %s (%p)\n", v, k, &currTable)
		}
		println("setup columns for each table (f.E a,b,c,d). The rows represent the VIA")
		cols = strings.Split(MustReadString(r, '\n'), ",")
		for r, v := range currTable {
			for _, c := range cols {
				v[c] = inf
				fmt.Printf("Added col %s to row %s on table %s\n", c, r, k)
			}

		}
		iter0[k] = currTable
	}
	println("now setup the default values like: A C B 5 (Table a, to c, via b has 5 costs)")
	println("type break to end")
	for {
		print(">")
		input := MustReadString(r, '\n')
		if input == "break" {
			break
		}
		args := strings.Split(input, " ")
		if len(args) != 4 {
			println("invalid input, try again")
			continue
		}
		table, ok := iter0[args[0]]
		if !ok {
			fmt.Printf("Table %s does not exists", args[0])
			continue
		}
		row, ok := table[args[1]]
		if !ok {
			fmt.Printf("Table %s does not have row %s", args[0], args[1])
			continue
		}
		_, ok = row[args[2]]
		if !ok {
			println("Table", table, "does not have col", args[2])
			continue
		}
		i, err := strconv.Atoi(strings.TrimSpace(args[3]))
		if err != nil {
			println(err.Error())
			continue
		}
		iter0[args[0]][args[1]][args[2]] = i
	}
	Calculate(iter0)
}
