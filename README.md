# Erase una vez 1

Aplicación Golang utilizada en los ejercicios del libro [Érase una vez Kubernetes](https://leanpub.com/erase-una-vez-kubernetes).

## Descripción

Inicia un servidor web y expone los siguientes puntos de acceso:

* `/echo`: muestra un mensaje junto con el nombre de la máquina.
* `/healthz`: muestra el mensaje OK si el servidor está funcionando correctamente.

Todos los puntos de acceso muestran la información en formato JSON.

## Funcionamiento

```
docker container run --rm \
  -p 8000:8000 \
  mmorejon/erase-una-vez-2:latest

2020/01/07 23:25:32 Servidor iniciado
```

```
curl http://localhost:8000/echo

{
  "hostname": "da1ebc56480a",
  "message": "érase una vez ..."
}%
```

```
curl http://localhost:8000/healthz
{
  "status": "OK"
}%
```

## Variables de entorno

El funcionamiento de la aplicación puede ser modificado a través de variables de entorno:

* `CHARACTER` modifica el final del mensaje impreso en pantalla.

```
docker container run --rm \
  -p 8000:8000 \
  --env CHARACTER="un castillo." \
  mmorejon/erase-una-vez-2:latest

2020/01/07 23:25:32 Servidor iniciado
```

```
curl http://localhost:8000/echo

{
  "hostname": "da1ebc56480a",
  "message": "érase una vez un castillo."
}%
```