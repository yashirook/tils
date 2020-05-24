#include <stdio.h>
#include <stdlib.h>

static void word_count(FILE *f);

int
main(int argc, char *argv[])
{
    int i;

    if (argc == 1) {
        word_count(stdin);
    }
    else {
        for (i = 1; i < argc; i++) {
            FILE *f;

            f = fopen(argv[i], "r");
            if (!f) {
                perror(argv[i]);
                exit(1);
            }

            word_count(f);
            fclose(f);
        }
    }
    exit(0);
}

static void
word_count(FILE *f)
{
    unsigned long n;
    int c;
    int prev = '\n';

    n = 0;

    while ((c = fgetc(f)) != EOF) {
        if (c == '\n') {
            n++;
        }
        prev = c;
    }
    if (prev != '\n') {
        n++;
    }
    printf("%lu\n", n);
}
