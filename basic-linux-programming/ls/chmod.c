#include <stdio.h>
#include <stdlib.h>
#include <sys/stat.h>


int
main(int argc, char *argv[])
{
    int i;
    int mode;

    if (argc < 2) {
        fprintf(stderr, "%s: no arguments\n", argv[0]);
        exit(1);
    }
    mode = strtol(argv[1], NULL, 0);
    for (i =2; i < argc; i++) {
        if (chmod(argv[i], mode) < 0) {
            perror(argv[i]);
        }
    }
    exit(0);
}

