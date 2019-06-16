package main

import (
	"fmt"
	"unsafe" //Saber tamaño de una variable en bytes
)

func main() {
	//Enteros con signo
	//var entero8 int8 //todo los enteros de 8 bit con signo (-128 hasta 127)
	//var entero16 int16 //(-32768 hasta 32767)
	var entero32 int32
	var entero64 int64

	// Enteros sin signo
	//var entero8 uint8 //Todos los enteros positivos (0 hasta 255)
	//var entero uint16 //(0 hasta 65535)
	//var entero32 uint32
	//var entero64 uint64

	//Alias
	//var enteroByte byte //Alias para uint8
	var enteroRune rune //para int32

	//Dependiendo de la plataforma
	//var enterouint uint //32 o 64 bytes
	var enteroInt int //32 o 64 bytes
	//var enteroUintptr uintptr //Entero que almacena un puntero
	//Convercion de datos

	entero32 = 20
	entero64 = 40
	enteroRune = 60
	enteroInt = 10

	fmt.Println(entero32 + int32(entero64)) //Convierte entero64 a int32
	fmt.Println(entero32 + enteroRune)
	fmt.Println(entero32 + int32(enteroInt))
	fmt.Println(unsafe.Sizeof(entero32), unsafe.Sizeof(enteroInt)) //Tamaño en Byte
}
