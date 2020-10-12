package env

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

const (
	// ValidOk valor retornado cuando Valid() no encuentre errores.
	ValidOk = iota
	// ValidRequiredValue la variable de entorno no ha sido definida.
	ValidRequiredValue
	// ValidWrongValue la variable fue definida con un valor incorrecto.
	ValidWrongValue
)

// SetupEnvHandler la configuración necesaria para validar las variables de entorno.
type SetupEnvHandler struct {
	// Result de resultado obtenido: ValidOk, ValidWrongValue, etc.
	Result int
	// Value el valor real obtenido.
	Value interface{}
	// Error el error detectado al intentar obtener el valor.
	Error error
}

// -----------------------------------------------------------------------------

// EnvironmentVarsHandler objeto para la gestión de variables de entorno.
type EnvironmentVarsHandler struct {
	// Vars el mapa de variables obtenidas.
	Vars map[string]SetupEnvHandler

	conf *interface{}
}

// Valid permite validar los valores de las variables de entorno existentes.
func (evh *EnvironmentVarsHandler) Valid() int {
	dataStruct := reflect.Indirect(reflect.ValueOf(*evh.conf).Elem())
	evh.Vars = make(map[string]SetupEnvHandler)

	for i := 0; i < dataStruct.NumField(); i++ {
		prop := dataStruct.Type().Field(i)
		name := prop.Tag.Get("name")

		if name == "" {
			name = prop.Name
		}

		for _, e := range strings.Split(name, ",") {
			if value := os.Getenv(e); value != "" {
				switch prop.Type.Name() {
				case "string":
					evh.Vars[prop.Name] = SetupEnvHandler{ValidOk, value, nil}

				case "int":
					if realValue, err := strconv.Atoi(value); err == nil {
						evh.Vars[prop.Name] = SetupEnvHandler{ValidOk, realValue, nil}
					} else {
						evh.Vars[prop.Name] = SetupEnvHandler{ValidWrongValue, value, err}
					}

				case "float":
					if realValue, err := strconv.ParseFloat(value, 32); err == nil {
						evh.Vars[prop.Name] = SetupEnvHandler{ValidOk, realValue, nil}
					} else {
						evh.Vars[prop.Name] = SetupEnvHandler{ValidWrongValue, value, err}
					}

				case "bool":
					if realValue, err := strconv.ParseBool(value); err == nil {
						evh.Vars[prop.Name] = SetupEnvHandler{ValidOk, realValue, nil}
					} else {
						evh.Vars[prop.Name] = SetupEnvHandler{ValidWrongValue, value, err}
					}

				default:
					evh.Vars[prop.Name] = SetupEnvHandler{ValidWrongValue, value, fmt.Errorf("Wrong value (%s=\"%s\")", prop.Name, value)}
				}
			} else {
				evh.Vars[prop.Name] = SetupEnvHandler{ValidRequiredValue, value, fmt.Errorf("Required value (%s)", prop.Name)}
			}
		}
	}

	return len(evh.Vars)
}

// -----------------------------------------------------------------------------

// New retorna un objeto EnvVars.
func New(confVars interface{}) *EnvironmentVarsHandler {
	return &EnvironmentVarsHandler{conf: &confVars}
}
