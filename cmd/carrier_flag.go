package cmd

type CarrierFlag string

func (f *CarrierFlag) UnmarshalFlag(value string) error {
	return nil
}
