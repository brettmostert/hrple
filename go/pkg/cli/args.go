package cli

type ArgSet struct {
	args map[string]*Arg
}

func (cmd *Command) Args() *ArgSet {
	if cmd.argSet == nil {
		cmd.argSet = &ArgSet{}
		cmd.argSet.args = make(map[string]*Arg)
	}

	return cmd.argSet
}

type Arg struct {
	Name  string // name as it appears on command line
	value string // value as set
}

func (set *ArgSet) Get(name string) string {
	value := ""
	for _, arg := range set.args {
		if arg.Name == name {
			value = arg.value
			break
		}
	}
	return value
}

func (set *ArgSet) Set(name string) {
	set.args[name] = &Arg{Name: name}
}
