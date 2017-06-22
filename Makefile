BIN=antlr4
build:
	cd sqlite && $(BIN) -Dlanguage=Go SQLite.g4
	cd golang && $(BIN) -Dlanguage=Go Golang.g4
