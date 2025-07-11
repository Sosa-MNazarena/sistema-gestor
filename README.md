# Sistema Gestor de Datos

## Breve explicación del proyecto

Este proyecto consiste en una **API RESTful** desarrollada en **Go** con **Gin** y **GORM**, que permite gestionar productos y stocks como parte de un sistema mayor orientado a la integración de datos provenientes de fuentes heterogéneas (archivos, APIs y archivos Excel). La API permite registrar productos, eliminarlos, editarlos y buscarlos con toda la información disponible. Toda la documentación de los endpoints está integrada en Swagger.

---

## Tecnologías utilizadas en el proyecto

Lenguaje: **Go (Golang)**
Web framework: **Gin**
ORM: **GORM**
Base de datos: **MySQL**
Documentación: **Swagger (Swaggo)**
Control de versiones: **Git + GitHub**

---

## Instrucciones para correrlo localmente

### Requisitos
- Tener Go instalado: https://go.dev/dl/
- Tener MySQL instalado y corriendo (puede ser también DBeaver)
- Git y GitHub

1. **Clonar el repositorio**

```bash
git clone https://github.com/Sosa-MNazarena/sistema-gestor.git
cd sistema-gestor
```

2. **Crear la base de datos**

Deberás ejecutar el archivo `crear_base.sql` para crear la base de datos.
    
Primero: Abrí el Workbench de MySQL o el DBeaver.
    **nota**: Si no tenés instalado MySQL, necesitás correr el XAMPP(https://www.apachefriends.org) que es esencial para que MySQL corra localmente.

Segundo: Una vez que estés conectado a tu servidor MySQL, abrí la base de datos y ejecutá el archivo `crear_base.sql` para crear la base de datos(File > Open > crear_base.sql).
O también podés hacerlo manualmente, abriendo una nueva pestaña query y ejecutando lo siguiente:

    
    CREATE DATABASE IF NOT EXISTS sistemagestor;
    USE sistemagestor;

Tercero: Ejecutá el script y revisá que la base de datos sistemagestor haya sido creada correctamente. 

3. **Instalar las dependencias**

Si es la primera vez que se instala el proyecto, se recomienda ejecutar el siguiente comando para crear un archivo `go.mod` con la información de dependencias.

```bash
go mod init sistema-gestor
```
Ahora se instalan las dependencias.

```bash
go mod tidy
```

4. **Generar la documentación Swagger**
En caso de no tener instalado `swag`, se recomienda instalarlo con el siguiente comando.
 ```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

```bash
swag init
```

5. **Ejecutar la aplicación**

```bash
go run main.go
``` 

---

## Pruebas con endpoints documentados

Para probar la aplicación, se pueden utilizar los siguientes endpoints:

1. **Acceder a la documentación Swagger**

```bash
http://localhost:8080/swagger/index.html
```

2. **Crear un producto**

    - Método: **POST**
    - Endpoint: /products
    - Body de ejemplo:

        ```json
        {
            "nombre": "Producto 1",
            "descripcion": "Descripción del producto 1",
            "categoria": "Categoria del producto 1",
            "proveedor": "Proveedor del producto 1",
            "precio": 10.0,
            "stocks": [
                {
                    "sucursal": "Sucursal 1",
                    "cantidad": 10
                },
                {
                    "sucursal": "Sucursal 2",
                    "cantidad": 20
                }
            ]
        }
        ```

3. **Obtener todos los productos**

    - Método: **GET**
    - Endpoint: /products

## Pruebas con archivos Excel y API

Para el correcto funcionamiento de estas funciones, y tener la experiencia completa, se recomienda cambiar el email destinatario configura para recibir los reportes de cargas de datos.
Para ello:

- Abrir el archivo events/emailEvent.go
- En la línea 13, cambiar el email de ejemplo por tu mail (por default, está el mail de la autora del proyecto) como el ejempo.

```bash
m.SetHeader("To", "email@ejemplo.com") 
```

- Guardar el archivo y ejecutar la aplicación nuevamente con el comando `go run main.go`.

4. **Cargar productos de prueba desde un archivo Excel**

    - Método: **POST**
    - Endpoint: /excelReader

    **nota**: El archivo debe estar en el mismo directorio que el proyecto. Se pega el path del archivo y sin comillas.

    En el proyecto se encuentra un archivos excels de ejemplo llamado `ejemploProductos.xlsx` y `ejemploProductosConDuplicados.xlsx` que contienen datos de ejemplo para probar la API.

5. **Cargar productos de prueba desde una API**

    - Método: **POST**
    - Endpoint: /apiReader

    Podés probar la API de ejemplo en la siguiente URL: https://686d9a71c9090c495386c173.mockapi.io/APIproductos/productos
        
---

## Estructura del proyecto

```bash
sistema-gestor/
├── controllers
│   ├── apiController.go
│   ├── productController.go
│   └── excelController.go
├── docs
├── events
│   └── emailEvent.go
├── models
│   ├── product.go
│   └── stock.go
├── services
│   ├── apiImportService.go
│   ├── excelImportService.go
│   └── productService.go
├── repositories
│   ├── productImportRepository.go
│   └── productRepository.go
├── strategy
│   ├── apiReader.go
│   ├── dataReader.go
│   ├── excelReader.go
│   └── readerContexto.go
├── main.go
├── go.mod
├── go.sum
├── crear_base.sql
├── ejemploProductos.xlsx
├── ejemploProductosConDuplicados.xlsx
└── README.md
```
---

## Autoría del proyecto

- Sosa, Mariana Nazarena
