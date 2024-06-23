package rutinaDeEjercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/gocarina/gocsv"
)


func CargarEjercicios(lista *ListaDeEjercicios) (*ListaDeEjercicios, error) {
	archivo, err := os.OpenFile("ejercicios.csv", os.O_RDONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer archivo.Close()

	var ejercicios []*Ejercicio
	if err := gocsv.UnmarshalFile(archivo, &ejercicios); err != nil {
		// Si el archivo está vacío, no es un error crítico.
		if err.Error() != "EOF" {
			return lista, err
		}
	}

	for _, ejercicio := range ejercicios {
		lista.AgregarEjercicio(ejercicio.Nombre, ejercicio.Descripcion, ejercicio.Tiempo, ejercicio.Calorias, ejercicio.TipoDeEjercicio, ejercicio.PuntosPorTipoDeEjercicio, string(ejercicio.Dificultad))
	}
	return lista, nil
}

func AgregarEjercicioMenu(lista *ListaDeEjercicios) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nombre: ")
	nombre, _ := reader.ReadString('\n')
	nombre = strings.TrimSpace(nombre)

	fmt.Print("Descripción: ")
	descripcion, _ := reader.ReadString('\n')
	descripcion = strings.TrimSpace(descripcion)

	fmt.Print("Tiempo (en segundos): ")
	tiempoStr, _ := reader.ReadString('\n')
	tiempo, _ := strconv.Atoi(strings.TrimSpace(tiempoStr))

	fmt.Print("Calorías: ")
	caloriasStr, _ := reader.ReadString('\n')
	calorias, _ := strconv.Atoi(strings.TrimSpace(caloriasStr))

	fmt.Print("Tipos de Ejercicio (separados por coma,(fuerza,cardio,balance): ")
	tipoStr, _ := reader.ReadString('\n')
	tipoStr = strings.TrimSpace(tipoStr)
	tipos := strings.Split(tipoStr, ",")
	var tiposDeEjercicio []TipoEjercicio
	for _, t := range tipos {
		tiposDeEjercicio = append(tiposDeEjercicio, TipoEjercicio(strings.TrimSpace(t)))
	}

	fmt.Print("Puntos por Tipo de Ejercicio (separados por coma, en el mismo orden que los tipos): ")
	puntosStr, _ := reader.ReadString('\n')
	puntosStr = strings.TrimSpace(puntosStr)
	puntosStrArr := strings.Split(puntosStr, ",")
	var puntosPorTipoDeEjercicio []int
	for _, p := range puntosStrArr {
		puntos, _ := strconv.Atoi(strings.TrimSpace(p))
		puntosPorTipoDeEjercicio = append(puntosPorTipoDeEjercicio, puntos)
	}

	fmt.Print("Dificultad (principiante, intermedio, avanzado): ")
	dificultad, _ := reader.ReadString('\n')
	dificultad = strings.TrimSpace(dificultad)

	err := lista.AgregarEjercicio(nombre, descripcion, tiempo, calorias, tiposDeEjercicio, puntosPorTipoDeEjercicio, dificultad)
	if err != nil {
		fmt.Println("Error al agregar ejercicio:", err)
	} else {
		fmt.Println("Ejercicio agregado exitosamente.")
		if err := GuardarEjercicios(lista); err != nil {
			fmt.Println("Error al guardar el ejercicio:", err)
		}
	}
}

func ListarEjerciciosMenu(lista *ListaDeEjercicios) {
	ejercicios, err := lista.ListarEjercicios()
	if err != nil {
		fmt.Println("Error al listar ejercicios:", err)
		return
	}
	for _, ejercicio := range ejercicios {
		fmt.Printf("Nombre: %s, Descripción: %s, Tiempo: %d, Calorías: %d, Tipo de Ejercicio: %v, Puntos por Tipo de Ejercicio: %v, Dificultad: %s\n",
			ejercicio.Nombre, ejercicio.Descripcion, ejercicio.Tiempo, ejercicio.Calorias, ejercicio.TipoDeEjercicio, ejercicio.PuntosPorTipoDeEjercicio, ejercicio.Dificultad)
	}
}

func BorrarEjercicioMenu(lista *ListaDeEjercicios) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Ingrese el nombre del ejercicio a borrar: ")
	nombre, _ := reader.ReadString('\n')
	nombre = strings.TrimSpace(nombre)

	err := lista.BorrarEjercicio(nombre)
	if err != nil {
		fmt.Println("Error al borrar ejercicio:", err)
	} else {
		fmt.Println("Ejercicio borrado exitosamente.")
		if err := GuardarEjercicios(lista); err != nil {
			fmt.Println("Error al guardar los cambios:", err)
		}
	}
}

func ConsultarEjercicioPorNombreMenu(lista *ListaDeEjercicios) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Ingrese el nombre del ejercicio a consultar: ")
	nombre, _ := reader.ReadString('\n')
	nombre = strings.TrimSpace(nombre)

	ejercicio, err := lista.ConsultarEjercicioPorNombre(nombre)
	if err != nil {
		fmt.Println("Error al consultar ejercicio:", err)
	} else {
		fmt.Printf("Nombre: %s, Descripción: %s, Tiempo: %d, Calorías: %d, Tipo de Ejercicio: %v, Puntos por Tipo de Ejercicio: %v, Dificultad: %s\n",
			ejercicio.Nombre, ejercicio.Descripcion, ejercicio.Tiempo, ejercicio.Calorias, ejercicio.TipoDeEjercicio, ejercicio.PuntosPorTipoDeEjercicio, ejercicio.Dificultad)
	}
}

// Función para normalizar Dificultad
func NormalizeDificultad(dificultad string) Dificultad {
	switch NormalizeString(dificultad) {
	case NormalizeString(string(Principiante)):
		return Principiante
	case NormalizeString(string(Intermedio)):
		return Intermedio
	case NormalizeString(string(Avanzado)):
		return Avanzado
	default:
		return ""
	}
}

// Función FiltrarEjerciciosMenu
func FiltrarEjerciciosMenu(lista *ListaDeEjercicios) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Ingrese el tipo de ejercicio a filtrar (deje vacío para omitir): ")
	tipo, _ := reader.ReadString('\n')
	tipo = strings.TrimSpace(tipo)

	var tipoEjercicio TipoEjercicio
	if tipo != "" {
		tipoEjercicio = TipoEjercicio(tipo)
	}

	fmt.Print("Ingrese la dificultad a filtrar (deje vacío para omitir): ")
	dificultadStr, _ := reader.ReadString('\n')
	dificultadStr = strings.TrimSpace(dificultadStr)
	dificultad := NormalizeDificultad(dificultadStr)

	fmt.Print("Ingrese la cantidad mínima de calorías (0 para omitir): ")
	minCaloriasStr, _ := reader.ReadString('\n')
	minCalorias, _ := strconv.Atoi(strings.TrimSpace(minCaloriasStr))

	ejercicios, err := lista.FiltrarEjercicios(tipoEjercicio, dificultad, minCalorias)
	if err != nil {
		fmt.Println("Error al filtrar ejercicios:", err)
	} else {
		for _, ejercicio := range ejercicios {
			fmt.Printf("Nombre: %s, Descripción: %s, Tiempo: %d, Calorías: %d, Tipo de Ejercicio: %v, Puntos por Tipo de Ejercicio: %v, Dificultad: %s\n",
				ejercicio.Nombre, ejercicio.Descripcion, ejercicio.Tiempo, ejercicio.Calorias, ejercicio.TipoDeEjercicio, ejercicio.PuntosPorTipoDeEjercicio, ejercicio.Dificultad)
		}
	}
}

func ModificarEjercicioMenu(lista *ListaDeEjercicios) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Ingrese el nombre del ejercicio a modificar: ")
	nombre, _ := reader.ReadString('\n')
	nombre = strings.TrimSpace(nombre)

	fmt.Print("Nueva Descripción: ")
	nuevaDescripcion, _ := reader.ReadString('\n')
	nuevaDescripcion = strings.TrimSpace(nuevaDescripcion)

	fmt.Print("Nuevo Tiempo (en segundos): ")
	nuevoTiempoStr, _ := reader.ReadString('\n')
	nuevoTiempo, _ := strconv.Atoi(strings.TrimSpace(nuevoTiempoStr))

	fmt.Print("Nuevas Calorías: ")
	nuevasCaloriasStr, _ := reader.ReadString('\n')
	nuevasCalorias, _ := strconv.Atoi(strings.TrimSpace(nuevasCaloriasStr))

	fmt.Print("Nuevos Tipos de Ejercicio (separados por coma,(fuerza,cardio,balance): ")
	nuevosTiposStr, _ := reader.ReadString('\n')
	nuevosTiposStr = strings.TrimSpace(nuevosTiposStr)
	nuevosTipos := strings.Split(nuevosTiposStr, ",")
	var nuevosTiposDeEjercicio []TipoEjercicio
	for _, t := range nuevosTipos {
		nuevosTiposDeEjercicio = append(nuevosTiposDeEjercicio, TipoEjercicio(strings.TrimSpace(t)))
	}

	fmt.Print("Nuevos Puntos por Tipo de Ejercicio (separados por coma, en el mismo orden que los tipos): ")
	nuevosPuntosStr, _ := reader.ReadString('\n')
	nuevosPuntosStr = strings.TrimSpace(nuevosPuntosStr)
	nuevosPuntosStrArr := strings.Split(nuevosPuntosStr, ",")
	var nuevosPuntosPorTipoDeEjercicio []int
	for _, p := range nuevosPuntosStrArr {
		puntos, _ := strconv.Atoi(strings.TrimSpace(p))
		nuevosPuntosPorTipoDeEjercicio = append(nuevosPuntosPorTipoDeEjercicio, puntos)
	}

	fmt.Print("Nueva Dificultad (principiante, intermedio, avanzado): ")
	nuevaDificultad, _ := reader.ReadString('\n')
	nuevaDificultad = strings.TrimSpace(nuevaDificultad)

	err := lista.ModificarEjercicio(nombre, nuevaDescripcion, nuevoTiempo, nuevasCalorias, nuevosTiposDeEjercicio, nuevosPuntosPorTipoDeEjercicio, nuevaDificultad)
	if err != nil {
		fmt.Println("Error al modificar ejercicio:", err)
	} else {
		fmt.Println("Ejercicio modificado exitosamente.")
		if err := GuardarEjercicios(lista); err != nil {
			fmt.Println("Error al guardar los cambios:", err)
		}
	}
}

// normalizeString convierte una cadena a minúsculas y elimina acentos
func NormalizeString(s string) string {
	var sb strings.Builder
	sb.Grow(len(s))

	for _, r := range s {
		// Convertir a minúscula
		lr := unicode.ToLower(r)
		// Eliminar acentos
		if lr >= 'a' && lr <= 'z' {
			sb.WriteRune(lr)
		} else {
			switch lr {
			case 'á':
				sb.WriteRune('a')
			case 'é':
				sb.WriteRune('e')
			case 'í':
				sb.WriteRune('i')
			case 'ó':
				sb.WriteRune('o')
			case 'ú':
				sb.WriteRune('u')
			case 'ü':
				sb.WriteRune('u')
			default:
				sb.WriteRune(r)
			}
		}
	}

	return sb.String()
}

func GuardarEjercicios(lista *ListaDeEjercicios) error {
	ejercicios, err := lista.ListarEjercicios()
	if err != nil {
		// No hay ejercicios para guardar, pero no es un error crítico
		if err.Error() == "no hay ningún ejercicio para listar" {
			return nil
		}
		return err
	}
	// Normalizar
	for _, ejercicio := range ejercicios {
		ejercicio.Nombre = NormalizeString(ejercicio.Nombre)
		ejercicio.Descripcion = NormalizeString(ejercicio.Descripcion)
		ejercicio.Dificultad = Dificultad(NormalizeString(string(ejercicio.Dificultad)))
		for i, tipo := range ejercicio.TipoDeEjercicio {
			ejercicio.TipoDeEjercicio[i] = TipoEjercicio(NormalizeString(string(tipo)))
		}
	}
	// Verificar si el archivo ya existe
	existe, err := ArchivoExiste("ejercicios.csv")
	if err != nil {
		return err
	}

	archivo, err := os.OpenFile("ejercicios.csv", os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer archivo.Close()

	// Escribir encabezado solo si el archivo no existía
	if !existe {
		if _, err := archivo.WriteString("Nombre,Descripción,Tiempo,Calorías,Tipo de Ejercicio,Puntos por Tipo de Ejercicio,Dificultad\n"); err != nil {
			return err
		}
	}

	return gocsv.MarshalFile(&ejercicios, archivo)
}

func ArchivoExiste(nombreArchivo string) (bool, error) {
	_, err := os.Stat(nombreArchivo)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
