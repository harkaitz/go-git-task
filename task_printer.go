package gtask

import (
	"fmt"
	"strings"
)

// -------------------------------------------------------------------

func (t Task) PrintRow(s Settings) {
	var field	string
	var format	string
	for _, field = range strings.Split(s.LsFields, ",") {
		format = fieldFormat(field)
		switch field {
		case "SubjectSlug": fmt.Printf("%s (%s)", t.Subject, t.Slug)
		//
		case "ID":          fmt.Printf(format, str2(t.ID))
		case "Slug":        fmt.Printf(format, str2(t.Slug))
		case "Status":      fmt.Printf(format, str2(t.Status))
		//
		case "Project":     fmt.Printf(format, str2(t.Project))
		case "Type":        fmt.Printf(format, str2(t.Type))
		case "Subject":     fmt.Printf(format, str2(t.Subject))
		case "Public":      fmt.Printf(format, bool2(t.Public))
		case "Prio", "Priority": fmt.Printf(format, t.Priority)
		case "Assignee":    fmt.Printf(format, str2(t.Assignee))
		case "Reporter":    fmt.Printf(format, str2(t.Reporter))
		case "Changelog":   fmt.Printf(format, str2(t.Changelog))
		case "Version":     fmt.Printf(format, str2(t.Version))
		//
		default:            fmt.Printf(format, "???")
		}
	}
	fmt.Println()
}

// -------------------------------------------------------------------

func fieldFormat(field string) string {
	switch field {
	case "ID":           return "%-8v "
	case "Prio":         return "%-5v "
	case "Status":       return "%-10v "
	case "Project":      return "%-10v "
	case "Reporter":     return "%-10v "
	case "Assignee":     return "%-10v "
	case "Subject":      return "%v "
	default:             return "%v "
	}
}

func str2(s string) string {
	switch s {
	case "": return "-"
	default: return s
	}
}

func bool2(b bool) string {
	switch b {
	case true: return "yes"
	default:   return "no"
	}
}

