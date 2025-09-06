package engine

import (
	"fmt"
	"log"
)

func Log(values ...any) (any, error) {
	if len(values) == 0 {
		return nil, fmt.Errorf("NOTHING TO LOG")
	}

	switch len(values) {
	case 0:
		return nil, fmt.Errorf("NOTHING TO LOG")
	case 1:
		log.Println(values...)

		return values[0], nil
	case 2:
		format, ok := values[0].(string)
		if !ok {
			log.Println(values...)
		} else {
			log.Printf(format, values[1])
		}

		return values[1], nil
	default:
		return nil, fmt.Errorf("too many args")
	}
}
