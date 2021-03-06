#include <stdio.h>
#include <string.h>

int main(int argc, char *argv[])
{
    if (argc != 2) {
        printf("You must enter in a string to reverse!\n");
        return 1;
    }
    char * string = argv[1];
    int len = strlen(string);
    char reversed[len];
    int idx;
    for (idx = 0; idx < len; idx++) {
        // Subtract 1 from the idx because the last char is \0 and we
        // do not want that to be in the first position
        // Ensure the algorithm does 5, 4, 3, 2, 1, 0 instead of
        // 6, 5, 4, 3, 2, 1
        reversed[idx] = string[len - idx - 1];
    }
    reversed[len] = '\0';
    printf("%s\n", reversed);
}
