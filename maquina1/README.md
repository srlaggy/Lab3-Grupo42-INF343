# Grupo 42

## Integrantes:
- Darinka Quiñones  201504159-8
- Ignacio Ulloa     201611004-8


## Ejecución de la tarea

### -> Abrir consolas:
Máquina 1:
- 1) Abrir una consola en el directorio de **FulcrumEntity**
- 2) Abrir una consola en el directorio de **AlmiranteEntity**

Máquina 2:
- 3) Abrir una consola en el directorio de **FulcrumEntity**
- 4) Abrir una consola en el directorio de **AhsokaEntity**

Máquina 3:
- 5) Abrir una consola en el directorio de **FulcrumEntity**
- 6) Abrir una consola en el directorio de **PrincesaEntity**

Máquina 4:
- 7) Abrir una consola en el directorio de **BrokerEntity**

### -> Construir los archivos:
En las 7 consolas ejecutar el comando `make`.

### -> Ejecutar los archivos:
En las 7 consolas ejecutar el comando `make run` y  ya se pueden probar los distintos comandos del enunciado de la tarea en los menú respectivos de los Informantes y la Princesa Leia.

## Consideraciones generales:
- El programa tiene pausas de un segundo en algunos menús para evitar que el menú avance rápidamente.
- Se crearon menús en los informantes y en la princesa para manejar el paso de comandos más limpiamente, aún así, se ruega enviar inputs correctos para evitar comportamientos inesperados.
- Los relojes vectoriales fueron almacenados en estructuras creadas para ello, al igual que los registros que contienen además el servidor, planeta y ciudad en el caso de informantes, Leia y Fulcrums.
- Los inputs de Planeta y ciudad, además de cualquier valor agregado en los menús, no pueden contener espacios.

## Consideraciones *Monotonic Reads* y *Read your writes*
- Monotonic reads fue manejado para aceptar relojes que cuenten con la misma o una mayor cantidad de cambios respecto al servidor que Leia almacena en memoria.
- Read your writes fue manejado para que priorice el aceptar relojes que tengan la misma cantidad de cambios que almacena el informante, y en el peor de los casos, aceptaría relojes que contengan más cambios de los almacenados por el informante, es decir, donde otro informante haya realizado cambios.
- Ambos tienen cierto nivel de aleatoriedad que dependiendo del caso, deja de serlo estrictamente.

## Consideraciones Merge y consistencia eventual
- La consistencia eventual se realiza cada 2 minutos propagando cambios y realizando la limpieza de los log de registros.
- El merge en caso de conflictos, se quedará con el último cambio.
- Se eligió el servidor 1 como nodo dominante para manejar la propagación de cambios.