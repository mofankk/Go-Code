package main

func main() {
	e := 6
	f := 7
	m := make(map[string]*int)
	a, b, c, d := 1, 2, 3, 4
	m["a"] = &a
	m["b"] = &b
	m["c"] = &c
	m["d"] = &d
	i := 0
	//	for k := range m {
	//		if i%2 == 0 {
	//			*m[k] = f
	//		} else {
	//
	//			m[k] = &e
	//		}
	//		i++
	//	}
	//	println("-------------------")
	//
	//	for k, _ := range m {
	//		println(*m[k])
	//	}
	//
	println("-------------------")
	for _, v := range m {
		if i%2 == 0 {
			*v = f
		} else {

			*v = e
		}
		i++
	}
	println("-------------------")

	for _, v := range m {
		println(*v)
	}

}
