package rutinaDeEjercicios

import (
	"errors"
	"strings"
)

// Estructura de Rutina
type Rutina struct {
	Nombre                  string `csv:"Nombre"`
	Duracion                int `csv:"Duración"`
	EjerciciosTotales       []*Ejercicio `csv:"Ejercicios:"`
	TipoDeEjercicios        TipoEjercicio `csv:"Tipo de Ejercicios"`
	CaloriasQuemadasTotales int `csv:"Calorías totales"`
	Dificultad              Dificultad `csv:"Dificultad"`
}

// Función auxiliar para calcular la duración total de una rutina
func calcularDuracion(ejercicios []*Ejercicio) int {
	duracion := 0
	for _, ejercicio := range ejercicios {
		duracion += ejercicio.Tiempo
	}
	return duracion
}

// Función auxiliar para calcular las calorías quemadas totales de una rutina
func calcularCaloriasTotales(ejercicios []*Ejercicio) int {
	calorias := 0
	for _, ejercicio := range ejercicios {
		calorias += ejercicio.Calorias
	}
	return calorias
}

// Función auxiliar para calcular el tipo de ejercicios más frecuentes
func calcularTipoEjercicios(ejercicios []*Ejercicio) TipoEjercicio {
	// Usamos un map para registrar la cantidad de veces que aparece cada tipo de ejercicio
	frecuenciaTipoEjercicio := make(map[TipoEjercicio]int)
	// Registrar la cantidad de veces que aparece cada tipo de ejercicio
	for _, ejercicio := range ejercicios {
		for _, tipo := range ejercicio.TipoDeEjercicio {
			frecuenciaTipoEjercicio[tipo]++
		}
	}
	// Encontrar el tipo de ejercicio más frecuente
	frecuenciaMaxima := 0
	tipoMasFrecuente := TipoEjercicio("")
	for tipo, frecuencia := range frecuenciaTipoEjercicio {
		if frecuencia > frecuenciaMaxima {
			frecuenciaMaxima = frecuencia
			tipoMasFrecuente = tipo
		}
	}
	return tipoMasFrecuente
}

// Función auxiliar para calcular la dificultad más frecuentes
func calcularDificultadEjercicios(ejercicios []*Ejercicio) Dificultad {
	// Usamos un map para registrar la cantidad de veces que aparece cada dificultad
	dificultades := make(map[Dificultad]int)
	// Registrar la cantidad de veces que aparece cada dificultad
	for _, ejercicio := range ejercicios {
		dificultades[ejercicio.Dificultad]++
	}
	// Encontrar la dificultad más frecuente
	frecuenciaMaxima := 0
	dificultadMasFrecuente := Dificultad("")
	for dificultad, frecuencia := range dificultades {
		if frecuencia > frecuenciaMaxima {
			frecuenciaMaxima = frecuencia
			dificultadMasFrecuente = dificultad
		}
	}
	return dificultadMasFrecuente
}

// Estructura para almacenar las rutinas
type ListaDeRutinas struct {
	listaDeRutinas map[string]*Rutina
}

// Inicializa una ListaDeRutinas y crea el map vacío
func NewListaDeRutinas() *ListaDeRutinas {
	return &ListaDeRutinas{listaDeRutinas: make(map[string]*Rutina)}
}

// AgregarRutina crea una rutina y la agrega al map de listaDeRutinas
func (lista *ListaDeRutinas) AgregarRutina(nombre string, ejerciciosTotales []*Ejercicio) error {
	if len(ejerciciosTotales) == 0 {
		return errors.New("una rutina debe contener al menos 1 ejercicio")
	}
	// Normalizar nombre
	nombreNormalizado := NormalizeString(nombre)
	duracionRutina := calcularDuracion(ejerciciosTotales)
	caloriasRutina := calcularCaloriasTotales(ejerciciosTotales)
	tipoEjerciciosRutina := calcularTipoEjercicios(ejerciciosTotales)
	dificultadRutina := calcularDificultadEjercicios(ejerciciosTotales)
	rutina := &Rutina{
		Nombre:                  nombreNormalizado,
		Duracion:                duracionRutina,
		EjerciciosTotales:       ejerciciosTotales,
		TipoDeEjercicios:        tipoEjerciciosRutina,
		CaloriasQuemadasTotales: caloriasRutina,
		Dificultad:              dificultadRutina,
	}
	lista.listaDeRutinas[nombre] = rutina
	return nil
}

// BorrarRutina elimina el par key value, a partir de la key indicada
func (lista *ListaDeRutinas) BorrarRutina(nombre string) error {
	// Validar que la rutina exista
	_, error := lista.ConsultarRutina(nombre)
	if error != nil {
		return error
	}
	// Normalizar el nombre de búsqueda
	nombreNormalizado := NormalizeString(nombre)
	delete(lista.listaDeRutinas, nombreNormalizado)
	return nil
}

// ConsultarRutina busca la rutina a partir de la key indicada y devuelve la Rutina
func (lista *ListaDeRutinas) ConsultarRutina(nombre string) (*Rutina, error) {
		// Normalizar el nombre de búsqueda
		nombreNormalizado := NormalizeString(nombre)

		// Iterar sobre los ejercicios y buscar coincidencia parcial
		for _, rutina := range lista.listaDeRutinas {
			nombreRutinaNormalizado := NormalizeString(rutina.Nombre)
			if strings.Contains(nombreRutinaNormalizado, nombreNormalizado) {
				return rutina, nil
			}
		}
	
		return nil, errors.New("la rutina no existe")
	}

// ModificarRutina permite modificar los valores de una rutina,
// a partir de identificar la misma a partir de la key indicada
func (lista *ListaDeRutinas) ModificarRutina(nombre string, nuevosEjerciciosTotales []*Ejercicio) error {
	// Validar que la rutina exista
	if _, existe := lista.listaDeRutinas[nombre]; !existe {
		return errors.New("la rutina no existe")
	}
	// Normalizar el nombre
	nombreNormalizado := NormalizeString(nombre)
		lista.AgregarRutina(nombreNormalizado, nuevosEjerciciosTotales)
		return nil
	}

// ListarRutinas permite listar todas las rutinas contenidas dentro del map
// de listaDeRutinas
func (lista *ListaDeRutinas) ListarRutinas() ([]*Rutina, error) {
	if len(lista.listaDeRutinas) == 0 {
		return nil, errors.New("no hay ninguna rutina para listar")
	}
	rutinas := make([]*Rutina, 0, len(lista.listaDeRutinas))
	for _, rutina := range lista.listaDeRutinas {
		rutinas = append(rutinas, rutina)
	}
	return rutinas, nil
}

// Método para agregar un ejercicio al map de ejerciciosTotales de una Rutina en particular
func (lista *ListaDeRutinas) AgregarEjercicioARutina(nombre string, ejercicio *Ejercicio) error {
	// Verificar si la rutina existe en la lista
	rutina, error := lista.ConsultarRutina(nombre)
	if error != nil {
		return error
	}
	// Verificar si el ejercicio ya está dentro del slice de la Rutina
	for _, ejer := range rutina.EjerciciosTotales {
		if ejer.Nombre == ejercicio.Nombre {
			return errors.New("el ejercicio ya está dentro de la rutina")
		}
	}
	// Agregar el ejercicio al slice de ejerciciosTotales de la Rutina
	rutina.EjerciciosTotales = append(rutina.EjerciciosTotales, ejercicio)
	// Actualizar la rutina en el map de ListaDeRutinas
		// Normalizar el nombre
		nombreNormalizado := NormalizeString(nombre)
	lista.AgregarRutina(nombreNormalizado, rutina.EjerciciosTotales)
	return nil
}

// Método para eliminar un ejercicio al map de ejerciciosTotales de una Rutina en particular
func (lista *ListaDeRutinas) EliminaEjercicioDeRutina(nombre string, ejercicio *Ejercicio) error {
	// Verificar si la rutina existe en la lista
	rutina, error := lista.ConsultarRutina(nombre)
	if error != nil {
		return error
	}
	// Encontrar el índice dentro del slice donde está el Ejercicio buscado
	indice := -1
	for i := range rutina.EjerciciosTotales {
		if rutina.EjerciciosTotales[i].Nombre == ejercicio.Nombre {
			indice = i
		}
	}
	// Si no encontró el ejercicio dentro del slice, devolver un error
	if indice == -1 {
		return errors.New("el ejercicio no existe dentro de la rutina")
	}
	// Si encontró el ejercicio, se crea un nuevo slice con todos los elementos menos
	// el identificado
	nuevosEjerciciosTotales := make([]*Ejercicio, 0)

	for _, ejer := range rutina.EjerciciosTotales {
		if ejer.Nombre != rutina.EjerciciosTotales[indice].Nombre {
			nuevosEjerciciosTotales = append(nuevosEjerciciosTotales, ejer)
		}
	}
	// Se reemplaza el slice existente por el nuevo
	rutina.EjerciciosTotales = nuevosEjerciciosTotales

	return nil
}

// QuickSort
func QuickSort(ejercicios []*Ejercicio, low, high int) {
    if low < high {
        pi := partition(ejercicios, low, high)

        QuickSort(ejercicios, low, pi-1)
        QuickSort(ejercicios, pi+1, high)
    }
}

// Función auxiliar para particionar el slice de ejercicios
func partition(ejercicios []*Ejercicio, low, high int) int {
    pivot := ejercicios[high].Tiempo
    i := low - 1

    for j := low; j < high; j++ {
        if ejercicios[j].Tiempo < pivot {
            i++
            ejercicios[i], ejercicios[j] = ejercicios[j], ejercicios[i]
        }
    }

    ejercicios[i+1], ejercicios[high] = ejercicios[high], ejercicios[i+1]
    return i+1
}

func (lista *ListaDeRutinas) GeneracionAutomagica(nombre string, duracionTotal int, tipo TipoEjercicio, dificultad Dificultad, listaEjercicios *ListaDeEjercicios) (*Rutina, error) {
	// Filtrar los ejercicios que cumplan con los criterios especificados
	ejerciciosFiltrados, err := listaEjercicios.FiltrarEjercicios(NormalizeTipoEjercicio(tipo), dificultad, 0)
	if err != nil {
		return nil, err
	}

	// Ordenar los ejercicios por tiempo ascendente usando Quicksort
	QuickSort(ejerciciosFiltrados, 0, len(ejerciciosFiltrados)-1)

	// Seleccionar ejercicios hasta completar la duración total
	ejerciciosSeleccionados := make([]*Ejercicio, 0)
	tiempoAcumulado := 0
	for i := 0; i < len(ejerciciosFiltrados); i++ {
		if tiempoAcumulado+ejerciciosFiltrados[i].Tiempo <= duracionTotal {
			ejerciciosSeleccionados = append(ejerciciosSeleccionados, ejerciciosFiltrados[i])
			tiempoAcumulado += ejerciciosFiltrados[i].Tiempo
		} else {
			break
		}
	}

	// Verificar si se alcanzó la duración total deseada
	if tiempoAcumulado < duracionTotal {
		return nil, errors.New("no se puede alcanzar el tiempo deseado con los ejercicios existentes")
	}

	// Normalizar el nombre de la rutina antes de agregarla
	nombreNormalizado := NormalizeString(nombre)

	// Agregar la rutina a la lista de rutinas
	lista.AgregarRutina(nombreNormalizado, ejerciciosSeleccionados)
	
	// Consultar y devolver la rutina recién agregada
	rutina, err := lista.ConsultarRutina(nombreNormalizado)
	if err != nil {
		return nil, err
	}

	return rutina, nil
}

// Generación Automágica de Rutinas 2
func (lista *ListaDeRutinas) GeneracionAutomagica2(nombre string, caloriasObjetivo int, listaEjercicios *ListaDeEjercicios) (*Rutina, error) {
	// Obtener todos los ejercicios
	ejerciciosFiltrados, err := listaEjercicios.ListarEjercicios()
	if err != nil {
		return nil, err
	}
	// Verificar si hay ejercicios disponibles
	if len(ejerciciosFiltrados) == 0 {
		return nil, errors.New("no hay ejercicios disponibles")
	}
	// Ordenar los ejercicios por tiempo ascendente usando Quicksort
	QuickSort(ejerciciosFiltrados, 0, len(ejerciciosFiltrados)-1)

	// Seleccionar ejercicios hasta que alcancen las calorías objetivo
	ejerciciosSeleccionados := make([]*Ejercicio, 0)
	caloriasAcumuladas := 0
	for i := 0; i < len(ejerciciosFiltrados); i++ {
		if caloriasAcumuladas+ejerciciosFiltrados[i].Calorias <= caloriasObjetivo {
			ejerciciosSeleccionados = append(ejerciciosSeleccionados, ejerciciosFiltrados[i])
			caloriasAcumuladas += ejerciciosFiltrados[i].Calorias
		} else {
			break
		}
	}

	// Verificar si se alcanzaron las calorías objetivo
	if caloriasAcumuladas < caloriasObjetivo {
		return nil, errors.New("no es posible alcanzar las calorías objetivo con los ejercicios disponibles")
	}
		
	nombreNormalizado := NormalizeString(nombre)
	lista.AgregarRutina(nombreNormalizado,ejerciciosSeleccionados)
	rutina,_:= lista.ConsultarRutina(nombre)
	return rutina, nil
}

func (lista *ListaDeRutinas) GeneracionAutomagica3(nombre string, duracionTotal int, tipo TipoEjercicio, listaEjercicios *ListaDeEjercicios) (*Rutina, error) {
	// Normalizar los campos
	nombreNormalizado := NormalizeString(nombre)
	tipoNormalizado := NormalizeTipoEjercicio(tipo)

	// Filtrar los ejercicios por tipo y que tengan una duración menor o igual a la duración total
	ejerciciosFiltrados, err := listaEjercicios.FiltrarEjercicios(tipoNormalizado, "", duracionTotal)
	if err != nil {
		return nil, err
	}

	// Crear una tabla para almacenar los máximos puntos
	n := len(ejerciciosFiltrados)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, duracionTotal+1)
	}

	// Llenar la tabla de forma dinámica para maximizar los puntos
	for i := 1; i <= n; i++ {
		ejercicio := ejerciciosFiltrados[i-1]
		puntos := 0
		for j, t := range ejercicio.TipoDeEjercicio {
			if t == tipoNormalizado {
				puntos = ejercicio.PuntosPorTipoDeEjercicio[j]
				break
			}
		}
		for j := 1; j <= duracionTotal; j++ {
			if ejercicio.Tiempo <= j {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-ejercicio.Tiempo]+puntos)
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	// Recuperar los ejercicios seleccionados que maximizan los puntos dentro de la duración máxima
	tiempoRestante := duracionTotal
	rutinaEjercicios := []*Ejercicio{}
	for i := n; i > 0 && tiempoRestante > 0; i-- {
		if dp[i][tiempoRestante] != dp[i-1][tiempoRestante] {
			ejercicio := ejerciciosFiltrados[i-1]
			rutinaEjercicios = append(rutinaEjercicios, ejercicio)
			tiempoRestante -= ejercicio.Tiempo
		}
	}

	// Validar que se pudieron seleccionar ejercicios
	if len(rutinaEjercicios) == 0 {
		return nil, errors.New("no se pudieron seleccionar ejercicios que cumplan con los criterios")
	}

	// Normalizar el nombre antes de agregar la rutina a la lista
	lista.AgregarRutina(nombreNormalizado, rutinaEjercicios)

	// Consultar la rutina agregada para devolverla
	rutina,_ := lista.ConsultarRutina(nombreNormalizado)
	return rutina, nil
}