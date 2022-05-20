package main

func main() {

	names := make([]int, 0)
	names2 := make([]int, 1)

	for i := 1; i < 3; i++ {
		names = append(names, i)
		names2 = append(names2, i)
	}

	println(len(names), names)
	println(len(names2), names2)

}
