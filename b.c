#include <stdlib.h>
#include <stdio.h>
void trstr(char *c_str, char orig, char rep) {
  while (*c_str != '\0') {
    if (*c_str == orig) {
      *c_str = rep;
    }
    ++c_str;
  }
}
void func(void) {
  char *env = getenv("TEST_ENV");

  trstr(env, 'a', 'A');
}

int main() {
    // 输出 "Hello, World!"
    printf("Hello, World!\n");
    int i,j=2,*p=&i; 
    // i=*p;
    *p=*&j;
    printf("%d\n",i);

    // 程序执行成功，返回0
    return 0;
}