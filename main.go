package main

/*
#include "Safe_stub.h"
#include <stdio.h>
#include <stdlib.h>
#cgo CFLAGS: -I /Users/rustam/.ghcup/ghc/8.10.7/lib/ghc-8.10.7/include/
#cgo LDFLAGS: Safe.o -L/Users/rustam/.ghcup/ghc/8.10.7/lib/ghc-8.10.7/base-4.14.3.0 -L/Users/rustam/.ghcup/ghc/8.10.7/lib/ghc-8.10.7/integer-gmp-1.0.3.0 -L/Users/rustam/.ghcup/ghc/8.10.7/lib/ghc-8.10.7/ghc-prim-0.6.1 -L/Users/rustam/.ghcup/ghc/8.10.7/lib/ghc-8.10.7/rts -lHSbase-4.14.3.0 -lHSinteger-gmp-1.0.3.0 -lHSghc-prim-0.6.1 -lHSrts -lCffi -liconv -lm -ldl

static void* allocArgv(int argc) {
    return malloc(sizeof(char *) * argc);
}

static void printArgs(int argc, char** argv) {
    int i;
    for (i = 0; i < argc; i++) {
        printf("%s\n", argv[i]);
    }
}

// init_hs is a convenience function not to mess around with pointers
// within Go itself.
static void init_hs(int argc, char** argv) {
	hs_init(&argc, &argv);
}
*/
import "C"
import (
	"fmt"
	"os"
	"unsafe"
)

func main() {
	argv := os.Args
	argc := C.int(len(argv))
	c_argv := (*[0xfff]*C.char)(C.allocArgv(argc))
	defer C.free(unsafe.Pointer(c_argv))

	for i, arg := range argv {
		c_argv[i] = C.CString(arg)
		defer C.free(unsafe.Pointer(c_argv[i]))
	}

	// debug output of the args.
	C.printArgs(argc, (**C.char)(unsafe.Pointer(c_argv)))

	// calling Haskell.
	C.init_hs(argc, (**C.char)(unsafe.Pointer(c_argv)))
	defer C.hs_exit()
	n := C.fibonacci_hs(42)
	fmt.Println(n)
}
