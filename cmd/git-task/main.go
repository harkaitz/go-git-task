package main

import (
	"os"
	"fmt"
	"strings"
	"github.com/harkaitz/go-git-task"
)

const help string =
`Usage: git task SUBCOMMAND [ARGS...]

This program is a task manager for git. It uses a directory ".task" to
store the tasks as ".task/@STATUS/@ID_SLUG.task":

  show                 Show configuration.
  new                  Create new task.
  edit [ID]            Edit task (the new one or the ongoing if none specified)
  ls [--help]          List tasks.
  lst                  List @todo,@ongoing and @done.
  lsn                  List @new tasks.
  @STATUS ID...        Move tasks between different status.
  rename ID SLUG       Set a slug to a task.
  remove ID            Remove a task.
  view [ID]            View task (by default ongoing).
  changelog CHANGELOG  Print changelog section.

Statuses: @new,@todo,@done,@closed,@invalid,@ongoing,@back
Fields: ID,Prio,Status,Project,Reporter,Assignee,SubjectSlug

Copyright (c) 2024 - Harkaitz Agirre - All rights reserved.`

const lsHelp string =
`git task ls [OPTIONS...]

List tasks.

  f=FIELDS,...  Fields to show.     r=REPORTER   Reporter to show.
  s=STATUS,...  Status to show.     a=ASSIGNEE   Assignee to show.
  p=PROJECT     Project to show.    c=CHANGELOG  Repo name for changelog.

Statuses: %v
Fields: %v
`

var S gtask.Settings

func main() {
	var err           error
	var cmd           string
	var args        []string

	// Error manager.
	defer func() {
		if err != nil {
			fmt.Fprintf(os.Stderr, "git-task: error: %v.\n", err.Error())
			os.Exit(1)
		}
	}()

	// Initialize the settings.
	err = S.Init()
	if err != nil { return }

	// Parse the command line.
	if len(os.Args) <= 1 || os.Args[1] == "--help" || os.Args[1] == "-h" {
		fmt.Println(help)
		return
	}
	cmd = os.Args[1]
	args = os.Args[2:]

	// Execute the command.
	switch {
	case cmd == "show":  err = Show()
	case cmd == "new":   err = New()
	case cmd == "edit":  err = Edit(args)
	case cmd == "ls":    err = Ls(args)
	case cmd == "lst":   err = Ls(append([]string {"s=@todo,@ongoing,@done"}, args...))
	case cmd == "lsn":   err = Ls(append([]string {"s=@new"}, args...))
	case len(cmd) > 1 && cmd[0] == '@': err = Status(cmd, args)
	case cmd == "rename":    err = Rename(args)
	case cmd == "remove":    err = Remove(args)
	case cmd == "view":      err = View(args)
	case cmd == "changelog": err = Changelog(args)
	default:                 err = fmt.Errorf("Unknown command: %s", cmd)
	}

	return
}

func Show() (err error) {
	S.Println()
	return
}

func New() (err error) {
	var task          gtask.Task
	var filename      string

	task.Init(S)
	filename, err = task.Save(&S)
	if err != nil { return }
	fmt.Println("Task created:", filename)

	return
}

func Edit(args []string) (err error) {
	var tasks         gtask.Tasks
	var task         *gtask.Task
	var found         bool
	var arg           string

	tasks, err = S.ListTasks()
	if err != nil { return }

	if len(args) == 0 {
		task, found, err = tasks.FilterByStatus("@new").First("No new tasks found, create with 'new'")
		if err != nil { return }
		if !found {
			err = fmt.Errorf("Task not found")
			return
		}
		fmt.Print(task)
		err = task.Edit(&S)
		if err != nil { return }
		return
	}

	for _, arg = range args {
		task, found, err = tasks.SearchByID(arg)
		if err != nil { return }
		if !found {
			err = fmt.Errorf("Task not found")
			return
		}
		fmt.Print(task)
		err = task.Edit(&S)
		if err != nil { return }
	}

	return
}

func Ls(args []string) (err error) {
	var tasks         gtask.Tasks
	var arg           string
	var parts       []string

	S.LsProject = S.GetProject()

	for _, arg = range args {
		parts = strings.SplitN(arg, "=", 2)
		if len(parts) != 2 || parts[0] == "help" || parts[0] == "-h" {
			fmt.Printf(lsHelp, S.GetLsStates(), S.GetLsFields())
			return
		}
		switch parts[0] {
		case "f": S.LsFields   = parts[1]
		case "s": S.LsStates   = parts[1]
		case "p": S.LsProject  = parts[1]
		case "r": S.LsReporter = parts[1]
		case "a": S.LsAssignee = parts[1]
		case "c": S.LsChangelog = parts[1]
		default: err = fmt.Errorf("Unknown option: %s", arg); return
		}
	}

	tasks, err = S.ListTasks()
	if err != nil { return }

	S.PrintTasksTableHeader()
	tasks.FilterBySettings(&S).PrintTable(&S)

	return
}

func Status(status string, args []string) (err error) {
	var tasks         gtask.Tasks
	var task         *gtask.Task
	var arg           string

	tasks, err = S.ListTasks()
	if err != nil { return }

	for _, arg = range args {

		task, _, err = tasks.SearchByID(arg)
		if err != nil { return }

		err = task.MoveStatus(&S, status)
		if err != nil { return }

	}
	
	return
}

func Rename(args []string) (err error) {
	var tasks         gtask.Tasks
	var task         *gtask.Task

	if len(args) < 2 {
		err = fmt.Errorf("Not enough arguments")
		return
	}

	tasks, err = S.ListTasks()
	if err != nil { return }

	task, _, err = tasks.SearchByID(args[0])
	if err != nil { return }

	err = task.MoveRename(&S, strings.Join(args[1:], "_"))
	if err != nil { return }

	return
}

func Remove(args []string) (err error) {
	var tasks         gtask.Tasks
	var task         *gtask.Task
	var arg           string

	for _, arg = range args {
		tasks, err = S.ListTasks()
		if err != nil { return }
		task, _, err = tasks.SearchByID(arg)
		if err != nil { return }
		err = task.Remove(&S)
		if err != nil { return }
	}

	return
}

func View(args []string) (err error) {
	var tasks         gtask.Tasks
	var task         *gtask.Task

	tasks, err = S.ListTasks()
	if err != nil { return }

	switch len(args) {
	case 0: task, _, err = tasks.FilterByStatus("@ongoing").First("No @ongoing tasks found")
	case 1: task, _, err = tasks.SearchByID(args[0])
	default:         err = fmt.Errorf("Too many arguments")
	}
	if err != nil { return }
	
	task.PrintTable()
	return
}

func Changelog(args []string) (err error) {
	var tasks         gtask.Tasks
	var task          gtask.Task

	if len(args) < 1 {
		err = fmt.Errorf("Not enough arguments")
		return
	}

	S.LsProject = S.GetProject()
	S.LsChangelog = args[0]
	S.LsStates = "@done"

	tasks, err = S.ListTasks()
	if err != nil { return }

	for _, task = range tasks.FilterBySettings(&S).Tasks {
		if task.Public {
			fmt.Printf("- (%v) %v (public)\n", task.ID, task.Subject)
		} else {
			fmt.Printf("- (%v) %v\n", task.ID, task.Subject)
		}
	}

	return
}
