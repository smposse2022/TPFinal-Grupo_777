package rutinaDeEjercicios

import (
	"errors"
)

// Estructura de Ejercicio
// Debe considerarse que los ejercicios tendrán etiquetas para definir su tipo de ejercicio y dificultad. Las etiquetas serán un conjunto de palabras clave que permitirán clasificar los ejercicios. Por ejemplo, un ejercicio de sentadillas podría tener las etiquetas "fuerza" y "piernas". Sin embargo, las dificultades son únicas; por ejemplo, las sentadillas serán de dificultad "media" sólamente.
// Estructura de Ejercicio
type Ejercicio struct {
	Nombre                  string         `csv:"Nombre"`
	Descripcion             string         `csv:"Descripción"`
	Tiempo                  int            `csv:"Tiempo"`
	Calorias                int            `csv:"Calorías"`
	TipoDeEjercicio         []TipoEjercicio `csv:"Tipo de Ejercicio"`
	PuntosPorTipoDeEjercicio []int          `csv:"Puntos por Tipo de Ejercicio"`
	Dificultad              Dificultad     `csv:"Dificultad"`
}

type Dificultad string

const (
	Principiante Dificultad = "principiante"
	Intermedio   Dificultad = "intermedio"
	Avanzado     Dificultad = "avanzado"
)

type TipoEjercicio string

const (
	Fuerza  TipoEjercicio = "fuerza"
	Balance TipoEjercicio = "balance"
	Cardio  TipoEjercicio = "cardio"
)

func validarTipoDeEjercicio(tipo []TipoEjercicio) error {
	for _, t := range tipo {
		switch t {
		case Fuerza, Balance, Cardio:
			// Tipo de ejercicio válido
		default:
			return errors.New("tipo de ejercicio no válido: " + string(t))
		}
	}
	return nil
}

func validarDificultad(dificultad Dificultad) error {
	switch dificultad {
	case Principiante, Intermedio, Avanzado:
		// Dificultad válida
		return nil
	default:
		return errors.New("dificultad no válida: " + string(dificultad))
	}
}

// Estructura para almacenar los ejercicios
type ListaDeEjercicios struct {
	listaDeEjercicios map[string]*Ejercicio
}

// Inicializa una ListaDeEjercicios y crea el map vacío
func NewListaDeEjercicios() *ListaDeEjercicios {
	return &ListaDeEjercicios{listaDeEjercicios: make(map[string]*Ejercicio)}
}

// AgregarEjercicio crea un ejercicio y lo agrega al map de listaDeEjercicios
func (lista *ListaDeEjercicios) AgregarEjercicio(nombre string, descripcion string, tiempo int, calorias int, tipoDeEjercicio []TipoEjercicio, puntosPorTipoDeEjercicio []int, dificultad string) error {
	// Validar la longitud de los slices tipoDeEjercicio y puntosPorTipoDeEjercicio
	if len(tipoDeEjercicio) != len(puntosPorTipoDeEjercicio) {
		return errors.New("los slices de tipoDeEjercicio y puntosPorTipoDeEjercicio deben tener la misma longitud")
	}
	// Validar los tipos de ejercicio
	if err := validarTipoDeEjercicio(tipoDeEjercicio); err != nil {
		return err
	}
	// Validar la dificultad
	dificultadValida := Dificultad(dificultad)
	if err := validarDificultad(dificultadValida); err != nil {
		return err
	}

	ejercicio := &Ejercicio{
		Nombre:                  nombre,
		Descripcion:             descripcion,
		Tiempo:                  tiempo,
		Calorias:                calorias,
		TipoDeEjercicio:         tipoDeEjercicio,
		PuntosPorTipoDeEjercicio: puntosPorTipoDeEjercicio,
		Dificultad:              dificultadValida,
	}
	lista.listaDeEjercicios[nombre] = ejercicio
	return nil
}

// BorrarEjercicio elimina el par key value, a partir de la key indicada
func (lista *ListaDeEjercicios) BorrarEjercicio(nombre string) error {
	// Validar que el ejercicio exista
	if _, existe := lista.listaDeEjercicios[nombre]; !existe {
		return errors.New("el ejercicio no existe")
	}
	delete(lista.listaDeEjercicios, nombre)
	return nil
}

// ConsultarEjercicioPorNombre busca el ejercicio a partir del nombre indicado y devuelve el Ejercicio
func (lista *ListaDeEjercicios) ConsultarEjercicioPorNombre(nombre string) (*Ejercicio, error) {
	// Validar que el ejercicio exista
	ejercicio, existe := lista.listaDeEjercicios[nombre]
	if !existe {
		return nil, errors.New("el ejercicio no existe")
	}
	return ejercicio, nil
}

// FiltrarEjercicios permite filtrar los ejercicios que cumplan con los criterios indicados por parámetro
// y devuleve un slice con los ejercicios que cumplan
func (lista *ListaDeEjercicios) FiltrarEjercicios(tipo TipoEjercicio, dificultad Dificultad, minCalorias int) ([]*Ejercicio, error) {
	ejerciciosFiltrados := make([]*Ejercicio, 0)
	// Recorrer todos los ejercicios
	for _, ejercicio := range lista.listaDeEjercicios {
		// Se crea un booleano para ver si el ejercicio cumple los filtros o no
		// se inicializa en true y luego las comprobaciones van pasando a false los que no cumplan
		cumpleFiltro := true
		// Se verifica el tipo de ejercicio, si es que se pasa por parámetro
		if tipo != "" {
			tipoEncontrado := false
			for _, t := range ejercicio.TipoDeEjercicio {
				if t == tipo {
					tipoEncontrado = true
				}
			}
			if !tipoEncontrado {
				cumpleFiltro = false
			}
		}
		// Se verifica la dificultad, si es que se pasa por parámetro
		if dificultad != "" && ejercicio.Dificultad != Dificultad(dificultad) {
			cumpleFiltro = false
		}
		// Se verifican las calorías mínimas, si es que se pasa por parámetro
		if minCalorias > 0 && ejercicio.Calorias < minCalorias {
			cumpleFiltro = false
		}
		// Si el ejercicio pasa los filtros, se agrega al slice
		if cumpleFiltro {
			ejerciciosFiltrados = append(ejerciciosFiltrados, ejercicio)
		}
	}
	// Chequeamos que el slice tenga algún elemento o esté vacío
	if len(ejerciciosFiltrados) == 0 {
		return nil, errors.New("no hay ejercicios que cumplan esas condiciones")
	}
	return ejerciciosFiltrados, nil
}

// ModificarEjercicio permite modificar los valores de un ejercicio,
// a partir de identificar al mismo a partir de la key indicada
func (lista *ListaDeEjercicios) ModificarEjercicio(nombre string, nuevaDescripcion string, nuevoTiempo int, nuevasCalorias int, nuevoTipoDeEjercicio []TipoEjercicio, nuevosPuntosPorTipoDeEjercicio []int, nuevaDificultad string) error {
	// Validar la longitud de los slices tipoDeEjercicio y puntosPorTipoDeEjercicio
	if len(nuevoTipoDeEjercicio) != len(nuevosPuntosPorTipoDeEjercicio) {
		return errors.New("los slices de tipoDeEjercicio y puntosPorTipoDeEjercicio deben tener la misma longitud")
	}
	// Validar los tipos de ejercicio
	if err := validarTipoDeEjercicio(nuevoTipoDeEjercicio); err != nil {
		return err
	}
	// Validar la dificultad
	nuevaDificultadValida := Dificultad(nuevaDificultad)
	if err := validarDificultad(nuevaDificultadValida); err != nil {
		return err
	}
	// Validar que el ejercicio exista
	_, existe := lista.listaDeEjercicios[nombre]
	if !existe {
		return errors.New("el ejercicio no existe")
	}

	// Llamamos a AgregarEjercicio para actualizar los valores del ejercicio existente
	err := lista.AgregarEjercicio(nombre, nuevaDescripcion, nuevoTiempo, nuevasCalorias, nuevoTipoDeEjercicio, nuevosPuntosPorTipoDeEjercicio, string(nuevaDificultadValida))
	if err != nil {
		return err
	}

	return nil
}

// ListarEjercicios permite listar todos los ejercicios contenidos dentro del map
// de listaDeEjercicios
func (lista *ListaDeEjercicios) ListarEjercicios() ([]*Ejercicio, error) {
    if len(lista.listaDeEjercicios) == 0 {
        return nil, errors.New("no hay ningún ejercicio para listar")
    }
    ejercicios := make([]*Ejercicio, 0, len(lista.listaDeEjercicios))
	for _, ejercicio := range lista.listaDeEjercicios {
		ejercicios = append(ejercicios, ejercicio)
	}
	return ejercicios, nil
}