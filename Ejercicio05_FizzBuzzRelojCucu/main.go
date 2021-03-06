package main

import (
	"fmt"
	"time"
)

/*
	* Cuando un minuto es divisible por tres, el reloj dirá la palabra "Fizz".
	* Cuando un minuto es divisible por cinco, el reloj dirá la palabra "Buzz".
	* Cuando un minuto es divisible por ambos, el reloj dirá "Fizz Buzz", con dos excepciones:
		* A la hora en punto, en lugar de "Fizz Buzz", el pájaro cucu saldrá y dira "Cucuu" entre una y doce veces, según la hora (01:00 = 1 vez && 13:00 = 1 vez).
		* En la media hora, en lugar de "Fizz Buzz", el cucu saldrá y "Cucuu" solo una vez.
	* Con minutos que no son divisibles entre tres o cinco, solo hara un sonido "tick".

	"13:34"       "tick"
	"21:00"       "Cucuu Cucuu Cucuu Cucuu Cucuu Cucuu Cucuu Cucuu Cucuu"
	"11:15"       "Fizz Buzz"
	"03:03"       "Fizz"
	"14:30"       "Cucuu"
	"08:55"       "Buzz"
	"00:00"       "Cucuu Cucuu Cucuu Cucuu Cucuu Cucuu Cucuu Cucuu Cucuu Cucuu Cucuu Cucuu"
	"12:00"       "Cucuu Cucuu Cucuu Cucuu Cucuu Cucuu Cucuu Cucuu Cucuu Cucuu Cucuu Cucuu"
*/

func main() {
	for {
		fmt.Println(RunCucu(DateNow()))
		time.Sleep(1 * time.Minute)
	}
}

// DateNow devuelve la HH:MM actual en formato de string (24hs)
func DateNow() string {
	hora := time.Now().Format("15:04")
	fmt.Println(hora)
	return hora
}

// RunCucu recibe HH:MM y devuelve lo que corresponda
func RunCucu(hora string) string {

	// SOLUCIÓN
	return "Cucuu"
}
