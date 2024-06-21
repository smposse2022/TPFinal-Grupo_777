package rutinaDeEjercicios

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcularDuracion(t *testing.T) {
	// Prueba con Lista de ejercicios vacía
	ejercicios := []*Ejercicio{}
	duracion := calcularDuracion(ejercicios)
	assert.Equal(t, 0, duracion, "La duración debe ser 0 para una lista vacía")

	// Prueba con Lista de ejercicios con elementos
	lista := NewListaDeEjercicios()

	// Agregar 2 ejercicios
	lista.AgregarEjercicio("Flexiones de brazos", "Descripcion de flexiones de brazos", 20, 40, []TipoEjercicio{"fuerza"}, []int{50}, "intermedio")
	lista.AgregarEjercicio("Sentadillas", "Descripcion de sentadillas", 10, 100, []TipoEjercicio{"fuerza"}, []int{50}, "principiante")

	ejerciciosFiltrados,_:= lista.ListarEjercicios()
	duracion = calcularDuracion(ejerciciosFiltrados)
	assert.Equal(t, 30, duracion, "La duración coincide con la esperada")
}

// COMPRUEBA LA FUNCION QUE CALCULA EL TIPO DE EJERCICIO DE UNA RUTINA
// TENIENDO EN CUENTA QUE EL MAS FRECUENTE ES “FUERZA”
func TestCalcularTipoEjerciciosConMayoriaDeFuerza(t *testing.T) {
	lista := NewListaDeEjercicios()

	// Agregar 3 ejercicios
	lista.AgregarEjercicio("Flexiones de brazos", "Descripcion de flexiones de brazos", 20, 40, []TipoEjercicio{"fuerza"}, []int{50}, "intermedio")
	lista.AgregarEjercicio("Sentadillas", "Descripcion de sentadillas", 10, 100, []TipoEjercicio{"fuerza"}, []int{50}, "principiante")
	lista.AgregarEjercicio("Estocadas", "Descripcion de estocadas", 30, 300, []TipoEjercicio{"balance"}, []int{50}, "principiante")

	ejerciciosFiltrados,_:= lista.ListarEjercicios()
	resultado := calcularTipoEjercicios(ejerciciosFiltrados)

	assert.Equal(t, "fuerza", string(resultado))
}


func TestAgregarRutina(t *testing.T) {
	lista := NewListaDeRutinas()
	listaEjercicios := NewListaDeEjercicios()

	// Agregar 3 ejercicios
	listaEjercicios.AgregarEjercicio("Flexiones de brazos", "Descripcion de flexiones de brazos", 20, 40, []TipoEjercicio{"fuerza"}, []int{50}, "intermedio")
	listaEjercicios.AgregarEjercicio("Sentadillas", "Descripcion de sentadillas", 10, 100, []TipoEjercicio{"fuerza"}, []int{50}, "principiante")
	listaEjercicios.AgregarEjercicio("Estocadas", "Descripcion de estocadas", 30, 300, []TipoEjercicio{"balance"}, []int{50}, "principiante")

	ejerciciosFiltrados,_:= listaEjercicios.ListarEjercicios()
	// prueba con nombre de rutina duplicado
	lista.AgregarRutina("Rutina1", ejerciciosFiltrados)
	rutinaAConsultar,_:=lista.ConsultarRutina("Rutina1")
	assert.Equal(t, "Rutina1", rutinaAConsultar.Nombre)
}

func TestBorrarRutina(t *testing.T) {
	lista := NewListaDeRutinas()
	listaEjercicios := NewListaDeEjercicios()

	// Agregar 3 ejercicios
	listaEjercicios.AgregarEjercicio("Flexiones de brazos", "Descripcion de flexiones de brazos", 20, 40, []TipoEjercicio{"fuerza"}, []int{50}, "intermedio")
	listaEjercicios.AgregarEjercicio("Sentadillas", "Descripcion de sentadillas", 10, 100, []TipoEjercicio{"fuerza"}, []int{50}, "principiante")
	listaEjercicios.AgregarEjercicio("Estocadas", "Descripcion de estocadas", 30, 300, []TipoEjercicio{"balance"}, []int{50}, "principiante")

	ejerciciosFiltrados,_:= listaEjercicios.ListarEjercicios()
	// prueba con nombre de rutina duplicado
	lista.AgregarRutina("Rutina1", ejerciciosFiltrados)
	lista.AgregarRutina("Rutina2", ejerciciosFiltrados)

	// eliminando una rutina existente
	lista.BorrarRutina("Rutina1")
	// Listar las rutinas
	rutinasListadas,_:=lista.ListarRutinas()
	assert.Equal(t, 1,len(rutinasListadas))
	assert.Equal(t, "Rutina2",rutinasListadas[0].Nombre)
}

func TestConsultarRutina(t *testing.T) {
	lista := NewListaDeRutinas()
	listaEjercicios := NewListaDeEjercicios()

	// Agregar 3 ejercicios
	listaEjercicios.AgregarEjercicio("Flexiones de brazos", "Descripcion de flexiones de brazos", 20, 40, []TipoEjercicio{"fuerza"}, []int{50}, "intermedio")
	listaEjercicios.AgregarEjercicio("Sentadillas", "Descripcion de sentadillas", 10, 100, []TipoEjercicio{"fuerza"}, []int{50}, "principiante")
	listaEjercicios.AgregarEjercicio("Estocadas", "Descripcion de estocadas", 30, 300, []TipoEjercicio{"balance"}, []int{50}, "principiante")

	ejerciciosFiltrados,_:= listaEjercicios.ListarEjercicios()
	// prueba con nombre de rutina duplicado
	lista.AgregarRutina("Rutina1", ejerciciosFiltrados)

	// consultar una rutina existente
	rutina, err := lista.ConsultarRutina("Rutina1")
	assert.NoError(t, err, "No se esperaba un error al consultar la rutina 1")
	assert.Equal(t, "Rutina1", rutina.Nombre, "El nombre de la rutina no coincide")
}

func TestModificarRutina(t *testing.T) {
	lista := NewListaDeRutinas()
	listaEjercicios := NewListaDeEjercicios()

	// Agregar 3 ejercicios
	listaEjercicios.AgregarEjercicio("Flexiones de brazos", "Descripcion de flexiones de brazos", 20, 40, []TipoEjercicio{"fuerza"}, []int{50}, "intermedio")
	listaEjercicios.AgregarEjercicio("Sentadillas", "Descripcion de sentadillas", 10, 100, []TipoEjercicio{"fuerza"}, []int{50}, "principiante")
	listaEjercicios.AgregarEjercicio("Estocadas", "Descripcion de estocadas", 30, 300, []TipoEjercicio{"balance"}, []int{50}, "principiante")

	ejerciciosFiltrados,_:= listaEjercicios.ListarEjercicios()
	// prueba con nombre de rutina duplicado
	lista.AgregarRutina("Rutina1", ejerciciosFiltrados)
	// Borramos un sjercicio de la lista de ejercicios y modificamos la rutina con la nuevaListaDeEjercicios
	listaEjercicios.BorrarEjercicio("Sentadillas")
	nuevosEjercicios,_:=listaEjercicios.ListarEjercicios()

	lista.ModificarRutina("Rutina1", nuevosEjercicios)
	rutinaModificada,_ :=lista.ConsultarRutina("Rutina1")
	assert.Equal(t, 2, len(rutinaModificada.EjerciciosTotales)) // Debería ser 2, ya que eran 3, pero luego se modificó con sólo 2 ejercicios
}

func TestListarRutinas(t *testing.T) {
	lista := NewListaDeRutinas()
	listaEjercicios := NewListaDeEjercicios()

	// Agregar 3 ejercicios
	listaEjercicios.AgregarEjercicio("Flexiones de brazos", "Descripcion de flexiones de brazos", 20, 40, []TipoEjercicio{"fuerza"}, []int{50}, "intermedio")
	listaEjercicios.AgregarEjercicio("Sentadillas", "Descripcion de sentadillas", 10, 100, []TipoEjercicio{"fuerza"}, []int{50}, "principiante")
	listaEjercicios.AgregarEjercicio("Estocadas", "Descripcion de estocadas", 30, 300, []TipoEjercicio{"balance"}, []int{50}, "principiante")

	ejerciciosFiltrados,_:= listaEjercicios.ListarEjercicios()
	// prueba con nombre de rutina duplicado
	lista.AgregarRutina("Rutina1", ejerciciosFiltrados)
	lista.AgregarRutina("Rutina2", ejerciciosFiltrados)

	// listar con 2 rutinas
	rutinas,_:= lista.ListarRutinas()
	assert.Equal(t, 2, len(rutinas), "Se esperaba una lista de rutinas con dos elementos")
	assert.Equal(t, "Rutina1", rutinas[0].Nombre)
}

// VERIFICA EL BORRADO DE UNA RUTINA ACTUAL
func TestBorrarRutinaExistente(t *testing.T) {
	lista := NewListaDeRutinas()
	listaEjercicios := NewListaDeEjercicios()

	// Agregar 3 ejercicios
	listaEjercicios.AgregarEjercicio("Flexiones de brazos", "Descripcion de flexiones de brazos", 20, 40, []TipoEjercicio{"fuerza"}, []int{50}, "intermedio")
	listaEjercicios.AgregarEjercicio("Sentadillas", "Descripcion de sentadillas", 10, 100, []TipoEjercicio{"fuerza"}, []int{50}, "principiante")
	listaEjercicios.AgregarEjercicio("Estocadas", "Descripcion de estocadas", 30, 300, []TipoEjercicio{"balance"}, []int{50}, "principiante")

	ejerciciosFiltrados,_:= listaEjercicios.ListarEjercicios()
	// prueba con nombre de rutina duplicado
	lista.AgregarRutina("Rutina1", ejerciciosFiltrados)
	lista.AgregarRutina("Rutina2", ejerciciosFiltrados)
	// Borrar la Rutina2
	lista.BorrarRutina("Rutina2")
	rutinas,_:= lista.ListarRutinas()
	assert.Equal(t, 1, len(rutinas), "Se esperaba una lista de rutinas con un elemento")
	assert.Equal(t, "Rutina1", rutinas[0].Nombre)

}

// VERIFICA QUE NO se puede BORRAR RUTINA INEXISTENTE
func TestBorrarRutinaNoExistente(t *testing.T) {
	rutinas := make(map[string]*Rutina)
	lista := ListaDeRutinas{listaDeRutinas: rutinas}

	error := lista.BorrarRutina("Rutina de martes")
	if error == nil || error.Error() != "la rutina no existe" {
		t.Errorf("Expected error 'la rutina no existe', got %v", error)
	}
}

// VERIFICA EN EL CASO DE AGREGAR UNA RUTINA DUPICADA
func TestAgregarEjercicioARutinaExistente(t *testing.T) {
	lista := NewListaDeRutinas()
	listaEjercicios := NewListaDeEjercicios()

	// Agregar 3 ejercicios
	listaEjercicios.AgregarEjercicio("Flexiones de brazos", "Descripcion de flexiones de brazos", 20, 40, []TipoEjercicio{"fuerza"}, []int{50}, "intermedio")
	listaEjercicios.AgregarEjercicio("Sentadillas", "Descripcion de sentadillas", 10, 100, []TipoEjercicio{"fuerza"}, []int{50}, "principiante")
	listaEjercicios.AgregarEjercicio("Estocadas", "Descripcion de estocadas", 30, 300, []TipoEjercicio{"balance"}, []int{50}, "principiante")

	ejerciciosFiltrados,_:= listaEjercicios.ListarEjercicios()

	lista.AgregarRutina("Rutina1", ejerciciosFiltrados)
	rutinaAConsultar,_:= lista.ConsultarRutina("Rutina1")
	assert.Equal(t, 3,len(rutinaAConsultar.EjerciciosTotales)) // Debería tener 3 ejercicios
	listaEjercicios.AgregarEjercicio("Saltar la soga", "Descripcion de saltar la soga", 30, 300, []TipoEjercicio{"balance"}, []int{50}, "principiante")
	ejercicioAAgregar,_ := listaEjercicios.ConsultarEjercicioPorNombre("Saltar la soga")
	lista.AgregarEjercicioARutina("Rutina1",ejercicioAAgregar)
	assert.Equal(t, 4,len(rutinaAConsultar.EjerciciosTotales)) // Debería tener 4 ejercicios, al haber agregado Saltar la soga
	assert.Equal(t, "Saltar la soga",rutinaAConsultar.EjerciciosTotales[3].Nombre) // Debería tener 4 ejercicios, al haber agregado Saltar la soga

}


// Test Generación Automágica de Rutinas
func TestGeneracionAutomagica_Exito(t *testing.T) {
	lista := NewListaDeRutinas()
	listaEjercicios := NewListaDeEjercicios()

	// Agregar 3 ejercicios
	listaEjercicios.AgregarEjercicio("Flexiones de brazos", "Descripcion de flexiones de brazos", 20, 40, []TipoEjercicio{"fuerza"}, []int{50}, "principiante")
	listaEjercicios.AgregarEjercicio("Sentadillas", "Descripcion de sentadillas", 10, 100, []TipoEjercicio{"fuerza"}, []int{50}, "principiante")
	listaEjercicios.AgregarEjercicio("Estocadas", "Descripcion de estocadas", 30, 300, []TipoEjercicio{"balance"}, []int{50}, "principiante")

	// Llamar a la función GeneracionAutomagica
	rutina, _ := lista.GeneracionAutomagica("RutinaAutomagica1", 30, "fuerza", "principiante", listaEjercicios)
	// Verificar que la duración de la rutina sea correcta: 20+10 y que tenga 2 ejercicios:Flexiones de brazos y Sentadillas
	assert.Equal(t, 30, rutina.Duracion, "La duración no es igual a la duración total de los ejercicios disponibles")
	assert.Equal(t, "Flexiones de brazos", rutina.EjerciciosTotales[0].Nombre)
	assert.Equal(t, "Sentadillas", rutina.EjerciciosTotales[1].Nombre)
}

func TestGeneracionAutomagica_Error_TipoEjercicioInexistente(t *testing.T) {
	lista := NewListaDeRutinas()
	listaEjercicios := NewListaDeEjercicios()

	// Agregar 3 ejercicios
	listaEjercicios.AgregarEjercicio("Flexiones de brazos", "Descripcion de flexiones de brazos", 20, 40, []TipoEjercicio{"fuerza"}, []int{50}, "principiante")
	listaEjercicios.AgregarEjercicio("Sentadillas", "Descripcion de sentadillas", 10, 100, []TipoEjercicio{"fuerza"}, []int{50}, "principiante")
	listaEjercicios.AgregarEjercicio("Estocadas", "Descripcion de estocadas", 30, 300, []TipoEjercicio{"balance"}, []int{50}, "principiante")

	// Llamar a la función GeneracionAutomagica
	_, err := lista.GeneracionAutomagica("RutinaAutomagica1", 30, "cardio", "principiante", listaEjercicios)
	assert.Error(t, err, "Se esperaba un error debido al tipo de ejercicio inexistente, pero no se recibió ningún error.")
}

func TestGeneracionAutomagica_Error_DuracionNoAlcanzada(t *testing.T) {
	lista := NewListaDeRutinas()
	listaEjercicios := NewListaDeEjercicios()

	// Agregar 3 ejercicios
	listaEjercicios.AgregarEjercicio("Flexiones de brazos", "Descripcion de flexiones de brazos", 20, 40, []TipoEjercicio{"fuerza"}, []int{50}, "principiante")
	listaEjercicios.AgregarEjercicio("Sentadillas", "Descripcion de sentadillas", 10, 100, []TipoEjercicio{"fuerza"}, []int{50}, "principiante")
	listaEjercicios.AgregarEjercicio("Estocadas", "Descripcion de estocadas", 30, 300, []TipoEjercicio{"balance"}, []int{50}, "principiante")

	// Llamar a la función GeneracionAutomagica
	_, err := lista.GeneracionAutomagica("RutinaAutomagica1", 70, "cardio", "principiante", listaEjercicios)
	assert.Error(t, err, "Se esperaba un error debido a que no se alcanza el tiempo deseado")
}

// Test Generación Automágica2 de Rutinas
func TestGeneracionAutomagica2_Exito(t *testing.T) {
	lista := NewListaDeRutinas()
	listaEjercicios := NewListaDeEjercicios()

	// Agregar 3 ejercicios
	listaEjercicios.AgregarEjercicio("Flexiones de brazos", "Descripcion de flexiones de brazos", 20, 40, []TipoEjercicio{"fuerza"}, []int{50}, "principiante")
	listaEjercicios.AgregarEjercicio("Sentadillas", "Descripcion de sentadillas", 10, 100, []TipoEjercicio{"fuerza"}, []int{50}, "principiante")
	listaEjercicios.AgregarEjercicio("Estocadas", "Descripcion de estocadas", 30, 300, []TipoEjercicio{"balance"}, []int{50}, "principiante")

	// Llamar a la función GeneracionAutomagica2
	rutina, _ := lista.GeneracionAutomagica2("RutinaAutomagica2", 140, listaEjercicios)
	// Verificar que la duración de la rutina sea correcta: 20+10 y que tenga 2 ejercicios:Flexiones de brazos y Sentadillas
	assert.Equal(t, 30, rutina.Duracion, "La duración no es igual a la duración total de los ejercicios disponibles")
	assert.Equal(t, "Flexiones de brazos", rutina.EjerciciosTotales[0].Nombre)
	assert.Equal(t, "Sentadillas", rutina.EjerciciosTotales[1].Nombre)
	assert.Equal(t, 140, rutina.CaloriasQuemadasTotales)
}

func TestGeneracionAutomagica2_Error_CaloriasInsuficientes(t *testing.T) {
	lista := NewListaDeRutinas()
	listaEjercicios := NewListaDeEjercicios()

	// Agregar 3 ejercicios
	listaEjercicios.AgregarEjercicio("Flexiones de brazos", "Descripcion de flexiones de brazos", 20, 40, []TipoEjercicio{"fuerza"}, []int{50}, "principiante")
	listaEjercicios.AgregarEjercicio("Sentadillas", "Descripcion de sentadillas", 10, 100, []TipoEjercicio{"fuerza"}, []int{50}, "principiante")
	listaEjercicios.AgregarEjercicio("Estocadas", "Descripcion de estocadas", 30, 300, []TipoEjercicio{"balance"}, []int{50}, "principiante")

	// Llamar a la función GeneracionAutomagica2
	_, err := lista.GeneracionAutomagica2("RutinaAutomagica2", 900, listaEjercicios)
	assert.Error(t, err, "Debería dar error por calorías inalcanzables")
}

/////// Automagica3 v2
func TestGeneracionAutomagica3v2(t *testing.T) {
	lista := NewListaDeRutinas()
	listaEjercicios := NewListaDeEjercicios()

	// Agregar 3 ejercicios
	listaEjercicios.AgregarEjercicio("Flexiones de brazos", "Descripcion de flexiones de brazos", 10, 40, []TipoEjercicio{"fuerza","balance"}, []int{100,50}, "principiante")
	listaEjercicios.AgregarEjercicio("Sentadillas", "Descripcion de sentadillas", 20, 100, []TipoEjercicio{"fuerza"}, []int{30}, "principiante")
	listaEjercicios.AgregarEjercicio("Estocadas", "Descripcion de estocadas", 20, 300, []TipoEjercicio{"fuerza"}, []int{20}, "principiante")

	// Llamar a la función GeneracionAutomagica2
	lista.GeneracionAutomagica3v2("RutinaAutomagica3", 40, "fuerza",listaEjercicios)
	rutina,_ := lista.ConsultarRutina("RutinaAutomagica3")
	assert.Equal(t, 2, len(rutina.EjerciciosTotales), "Debería tener 2 ejercicios, Flexiones de brazos y Sentadillas")
}