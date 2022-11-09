# Kubernetes, hablemos de Kubernetes

## Laboratorio

Versión utilizada de Kind.

```
kind version
```

Crear nuevo cluster.

```
kind create cluster
```

Listar los nodos.

```
k get nodes
```

Comprobar el estado de los pods.

```
k get po -A
```

Mostar los contenedores en funcionamiento. Ahora un clustes está dentro de un pod.

```
docker container ls
```

Desplegar el servidor y el cliente.

```
k apply -f kubernetes
```

Comprobar el funcionamiento de los pods.

```
k get po
```

Confirmar que el nombre del host.

```
k logs client-.....
```

## Construcción de imágenes

El comando build será utilizado posteriormente.
Descargar la imagen desplegada en el cluster

```
docker image pull ghcr.io/mmorejon/erase-una-vez-2:v0.3.0
```

Analizar el tamaño.

```
docker image ls --filter=reference=ghcr.io/mmorejon/erase-una-vez-2:v0.3.0
```

Explorar el contenido de la imagen.
Se va a utilizar la herramienta dive.

```
dive --help | less
```

```
dive ghcr.io/mmorejon/erase-una-vez-2:v0.3.0
```

Utitlizar Hadolint para analizar la calidad del contenedor.

```
hadolint --help | less
```

Realizar modificaciones en el fichero para ver nuevos mensajes.

```
hadolint Dockerfile
```

Utilizar trivy para identificar vulnerabilidades con Docker Desktop

```
trivy --help | less
```

## Despliegue automático

Pensemos que queremos realizar cambios y desplegar una
nueva versión.

```
skaffold --help | less
```

Crear la configuración de skaffold

```
skaffold init
```

Borrar el fichero de despliegue del cliente
Modificar el mensaje del servidor
Abrir nueva consola para ver los logs del cliente
Ejecutar el comando run

```
skaffold run
```

Listar los pods

```
k get po
```

Utilizar el modo dev

```
skaffold dev
```

## Análisis de logs

Escalar el número de clientes en el fichero.
Listar los pods junto con las etiquetas.
Mostrar los logs de todas las réplicas.

```
k logs --selector app=client --follow
```

## Depuración de errores

La depuración de errores se va a realizar con Telepresence.
Instalar

```
telepresence helm install
```

Mostrar el agente instalado

```
k get po -A
```

Nueva ventana para ver cambios del pod.

```
telepresence list
```

Interceptar el tráfico del servicio server
Recordar el mensaje del cliente

```
k logs --selector app=client
```
```
telepresence intercept server --port 8000:http
```
