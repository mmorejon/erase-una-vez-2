# erase-una-vez-2

[![English](https://img.shields.io/badge/Read_in-English-blue?style=flat-square)](README.en.md)

<div align="center">

<img src="./assets/book-cover.jpg" alt="Portada Libro Érase una vez Kubernetes" width="300"/>

Aplicación Golang utilizada en los ejercicios del libro Érase una vez Kubernetes.

👇 **Consigue la edición actualizada 2026 aquí:** 👇

[![Amazon](https://img.shields.io/badge/Amazon-Comprar_en_Tapa_Blanda-orange?style=for-the-badge&logo=amazon)](https://www.amazon.es/dp/8409212765)
[![LeanPub](https://img.shields.io/badge/LeanPub-Descargar_Ebook-blue?style=for-the-badge&logo=leanpub)](https://leanpub.com/erase-una-vez-kubernetes)

</div>

---

## Descripción

La código del repositorio se utiliza para realizar ejemplos de comunicación entre dos aplicaciones desplegadas en el cluster: un cliente y un servidor. Ambas aplicaciones han sido compiladas en la misma imagen Docker, pero el comando de inicio es diferente.

### Características del servidor

Utilice el comando `server` para iniciar el sistema como servidor web. La aplicación brinda los siguientes puntos de acceso:

|Punto de acceso |Descripción|
|-----|-----------|
|`/echo`| Muestra el mensaje `érase una vez ...` junto con el nombre de la máquina.|
|`/healthz`| muestra el mensaje OK si el servidor está funcionando correctamente.|

Todos los puntos de acceso muestran la información en formato JSON.

|Variable de entorno|Descripción|Tipo|Valor por defecto|
|-----|-----------|------|----|
|`CHARACTER`| Modifica el mensaje de respuesta del servidor.| String | `...` |

### Características del cliente

Utilice el comando `client` para iniciar el sistema como cliente web. La aplicación realiza peticiones hacia el enlace configurado por intervalos de tiempo. Para configurar la aplicación cliente utilice las siguientes variables de entorno:

|Variable de entorno|Descripción|Tipo|Valor por defecto|
|-----|-----------|------|---|
|`SLEEP_TIME`| Intervalo de tiempo entre peticiones. Es una cadena. | String | `1s` |
|`ENDPOINT`| Punto de acceso del servidor web.| String | `""` |
|`HTTP_HEADERS`| Encabezados de la petición HTTP. | String | `""` |

## Funcionamiento en una máquina

```bash
# iniciar servidor web
docker container run --rm \
  --name server \
  --entrypoint server \
  -p 8000:8000 \
  --detach \
  ghcr.io/mmorejon/erase-una-vez-2:v0.5.0

2020/01/20 23:02:24 Servidor iniciado
```

```bash
curl http://localhost:8000/echo

{
  "hostname": "da1ebc56480a",
  "message": "érase una vez ..."
}%
```

```bash
curl http://localhost:8000/healthz
{
  "status": "OK"
}%
```

```bash
# obtener el ip del servidor
docker container inspect \
  --format='{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' \
  server

172.17.0.2
```

```bash
# iniciar servidor web
docker container run --rm \
  --name client \
  --entrypoint client \
  --detach \
  --env ENDPOINT="http://172.17.0.2:8000/echo" \
  ghcr.io/mmorejon/erase-una-vez-2:v0.5.0
```

```bash
# consultar los logs del cliente
docker container logs client
{
  "hostname": "9e65a5deb2f7",
  "message": "érase una vez ..."
}
{
  "hostname": "9e65a5deb2f7",
  "message": "érase una vez ..."
}
```

## Funcionamiento en un cluster de Kubernetes

```bash
kubectl apply -f kubernetes/

deployment.apps/client created
deployment.apps/server created
service/server created
```

```bash
kubectl get all -l app=server

NAME                        READY   STATUS    RESTARTS   AGE
pod/server-b4746684-wgfxm   1/1     Running   0          6m46s

NAME             TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)    AGE
service/server   ClusterIP   10.109.117.210   <none>        8000/TCP   6m46s

NAME                     READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/server   1/1     1            1           6m46s

NAME                              DESIRED   CURRENT   READY   AGE
replicaset.apps/server-b4746684   1         1         1       6m46s
```

```bash
kubectl get all -l app=client

NAME                         READY   STATUS    RESTARTS   AGE
pod/client-7dd5bc4b4-j6qg4   1/1     Running   1          7m41s

NAME                     READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/client   1/1     1            1           7m41s

NAME                               DESIRED   CURRENT   READY   AGE
replicaset.apps/client-7dd5bc4b4   1         1         1       7m41s
```

---

## 🤝 Comunidad y Feedback

1.  ⭐ **¿Te ha sido útil?** Dale una **estrella** al repositorio (arriba a la derecha). Nos ayuda a llegar a más ingenieros.
2.  📚 **¿Aún no tienes el libro?** Compra el libro en Amazon o Leanpub.

<div align="center">
    <a href="https://www.amazon.es/dp/8409212765">
        <img src="https://img.shields.io/badge/Amazon-Ver_Precio_y_Opiniones-orange?style=for-the-badge&logo=amazon" />
    </a>
</div>
