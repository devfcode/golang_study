cc=clang

dylib:
	@$(cc) -fPIC -shared callee.c -o libcallee.so

program: dylib
	@go build caller.go

all: program
	@./caller

clean:
	@-rm -rf caller *.so