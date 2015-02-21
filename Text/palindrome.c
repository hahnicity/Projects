#include <stdio.h>
#include <string.h>


int main( int argc, char * argv[] )
{
    if ( argc != 2 ) {
        printf("You must input a string to check if it's a palindrome\n");
        return 1;
    }
    char * string = argv[1];
    int len = strlen(string);
    int idx;
    for ( idx = 0; idx < len / 2; idx++ ) {
        if ( string[idx] != string[len - idx - 1]) {
            printf("%s is not a palindrome!\n", string);
            return 0;
        }
    }
    printf("%s is a palindrome!\n", string);
    return 0;
}
