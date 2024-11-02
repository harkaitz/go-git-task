GIT TASKS
=========

Simple task manager integrated with git.

## The ".task" file format.

The tasks are saved in files that match: ".task/@STATUS/@ID_SLUG.task" with
the following format:

    Project: git-task
    Type: bug
    Subject: We shall fix this bug.
    Public: no
    Priority: 0
    Assignee: Harkaitz
    Reporter: Sandra
    Changelog: The text to put in changelog.
    Version: 0.1
    
    This is the description

## Go programs

    Usage: git-task SUBCOMMAND [ARGS...]
    
    This program is a task manager for git. It uses a directory ".task" to
    store the tasks as ".task/@STATUS/@ID_SLUG.task":
    
      show                Show configuration.
      new                 Create new task.
      edit [ID]           Edit task (the new one or the ongoing if none specified)
      ls [--help]         List tasks.
      @STATUS ID...       Move tasks between different status.
      rename ID SLUG      Set slug for a task.
      view [ID]           View task (by default ongoing).
      changelog VER LINE  Print changelog section for version.
    
    Statuses: %v
    Fields: %v

## Go documentation

    package gtask // import "github.com/harkaitz/go-git-task"
    
    type Settings struct{ ... }
    type Task struct{ ... }
    type Tasks struct{ ... }

## Collaborating

For making bug reports, feature requests and donations visit
one of the following links:

1. [gemini://harkadev.com/oss/](gemini://harkadev.com/oss/)
2. [https://harkadev.com/oss/](https://harkadev.com/oss/)
