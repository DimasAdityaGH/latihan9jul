package main

import "fmt"

func main() {
	//hello world
	fmt.Println("hello, world!")
	//var const string
	var name string = "jon"
	fmt.Println(name)

	const nama string = "john"
	fmt.Println(nama)

	//perbandingan

	var a = 10
	var b = 20
	var c = a == b
	fmt.Println(c)

	//operasi matematika
	var i = 10
	i++
	i += 9
	fmt.Println(i)

	var hasil = 10 * 10
	fmt.Println(hasil)

	//boolean

	var pelajar bool = true
	fmt.Println("are you student?", pelajar)

	//type data number

	var nilai8 int8 = 127
	var nilai16 int16 = 32767
	var nilai32 int32 = 2147483647
	var nilai64 int64 = 9223372036854775807
	fmt.Println(nilai8)
	fmt.Println(nilai16)
	fmt.Println(nilai32)
	fmt.Println(nilai64)

	//type declaration
	type str string
	type numb int8

	var produk str = "apple"
	var harga numb = 10
	fmt.Println(produk)
	fmt.Println(harga)

	//conversion

	var ni32 int32 = 100000
	var ni64 int64 = int64(nilai32)
	var ni8 int8 = int8(nilai32)
	fmt.Println(ni32)
	fmt.Println(ni64)
	fmt.Println(ni8)

	var example = 10
	var e = example
	var eInt = int8(e)
	fmt.Println(example)
	fmt.Println(eInt)

	//operasi boolean //array

	var dimas [03]int8
	dimas[0] = 80 //matematika[0]
	dimas[1] = 80 //ipa[1]
	dimas[2] = 80 //penjas[2]

	var hasilMtk = dimas[0] >= 80
	var hasilIpa = dimas[1] >= 80
	var hasilPenjas = dimas[2] >= 70

	var lulus = hasilMtk && hasilIpa && hasilPenjas
	fmt.Println(lulus)
}
