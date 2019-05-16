package main

import (
	"fmt"
	"time"
)

/*
	* Cuando un minuto es divisible por tres, el reloj dirá la palabra "Fizz".
	* Cuando un minuto es divisible por cinco, el reloj dirá la palabra "Buzz".
	* Cuando un minuto es divisible por ambos, el reloj dirá "Fizz Buzz", con dos excepciones:
		* En la hora, en lugar de "Fizz Buzz", la puerta del reloj se abrirá, y el pájaro cucu saldrá y "Cucuu" entre una y doce veces, según la hora.
		* En la media hora, en lugar de "Fizz Buzz", la puerta del reloj se abrirá, y el cuco saldrá y "Cucuu" solo una vez.
	* Con minutos que no son divisibles entre tres o cinco, un sonido "tick".

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
		fmt.Println(RunCucu(dateNow()))
		time.Sleep(1 * time.Minute)
	}
}

func dateNow() string {
	hora := time.Now().Format("15:04")
	fmt.Println(hora)
	return hora
}

func RunCucu(hora string) string {

	// SOLUCIÓN
	return "Cucuu"
}
