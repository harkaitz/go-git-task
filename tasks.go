package gtask

import (
	"os/exec"
	"strings"
	"errors"
	"fmt"
)

type Tasks struct {
	Tasks	[]Task
}


func (s Settings) ListTasks() (t Tasks, err error) {
	var cmd          *exec.Cmd
	var out         []byte
	var task          Task
	var filename      string
	var errl        []error

	cmd = exec.Command("find", s.Directory, "-type", "f", "-name", "*.task")
	out, err = cmd.Output()
	if err != nil { return }

	t.Tasks = []Task{}
	for _, filename = range strings.Split(string(out), "\n") {
		if filename == "" { continue }
		task = Task{}
		err = task.ParseFile(filename)
		if err != nil { errl = append(errl, err) }
		t.Tasks = append(t.Tasks, task)
	}

	if errl != nil { err = errors.Join(errl...) }

	return
}

func (i Tasks) FilterByStatus(status string) (o Tasks) {
	var task          Task

	o.Tasks = []Task{}

	for _, task = range i.Tasks {
		if task.Status == status {
			o.Tasks = append(o.Tasks, task)
		}
	}

	return
}

func (i Tasks) FilterByVersionPublic(version string) (o Tasks) {
	var task          Task

	o.Tasks = []Task{}

	for _, task = range i.Tasks {
		if task.Version == version {
			o.Tasks = append(o.Tasks, task)
		}
	}

	return
}

func (i Tasks) FilterBySettings(s Settings) (o Tasks) {
	var status        string
	var task          Task

	o.Tasks = []Task{}

	for _, status = range strings.Split(s.LsStates, ",") {
		for _, task = range i.Tasks {
			if s.LsProject != "" && task.Project != s.LsProject {
				continue
			}
			if s.LsReporter != "" && task.Reporter != s.LsReporter {
				continue
			}
			if s.LsAssignee != "" && task.Assignee != s.LsAssignee {
				continue
			}
			if task.Status == status {
				o.Tasks = append(o.Tasks, task)
			}
		}
	}

	return
}

func (i Tasks) First(errm string, a ...any) (task *Task, found bool, err error) {
	if len(i.Tasks) >= 1 {
		task = &i.Tasks[0]
		found = true
	}
	if errm != "" {
		err = fmt.Errorf(errm, a...)
	}
	return
}

func (i Tasks) SearchByID(id string) (t *Task, found bool, err error) {
	var n             int

	for n = range i.Tasks {
		if i.Tasks[n].ID == id {
			return &i.Tasks[n], true, nil
		}
	}

	err = fmt.Errorf("Task not found: %s", id)
	return
}
