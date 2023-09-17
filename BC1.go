package main

import "fmt"

func main() {
	sisip := 4
	luasp := sisip * 4
	fmt.Println("luas lingkaran adalah", luasp)

	alas := 4
	tinggi := 6
	luasst := alas * tinggi / 2
	fmt.Println("luas segitiga adalah", luasst)

	rusukl := 7.0
	π := 3.14
	luasling := rusukl * rusukl * π
	fmt.Println("luas lingkatan adalah", luasling)

	massa := 10.0
	gravitasi := 10.0
	height := 2.0
	Ep := massa * gravitasi * height

	fmt.Println("energi potensial adalah", Ep)

	Velocity := 2.0
	Nolkomalima := 0.5
	Ek := Nolkomalima * massa * Velocity * Velocity
	fmt.Println("energi kinetik adalah", Ek)

	frekuensi := 50
	getaran := 500
	time := 10
	frekuensi1 := getaran / time
	getaran1 := frekuensi * time

	println("frekuensinya adalah", frekuensi1)
	println("getarannya adalah", getaran1)

	celcius := 50
	Kelvin := celcius + 273
	reamur := celcius * 4 / 5
	fahrenheit := celcius*9/5 + 2

	println("suhu", "c", celcius, "k", Kelvin, "k", reamur, "f", fahrenheit)
}
