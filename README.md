# Prueba Tecnica


### Backend

En la carpeta back, encontraras un back desarollado en GO. Para configurar el back modificar el archivo .env
```
 PORT=<Puerto de la APP>
```
 Encontraras un Makefile con los siguientes comandos
```
run <Lanza el back>

install <Instala los paquetes necesarios>

build <Compila el projecto>

unit-test <Ejecuta las pruebas unitarias>

docker-build <Crea un docker>

docker-run <Lanza el docker>
```

### FrontEnd
En la carpeta front, encontraras un front desarollado en VueJS (Vue-Cli). Para configurar el front modificar el archivo .env

```
PORT=<Puerto del front>
VUE_APP_URL_API= <Url del back/EJEMPLO:http://localhost:3000/search_tracks>
```
 Encontraras un Makefile con los siguientes comandos
```
install <Instala los node_modules>

run <Instala y ejecuta la app>

build <Hace un build de la app>

lint < ejecuta el lint>

docker-build <Levanta un docker para Apple Silicon(M1/M1PRO/M1MAX)>

docker-run <Lanza el docker>
```

### Para probar
Para probar ejecuta Make run en ambas carpetas posterior a la configuracion.
