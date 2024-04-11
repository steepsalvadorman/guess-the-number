package main

import (
	"fmt"
	"math/rand"
)

func main(){
	jugar()
}

func jugar(){
	numAleatorio := rand.Intn(100)
	var numIngresado int
	var intentos int
	const maxIntentos = 10

	for intentos < maxIntentos {
		intentos++
		fmt.Printf("Ingresa un numero (intentos restantes: %d): " , maxIntentos - intentos+1)
		fmt.Scanln(&numIngresado)

		if numIngresado == numAleatorio {
			fmt.Println("Felicitaciones, adivinaste el numero!")
			jugarNuevamente()
			return
		} else if numIngresado < numAleatorio {
			fmt.Println("El numero a adivinar es mayor")
		} else if numIngresado > numAleatorio {
			fmt.Println("El numero a adivinar es menor")
		}
	}

	fmt.Println("Se acabaron los intentos. El numero era: " , numAleatorio)
	jugarNuevamente()
}


func jugarNuevamente(){
	var eleccion string
	fmt.Println("¿Quieres jugar nuevamente? (s/n): ")
	fmt.Scanln(&eleccion)

	switch eleccion {
	case "s":
		jugar()
	case "n":
		fmt.Println("Gracias por jugar")
	default:
		fmt.Println("Eleccion invalida. Intentalo nuevamente.")
		jugarNuevamente()
	}
}