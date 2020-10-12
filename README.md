# Simple gestión de variables de entorno

Permite gestionar variables de entornos pasando una estructura de forma similar
en como se "traduce" una estructura de valores a (o desde) el formato JSON.

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
```

```go
// Validar...
envVars := &misVariables{}
checkEnvVars := env.New(envVars)

fmt.Println(checkEnvVars.Valid())
fmt.Println(envVars)                // Imprimir la estructura con los datos obtenidos
fmt.Println(checkEnvVars.Result)    // Imprimir los resultados de cada variable verificada
```

## Funcionamiento

El método **Valid()** retorna TRUE si todas las variables verificadas fueron validadas.

La propiedad **Results** tendrá un registro por cada variable con el estado de
la validación según la estructura **env.Result**.

**env.Result.Status** puede contener uno de los siguientes valores:

-   **ValidOk** (0) el valor ha sido validado.
-   **ValidRequiredValue** (1) la variable no ha sido definida o su valor es "vacío".
-   **ValidWrongValue** (2) la variable posee un valor incorrecto.

## License

[Apache 2.0 license](LICENSE).

## Contacto

Ante alguna catástrofe por favor escriba.

-   https://github.com/ahbenevento/env/issues
