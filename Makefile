BIN=antlr4
build:
	cd sqlite && $(BIN) -Dlanguage=Go SQLite.g4
	cd golang && $(BIN) -Dlanguage=Go Golang.g4
	cd json && $(BIN) -Dlanguage=Go JSON.g4
	cd python3 && $(BIN) -Dlanguage=Go Python3.g4
