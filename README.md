# Calling Haskell From Go Example
> I did it for the lulz.

<a rel="license" href="http://creativecommons.org/licenses/by-sa/4.0/"><img alt="Creative Commons License" style="border-width:0" src="https://i.creativecommons.org/l/by-sa/4.0/88x31.png" /></a><br />This work is licensed under a <a rel="license" href="http://creativecommons.org/licenses/by-sa/4.0/">Creative Commons Attribution-ShareAlike 4.0 International License</a>.

## Building

#### Updating the source paths
Make sure to update the libraries path in `main.go` on the line
```
// #cgo LDFLAGS: Safe.o -L <...> -lHSbase-4.14.3.0 -lHSinteger-gmp-1.0.3.0 -lHSghc-prim-0.6.1 -lHSrts -lCffi -liconv -lm -ldl
```

`-l` parameters should remain - they specify the required Haskell libraries
for linking.  You may need to update the libraries versions.

#### Easy (Makefile)
Run:
```
make
```
and it should do the job.

#### Manual build
```
ghc -c -O Safe.hs
go build
```

## If something goes wrong
Most likely some libraries are missing or not specified in `main.go`.

To verify which libraries and paths are in use on your system, follow the steps below.

Paste this source to a file named `test.c`
```c
#include <HsFFI.h>
#ifdef __GLASGOW_HASKELL__
#include "Safe_stub.h"
#endif
#include <stdio.h>

int main(int argc, char *argv[])
{
    int i;
    hs_init(&argc, &argv);

    i = fibonacci_hs(42);
    printf("Fibonacci: %d\n", i);

    hs_exit();
    return 0;
}
```

Run:

```sh
ghc -c -O Safe.hs
ghc --make -no-hs-main -v -optc-O test.c Safe -o test
```
This will generate the compile and link commands for your version of C compiler.
Check the output and use it as the guidance, and let the sun enlighten your treacherous path.

## References

Based on the following materials:
* [Calling Haskell From C](https://wiki.haskell.org/Calling_Haskell_from_C)
* [Go: passing argv to C function](https://stackoverflow.com/questions/37657326/go-passing-argv-to-c-function)
