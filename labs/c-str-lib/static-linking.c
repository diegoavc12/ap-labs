#include <stdio.h>
#include <string.h>
#include <stdlib.h>
int main(int argc, char** argv) {

    if(argv[1]==NULL || argv[2]==NULL || argv[3]==NULL){
        printf("Error, missing parameters\n");
        exit(1);
    }

    if(strcmp("-add",argv[1])==0){
        mystradd(argv[2],argv[3]);
    }else if(strcmp("-find",argv[1])==0){
        mystrfind(argv[2], argv[3]);
    }else{
        printf("Error, invalid parameters\n");
    }
}
