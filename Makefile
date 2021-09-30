haskelgo: Safe.o main.go
	go build

Safe.o: Safe.hs
	ghc -c -O $<
