package cmd

type GifBaseCmd struct {
	Encrypt EncryptCmd `command:"encrypt" alias:"e" description:"encrypt a message and hide it in a gif"`
	Decrypt DecryptCmd `command:"decrypt" alias:"d" description:"decrypt a message hidden in a gif"`
}
