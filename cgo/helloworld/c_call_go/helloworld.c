#include <string.h>
#include "helloworld.h"

// go build -buildmode=c-shared -o helloworld.so helloworld.go
// gcc -o helloworld helloworld.c ./helloworld.so
// ./helloworld

// go build -buildmode=c-archive -o helloworld.so helloworld.go
// gcc -o helloworld helloworld.c ./helloworld.so -lpthread
// ./helloworld
int main(int argc, char *argv[]) {
    hello();
    char s[] = "bye bye\n";
    GoString gs = {s, strlen(s)};
    say(gs);
}