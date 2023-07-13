#include <stdio.h>

#define EXPORT __attribute__((visibility("default")))

EXPORT
int main(void) 
{
  return 0;
}

EXPORT
int sum(int x, int y) 
{
  return x + y;
}
