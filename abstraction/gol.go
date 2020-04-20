package abstraction

type IGol interface {
	Start()
	StartFromFile(string)
	Next()
	Print()
}
