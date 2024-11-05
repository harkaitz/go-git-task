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

## Programs

    Usage: git task SUBCOMMAND [ARGS...]
    
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
    
    Statuses: @new,@todo,@done,@closed,@invalid,@ongoing,@back
    Fields: ID,Prio,Status,Project,Reporter,Assignee,SubjectSlug

## Installation on MS Windows machines (v10 and above)

Download and extract the following zip and place the "bin" directory
into your path.

- https://github.com/harkaitz/go-git-task/releases/download/v1.0.5/git-task-1.0.5_Windows_NT_x86_64.zip
- [How to add a directory to your Path](https://learn.microsoft.com/en-us/previous-versions/office/developer/sharepoint-2010/ee537574(v=office.14))

We provide a simple ".bat" file that downloads and installs the program
in your local user directory (%LOCALAPPDATA%).

- https://github.com/harkaitz/go-git-task/releases/download/v1.0.5/git-task-1.0.5_Windows_NT_x86_64.bat

Simply execute and type enter when asked.

## Installation on POSIX machines

Execute the following commands (or equivalent) to install the program
in "/usr/local":

    $ v="1.0.5"
    $ u="https://github.com/harkaitz/go-git-task/releases/download/v${v}/git-task-${v}_$(uname -s)_$(uname -m).tar.gz"
    $ curl -L -o /tmp/git-task.tar.gz "${u}"
    $ sudo tar xf /tmp/git-task.tar.gz -C /

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
