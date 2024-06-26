package main

import (
	"fmt"

	rutinaDeEjercicios "TP-2024-grupo_777"
)

func main() {
	listaEjercicios, err := rutinaDeEjercicios.CargarEjercicios(rutinaDeEjercicios.NewListaDeEjercicios())
	if err != nil {
		fmt.Println("Error al cargar los ejercicios:", err)
	}

	listaRutinas, err := rutinaDeEjercicios.CargarRutinas(rutinaDeEjercicios.NewListaDeRutinas())
	if err != nil {
		fmt.Println("Error al cargar las rutinas:", err)
	}

	for {
		fmt.Println("Seleccione una opción:")
		fmt.Println("1. Menú Ejercicios")
		fmt.Println("2. Menú Rutinas")
		fmt.Println("3. Salir")

		var opcion int
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			menuEjercicios(listaEjercicios)
		case 2:
			menuRutinas(listaRutinas, listaEjercicios)
		case 3:
			if err := rutinaDeEjercicios.GuardarEjercicios(listaEjercicios); err != nil {
				fmt.Println("Error al guardar los ejercicios:", err)
			}
			fmt.Println("Saliendo...")
			return
		default:
			fmt.Println("Opción no válida, intente nuevamente.")
		}
	}
	}
	
	func menuEjercicios(lista *rutinaDeEjercicios.ListaDeEjercicios)() {
		for {
			fmt.Println("Seleccione una opción:")
			fmt.Println("1. Agregar un ejercicio")
			fmt.Println("2. Listar ejercicios")
			fmt.Println("3. Borrar ejercicio")
			fmt.Println("4. Consultar ejercicio por nombre")
			fmt.Println("5. Filtrar ejercicios")
			fmt.Println("6. Modificar ejercicio")
			fmt.Println("7. Salir")
	
			var opcion int
			fmt.Scanln(&opcion)

			switch opcion {
			case 1:
				rutinaDeEjercicios.AgregarEjercicioMenu(lista)
			case 2:
				rutinaDeEjercicios.ListarEjerciciosMenu(lista)
			case 3:
				rutinaDeEjercicios.BorrarEjercicioMenu(lista)
			case 4:
				rutinaDeEjercicios.ConsultarEjercicioPorNombreMenu(lista)
			case 5:
				rutinaDeEjercicios.FiltrarEjerciciosMenu(lista)
			case 6:
				rutinaDeEjercicios.ModificarEjercicioMenu(lista)
			case 7:
				if err := rutinaDeEjercicios.GuardarEjercicios(lista); err != nil {
					fmt.Println("Error al guardar los ejercicios:", err)
				}
				fmt.Println("Saliendo del programa...")
				return
			default:
				fmt.Println("Opción no válida, intente nuevamente.")
			}
		}
	}

	func menuRutinas(lista *rutinaDeEjercicios.ListaDeRutinas, listaEjercicios *rutinaDeEjercicios.ListaDeEjercicios)() {
		for {
			for {
				fmt.Println("Seleccione una opción:")
				fmt.Println("1. Agregar una rutina")
				fmt.Println("2. Listar rutinas")
				fmt.Println("3. Borrar rutina")
				fmt.Println("4. Consultar rutina por nombre")
				fmt.Println("5. Agregar ejercicio de una rutina")
				fmt.Println("6. Eliminar ejercicio de una rutina")
				fmt.Println("7. GeneracionAutomagica")
				fmt.Println("8. GeneracionAutomagica 2")
				fmt.Println("9. GeneracionAutomagica 3")
				fmt.Println("10. Salir")
		
				var opcion int
				fmt.Scanln(&opcion)
	
				switch opcion {
				case 1:
					rutinaDeEjercicios.AgregarRutinaMenu(lista, listaEjercicios)
				case 2:
					rutinaDeEjercicios.ListarRutinasMenu(lista)
				case 3:
					rutinaDeEjercicios.BorrarRutinaMenu(lista)
				case 4:
					rutinaDeEjercicios.ConsultarRutinaPorNombreMenu(lista)
				case 5:
					rutinaDeEjercicios.AgregarEjercicioARutinaMenu(lista, listaEjercicios)
				case 6:
					rutinaDeEjercicios.EliminarEjercicioDeRutinaMenu(lista, listaEjercicios)
				case 7:
					rutinaDeEjercicios.GeneracionAutomagicaMenu(lista, listaEjercicios)
				case 8:
					//rutinaDeEjercicios.GeneracionAutomagica2Menu(lista, listaEjercicios)
				case 9:
				//	rutinaDeEjercicios.GeneracionAutomagica3v2Menu(lista, listaEjercicios)
				case 10:
					if err := rutinaDeEjercicios.GuardarRutinas(lista); err != nil {
						fmt.Println("Error al guardar las rutinas:", err)
					}
					fmt.Println("Saliendo del programa...")
					return
				default:
					fmt.Println("Opción no válida, intente nuevamente.")
				}
			}

		}
	}