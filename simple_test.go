package env

import (
	"fmt"
	"testing"
)

type misVariables struct {
	temp   string `name:"TEMP,TMP"`
	prueba int
}

func TestGestor(t *testing.T) {
	evh := New(&misVariables{})

	fmt.Println(evh.Valid())
	fmt.Println(evh.Vars)
	fmt.Println(evh.Vars["prueba"])
}
