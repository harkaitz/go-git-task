package gtask

import (
	"fmt"
	"strings"
	"time"
)

func (s Settings) PrintTasksTableHeader() {
	var field         string

	for _, field = range strings.Split(s.LsFields, ",") {
		fmt.Printf(fieldFormat(field), field)
	}
	fmt.Printf("\n\n")
}


func (i Tasks) PrintTable(s Settings) {
	var task          Task

	for _, task = range i.Tasks {
		task.PrintRow(s)
	}
}

func (i Tasks) PrintChangelog(version, line string) (errl []error) {
	var task          Task

	fmt.Printf("%v  %v  %v\n\n", time.Now().Format("2006-01-02"), version, line)
	for _, task = range i.Tasks {
		if task.Changelog == "" {
			errl = append(errl, fmt.Errorf("%v: No changelog", task.ID))
			continue
		}
		if task.Changelog != "" {
			fmt.Printf("	- %v (%v)\n", task.Changelog, task.ID)
		}
	}
	fmt.Printf("\n")

	return
}
