package main

type man struct {
	name string
	age  int
}

type woman struct {
	name string
	age  int
}

type skillz interface {
	work()
	rest()
}

func (m man) work(s string) {
	return
}

func (w woman) rest(r string) {
	return
}

func main() {
	dima := man{name: "Дима", age: 31}
	lena := woman{name: "Лена", age: 30}
	dima.work()
	lena.rest()

}
