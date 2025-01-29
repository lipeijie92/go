package main

func fun3() {
	println("Hello, world")
}

func fun4() {
	fn := func() {
		println("Hello, world")
	}
	fn()
}
func adder() func(int) int {
	sum := 1
	return func(x int) int {
		sum += x
		return sum
	}
}

// 匿名方法
func fun5() {
	hello := func(name string) {
		println("Hello, ", name)
	}
	hello("lpj")
}

func fun6(age int, str ...string) {
	/*for _, v := range alias {
		println(v)
	}*/
	println(str[0])
	defer func() {
		println(str[0])
	}()

}
