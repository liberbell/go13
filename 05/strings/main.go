package main

import "fmt"

func main() {
	var s string
	s = "How are you."
	s = `How are you?`
	fmt.Println(s)

	s = "<html>\n\t<body>\"Hello\"</body>\n</html>"
}
