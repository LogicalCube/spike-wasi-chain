#include <stdio.h>

#define EXPORT __attribute__((visibility("default")))

EXPORT
int main(void) 
{
  // puts("hi there");
  return 0;
}

EXPORT
int sum(int x, int y) 
{
  // printf("%i\n", x + y);
  return x + y;
}

// _Noreturn void __wasi_proc_exit(int)
// {
//     __builtin_trap();
// }