package main

func main() {
	age := 0
	println("adding 5")
	addAge(&age, 5)
	println("adding 10")
	addAge(&age, 10)

	println("now age:", age)

}

func addAge(age *int, grow int) {

	*age = *age + grow
}
