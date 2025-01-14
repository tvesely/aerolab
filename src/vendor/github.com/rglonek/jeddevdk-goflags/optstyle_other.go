//go:build !windows || forceposix
// +build !windows forceposix

package flags

import (
	"strings"
)

const (
	defaultShortOptDelimiter = '-'
	defaultLongOptDelimiter  = "--"
	defaultNameArgDelimiter  = '='
)

func argumentStartsOption(arg string) bool {
	return len(arg) > 0 && arg[0] == '-'
}

func argumentIsOption(arg string) bool {
	if len(arg) > 1 && arg[0] == '-' && arg[1] != '-' {
		return true
	}

	if len(arg) > 2 && arg[0] == '-' && arg[1] == '-' && arg[2] != '-' {
		return true
	}

	return false
}

// stripOptionPrefix returns the option without the prefix and whether or
// not the option is a long option or not.
func stripOptionPrefix(optname string) (prefix string, name string, islong bool) {
	if strings.HasPrefix(optname, "--") {
		return "--", optname[2:], true
	} else if strings.HasPrefix(optname, "-") {
		return "-", optname[1:], false
	}

	return "", optname, false
}

// splitOption attempts to split the passed option into a name and an argument.
// When there is no argument specified, nil will be returned for it.
func splitOption(prefix string, option string, islong bool) (string, string, *string) {
	pos := strings.Index(option, "=")

	if (islong && pos >= 0) || (!islong && pos == 1) {
		rest := option[pos+1:]
		return option[:pos], "=", &rest
	}

	return option, "", nil
}

var globalsAdded = false

// addHelpGroup adds a new group that contains default help parameters.
func (c *Command) addHelpGroup(showHelp func() error) *Group {
	var help struct {
		ShowHelp func() error `hidden:"true" short:"h" long:"help" description:"Show this help message"`
	}

	var globals struct {
		Beep  bool `long:"beep" description:"cause the terminal to beep on exit; if specified multiple times, will be once on success and >1 on failure"`
		Beepf bool `long:"beepf" description:"like beep, but does not trigger beep on success, only failures"`
	}

	help.ShowHelp = showHelp
	ret, _ := c.AddGroup("Help Options", "", &help)
	ret.isBuiltinHelp = true
	ret.Hidden = true
	if !globalsAdded {
		globalsAdded = true
		c.AddGroup("Global Options", "", &globals)
	}
	return ret
}
