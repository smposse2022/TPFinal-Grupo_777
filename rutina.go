package rutinaDeEjercicios

import (
	"errors"
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
	duracionRutina := calcularDuracion(ejerciciosTotales)
	caloriasRutina := calcularCaloriasTotales(ejerciciosTotales)
	tipoEjerciciosRutina := calcularTipoEjercicios(ejerciciosTotales)
	dificultadRutina := calcularDificultadEjercicios(ejerciciosTotales)
	rutina := &Rutina{
		Nombre:                  nombre,
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
	delete(lista.listaDeRutinas, nombre)
	return nil
}

// ConsultarRutina busca la rutina a partir de la key indicada y devuelve la Rutina
func (lista *ListaDeRutinas) ConsultarRutina(nombre string) (*Rutina, error) {
	// Validar que la rutina exista
	rutina, existe := lista.listaDeRutinas[nombre]
	if !existe {
		return nil, errors.New("la rutina no existe")
	}
	return rutina, nil
}

// ModificarRutina permite modificar los valores de una rutina,
// a partir de identificar la misma a partir de la key indicada
func (lista *ListaDeRutinas) ModificarRutina(nombre string, nuevosEjerciciosTotales []*Ejercicio) error {
	// Validar que la rutina exista
	if _, existe := lista.listaDeRutinas[nombre]; !existe {
		return errors.New("la rutina no existe")
	}
		lista.AgregarRutina(nombre, nuevosEjerciciosTotales)
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
	lista.AgregarRutina(nombre, rutina.EjerciciosTotales)
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

// Heapsort para ordenamiento por duración
func HeapSort(ejercicios []*Ejercicio) []*Ejercicio {
	size := len(ejercicios)
	heapify(ejercicios) // Construye un heap máximo a partir del arreglo

	// En cada iteración, se extrae el elemento máximo del heap y se coloca al final del arreglo.
	// Luego, se ajusta el heap hacia abajo para mantener la propiedad del heap.
	end := size - 1
	for end > 0 {
		// Intercambia el máximo actual con el último elemento del arreglo
		ejercicios[end], ejercicios[0] = ejercicios[0], ejercicios[end]
		// Ajusta el heap hacia abajo (restaura la propiedad del heap)
		downHeap(ejercicios, 0, end-1)
		// Reduce el tamaño efectivo del arreglo en 1 para excluir el elemento ya ordenado
		end--
	}
	return ejercicios
}

func heapify(ejercicios []*Ejercicio) {
	size := len(ejercicios)
	// El primer nodo que tiene hijos se encuentra en la posición (size - 2) / 2.
	start := (size - 2) / 2

	// Comienza desde el último padre y ajusta cada subárbol hacia abajo para cumplir la propiedad del heap.
	for start >= 0 {
		downHeap(ejercicios, start, size-1)
		start--
	}
}

func downHeap(ejercicios []*Ejercicio, start, end int) {
	father := start
	leftSon := father*2 + 1
	rightSon := leftSon + 1

	// Mientras el padre tenga al menos un hijo
	for leftSon <= end {
		// Si el padre tiene dos hijos, nos quedamos con el menor
		if rightSon <= end && ejercicios[rightSon].Tiempo < ejercicios[leftSon].Tiempo {
			leftSon = rightSon
		}
		// Si el hijo es menor que el padre, los intercambiamos
		if ejercicios[leftSon].Tiempo < ejercicios[father].Tiempo {
			ejercicios[leftSon], ejercicios[father] = ejercicios[father], ejercicios[leftSon]
			// El hijo se convierte en el padre
			father = leftSon
			leftSon = father*2 + 1
			rightSon = leftSon + 1
		} else {
			return
		}
	}
}

// Generación Automágica de Rutinas 1
func (lista *ListaDeRutinas) GeneracionAutomagica(nombre string, duracionTotal int, tipo TipoEjercicio, dificultad Dificultad, listaEjercicios *ListaDeEjercicios) (*Rutina, error) {
	// Filtrar los ejercicios que cumplan con los criterios especificados. Filtrar filtra por tipo, dificultad y mincalorias
	ejerciciosFiltrados, err := listaEjercicios.FiltrarEjercicios(tipo, dificultad, 0)
	if err != nil {
		return nil, err
	}
	// Ordenar los ejercicios filtrados por duración ascendente usando HeapSort
	ejerciciosOrdenados := HeapSort(ejerciciosFiltrados)

	var rutinaEjerciciosOrdenados []*Ejercicio
	tiempoAcumulado := 0

	// Seleccionar los ejercicios de manera greedy
	for _, ejercicio := range ejerciciosOrdenados {
		if tiempoAcumulado+ejercicio.Tiempo <= duracionTotal {
			rutinaEjerciciosOrdenados = append(rutinaEjerciciosOrdenados, ejercicio)
			tiempoAcumulado += ejercicio.Tiempo
		}
	}

	// Agregar la rutina a la lista de rutinas
	lista.AgregarRutina(nombre,rutinaEjerciciosOrdenados)
	rutina,err:= lista.ConsultarRutina(nombre)

	return rutina, err
}

// Generación Automágica de Rutinas 2


func (lista *ListaDeRutinas) GeneracionAutomagica2(nombre string, caloriasObjetivo int, listaEjercicios *ListaDeEjercicios) (*Rutina, error) {
	// Obtener todos los ejercicios
	ejerciciosDisponibles, err := listaEjercicios.ListarEjercicios()
	if err != nil {
		return nil, err
	}

	// Verificar si hay ejercicios disponibles
	if len(ejerciciosDisponibles) == 0 {
		return nil, errors.New("no hay ejercicios disponibles")
	}

	// Ordenar los ejercicios por duración usando HeapSort
	ejerciciosOrdenados := HeapSort(ejerciciosDisponibles)

	var rutinaEjercicios []*Ejercicio
	caloriasAcumuladas := 0

	for _, ejercicio := range ejerciciosOrdenados {
		// Verificar si agregar el ejercicio excede las calorías objetivo
		if caloriasAcumuladas+ejercicio.Calorias <= caloriasObjetivo {
			// Agregar el ejercicio a la rutina
			rutinaEjercicios = append(rutinaEjercicios, ejercicio)
			caloriasAcumuladas += ejercicio.Calorias

			// Verificar si se alcanzaron las calorías objetivo
			if caloriasAcumuladas >= caloriasObjetivo {
				break
			}
		}
	}

	lista.AgregarRutina(nombre,rutinaEjercicios)
	rutina,_:= lista.ConsultarRutina(nombre)

	return rutina, nil
}

// Versión 2 de Automagicas3
func (lista *ListaDeRutinas) GeneracionAutomagica3v2(nombre string, duracionTotal int, tipo TipoEjercicio, listaEjercicios *ListaDeEjercicios) (*Rutina, error) {
	// Filtrar los ejercicios por tipo
	ejerciciosFiltrados, err := listaEjercicios.FiltrarEjercicios(tipo, "", 0)
	if err != nil {
		return nil, err
	}

	// Crear una tabla para almacenar los máximos puntos
	n := len(ejerciciosFiltrados)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, duracionTotal+1)
	}

	// Llenar la tabla de forma dinámica
	for i := 1; i <= n; i++ {
		ejercicio := ejerciciosFiltrados[i-1]
		puntos := 0
		for j, t := range ejercicio.TipoDeEjercicio {
			if t == tipo {
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

	// Recuperar los ejercicios seleccionados
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
		return nil, errors.New("no se pudieron seleccionar ejercicios")
	}
	lista.AgregarRutina(nombre,rutinaEjercicios)
	rutina,err:= lista.ConsultarRutina(nombre)
	return rutina, err
}
