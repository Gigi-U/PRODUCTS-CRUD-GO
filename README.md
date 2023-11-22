# PRODUCTS-CRUD-GO

Inicia el proyecto en la clase 10 --> rama "clase10"

Clase 11 - organización de proyecto en paquetes

cmd/server = 
    - punto de entrada de nuestra api 
    - incluye: 
        - main
        - handler (o controlador) --> van los controllers de los productos (dentro de carpetas con su nombre)
        - router --> van todas las declaraciones de rutas de los endpoints

internal = 
    - todo lo de manejo interno va ahi. 
    - incluye:
        - models(o domain) --> va el modelo
        - products --> van el service y el repository

pkg = 
    - va todo lo que es comun al proyecto. 
    - puede incluir: 
        - middleware
        - utils 
        - web (en web van todas las salidas comunes de la api, o estructuras comunes de respuesta o request) .

Configuracion de servidor va tmb en pkg, si es una configuración interna del servidor piede ir en internal


<<<<<<<<<<<<<FLUJO>>>>>>>>>>

1) muevo el main de clase 10 cmd/server y le agrego una estructura Storage que va a ser un slice de producto con el que se va a trabajar en memoria

2) en models --> products defino las estructuras que van a ser de modelo , pueden haber varias con diferentes propiedades.

3) en repository creo la estructura que va a contener el slice de productos con los que va a trabajar en memoria. estamos simulando una base de datos.
 - en repository, necesitamos una interface que tiene que cumplir ciertos contratos



 A CONTINUAR ---> clase ante última - 17:19 



Clases 12 y 13