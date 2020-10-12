package env

import (
	"testing"
)

type misVariables struct {
	Temp   string `name:"TEMP,TMP"`
	Prueba int
}

func TestGestor(t *testing.T) {
	vars := &misVariables{}
	evh := New(vars)

	t.Log(evh.Valid())
	t.Log(evh.Results)
	t.Log("vars.Temp", vars.Temp)
	t.Log("vars.Prueba", vars.Prueba)
}
