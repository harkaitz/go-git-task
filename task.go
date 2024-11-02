package gtask

import (
	"math/rand"
	"strings"
	"strconv"
	"os"
	"fmt"
)

type Task struct {
	ID          string
	Slug        string
	Status      string
	
	Project     string
	Type        string
	Subject     string
	Public      bool
	Priority    int
	Assignee    string
	Reporter    string
	Changelog   string
	Version     string
	Description string
}


func (t *Task) Init(s Settings) {
	t.ID = getNewID(s, 6)
	t.Slug = "no_name"
	t.Status = "@new"
	//
	t.Project = s.Project
	t.Type = "task"
	t.Subject = "No subject"
	t.Public = false
	t.Priority = 0
	t.Assignee = "nobody"
	t.Reporter = s.Reporter
	t.Changelog = ""
	t.Version = ""
	t.Description = ""
}

func (t *Task) ParseString(data string) (err error) {
	var lines       []string
	var line          string
	var field         string
	var value         string
	var lineNo        int
	var found         bool

	// Split the string into lines
	lines = strings.Split(data, "\n")

	// Parse the lines
	for lineNo, line = range lines {

		if line == "" {
			t.Description = strings.Join(lines[lineNo+1:], "\n")
			break
		}

		field, value, found = getField(line)
		if !found { continue }

		switch field {
		case "Project":   t.Project = value
		case "Type":      t.Type = value
		case "Subject":   t.Subject = value
		case "Public":    t.Public = (value == "yes")
		case "Priority":  t.Priority, _ = strconv.Atoi(value)
		case "Assignee":  t.Assignee = value
		case "Reporter":  t.Reporter = value
		case "Changelog": t.Changelog = value
		case "Version":   t.Version = value
		}

	}

	return
}

func (t *Task) ParseFile(filename string) (err error) {
	var data        []byte
	var parts1      []string
	var parts2      []string

	parts1 = strings.Split(filename, "/")
	if len(parts1) < 2 {
		err = fmt.Errorf("%v: Invalid path", filename)
		return
	}
	parts2 = strings.SplitN(parts1[len(parts1)-1], "_", 2)
	if len(parts2) != 2 {
		err = fmt.Errorf("%v: Invalid filename", parts1[len(parts1)-1])
		return
	}

	t.Status = parts1[len(parts1)-2]
	t.ID = parts2[0]
	t.Slug = strings.SplitN(parts2[1], ".", 2)[0]

	if len(t.Status) < 4 || t.Status[0] != '@' || len(t.ID) < 4 || t.ID[0] != '@' {
		err = fmt.Errorf("%v: Invalid status or ID", filename)
		return
	}

	data, err = os.ReadFile(filename)
	if err != nil { return }
	return t.ParseString(string(data))
}

func (t *Task) String() (data string) {
	data  = "Project: " + t.Project + "\n"
	data += "Type: " + t.Type + "\n"
	data += "Subject: " + t.Subject + "\n"
	data += "Public: " + strconv.FormatBool(t.Public) + "\n"
	data += "Priority: " + strconv.Itoa(t.Priority) + "\n"
	data += "Assignee: " + t.Assignee + "\n"
	data += "Reporter: " + t.Reporter + "\n"
	data += "Changelog: " + t.Changelog + "\n"
	data += "Version: " + t.Version + "\n"
	data += "\n"
	data += t.Description
	return
}

func (t *Task) Save(s Settings) (filename string, err error) {
	var data          string
	var fp           *os.File

	filename = t.Filename(s)
	data = t.String()

	err = os.MkdirAll(s.Directory + "/" + t.Status, 0755)
	if err != nil { return }

	fp, err = os.Create(filename)
	if err != nil { return }
	defer fp.Close()

	_, err = fp.WriteString(data)
	return
}

// -------------------------------------------------------------------
// ---- File renaming operations -------------------------------------
// -------------------------------------------------------------------

func (t *Task) Directory(s Settings) (directory string) {
	directory = s.Directory + "/" + t.Status
	return
}

func (t *Task) Filename(s Settings) (filename string) {
	filename = t.Directory(s) + "/" + t.ID + "_" + t.Slug + ".task"
	return
}

func (t *Task) MoveStatus(s Settings, status string) (err error) {
	var fr, to, dir   string

	err = t.CheckNewStatus(s, status)
	if err != nil { return }

	fr = t.Filename(s)
	t.Status = status
	dir = t.Directory(s)
	to = t.Filename(s)

	err = os.MkdirAll(dir, 0755)
	if err != nil { return }

	err = os.Rename(fr, to)
	if err != nil { return }
	
	return
}

func (t *Task) MoveRename(s Settings, slug string) (err error) {
	var fr, to, dir   string

	fr = t.Filename(s)
	t.Slug = slug
	dir = t.Directory(s)
	to = t.Filename(s)

	err = os.MkdirAll(dir, 0755)
	if err != nil { return }

	err = os.Rename(fr, to)
	if err != nil { return }
	
	return
}

func (t *Task) CheckNewStatus(s Settings, status string) (err error) {
	var state         string

	for _, state = range strings.Split(s.States, ",") {
		if state == status { return }
	}

	err = fmt.Errorf("Invalid status: %s", status)
	return
}


// -------------------------------------------------------------------
// ---- Private functions --------------------------------------------
// -------------------------------------------------------------------

func getField(line string) (field, value string, found bool) {
	var parts       []string

	// Split by the first ":"
	parts = strings.SplitN(line, ":", 2)
	if len(parts) != 2 {
		return
	}

	// Remove everything after the first "#"
	parts[1] = strings.Split(parts[1], "#")[0]

	// Trim the parts
	field = strings.TrimSpace(parts[0])
	value = strings.TrimSpace(parts[1])
	found = true
	return
}

func getNewID(s Settings, n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = s.IDCharacters[rand.Intn(len(s.IDCharacters))]
    }
    return "@" + string(b)
}
