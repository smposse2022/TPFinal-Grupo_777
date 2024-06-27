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
	lista.AgregarRutina("rutina1", ejerciciosFiltrados)
	rutinaAConsultar,_:=lista.ConsultarRutina("Rutina1")
	assert.Equal(t, "rutina1", rutinaAConsultar.Nombre)
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
	lista.AgregarRutina("rutina1", ejerciciosFiltrados)
	lista.AgregarRutina("rutina2", ejerciciosFiltrados)

	// eliminando una rutina existente
	lista.BorrarRutina("rutina1")
	// Listar las rutinas
	rutinasListadas,_:=lista.ListarRutinas()
	assert.Equal(t, 1,len(rutinasListadas))
	assert.Equal(t, "rutina2",rutinasListadas[0].Nombre)
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
	lista.AgregarRutina("rutina1", ejerciciosFiltrados)

	// consultar una rutina existente
	rutina, err := lista.ConsultarRutina("Rutina1")
	assert.NoError(t, err, "No se esperaba un error al consultar la rutina 1")
	assert.Equal(t, "rutina1", rutina.Nombre, "El nombre de la rutina no coincide")
}

func TestModificarRutina(t *testing.T) {
	lista := NewListaDeRutinas()
	listaEjercicios := NewListaDeEjercicios()

	// Agregar 3 ejercicios
	listaEjercicios.AgregarEjercicio("flexiones de brazos", "descripcion de flexiones de brazos", 20, 40, []TipoEjercicio{"fuerza"}, []int{50}, "intermedio")
	listaEjercicios.AgregarEjercicio("sentadillas", "descripcion de sentadillas", 10, 100, []TipoEjercicio{"fuerza"}, []int{50}, "principiante")
	listaEjercicios.AgregarEjercicio("estocadas", "descripcion de estocadas", 30, 300, []TipoEjercicio{"balance"}, []int{50}, "principiante")

	ejerciciosFiltrados,_:= listaEjercicios.ListarEjercicios()
	// prueba con nombre de rutina duplicado
	lista.AgregarRutina("rutina1", ejerciciosFiltrados)
	// Borramos un sjercicio de la lista de ejercicios y modificamos la rutina con la nuevaListaDeEjercicios
	listaEjercicios.BorrarEjercicio("sentadillas")
	nuevosEjercicios,_:=listaEjercicios.ListarEjercicios()

	lista.ModificarRutina("rutina1", nuevosEjercicios)
	rutinaModificada,_ :=lista.ConsultarRutina("rutina1")
	assert.Equal(t, 2, len(rutinaModificada.EjerciciosTotales)) // Debería ser 2, ya que eran 3, pero luego se modificó con sólo 2 ejercicios
}

func TestListarRutinas(t *testing.T) {
	lista := NewListaDeRutinas()
	listaEjercicios := NewListaDeEjercicios()

	// Agregar 3 ejercicios
	listaEjercicios.AgregarEjercicio("flexiones de brazos", "descripcion de flexiones de brazos", 20, 40, []TipoEjercicio{"fuerza"}, []int{50}, "intermedio")
	listaEjercicios.AgregarEjercicio("sentadillas", "descripcion de sentadillas", 10, 100, []TipoEjercicio{"fuerza"}, []int{50}, "principiante")
	listaEjercicios.AgregarEjercicio("estocadas", "descripcion de estocadas", 30, 300, []TipoEjercicio{"balance"}, []int{50}, "principiante")

	ejerciciosFiltrados,_:= listaEjercicios.ListarEjercicios()
	// prueba con nombre de rutina duplicado
	lista.AgregarRutina("rutina1", ejerciciosFiltrados)
	lista.AgregarRutina("rutina2", ejerciciosFiltrados)

	// listar con 2 rutinas
	rutinas,_:= lista.ListarRutinas()
	assert.Equal(t, 2, len(rutinas), "Se esperaba una lista de rutinas con dos elementos")
	assert.Equal(t, "rutina1", rutinas[0].Nombre)
}

// VERIFICA EL BORRADO DE UNA RUTINA ACTUAL
func TestBorrarRutinaExistente(t *testing.T) {
	lista := NewListaDeRutinas()
	listaEjercicios := NewListaDeEjercicios()

	// Agregar 3 ejercicios
	listaEjercicios.AgregarEjercicio("flexiones de brazos", "descripcion de flexiones de brazos", 20, 40, []TipoEjercicio{"fuerza"}, []int{50}, "intermedio")
	listaEjercicios.AgregarEjercicio("sentadillas", "descripcion de sentadillas", 10, 100, []TipoEjercicio{"fuerza"}, []int{50}, "principiante")
	listaEjercicios.AgregarEjercicio("estocadas", "descripcion de estocadas", 30, 300, []TipoEjercicio{"balance"}, []int{50}, "principiante")

	ejerciciosFiltrados,_:= listaEjercicios.ListarEjercicios()
	// prueba con nombre de rutina duplicado
	lista.AgregarRutina("rutina1", ejerciciosFiltrados)
	lista.AgregarRutina("rutina2", ejerciciosFiltrados)
	// Borrar la Rutina2
	lista.BorrarRutina("rutina2")
	rutinas,_:= lista.ListarRutinas()
	assert.Equal(t, 1, len(rutinas), "Se esperaba una lista de rutinas con un elemento")
	assert.Equal(t, "rutina1", rutinas[0].Nombre)

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
	assert.Equal(t, "saltar la soga",rutinaAConsultar.EjerciciosTotales[3].Nombre) // Debería tener 4 ejercicios, al haber agregado Saltar la soga

}


// Test Generación Automágica de Rutinas
func TestGeneracionAutomagica_Exito(t *testing.T) {
	lista := NewListaDeRutinas()
	listaEjercicios := NewListaDeEjercicios()

	// Agregar ejercicios
	listaEjercicios.AgregarEjercicio("ejercicio1", "descripcion de ejercicio 1", 12, 40, []TipoEjercicio{"fuerza", "balance"}, []int{50,60}, "principiante")
	listaEjercicios.AgregarEjercicio("ejercicio2", "descripcion de ejercicio 2", 7, 100, []TipoEjercicio{"fuerza","balance"}, []int{24,20}, "intermedio")
	listaEjercicios.AgregarEjercicio("ejercicio3", "descripcion de ejercicio 3", 12, 300, []TipoEjercicio{"balance","cardio"}, []int{19,12}, "principiante")
	listaEjercicios.AgregarEjercicio("ejercicio4", "descripcion de ejercicio 4", 40, 40, []TipoEjercicio{"fuerza"}, []int{16}, "avanzado")
	listaEjercicios.AgregarEjercicio("ejercicio5", "descripcion de ejercicio 5", 350, 100, []TipoEjercicio{"fuerza"}, []int{17}, "principiante")
	listaEjercicios.AgregarEjercicio("ejercicio6", "descripcion de ejercicio 6", 20, 300, []TipoEjercicio{"fuerza"}, []int{14}, "principiante")
	listaEjercicios.AgregarEjercicio("ejercicio7", "descripcion de ejercicio 7", 14, 40, []TipoEjercicio{"fuerza"}, []int{55}, "principiante")
	listaEjercicios.AgregarEjercicio("ejercicio8", "descripcion de ejercicio 8", 7, 100, []TipoEjercicio{"balance"}, []int{63}, "principiante")
	listaEjercicios.AgregarEjercicio("ejercicio9", "descripcion de ejercicio 9", 8, 300, []TipoEjercicio{"balance","fuerza"}, []int{59,44}, "principiante")
	listaEjercicios.AgregarEjercicio("ejercicio10", "descripcion de ejercicio 10", 35, 40, []TipoEjercicio{"cardio","balance"}, []int{58,42}, "principiante")
	listaEjercicios.AgregarEjercicio("ejercicio11", "descripcion de ejercicio 11", 60, 100, []TipoEjercicio{"fuerza","cardio"}, []int{50,20}, "principiante")
	listaEjercicios.AgregarEjercicio("ejercicio12", "descripcion de ejercicio 12", 53, 300, []TipoEjercicio{"cardio"}, []int{50}, "avanzado")


	// Llamar a la función GeneracionAutomagica
	rutina,_ := lista.GeneracionAutomagica("rutinaAutomagica1", 60, "fuerza", "principiante", listaEjercicios)
	// Verificar que la duración de la rutina sea correcta, que es 54, no llega a 60 con los ejercicios disponibles. Y si pone el siguiente ejercicio de menor tiempo se pasa, por eso no lo agrega
	assert.Equal(t, 54, rutina.Duracion, "la duración no es igual a la duración total de los ejercicios disponibles")
	assert.Equal(t, 4, len(rutina.EjerciciosTotales)) // debería haber 4 ejercicios, ejercicio 9, ejercicio 1, ejercicio 7, ejercicio 6, ejercicio 11, ejercicio 5
	assert.Equal(t, "ejercicio9", rutina.EjerciciosTotales[0].Nombre) // El 1er ejercicio que se agrega es el ejercicio9 al ser el más corto
	assert.Equal(t, "ejercicio1", rutina.EjerciciosTotales[1].Nombre) // El 2do ejercicio que se agrega es el ejercicio1 al ser el segundo más corto
	assert.Equal(t, "ejercicio7", rutina.EjerciciosTotales[2].Nombre) // El 3ro ejercicio que se agrega es el ejercicio7 al ser el tercero más corto
	}

func TestGeneracionAutomagica_Error_TipoEjercicioInexistente(t *testing.T) {
	lista := NewListaDeRutinas()
	listaEjercicios := NewListaDeEjercicios()

	// Agregar ejercicios
	listaEjercicios.AgregarEjercicio("ejercicio1", "descripcion de ejercicio 1", 12, 40, []TipoEjercicio{"fuerza", "balance"}, []int{50,60}, "principiante")
	listaEjercicios.AgregarEjercicio("ejercicio2", "descripcion de ejercicio 2", 7, 100, []TipoEjercicio{"fuerza","balance"}, []int{24,20}, "intermedio")
	listaEjercicios.AgregarEjercicio("ejercicio3", "descripcion de ejercicio 3", 12, 300, []TipoEjercicio{"balance"}, []int{19}, "principiante")
	listaEjercicios.AgregarEjercicio("ejercicio4", "descripcion de ejercicio 4", 40, 40, []TipoEjercicio{"fuerza"}, []int{16}, "avanzado")
	listaEjercicios.AgregarEjercicio("ejercicio5", "descripcion de ejercicio 5", 350, 100, []TipoEjercicio{"fuerza"}, []int{17}, "principiante")
	listaEjercicios.AgregarEjercicio("ejercicio6", "descripcion de ejercicio 6", 20, 300, []TipoEjercicio{"fuerza"}, []int{14}, "principiante")
	listaEjercicios.AgregarEjercicio("ejercicio7", "descripcion de ejercicio 7", 14, 40, []TipoEjercicio{"fuerza"}, []int{55}, "principiante")
	listaEjercicios.AgregarEjercicio("ejercicio8", "descripcion de ejercicio 8", 7, 100, []TipoEjercicio{"balance"}, []int{63}, "principiante")
	listaEjercicios.AgregarEjercicio("ejercicio9", "descripcion de ejercicio 9", 8, 300, []TipoEjercicio{"balance","fuerza"}, []int{59,44}, "principiante")
	listaEjercicios.AgregarEjercicio("ejercicio10", "descripcion de ejercicio 10", 35, 40, []TipoEjercicio{"balance"}, []int{42}, "principiante")
	listaEjercicios.AgregarEjercicio("ejercicio11", "descripcion de ejercicio 11", 60, 100, []TipoEjercicio{"fuerza"}, []int{20}, "principiante")
	listaEjercicios.AgregarEjercicio("ejercicio12", "descripcion de ejercicio 12", 53, 300, []TipoEjercicio{"balance"}, []int{50}, "avanzado")


	// Llamar a la función GeneracionAutomagica
	_, err := lista.GeneracionAutomagica("rutinaAutomagica1", 60, "cardio", "principiante", listaEjercicios)
	// Debería dar error por tipo inexistente
	assert.Error(t, err, "Se esperaba un error debido a que se solicita un tipo inexistente entre los ejercicios disponibles (cardio)")
}

func TestGeneracionAutomagica2_Exito(t *testing.T) {
	lista := NewListaDeRutinas()
	listaEjercicios := NewListaDeEjercicios()

	// Agregar ejercicios
	listaEjercicios.AgregarEjercicio("ejercicio1", "descripcion de ejercicio 1", 12, 40, []TipoEjercicio{"fuerza", "balance"}, []int{50,60}, "principiante")
	listaEjercicios.AgregarEjercicio("ejercicio2", "descripcion de ejercicio 2", 7, 100, []TipoEjercicio{"fuerza","balance"}, []int{24,20}, "intermedio")
	listaEjercicios.AgregarEjercicio("ejercicio3", "descripcion de ejercicio 3", 12, 300, []TipoEjercicio{"balance","cardio"}, []int{19,12}, "principiante")
	listaEjercicios.AgregarEjercicio("ejercicio4", "descripcion de ejercicio 4", 40, 40, []TipoEjercicio{"fuerza"}, []int{16}, "avanzado")
	listaEjercicios.AgregarEjercicio("ejercicio5", "descripcion de ejercicio 5", 350, 100, []TipoEjercicio{"fuerza"}, []int{17}, "principiante")
	listaEjercicios.AgregarEjercicio("ejercicio6", "descripcion de ejercicio 6", 20, 300, []TipoEjercicio{"fuerza"}, []int{14}, "principiante")
	listaEjercicios.AgregarEjercicio("ejercicio7", "descripcion de ejercicio 7", 14, 40, []TipoEjercicio{"fuerza"}, []int{55}, "principiante")
	listaEjercicios.AgregarEjercicio("ejercicio8", "descripcion de ejercicio 8", 6, 100, []TipoEjercicio{"balance"}, []int{63}, "principiante")
	listaEjercicios.AgregarEjercicio("ejercicio9", "descripcion de ejercicio 9", 8, 300, []TipoEjercicio{"balance","fuerza"}, []int{59,44}, "principiante")
	listaEjercicios.AgregarEjercicio("ejercicio10", "descripcion de ejercicio 10", 35, 40, []TipoEjercicio{"cardio","balance"}, []int{58,42}, "principiante")
	listaEjercicios.AgregarEjercicio("ejercicio11", "descripcion de ejercicio 11", 60, 100, []TipoEjercicio{"fuerza","cardio"}, []int{50,20}, "principiante")
	listaEjercicios.AgregarEjercicio("ejercicio12", "descripcion de ejercicio 12", 53, 300, []TipoEjercicio{"cardio"}, []int{50}, "avanzado")


	// Llamar a la función GeneracionAutomagica2
	rutina,_ := lista.GeneracionAutomagica2("automagica2", 400, listaEjercicios)
	assert.Equal(t, 21, rutina.Duracion) // Agrega los 3 ejercicios de menor tiempo. Se pasa de las 400 calorías y corta. Es el mínimo tiempo en que llega a las 400 calorías

}
func TestGeneracionAutomagica2_Error_CaloriasNoAlcanzada(t *testing.T) {
	lista := NewListaDeRutinas()
	listaEjercicios := NewListaDeEjercicios()

	// Agregar ejercicios
	listaEjercicios.AgregarEjercicio("ejercicio1", "descripcion de ejercicio 1", 12, 40, []TipoEjercicio{"fuerza", "balance"}, []int{50,60}, "principiante")
	listaEjercicios.AgregarEjercicio("ejercicio2", "descripcion de ejercicio 2", 7, 100, []TipoEjercicio{"fuerza","balance"}, []int{24,20}, "intermedio")
	listaEjercicios.AgregarEjercicio("ejercicio3", "descripcion de ejercicio 3", 12, 300, []TipoEjercicio{"balance","cardio"}, []int{19,12}, "principiante")
	listaEjercicios.AgregarEjercicio("ejercicio4", "descripcion de ejercicio 4", 40, 40, []TipoEjercicio{"fuerza"}, []int{16}, "avanzado")
	listaEjercicios.AgregarEjercicio("ejercicio5", "descripcion de ejercicio 5", 350, 100, []TipoEjercicio{"fuerza"}, []int{17}, "principiante")
	listaEjercicios.AgregarEjercicio("ejercicio6", "descripcion de ejercicio 6", 20, 300, []TipoEjercicio{"fuerza"}, []int{14}, "principiante")
	listaEjercicios.AgregarEjercicio("ejercicio7", "descripcion de ejercicio 7", 14, 40, []TipoEjercicio{"fuerza"}, []int{55}, "principiante")
	listaEjercicios.AgregarEjercicio("ejercicio8", "descripcion de ejercicio 8", 7, 100, []TipoEjercicio{"balance"}, []int{63}, "principiante")
	listaEjercicios.AgregarEjercicio("ejercicio9", "descripcion de ejercicio 9", 8, 300, []TipoEjercicio{"balance","fuerza"}, []int{59,44}, "principiante")
	listaEjercicios.AgregarEjercicio("ejercicio10", "descripcion de ejercicio 10", 35, 40, []TipoEjercicio{"cardio","balance"}, []int{58,42}, "principiante")
	listaEjercicios.AgregarEjercicio("ejercicio11", "descripcion de ejercicio 11", 60, 100, []TipoEjercicio{"fuerza","cardio"}, []int{50,20}, "principiante")
	listaEjercicios.AgregarEjercicio("ejercicio12", "descripcion de ejercicio 12", 53, 300, []TipoEjercicio{"cardio"}, []int{50}, "avanzado")


	// Llamar a la función GeneracionAutomagica
	_,err := lista.GeneracionAutomagica2("rutinaAutomagica1", 10000, listaEjercicios)
	// Debería dar error por no alcanzar el tiempo deseado
	assert.Error(t, err, "debería dar error al no alcanzar las calorías deseadas")
}

func TestGeneracionAutomagica3_Exito(t *testing.T) {
	lista := NewListaDeRutinas()
	listaEjercicios := NewListaDeEjercicios()

	// Agregar ejercicios
	listaEjercicios.AgregarEjercicio("ejercicio1", "descripcion de ejercicio 1", 12, 40, []TipoEjercicio{"fuerza", "balance"}, []int{50,60}, "principiante")
	listaEjercicios.AgregarEjercicio("ejercicio2", "descripcion de ejercicio 2", 7, 100, []TipoEjercicio{"fuerza","balance"}, []int{24,20}, "intermedio")
	listaEjercicios.AgregarEjercicio("ejercicio3", "descripcion de ejercicio 3", 12, 300, []TipoEjercicio{"balance","cardio"}, []int{19,12}, "principiante")
	listaEjercicios.AgregarEjercicio("ejercicio4", "descripcion de ejercicio 4", 40, 40, []TipoEjercicio{"fuerza"}, []int{16}, "avanzado")
	listaEjercicios.AgregarEjercicio("ejercicio5", "descripcion de ejercicio 5", 35, 100, []TipoEjercicio{"fuerza"}, []int{17}, "principiante")
	listaEjercicios.AgregarEjercicio("ejercicio6", "descripcion de ejercicio 6", 20, 300, []TipoEjercicio{"fuerza"}, []int{14}, "principiante")
	listaEjercicios.AgregarEjercicio("ejercicio7", "descripcion de ejercicio 7", 14, 40, []TipoEjercicio{"fuerza"}, []int{55}, "principiante")
	listaEjercicios.AgregarEjercicio("ejercicio8", "descripcion de ejercicio 8", 6, 100, []TipoEjercicio{"balance"}, []int{63}, "principiante")
	listaEjercicios.AgregarEjercicio("ejercicio9", "descripcion de ejercicio 9", 8, 300, []TipoEjercicio{"balance","fuerza"}, []int{59,44}, "principiante")
	listaEjercicios.AgregarEjercicio("ejercicio10", "descripcion de ejercicio 10", 35, 40, []TipoEjercicio{"cardio","balance"}, []int{58,42}, "principiante")
	listaEjercicios.AgregarEjercicio("ejercicio11", "descripcion de ejercicio 11", 60, 100, []TipoEjercicio{"fuerza","cardio"}, []int{50,20}, "principiante")
	listaEjercicios.AgregarEjercicio("ejercicio12", "descripcion de ejercicio 12", 53, 300, []TipoEjercicio{"cardio"}, []int{50}, "avanzado")


	// Llamar a la función GeneracionAutomagica2
	rutina,_ := lista.GeneracionAutomagica3("automagica2", 90, "fuerza", listaEjercicios)
	assert.Equal(t, 3, len(rutina.EjerciciosTotales)) // Deberían ser 3 ejercicios, ya que luego de agregar el ejercicio 7, 1 y 11 ya no puede agregar otro sin pasarse del tiempo
	assert.Equal(t, "ejercicio7", rutina.EjerciciosTotales[0].Nombre)
	assert.Equal(t, "ejercicio1", rutina.EjerciciosTotales[1].Nombre)
	assert.Equal(t, "ejercicio11", rutina.EjerciciosTotales[2].Nombre)
}