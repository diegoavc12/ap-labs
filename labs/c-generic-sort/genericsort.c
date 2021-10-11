#include <stdio.h>
#include <stdlib.h>
#include <string.h>


void quicksort(void *lineptr[], int left, int right,
	   int (*comp)(void *, void *));

void mergesort(void *lineptr[], int left, int right,
	   int (*comp)(void *, void *));

int numcmp(char* s1, char* s2);

int numcmp(char* s1, char*s2){
	int v1, v2;
	v1=atoi(s1);
	v2=atoi(s2);

	if(v1<v2){
		return -1;
	}
	else if(v1>v2){
		return 1;
	}else{
		return 0;
	}

}

int main(int argc, char *argv[])
{
	

	if(argc<5){
		printf("Error: Missing parameters \n");
		exit(1);
	}
	FILE *inputFile;
	FILE *outputFile;
	int numeric;
	char *arr[100];
	size_t len = 0;
	if(strcmp("-n", argv[1])==0){
		inputFile=fopen(argv[2],"r");
		numeric=1; 
	}else{
		inputFile=fopen(argv[1],"r");
		numeric=0;
	}

	if(numeric && argc<6){
		printf("Error: Missing parameters \n");
		exit(1);
	}

	if(inputFile==NULL){
		printf("Error, cannot open file \n");
		exit(1);
	}else{
		int i=0;
		for(i=0; i<100; i++){
			arr[i]=NULL;
		}
		i=0;
		while(getline(&arr[i],&len,inputFile)!=-1){
			//printf("%s\n ", arr[i]);
			//arr[i][strlen(arr[i])-1]='\0';
			i++;
			
		}
		/*for(i=0; i<18; i++){
			printf("%s\n ", arr[i]);
		}*/
		fclose(inputFile);

		if(numeric==1){
			if(strcmp(argv[3], "-quicksort")==0){
				quicksort((void **)arr, 0, i-1,(int (*)(void* ,void*)) numcmp);
			}
			else if(strcmp(argv[3], "-mergesort")==0){
				mergesort((void **)arr,0,i-1,(int (*)(void* ,void*))numcmp);
			}else{
				printf("Error: invalid sorting algorithm parameter \n");
				exit(1);
			}
		}else{
			if(strcmp(argv[2], "-quicksort")==0){
				quicksort((void **)arr, 0, i-1,(int (*)(void* ,void*))strcmp);
			}
			else if(strcmp(argv[2], "-mergesort")==0){
				mergesort((void **)arr,0,i-1,(int (*)(void*, void*))strcmp);
			}else{
				printf("Error: invalid sorting algorithm parameter \n");
				exit(1);
			}
		}
		if(numeric==1){
			outputFile=fopen(argv[5],"w");
		}else{
			outputFile=fopen(argv[4], "w");
		}
		int j;
		for(j=0; j<i; j++){
			//printf("%s\n ", arr[j]);
			fputs(arr[j],outputFile);
		}
		fclose(outputFile);
		if(numeric==1){
			printf("Results file can be found at ./%s\n",argv[5]);
		}else{
			printf("Results file can be found at ./%s\n",argv[4]);
		}
		
	}

    return 0;

}
