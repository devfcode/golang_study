#include <stdio.h>
#include "callee.h"

void SayHello() {
    printf("Hello, world!\n");
}

int Foo(int a, int b) {
    printf("a:%d, b:%d\n", a, b);
    return a + b;
}
