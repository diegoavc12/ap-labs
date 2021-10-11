//#include <stdio.h>
void quicksort(void *v[], int left, int right,
	   int (*comp)(void *, void *))
{
    
    int k, last;
    void swap(void *v[], int, int);

    if (left >= right)
	return;
    swap(v, left, (left + right)/2);
    last = left;
    for (k = left+1; k <= right; k++)
	if ((*comp)(v[k], v[left]) < 0)
	    swap(v, ++last, k);
    swap(v, left, last);
    quicksort(v, left, last-1, comp);
    quicksort(v, last+1, right, comp);
}

