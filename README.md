# PROYECTO-FINAL-ARQUI-DE-SOFTWARE

integrantes: Julian Irusta, Ignacio Magoia , Vicente Monzo , Melinsky Julieta


# Instrucciones para correr el proyecto

Con Docker instalado, correr

```
docker compose up -d
```

Para reiniciar el backend, usar:

```
docker compose down backend
docker compose up -d
```

Para borrar todo y levantarlo desde cero:

````
docker compose down
rm -rf .database
docker compose up -d
```
