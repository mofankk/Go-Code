package main

import "fmt"

type People struct{}

func (p *People) ShowA() {
	fmt.Println("ShowA")
	p.ShowB()
	return
}

func (p *People) ShowB() {
	fmt.Println("ShowB")
	return
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher ShowB")
	return
}

func main() {
	t := Teacher{}
	t.ShowB()
	t.ShowA()
}

/*
打印结果:
teacher ShowB
ShowA
ShowB
*/

/*
知识点：
不要在golang里面谈父类子类和继承。类比都不要类比。纯属误导人！！！

golang没有类与继承的概念。有接口，内嵌，组合。

与面向对象的类对应的概念。不是struct，是interface！！

那些觉得struct对应类的。我问你，是不是觉得只有struct才能有方法？

事实上golang里面任何类型都可以有方法。接收器的类型远不止是struct。甚至连int都可以是接收器。

----------------------------

结构体嵌套。在嵌套结构体中，People 称为内部类型，Teacher 称为外部类型；通过嵌套，内部类型的属性、方法，可以为外部类型所有，就好像是外部类型自己的一样。此外，外部类型还可以定义自己的属性和方法，甚至可以定义与内部相同的方法，这样内部类型的方法就会被“屏蔽”。这个例子中的 ShowB() 就是同名方法。

*/
