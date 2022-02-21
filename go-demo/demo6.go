package main

import (
	"fmt"
	"unicode"
)

const (
	Left = iota
	Top
	Right
	Bottom
)

// 6. 机器人坐标问题
func main() {
	fmt.Println(move("R2(LF)", 0, 0, Top))
}

func move(cmd string, x0, y0, z0 int) (x, y, z int) {
	x, y, z = x0, y0, z0

	var repeat int
	var repeatCmd string

	for _, v := range cmd {
		switch {
		case unicode.IsNumber(v):
			repeat = repeat*10 + (int(v) - '0')
		case repeat > 0 && v != '(' && v != ')':
			repeatCmd += string(v)
		case v == ')':
			for i := 0; i < repeat; i++ {

				x, y, z = move(repeatCmd, x, y, z)
			}
			repeat = 0
			repeatCmd = ""
		case v == 'L':
			z = (z - 1 + 4) % 4
		case v == 'R':
			z = (z + 1) % 4
		case v == 'F':
			if z == Top || z == Bottom {
				y = y - z + 2
			}
			if z == Left || z == Right {
				x = x + z - 1
			}
		case v == 'B':
			if z == Top || z == Bottom {
				y = y + z - 2
			}
			if z == Left || z == Right {
				x = x - z + 1
			}

		}
	}
	return
}
