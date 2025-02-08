package gtask

import (
	"os"
	"io/fs"
	"strings"
	"errors"
	"fmt"
	"path/filepath"
)

type Tasks struct {
	Tasks	[]Task
}

func (s *Settings) ListTasks() (t Tasks, err error) {
	var task          Task
	var filesChan     chan string
	var filename      string
	var errl        []error

	filesChan = s.TaskFiles()

	t.Tasks = []Task{}
	for filename = range filesChan {
		task = Task{}
		err = task.ParseFile(filename)
		if err != nil { errl = append(errl, err); continue }
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

func (i Tasks) FilterBySettings(s *Settings) (o Tasks) {
	var status        string
	var task          Task

	o.Tasks = []Task{}

	s.GetLsProject()

	for _, status = range strings.Split(s.GetLsStates(), ",") {
		for _, task = range i.Tasks {
			if s.LsProject != "" && !strings.EqualFold(task.Project, s.LsProject) {
				continue
			}
			if s.LsReporter != "" && !strings.EqualFold(task.Reporter, s.LsReporter) {
				continue
			}
			if s.LsAssignee != "" && !strings.EqualFold(task.Assignee, s.LsAssignee) {
				continue
			}
			if s.LsChangelog != "" && !strings.EqualFold(task.Changelog, s.LsChangelog) {
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
		return
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

func (s *Settings) TaskFiles() (files chan string) {
	files = make(chan string, 20)

	go func(dd string) {
		var e1l, e2l     []fs.DirEntry
		var e1, e2        fs.DirEntry
		var err           error
		var p1            string

		e1l, err = os.ReadDir(s.GetDirectory())
		if err != nil { return }

		for _, e1 = range e1l {
			if e1.IsDir() && e1.Name()[0] == '@' {
				p1 = dd + "/" + e1.Name()
				e2l, err = os.ReadDir(p1)
				if err != nil { continue }
				for _, e2 = range e2l {
					if !e2.IsDir() && filepath.Ext(e2.Name()) == ".task" {
						files <- p1 + "/" + e2.Name()
					}
				}
			}
		}

		close(files)
	}(s.GetDirectory())

	return
}

func (s Settings) PrintTasksTableHeader() {
	var field         string

	for _, field = range strings.Split(s.GetLsFields(), ",") {
		fmt.Printf(fieldFormat(field), field)
	}
	fmt.Printf("\n\n")
}


func (i Tasks) PrintTable(s *Settings) {
	var task          Task

	for _, task = range i.Tasks {
		task.PrintRow(s)
	}
}

