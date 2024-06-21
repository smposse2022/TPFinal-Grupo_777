package rutinaDeEjercicios

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestAgregarEjercicio(t *testing.T) {
	lista := NewListaDeEjercicios()

	// Caso de prueba exitoso
	lista.AgregarEjercicio("Flexiones de brazos", "Se realiza estando en posición inclinada, recostado hacia abajo, levantando el cuerpo únicamente con los brazos y bajando de nuevo al suelo", 10, 100, []TipoEjercicio{"fuerza"}, []int{50}, "principiante")

	ejercicio,_ := lista.ConsultarEjercicioPorNombre("Flexiones de brazos")
	assert.Equal(t,"Flexiones de brazos", ejercicio.Nombre)

}

func TestBorrarEjercicio(t *testing.T) {
	lista := NewListaDeEjercicios()

	// Agregar 2 ejercicios para poder borrarlo uno después
	lista.AgregarEjercicio("Flexiones de brazos", "Descripcion de flexiones de brazos", 20, 40, []TipoEjercicio{"fuerza"}, []int{50}, "intermedio")
	lista.AgregarEjercicio("Sentadillas", "Descripcion de sentadillas", 10, 100, []TipoEjercicio{"fuerza"}, []int{50}, "principiante")

	// Caso de prueba exitoso
	lista.BorrarEjercicio("Flexiones de brazos")
	assert.Equal(t, 1, len(lista.listaDeEjercicios))
	ejercicio,_ := lista.ConsultarEjercicioPorNombre("Sentadillas")
	assert.Equal(t,"Sentadillas", ejercicio.Nombre)

}

func TestConsultarEjercicioPorNombre(t *testing.T) {
	lista := NewListaDeEjercicios()

	// Agregar 2 ejercicios para poder consultarlos
	lista.AgregarEjercicio("Flexiones de brazos", "Descripcion de flexiones de brazos", 20, 40, []TipoEjercicio{"fuerza"}, []int{50}, "intermedio")
	lista.AgregarEjercicio("Sentadillas", "Descripcion de sentadillas", 10, 100, []TipoEjercicio{"fuerza"}, []int{50}, "principiante")

	// Caso de prueba exitoso
	ejercicio1,_ := lista.ConsultarEjercicioPorNombre("Flexiones de brazos")
	ejercicio2,_ := lista.ConsultarEjercicioPorNombre("Sentadillas")
	assert.Equal(t, "Flexiones de brazos", ejercicio1.Nombre, "Se esperaban que los nombres fueran iguales")
	assert.Equal(t, 20, ejercicio1.Tiempo, "Los ejercicios deberían ser iguales")
	assert.Equal(t, "Descripcion de flexiones de brazos", ejercicio1.Descripcion, "Los ejercicios deberían ser iguales")
	assert.Equal(t, "Sentadillas", ejercicio2.Nombre, "Se esperaban que los nombres fueran iguales")
	assert.Equal(t, 10, ejercicio2.Tiempo, "Los ejercicios deberían ser iguales")
	assert.Equal(t, "Descripcion de sentadillas", ejercicio2.Descripcion, "Los ejercicios deberían ser iguales")
}

func TestModificarEjercicio(t *testing.T) {
	lista := NewListaDeEjercicios()

	// Agregar 2 ejercicios para poder modificarlos
	lista.AgregarEjercicio("Flexiones de brazos", "Descripcion de flexiones de brazos", 20, 40, []TipoEjercicio{"fuerza"}, []int{50}, "intermedio")
	lista.AgregarEjercicio("Sentadillas", "Descripcion de sentadillas", 10, 100, []TipoEjercicio{"fuerza"}, []int{50}, "principiante")

	// Caso de prueba exitoso
	lista.ModificarEjercicio("Flexiones de brazos", "Nueva descripción de flexiones de brazos", 15, 150, []TipoEjercicio{"fuerza"}, []int{60}, "intermedio")
	// Consultar ejercicio para ver que se haya modificado
	ejercicioModificado,_:=lista.ConsultarEjercicioPorNombre("Flexiones de brazos")
	assert.Equal(t,"Nueva descripción de flexiones de brazos", ejercicioModificado.Descripcion )
	assert.Equal(t,15, ejercicioModificado.Tiempo )

}

func TestListarEjercicios(t *testing.T) {
	lista := NewListaDeEjercicios()
	// Agregar ejercicios para poder listar
	lista.AgregarEjercicio("Flexiones de brazos", "Descripcion de flexiones de brazos", 20, 40, []TipoEjercicio{"fuerza"}, []int{50}, "intermedio")
	lista.AgregarEjercicio("Sentadillas", "Descripcion de sentadillas", 10, 100, []TipoEjercicio{"fuerza"}, []int{50}, "principiante")
	lista.AgregarEjercicio("Estocadas", "Descripcion de estocadas", 12, 200, []TipoEjercicio{"balance"}, []int{70}, "avanzado")

	// Caso de prueba exitoso
	ejercicios, _ := lista.ListarEjercicios()

	// Verificar la longitud del slice devuelto
	assert.Equal(t, 3, len(ejercicios), "Se espera que haya 2 ejercicios en la listaDeEjercicios")
	assert.Equal(t, "Flexiones de brazos", ejercicios[0].Nombre, "Los ejercicios deberían ser iguales")
	assert.Equal(t, "Descripcion de sentadillas", ejercicios[1].Descripcion, "Los ejercicios deberían ser iguales")
	assert.Equal(t, 12, ejercicios[2].Tiempo, "Los ejercicios deberían ser iguales")

}

func TestFiltrarEjercicios(t *testing.T) {
	lista := NewListaDeEjercicios()
	// Agregar ejercicios para filtrar
	lista.AgregarEjercicio("Flexiones de brazos", "Descripcion de flexiones de brazos", 20, 40, []TipoEjercicio{"fuerza"}, []int{50}, "intermedio")
	lista.AgregarEjercicio("Sentadillas", "Descripcion de sentadillas", 10, 100, []TipoEjercicio{"fuerza"}, []int{50}, "principiante")
	lista.AgregarEjercicio("Estocadas", "Descripcion de estocadas", 12, 200, []TipoEjercicio{"balance"}, []int{70}, "avanzado")


	ejerciciosFiltrados, _ := lista.FiltrarEjercicios("balance", "", 0)
	assert.Equal(t, "Estocadas", ejerciciosFiltrados[0].Nombre, "Los ejercicios deberían ser iguales")
	// Caso de prueba donde se filtra por mínimo de calorías
	ejerciciosFiltrados, _ = lista.FiltrarEjercicios("", "", 100) // Debería traer un slice con 2 ejercicios
	assert.Equal(t, 2, len(ejerciciosFiltrados))
	// Caso de prueba donde se filtra por dificultad
	ejerciciosFiltrados, _ = lista.FiltrarEjercicios("", "intermedio", 0) // Debería traer un slice con 1 ejercicio
	assert.Equal(t, 1, len(ejerciciosFiltrados), "Los ejercicios deberían ser iguales")
	// Caso de prueba donde no se paasan los filtros vacíos, debería traer un slice con 3 elementos
	ejerciciosFiltrados, _ = lista.FiltrarEjercicios("", "", 0)
	assert.Equal(t, 3, len(ejerciciosFiltrados))
}
