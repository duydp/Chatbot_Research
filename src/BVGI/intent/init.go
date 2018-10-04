package intent 

import (
	"fmt"
)

var engine Engine

func Init(name string, config map[string]string) (err error) {
	switch name {
	case "wit":
		engine, err = InitWit(config)
		if err != nil {
			return err
		}
	case "fptai":
		engine, err = NewFPTAI(config)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("engine %s does not exist\n", name)
	}

	return nil
}