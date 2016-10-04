# Generador de archivos LESS para automatizar inicio de proyecto web

La idea de este pequeño algoritmo escrito en *golang* es automatizar el inicio de mis proyectos web generando archivos less con algunas cosas predefinidas como la importación de archivos less, y archivos comentados.

### Características

Este algoritmo está diseñado para la pseudo-integración con el plugin *Easy LESS* del editor *Visual Studio Code*, donde el programa en ejecución lo primero que pregunta es si deseas hacer uso del plugin, para asistirte en la configuración de éste.

### Ejecutable

Para crear el ejecutable en go solo hay que hacer un *build* del archivo app.go.
    
    go build app.go

