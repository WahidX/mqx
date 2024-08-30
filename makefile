include .env
export

gin:
	@pp=$(shell python3 -c "print (${PORT} + 1)");  \
	gin --build cmd -i --bin build/server --port $${pp} --appPort ${PORT}
