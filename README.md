# Golang-Ecommerce-User

Este repositorio hace parte del ecommerce desarrollado en el lenguaje de programacion Golang, se ejecuta cuando el usuario confirma su correo electronico y se encarga de ir a guardar a la base de datos el nuevo usuario:

## Requisitos Previos

Asegúrate de tener instalado lo siguiente:

- Go: [Descargar Go](https://go.dev/)

**Lenguajes utilizados:** Go  
**Frameworks, herramientas o librerias utilizados:** aws-sdk-go-v2, context, database/sql, errors, fmt, go-sql-driver, os, time

## Paso a paso para ejecutar el repositorio
* Clonar el repositorio en el entorno local utilizando el comando 
```
git clone https://github.com/juparefe/Golang-Ecommerce-User.git
```
* Abrir la carpeta clonada utilizando algun editor de codigo
* El objetivo de este repositorio es obtener un archivo ejecutable en formato .zip para que sea el codigo fuente de la lambda creada en 'Amazon Linux 2023'. Ejecuta el siguiente comando para obtenerlo:
```
set GOOS=linux
set GOARCH=amd64
go build -tags lambda.norpc -o bootstrap main.go
del main.zip
tar.exe -a -cf main.zip bootstrap
```

