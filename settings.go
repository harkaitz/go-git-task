package gtask

import (
	"github.com/tcnksm/go-gitconfig"
	"os"
	"fmt"
	"strings"
)

type Settings struct {
	Directory         string
	Project           string
	Editor            string
	States            string
	Reporter          string
	LsFields          string
	LsStates          string
	LsProject         string
	LsReporter        string
	LsAssignee        string
	LsChangelog       string
	IDCharacters      string
}

func (s *Settings) Init() (err error) {
	s.LsReporter = ""
	s.LsAssignee = ""
	s.IDCharacters = "0123456789"
	return
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
		s.GetDirectory(),
		s.GetProject(),
		s.GetEditor(),
		s.GetStates(),
		s.GetReporter(),
		s.GetReporter(),
		s.GetLsFields(),
		s.GetLsStates(),
		s.GetLsProject(),
	)
}

func (s *Settings) GetDirectory() string {
	return getSetting(&s.Directory, "task.directory", "GIT_TASK_DIRECTORY", ".task")
}

func (s *Settings) GetProject() string {
	return getSetting(&s.Project, "task.project", "GIT_TASK_PROJECT", "")
}

func (s *Settings) GetEditor() string {
	return getSetting(&s.Editor, "core.editor", "EDITOR", "vi")
}

func (s *Settings) GetStates() string {
	return getSetting(&s.States, "task.states", "GIT_TASK_STATES", "@new,@todo,@done,@closed,@invalid,@ongoing,@back")
}

func (s *Settings) GetReporter() string {
	return getSetting(&s.Reporter, "user.reporter", "GIT_TASK_REPORTER,USER,USERNAME", "")
}

func (s *Settings) GetLsFields() string {
	return getSetting(&s.LsFields, "task.ls.fields",   "", "ID,Prio,Status,Project,Changelog,Assignee,SubjectSlug")
}

func (s *Settings) GetLsStates() string {
	return getSetting(&s.LsStates, "task.ls.states", "", "@back,@done,@todo,@ongoing")
}

func (s *Settings) GetLsProject() string {
	return getSetting(&s.LsProject, "task.ls.project",  "", "")
}

// -------------------------------------------------------------------
// ---- Private functions -------------------------------------------
// -------------------------------------------------------------------

func getSetting(sd *string, gitConfig, envConfig, defValue string) string {
	var err           error
	var env           string

	if *sd != "" {
		return *sd
	}

	if gitConfig != "" {
		*sd, err = gitconfig.Entire(gitConfig)
		if err == nil { return *sd }
	}

	if envConfig != "" {
		for _, env = range strings.Split(envConfig, ",") {
			*sd = os.Getenv(env)
			if *sd != "" { return *sd }
		}
	}

	*sd = defValue
	return *sd
}

