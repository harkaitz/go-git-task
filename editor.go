package gtask

import (
	"os"
	"os/exec"
)

func (s Settings) OpenEditor(filename string) (err error) {
	var cmd		*exec.Cmd

	cmd = exec.Command("sh", "-ec", s.GetEditor() + " \"$1\"", "--", filename)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil { return }

	return
}

func (t *Task) Edit(s *Settings) (err error) {
	var filename string

	filename = t.Filename(s)

	err = s.OpenEditor(filename)
	if err != nil { return }

	err = t.ParseFile(filename)
	if err != nil { return }

	return
}
