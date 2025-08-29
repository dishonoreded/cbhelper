package cmd

type Options interface {
	Options() []string
	RetrieveExecution(option string) interface{}
}
