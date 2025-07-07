# Sistema Gestor de Datos

## Breve explicación del proyecto

Este proyecto consiste en una **API RESTful** desarrollada en **Go** con **Gin** y **GORM**, que permite gestionar productos y stocks como parte de un sistema mayor orientado a la integración y limpieza de datos provenientes de fuentes heterogéneas (archivos, APIs, bases locales). La API permite registrar productos con su información, proveedores y disponibilidad en distintas sucursales. Toda la documentación de los endpoints está integrada en Swagger.

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

1. **Clonar el repositorio**

```bash
git clone https://github.com/Sosa-MNazarena/sistema-gestor.git
cd sistema-gestor
```

2. **Crear la base de datos**

```bash
mysql -u root -p < crear_base.sql

```
El archivo crear_base.sql contiene el script SQL necesario para crear la base de datos.

```sql
CREATE DATABASE IF NOT EXISTS sistemagestor;
USE sistemagestor;
```

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

---

## Estructura del proyecto

sistema-gestor/
│
├── controllers/       # Lógica de los endpoints
│   └── product.go
├── models/            # Modelos de dominio (ORM)
│   ├── product.go
│   └── stock.go
├── docs/              # Documentación generada por Swagger
│
├── main.go            # Punto de entrada principal
├── crear_base.sql     # Script de creación de base MySQL
└── go.mod             # Dependencias y módulos


---

## Autoría del proyecto

- Sosa, Mariana Nazarena