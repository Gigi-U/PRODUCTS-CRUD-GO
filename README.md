# PRODUCTS-CRUD-GO

Clase 9 (Servidor Web)
    command: go get -u github.com/gin-gonic/gin

Clase 10 (inicio CRUD) método GET --> rama "clase10"   
Clase 11 (cont. CRUD + ESTRUCTURACIÓN de proyecto en paquetes) --> rama "clase11" 

    PRODUCTS-CRUD-GO
        |__cmd/server (punto de entrada de nuestra api )
        |    |_ handler (o controlador) --> van los controllers de los productos (dentro de carpetas con su nombre)
        |    |_ router --> van todas las declaraciones de rutas de los endpoints
        |    - main
        |    
        |__internal (todo lo de manejo interno. lógica)
        |    |_ models(o domain) --> va el modelo
        |    |_ products --> van el service y el repository
        |
        |__pkg (va todo lo que es comun al proyecto)
            |_ middleware
            |_ utils 
            |_ web (en web van todas las salidas comunes de la api, o estructuras comunes de respuesta o request) .
            Configuracion de servidor va tmb en pkg, si es una configuración interna del servidor puede ir en internal



Clase 12 (VARIABLES DE ENTORNO) configuración de entorno --> rama "clase12y13" 
    command: go get -u github.com/joho/godotenv

Clase 13 (SEGURIDAD) Middleware --> rama "clase12y13" 

Clase 14 (DB) - capa de acceso a datos --> rama "clase14" 
    command: go get "github.com/go-sql-driver/mysql" +  crear archivo .env donde iran las variables
    +
    import (
        "database/sql"

        _ "github.com/go-sql-driver/mysql"
    )  // en main

<<<<<<<<<<<<<FLUJO>>>>>>>>>>

1) muevo el main de clase 10 cmd/server y le agrego una estructura Storage que va a ser un slice de producto con el que se va a trabajar en memoria

2) en models --> products defino las estructuras que van a ser de modelo , pueden haber varias con diferentes propiedades.

3) en repository creo la estructura que va a contener el slice de productos con los que va a trabajar en memoria. estamos simulando una base de datos.
 - en repository, necesitamos una interface que tiene que cumplir ciertos contratos



 CONTINUAR ---> grabación clase 11 - 17:19 