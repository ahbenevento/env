package env // Permite gestionar una lista de variables de entorno.

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

const (
	// ValidOk valor revisado y confirmado.
	ValidOk = iota
	// ValidRequiredValue la variable de entorno no ha sido definida.
	ValidRequiredValue
	// ValidWrongValue la variable fue definida con un valor incorrecto.
	ValidWrongValue
)

// -----------------------------------------------------------------------------

// Result estructura que informa la validación de cada una de las variables de entorno.
type Result struct {
	// Result resultado obtenido: ValidOk, ValidRequiredValue, ValidWrongValue.
	Result int
	// EnvValue el valor original obtenido.
	EnvValue string
	// Error el error detectado al intentar obtener el valor.
	Error error
}

// -----------------------------------------------------------------------------

// EnvironmentVarsHandler objeto para la gestión de variables de entorno.
type EnvironmentVarsHandler struct {
	// Vars el mapa de variables obtenidas.
	Result map[string]Result

	conf *interface{}
}

// Valid verifica una o más variables de entorno.
func (evh *EnvironmentVarsHandler) Valid() bool {
	evh.Result = make(map[string]Result)
	dataStruct := reflect.ValueOf(*evh.conf).Elem()
	validated := true

	for i := 0; i < dataStruct.NumField(); i++ {
		property := dataStruct.Type().Field(i)
		name := property.Tag.Get("name")

		if name == "" {
			name = property.Name
		}

		for _, e := range strings.Split(name, ",") {
			var valueError error

			status := ValidOk
			value := os.Getenv(e)

			if value != "" {
				switch property.Type.String() {
				case "string":
					dataStruct.Field(i).SetString(value)

					status = ValidOk
					valueError = nil

				case "int":
					if realValue, err := strconv.ParseInt(value, 10, 64); err == nil {
						dataStruct.Field(i).SetInt(realValue)
					} else {
						status = ValidWrongValue
						valueError = err
					}

				case "float":
					if realValue, err := strconv.ParseFloat(value, 32); err == nil {
						dataStruct.Field(i).SetFloat(realValue)
					} else {
						status = ValidWrongValue
						valueError = err
					}

				case "bool":
					if realValue, err := strconv.ParseBool(value); err == nil {
						dataStruct.Field(i).SetBool(realValue)
					} else {
						status = ValidWrongValue
						valueError = err
					}

				default:
					status = ValidWrongValue
					valueError = fmt.Errorf("Wrong value (%s=\"%s\")", property.Name, value)
				}
			} else {
				status = ValidRequiredValue
				valueError = fmt.Errorf("Required value (%s)", property.Name)
			}

			evh.Result[property.Name] = Result{status, value, valueError}

			if status != ValidOk && validated {
				validated = false
			}
		}
	}

	return validated
}

// -----------------------------------------------------------------------------

// New retorna un objeto EnvironmentVarsHandler.
func New(confVars interface{}) *EnvironmentVarsHandler {
	return &EnvironmentVarsHandler{conf: &confVars}
}
