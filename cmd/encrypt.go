package cmd

import ()

type EncryptCmd struct {
	Carrier        CarrierFlag `short:"c"  long:"carrier"        description:"Specify the gif file to use as the carrier for the message"`

	Message        string      `short:"m"  long:"message"      description:"Provide the message on the command line"`
	InFile         string      `short:"i"  long:"infile"        description:"Specify an input file containing the plaintext message"`
	OutFile        string      `short:"o"  long:"outfile"        description:"Specify an output file to save the gif containing the encrypted message"`

	NoSelf         bool        `long:"no-self"        description:"Don't encrypt for yourself"`
	HideRecipients bool        `long:"hide-recipients" description:"Don't include recipients in metadata"`
	Anonymous      bool        `long:"anonymous" description:"Don't include sender or recipients in metadata. Implies --hide-recipients"`

}

func (cmd *EncryptCmd) Execute([]string) error {
	return nil
}
