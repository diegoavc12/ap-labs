#include <stdio.h>
#include <stdlib.h>

int mystrlen(char *str){
    int len=0;
    while(str[len]!='\0'){
        len+=1;
    }
    return len;
}

char *mystradd(char *origin, char *addition){

    int i;
    char* result=NULL;
    int lenOrigin=mystrlen(origin);
    printf("Initial Length: %d \n",lenOrigin);

    int lenAddition=mystrlen(addition);
    result=(char*)malloc(lenOrigin+lenAddition);
    for(i=0;i<lenOrigin;i++){
        result[i]=origin[i];
    }
    for(i=lenOrigin;i<(lenOrigin+lenAddition);i++){
        result[i]=addition[i-lenOrigin];
    }
    printf("New String: %s \n",result);
    printf("New Length: %d \n",mystrlen(result));
    return result;
}


int mystrfind(char *origin, char *substr){
    int i;
    int count=0;
    int indexFound=0;
    for(i=0; i<mystrlen(origin);i++){
        if(origin[i]==substr[count]){
            count+=1;
        }else{
            count=0;
            indexFound=i;
        }
        if(count==mystrlen(substr)){
            printf("[%s] was found at [%d] position \n", substr,indexFound);
            return 1;
        }
    }
    printf("[%s] was not found \n",substr);
    return 0;
    
}

