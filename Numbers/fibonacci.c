#include <stdio.h>
#include <stdlib.h>


int main(int argc, char * argv[])
{
    if ( argc != 2 )
    {
        printf("You must input the Nth term in the fibonacci series you want to see!");
        return 1;
    }
    int n_term = atoi(argv[1]);
    if ( n_term < 2 ) 
    {
        printf("The first two terms are 0, 1. If you want to view more input a higher number");
    }
    long long a = 0;
    long long b = 1;
    long long next;
    int i;
    printf("0, 1");
    for (i=2; i <= n_term; i++)
    {
        next = a + b;
        a = b;
        b = next;
        printf(", %llu", b);
    }
    printf("\n");
    return 0;
}
        
