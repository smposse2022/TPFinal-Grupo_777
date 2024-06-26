package rutinaDeEjercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gocarina/gocsv"
)

func CargarRutinas(lista *ListaDeRutinas) (*ListaDeRutinas,error) {
	archivo, err := os.OpenFile("rutinas.csv", os.O_RDONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer archivo.Close()

	var rutinas []*Rutina
	if err := gocsv.UnmarshalFile(archivo, &rutinas); err != nil {
		// Si el archivo está vacío, no es un error crítico.
		if err.Error() != "EOF" {
			return lista, err
		}
	}

	for _, rutina := range rutinas {
		lista.AgregarRutina(rutina.Nombre, rutina.EjerciciosTotales)
	}
	return lista, nil
}

func AgregarRutinaMenu(lista *ListaDeRutinas, listaEjercicios *ListaDeEjercicios) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nombre de la rutina: ")
	nombre, _ := reader.ReadString('\n')
	nombre = strings.TrimSpace(nombre)

	// Verificar si la rutina ya existe
	_, err := lista.ConsultarRutina(nombre)
	if err == nil {
		fmt.Println("Error: la rutina ya existe.")
		return
	}

	// Solicitar los ejercicios para la rutina
	var ejercicios []*Ejercicio
	for {
		fmt.Print("Nombre del ejercicio a agregar (dejar en blanco para terminar): ")
		nombreEjercicio, _ := reader.ReadString('\n')
		nombreEjercicio = strings.TrimSpace(nombreEjercicio)
		if nombreEjercicio == "" {
			break
		}

		// Consultar el ejercicio por nombre
		ejercicio, err := listaEjercicios.ConsultarEjercicioPorNombre(nombreEjercicio)
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			continue
		}

		ejercicios = append(ejercicios, ejercicio)
	}

	// Agregar la rutina con los ejercicios seleccionados
	err = lista.AgregarRutina(nombre, ejercicios)
	if err != nil {
		fmt.Println("Error al agregar la rutina:", err)
	} else {
		fmt.Println("Rutina agregada exitosamente.")
		if err := GuardarRutinas(lista); err != nil {
			fmt.Println("Error al guardar el ejercicio:", err)
		}
	}
}

func ListarRutinasMenu(lista *ListaDeRutinas) {
	rutinas, err := lista.ListarRutinas()
	if err != nil {
		fmt.Println("Error al listar rutinas:", err)
		return
	}
	for _, rutina := range rutinas {
		// Construir una lista de nombres de ejercicios separados por coma
		var nombresEjercicios []string
		for _, ejercicio := range rutina.EjerciciosTotales {
			nombresEjercicios = append(nombresEjercicios, ejercicio.Nombre)
		}
		ejerciciosStr := strings.Join(nombresEjercicios, ", ")
		fmt.Printf("Nombre: %s, Duración: %d segundos, Ejercicios: %s, Calorías Quemadas: %d, Dificultad: %s\n",
			rutina.Nombre, rutina.Duracion, ejerciciosStr, rutina.CaloriasQuemadasTotales, rutina.Dificultad)
	}
}


func BorrarRutinaMenu(lista *ListaDeRutinas) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Ingrese el nombre de la rutina a borrar: ")
	nombre, _ := reader.ReadString('\n')
	nombre = strings.TrimSpace(nombre)

	err := lista.BorrarRutina(nombre)
	if err != nil {
		fmt.Println("Error al borrar la rutina:", err)
	} else {
		fmt.Println("Rutina borrada exitosamente.")
		if err := GuardarRutinas(lista); err != nil {
			fmt.Println("Error al guardar los cambios:", err)
		}
	}
}

func ConsultarRutinaPorNombreMenu(lista *ListaDeRutinas) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Ingrese el nombre de la rutina a consultar: ")
	nombre, _ := reader.ReadString('\n')
	nombre = strings.TrimSpace(nombre)

	rutina, err := lista.ConsultarRutina(nombre)
	
	if err != nil {
		fmt.Println("Error al consultar la rutina:", err)
	} else {
		var nombresEjercicios []string
		for _, ejercicio := range rutina.EjerciciosTotales {
			nombresEjercicios = append(nombresEjercicios, ejercicio.Nombre)
		}
		ejerciciosStr := strings.Join(nombresEjercicios, ", ")
		fmt.Printf("Nombre: %s, Duración: %d segundos, Ejercicios: %s, Calorías Quemadas: %d, Dificultad: %s\n",
			rutina.Nombre, rutina.Duracion, ejerciciosStr, rutina.CaloriasQuemadasTotales, rutina.Dificultad)
	}
}

func AgregarEjercicioARutinaMenu(lista *ListaDeRutinas, listaEjercicios *ListaDeEjercicios) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Ingrese el nombre de la rutina a modificar: ")
	nombre, _ := reader.ReadString('\n')
	nombre = strings.TrimSpace(nombre)
		// Solicitar los ejercicios para la rutina
		for {
			fmt.Print("Nombre del ejercicio a agregar (dejar en blanco para terminar): ")
			nombreEjercicio, _ := reader.ReadString('\n')
			nombreEjercicio = strings.TrimSpace(nombreEjercicio)
			if nombreEjercicio == "" {
				break
			}
	
			// Consultar el ejercicio por nombre
			ejercicio, _ := listaEjercicios.ConsultarEjercicioPorNombre(nombreEjercicio)
			lista.AgregarEjercicioARutina(nombre, ejercicio)
		}
	
		fmt.Println("Rutina agregada exitosamente.")
		if err := GuardarRutinas(lista); err != nil {
			fmt.Println("Error al guardar el ejercicio:", err)
		}
	}

func EliminarEjercicioDeRutinaMenu(lista *ListaDeRutinas, listaEjercicios *ListaDeEjercicios) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Ingrese el nombre de la rutina a modificar: ")
	nombre, _ := reader.ReadString('\n')
	nombre = strings.TrimSpace(nombre)
		// Solicitar los ejercicios para la rutina
		for {
			fmt.Print("Nombre del ejercicio a eliminar (dejar en blanco para terminar): ")
			nombreEjercicio, _ := reader.ReadString('\n')
			nombreEjercicio = strings.TrimSpace(nombreEjercicio)
			if nombreEjercicio == "" {
				break
			}
	
			// Consultar el ejercicio por nombre
			ejercicio, _ := listaEjercicios.ConsultarEjercicioPorNombre(nombreEjercicio)
			lista.EliminaEjercicioDeRutina(nombre, ejercicio)
		}
	
		fmt.Println("Rutina agregada exitosamente.")
		if err := GuardarRutinas(lista); err != nil {
			fmt.Println("Error al guardar el ejercicio:", err)
		}
}

// Función GeneracionAutomagicaMenu actualizada
func GeneracionAutomagicaMenu(lista *ListaDeRutinas, listaEjercicios *ListaDeEjercicios) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nombre: ")
	nombre, _ := reader.ReadString('\n')
	nombre = strings.TrimSpace(nombre)

	fmt.Print("Duración total de la rutina (en segundos): ")
	duracionStr, _ := reader.ReadString('\n')
	duracionStr = strings.TrimSpace(duracionStr)
	duracion, err := strconv.Atoi(duracionStr)
	if err != nil {
		fmt.Println("Duración inválida. Debe ser un número entero.")
		return
	}

	fmt.Print("Tipo de ejercicios a incluir (por ejemplo, cardio, fuerza, balance, etc.): ")
	tipoStr, _ := reader.ReadString('\n')
	tipoStr = strings.TrimSpace(tipoStr)
	tipo := TipoEjercicio(tipoStr)

	fmt.Print("Nivel de dificultad de los ejercicios a incluir (por ejemplo, principiante, intermedio, avanzado): ")
	dificultadStr, _ := reader.ReadString('\n')
	dificultadStr = strings.TrimSpace(dificultadStr)
	dificultad := NormalizeDificultad(dificultadStr)

	rutina, err := lista.GeneracionAutomagica(NormalizeString(nombre), duracion, NormalizeTipoEjercicio(tipo), dificultad, listaEjercicios)
	if err != nil {
		fmt.Println("Error al generar la rutina:", err)
		return
	}

	err = lista.AgregarRutina(rutina.Nombre, rutina.EjerciciosTotales)
	if err != nil {
		fmt.Println("Error al agregar la rutina:", err)
		return
	}

	fmt.Println("Rutina agregada exitosamente.")
	if err := GuardarRutinas(lista); err != nil {
		fmt.Println("Error al guardar la rutina:", err)
	}
}
/*
func GeneracionAutomagica2Menu(lista *ListaDeRutinas, listaEjercicios *ListaDeEjercicios) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nombre: ")
	nombre, _ := reader.ReadString('\n')
	nombre = strings.TrimSpace(nombre)

	fmt.Print("Calorías Objetivo: ")
	caloriasStr, _ := reader.ReadString('\n')
	caloriasStr = strings.TrimSpace(caloriasStr)
	calorias, err := strconv.Atoi(caloriasStr)
	if err != nil {
		fmt.Println("Calorías inválidas. Debe ser un número entero.")
		return
	}

	// Generar la rutina automáticamente
	rutina, err := lista.GeneracionAutomagica2(NormalizeString(nombre), calorias, listaEjercicios)
	if err != nil {
		fmt.Println("Error al generar la rutina:", err)
		return
	}

	// Agregar la rutina a la lista de rutinas
	err = lista.AgregarRutina(rutina.Nombre, rutina.EjerciciosTotales)
	if err != nil {
		fmt.Println("Error al agregar la rutina:", err)
	} else {
		fmt.Println("Rutina agregada exitosamente.")
		if err := GuardarRutinas(lista); err != nil {
			fmt.Println("Error al guardar la rutina:", err)
		}
	}
}

func GeneracionAutomagica3v2Menu(lista *ListaDeRutinas, listaEjercicios *ListaDeEjercicios) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nombre: ")
	nombre, _ := reader.ReadString('\n')
	nombre = strings.TrimSpace(nombre)

	fmt.Print("Duración total de la rutina (en segundos): ")
	duracionStr, _ := reader.ReadString('\n')
	duracionStr = strings.TrimSpace(duracionStr)
	duracion, err := strconv.Atoi(duracionStr)
	if err != nil {
		fmt.Println("Duración inválida. Debe ser un número entero.")
		return
	}
	
	fmt.Print("Tipo de ejercicios a incluir (por ejemplo, cardio, fuerza, flexibilidad, etc.): ")
	tipoStr, _ := reader.ReadString('\n')
	tipoStr = strings.TrimSpace(tipoStr)

	// Convertir el tipo de ejercicio a TipoEjercicio
	tipoEjercicio := TipoEjercicio(tipoStr)
	tipoNormalizado := NormalizeTipoEjercicio(tipoEjercicio)

	// Generar la rutina automágicamente
	rutina, err := lista.GeneracionAutomagica3v2(NormalizeString(nombre), duracion, tipoNormalizado, listaEjercicios)
	if err != nil {
		fmt.Println("Error al generar la rutina:", err)
		return
	}

	// Agregar la rutina a la lista de rutinas
	err = lista.AgregarRutina(rutina.Nombre, rutina.EjerciciosTotales)
	if err != nil {
		fmt.Println("Error al agregar la rutina:", err)
	} else {
		fmt.Println("Rutina agregada exitosamente.")
		if err := GuardarRutinas(lista); err != nil {
			fmt.Println("Error al guardar la rutina:", err)
		}
	}
}
	*/
func GuardarRutinas(lista *ListaDeRutinas) error {
	rutinas, err := lista.ListarRutinas()
	if err != nil {
		// No hay rutinas para guardar, pero no es un error crítico
		if err.Error() == "no hay ninguna rutina para listar" {
			return nil
		}
		return err
	}
	// Normalizar
	for _, rutina := range rutinas {
		rutina.Nombre = NormalizeString(rutina.Nombre)
		rutina.Dificultad = Dificultad(NormalizeString(string(rutina.Dificultad)))
	}
	// Verificar si el archivo ya existe
	existe, err := ArchivoExiste("rutinas.csv")
	if err != nil {
		return err
	}

	archivo, err := os.OpenFile("rutinas.csv", os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer archivo.Close()

	// Escribir encabezado solo si el archivo no existía
	if !existe {
		if _, err := archivo.WriteString("Nombre,Duración,Ejercicios,TipoDeEjercicios,Calorpias,Dificultad,\n"); err != nil {
			return err
		}
	}

	return gocsv.MarshalFile(&rutinas, archivo)
}