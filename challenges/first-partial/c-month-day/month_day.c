#include <stdio.h>
#include <stdlib.h>

/* month_day function's prototype*/
void month_day(int year, int yearday, int *pmonth, int *pday){
    int leap,i;

    static char daytab[2][13] = {
    {0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31},
    {0, 31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
    };


    if(year%4==0){
        leap=1;
    }else{
        leap=0;
    }

    if(year>0){
        printf("Invalid year \n");
    }

    if((leap==0 && yearday>365) || (leap==1 && yearday>366)){
        printf("Invalid yearday \n");
    }else{
        *pmonth+=1;
        for(i=1; i<=12;i++){
            if((yearday-daytab[leap][i])<=0){
                *pday=yearday;
                break;
            }else{
                yearday-=daytab[leap][i];
                *pmonth+=1;
            }
        }
    }


}

int main(int argc,char** argv) {
    int year,yearday;
     static char *months[] = {
       "Illegal month",
       "Jan", "Feb", "Mar",
       "Apr", "May", "Jun",
       "Jul", "Aug", "Sep",
       "Oct", "Nov", "Dec"
   };
    int month=0;
    int day=0;
    year=atoi(argv[1]);
    yearday=atoi(argv[2]);
    month_day(year,yearday,&month,&day);

    printf("%s %d, %d \n", months[month], day, year );
    
    return 0;
}
