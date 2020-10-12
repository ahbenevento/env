# Simple gestión de variables de entorno

Permite gestionar variables de entornos pasando una estructura de forma similar
en como se "traduce" una estructura de valores al o desde el formato JSON.

## Instalación

```
go get -u https://github.com/ahbenevento/env
```

## Uso

```go
// Las variables a comprobar...
type misVariables struct {
    TempFolder     string `name:"TEMP,TMP"`
	EnteroEsperado int    `name:"VALOR"`
	OtroValor      bool
}

// Validar...
envVars := &misVariables{}
checkEnvVars := env.New(envVars)

fmt.Println(checkEnvVars.Valid())
fmt.Println(envVars)
```

## Contacto

Ante alguna catástrofe por favor escriba.

- https://github.com/ahbenevento/env/issues
