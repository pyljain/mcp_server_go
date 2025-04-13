package tools

type Tool interface {
	Execute(arguments map[string]interface{}) (string, error)
}
