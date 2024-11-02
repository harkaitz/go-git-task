package gtask

import (
	"github.com/tcnksm/go-gitconfig"
	"encoding/json"
	"os"
	"fmt"
	"strings"
)

type Settings struct {
	Directory	string
	Project		string
	Editor		string
	States		string
	Reporter	string
	LsFields	string
	LsStates	string
	LsProject	string
	LsReporter	string
	LsAssignee	string
}

func (s *Settings) Init() (err error) {
	s.Directory  = getSetting("task.directory", "GIT_TASK_DIRECTORY", ".task")
	s.Project    = getSetting("task.project", "GIT_TASK_PROJECT", "")
	s.Editor     = getSetting("core.editor", "EDITOR", "vi")
	s.States     = getSetting("task.states", "GIT_TASK_STATES", "@new,@todo,@done,@closed,@invalid,@ongoing,@back")
	s.Reporter   = getSetting("user.reporter", "GIT_TASK_REPORTER,USER,USERNAME", "")
	s.LsFields   = getSetting("task.ls.fields",   "", "ID,Prio,Status,Project,Reporter,Assignee,SubjectSlug")
	s.LsStates   = getSetting("task.ls.states", "", "@back,@done,@todo,@ongoing")
	s.LsProject  = getSetting("task.ls.project",  "", "")
	s.LsReporter = ""
	s.LsAssignee = ""
	return
}

func (s Settings) String() string {
	var b	[]byte
	var err	error
	b, err = json.Marshal(s)
	if err != nil { return "" }
	json.MarshalIndent(b, "", "    ")
	return string(b)
}

func (s Settings) Println() {
	fmt.Printf(
		"GIT_TASK_DIRECTORY : task.directory  : %v\n" +
		"GIT_TASK_PROJECT   : task.project    : %v\n" +
		"EDITOR             : core.editor     : %v\n" +
		"GIT_TASK_STATES    : task.states     : %v\n" +
		"GIT_TASK_REPORTER  : user.reporter   : %v\n" +
		"USER[NAME]         : user.reporter   : %v\n" +
		"-                  : task.ls.fields  : %v\n" +
		"-                  : task.ls.states  : %v\n" +
		"-                  : task.ls.project : %v\n",
		s.Directory,
		s.Project,
		s.Editor,
		s.States,
		s.Reporter,
		s.Reporter,
		s.LsFields,
		s.LsStates,
		s.LsProject,
	)
}

// -------------------------------------------------------------------
// ---- Private functions -------------------------------------------
// -------------------------------------------------------------------

func getSetting(gitConfig, envConfig, defValue string) (value string) {
	var err error
	var env string

	if gitConfig != "" {
		value, err = gitconfig.Entire(gitConfig)
		if err == nil { return }
	}

	
	
	if envConfig != "" {
		for _, env = range strings.Split(envConfig, ",") {
			value = os.Getenv(env)
			if value != "" { return }
		}
	}

	value = defValue
	return
}

