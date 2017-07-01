BIN=antlr4
GRAMMERS= sqlite golang json python3 xml graphql

build:
	@for grammer in $(GRAMMERS) ; do \
		(echo $$grammer && cd $$grammer && $(BIN) -Dlanguage=Go *.g4) ;\
	done
