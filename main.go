package main

import (
	"flag"
	concrete "gol/concrete"
)

/*Constants loaded when respective arg is not passed*/
const (
	DEFAULT_ROWS_SIZE int    = 5
	DEFAULT_COLS_SIZE int    = 5
	DEFAULT_FILE_PATH string = ""
	DEFAULT_FACTOR    int    = 5
)

func main() {

	rows := flag.Int("rows", DEFAULT_ROWS_SIZE, "Number of rows.")
	cols := flag.Int("cols", DEFAULT_COLS_SIZE, "Number of columns.")
	factor := flag.Int("factor", DEFAULT_FACTOR, "Factor for initial Matrix generation (0 - 100)")
	filepath := flag.String("filepath", DEFAULT_FILE_PATH, "Absolute file path.")

	flag.Parse()

	gol := concrete.NewGol(*rows, *cols, *factor)

	if *filepath == "" {
		gol.Start()
	} else {
		gol.StartFromFile(*filepath)
	}

	for {
		gol.Print()
		gol.Next()
		ClearScreen()
	}
}
