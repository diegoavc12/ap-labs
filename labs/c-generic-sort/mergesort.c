//#include <stdio.h>

void swap(void *v[], int p, int q)
{
	//printf("Swapping %d and %d", p,q);
   
    void *temp;
    temp = v[p];
    v[p] = v[q];
    v[q] = temp;
}


void merge(void *lineptr[], int left, int right,  int (*comp)(void *, void *)) {
	//printf("Merge\n");

	int m,n;
	for(m=left; m<right; m++){
		for(n=left;n<right-m;n++){
			//printf("Comparando\n");
			if((*comp)(lineptr[n],lineptr[n+1])>0){
				//printf("Swap\n");
				swap(lineptr,n, n+1);
			}
		}
	}
}



void mergesort(void *lineptr[], int left, int right,  int (*comp)(void *, void *)) {
	if(left<right){
		//printf("Mergesort\n");
		int mid=(left+right)/2;
		mergesort(lineptr,left, mid,comp);
		mergesort(lineptr,mid+1,right,comp);
		merge(lineptr,left,right,comp);
	}

}


