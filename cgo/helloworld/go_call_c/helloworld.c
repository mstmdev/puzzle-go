#include <stdio.h>

void hello() {
    printf("from c: hello world\n");
}

void say(const char *s) {
    printf("from c: %s", s);
}