#include <stdio.h>

int main() {
     char *p = "hello"; 
 
 *(p+1) = 'w';

    // 程序执行成功，返回0
    return 0;
}