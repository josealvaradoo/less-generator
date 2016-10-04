// --------------------------------------------------------
// less-generator
//
// Se recomienda leer el archivo readme.md
// --------------------------------------------------------

package main

// Librerias utilizadas
import (
    "os"
    "fmt"
    "time" // Esta es opcional, se usa en el método Sleep() para darle el efecto progresivo 
)

// Declaracion de variables
var lessCompiler bool
var answer string
var wait string
var easyLessConfig string

// func main()
// Se ejecuta siempre al inicio del programa
func main() {

    // Configuracion rápida del plugin Easy LESS de Visual Studio Code
    easyLessConfig = "// main:styles.less, out: ../css/styles.min.css, compress: true, sourceMap: false"

    fmt.Println("¿Desea usar el plugin Easy LESS de Studio Code para compilar less? (Y/N)")
    fmt.Scanln(&answer)
    if answer == "Y" || answer == "y" {
        lessCompiler = true
    } else {
        lessCompiler = false
    }
    
    // Llamado secuencial de funciones
    createStylesDocument()
    createVariablesDocument()
    createResetDocument()
    createLessDocuments()

    // Para evitar que el programa se cierre, lo dejamos a la espera de una lectura
    fmt.Scanln(&wait)
    
}

// Esta funcion crea el archivo styles.less
// Este archivo contiene 15 líneas donde importa otros archivos less
// con el fin de dividir el código por diferentes componentes.
func createStylesDocument() {
    file, err := os.Create("styles.less")
    if err != nil {
        fmt.Println("Hubo un problema con la creación del archivo styles.Less")
        return
    }
    if lessCompiler == true {
        fmt.Fprintln(file, easyLessConfig)
    }
    fmt.Fprintln(file, "@import url('reset.less');")
    fmt.Fprintln(file, "@import url('variables.less');")
    fmt.Fprintln(file, "@import url('functions.less');")
    fmt.Fprintln(file, "@import url('main.less');")
    fmt.Fprintln(file, "@import url('header.less');")
    fmt.Fprintln(file, "@import url('aside.less');")
    fmt.Fprintln(file, "@import url('navbar.less');")
    fmt.Fprintln(file, "@import url('articles.less');")
    fmt.Fprintln(file, "@import url('footer.less');")
    fmt.Fprintln(file, "@import url('forms.less');")
    fmt.Fprintln(file, "@import url('buttons.less');")
    fmt.Fprintln(file, "@import url('tables.less');")
    fmt.Fprintln(file, "@import url('helpers.less');")
    fmt.Fprintln(file, "@import url('responsive.less');")
    fmt.Fprintln(file, "@import url('animations.less');")
    file.Close()
    fmt.Println("")
    fmt.Println("El archivo styles.less se ha creado satisfactoriamente.")
    fmt.Println("El archivo styles.less ha importado previamente sus componentes.")
    fmt.Println("")
}

// Esta funcion crea el archivo variables.less
// Este archivo viene previamente escrito con lineas de código comentadas
// para organizar las variables dependiendo de su tipo (atributos css).
func createVariablesDocument() {

    selectors := [9]string{"@imports","background","border","colors","font","height","margin","padding","width"}

    file, err := os.Create("variables.less")
    if err != nil {
        fmt.Println("Hubo un problema con la creación del archivo variables.less")
        return
    }
    if lessCompiler == true {
        fmt.Fprintln(file, easyLessConfig)
    }
    for i:=0; i<9; i++ {
        fmt.Fprintln(file, "/* =============================================================================")
        fmt.Fprintln(file, "* \t", selectors[i])
        fmt.Fprintln(file, "*  ========================================================================== */")
        fmt.Fprintln(file, "")
    }
    fmt.Println("El archivo variables.less se ha creado satisfactoriamente.")
    fmt.Println("El archivo variables.less ha prestablecido un orden.")
    fmt.Println("")
}

// Esta funcion crea el archivo reset.less
// La idea de este archivo es formatear todos los estilos para que los
// navegadores no usen sus estilos predefinidos.
func createResetDocument() {
    file, err := os.Create("reset.less")
    if err != nil {
        fmt.Println("Hubo un problema con la creación del archivo reset.Less")
        return
    }
    if lessCompiler == true {
        fmt.Fprintln(file, easyLessConfig)
    }
    fmt.Fprintln(file, "* {")
    fmt.Fprintln(file, "\t margin: 0;")
    fmt.Fprintln(file, "\t border: 0;")
    fmt.Fprintln(file, "\t padding: 0;")
    fmt.Fprintln(file, "\t outline: none;")
    fmt.Fprintln(file, "}")
    file.Close()
    fmt.Println("")
    fmt.Println("El archivo reset.less se ha creado satisfactoriamente.")
    fmt.Println("El archivo reset.less se ha predefinido.")
    fmt.Println("")
}

// Esta función crea todos los archivos que son importados en styles.less
// excepto el de reset y variables que son creados previamente.
func createLessDocuments() {

    files := [13]string{"functions.less","header.less","navbar.less","main.less","aside.less","articles.less","footer.less",                    "forms.less","buttons.less","tables.less","helpers.less","responsive.less","animations.less"}

    for i:=0; i<len(files); i++ {
        file, err := os.Create(files[i])
        if err != nil {
            fmt.Printf("Hubo un problema con la creación del archivo %v", files[i])
            panic(err)
        }
        if lessCompiler == true {
            fmt.Fprintln(file, easyLessConfig)
        }
        fmt.Printf("El archivo %v se ha creado satisfactoriamente", files[i])
        defer file.Close();
        fmt.Println("")
        time.Sleep(100 * time.Millisecond)
    }

}