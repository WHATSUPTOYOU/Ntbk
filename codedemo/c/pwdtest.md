此段代码为根据shadow和passwd文件测试密码
```C
#include <pwd.h>
#include <stddef.h>
#include <string.h>
#include <shadow.h>
#include <stdio.h>
#include <unistd.h>

char * crypt (const char *key,const char * salt);


int main(int argc, char *argv[]){

    if(argc < 2){
        printf("no usrname input");
        return 1;
    }

    if (geteuid() != 0){
        fprintf(stderr, "must be setuid root");
        return -1;
    }

    struct spwd *shd= getspnam(argv[1]);

    if(shd != NULL){
        static char crypt_char[80];
        strcpy(crypt_char, "$5$fc59262b60596e38$OiWrL77sj9JwO5h3vYWcg1GpeGjw0NU9wUYDtYpBft8");
        char salt[13];
        int i=0,j=0;
        while(crypt_char[i]!='\0'){
            salt[i]=crypt_char[i];
            if(salt[i]=='$'){
                j++;
                if(j==3){
                    salt[i+1]='\0';
                    break;
                }
            }

            i++;
        }

        if(j<3) perror("file error or user cannot use.");

        if(argc==3){
            printf("crypt: %s\n", crypt(argv[2], salt));
            printf("res:%d\n",strcmp(crypt(argv[2], salt),crypt_char));
            printf("shadowd passwd: %s\n", crypt_char);

            printf("res:%d",strcmp(crypt(argv[2], salt),crypt_char));
        }

    }

    return 0;

}
```
