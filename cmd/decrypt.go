package cmd

import ()

type DecryptCmd struct {
	InFile  string `short:"i"  long:"infile"        description:"Specify an input file."`
	OutFile string `short:"o"  long:"outfile"        description:"Specify an input file."`
}

func (cmd *DecryptCmd) Execute([]string) error {
	return nil
}
