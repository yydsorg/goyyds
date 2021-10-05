package tool

import (
	"fmt"
)

type Tool struct {
	Name string
	Type string
}

func (t *Tool) Init() error {
	switch t.Type {
	case "web":

	case "app":

	default:
		return fmt.Errorf("unkonwn tool type: %s", t.Type)
	}

	return nil
}
