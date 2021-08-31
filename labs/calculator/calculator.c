#include <stdio.h>
#include <stdlib.h>
#include <string.h>

// adds/subtracts/multiplies all values that are in the *values array.
// nValues is the number of values you're reading from the array
// operator will indicate if it's an addition (1), subtraction (2) or
// multiplication (3)
long calc(int operator, int nValues, int *values) {
    long result;
    int j;
    result=values[0];
    if(operator==1){
        for(j=1;j<nValues;j++){
            //printf("%d \n",values[j]);
            result+=values[j];
        }
        return result;
    }else if(operator==2){
        for(j=1;j<nValues;j++){
            result-=values[j];
        }
        return result;
    }else if(operator==3){
        for(j=1;j<nValues;j++){
            result*=values[j];
        }
        return result;
    }else{
        printf("Invalid input :( \n");
        return -1;
    }
}

int main(int argc, char** argv) {
    int operator,nValues,i;
    if(strcmp(argv[1],"add")==0){
        operator=1;
    }else if(strcmp(argv[1],"sub")==0){
        operator=2;

    }else if(strcmp(argv[1],"mult")==0){
        operator=3;
        
    }else{
         printf("Invalid operator input :( \n");
        return -1;
    }
    
    nValues=argc-2;
    int valuesInput[nValues];
    //valuesInput=malloc(nValues*sizeof(int));

    for(i=0; i<nValues;i++){
        valuesInput[i]=atoi(argv[i+2]);
        //printf("%d \n",valuesInput[i]);
    }

    printf("%ld \n",calc(operator,nValues,&valuesInput[0]));


    return 0;
}
