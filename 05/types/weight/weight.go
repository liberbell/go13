package weight

type (
	Gram     int
	Kilogram Gram
	Ton      Kilogram
)

var (
	salt  Gram     = 100
	apple Kilogram = 5
	truck Ton      = 10
)

fmt.Printf("salt: %d, apple: %d, truck: %d", salt, apple, truck)