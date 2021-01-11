package b

import (
	c "github.com/jeffgreenca/gd-c/pkg/c"
)

func B() string {
	return "b" + c.C()
}
