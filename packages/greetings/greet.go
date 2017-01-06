package greetings

func GreetMe(greeting, name string) string {
	return greeting + ", " + greet(name)
}

func greet(name string) string {
	return name
}

