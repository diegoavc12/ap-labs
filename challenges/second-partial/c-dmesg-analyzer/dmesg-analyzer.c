#include <stdio.h>
#include <fcntl.h>
#include <unistd.h>
#include <string.h>
#include <stdlib.h>

#define REPORT_FILE "report.txt"
#define HASHSIZE 64997

struct nlist{
    struct nlist *next;
    char *name;
    char *defn[2300];
    int size;
};

static struct nlist *hashtab [HASHSIZE];

unsigned hash(char* s){
    unsigned hashval;
    for(hashval=0; *s!= '\0';s++){
        hashval=*s + 31*hashval;
    }
    return hashval%HASHSIZE;
}

struct nlist *lookup(char *s){
    struct nlist *np;
    for(np= hashtab[hash(s)]; np!= NULL; np->next){
        if(strcmp(s,np->name)==0){
            return np;
        }
    }
    return NULL;
}


struct nlist *install(char* name, char*defn){
    struct nlist *np; 
    unsigned hashval;
    if((np=lookup(name))== NULL){
       
        np= (struct nlist *) malloc(sizeof(*np));
        if(np==NULL || (np->name =strdup(name))==NULL){
            return NULL;
        }
        
        hashval=hash(name);
        np->next=hashtab[hashval];
        hashtab[hashval]=np;
        np->defn[0]=strdup(defn);
        np->size=1;
    }else{
        np->defn[np->size]=strdup(defn);
        np->size++; 
    }
    return np;
}



void analizeLog(char *logFile, char *report);

int main(int argc, char **argv) {

    if (argc < 2) {
	printf("Usage:./dmesg-analizer logfile.txt\n");
	return 1;
    }

    analizeLog(argv[1], REPORT_FILE);

    return 0;
}

void analizeLog(char *logFile, char *report) {
    printf("Generating Report from: [%s] log file\n", logFile);

    // Implement your solution here.
    int fd;
    ssize_t n;
    int count=0;
    unsigned char buf[1005]="";
    char c;
    if((fd=open(logFile,O_RDONLY))<0){
        printf("Error: Cannot open file \n");
        exit(1);
    }


    while((n=read(fd,&c,1))>0){
        
        buf[count]=c;
        
        if(buf[count]=='\n'){
            buf[count+1]='\0';
            char* line=buf;
            char* delim1="]";
            char *timestamp=strtok_r(line,delim1,&line);
            char* ptrCategory=strstr(line,": ");
            char* category;
            int posCategory;
            if(ptrCategory==NULL){
                category="General";
                posCategory=0;
            }else{
                posCategory=ptrCategory-line;
                category=(char*)calloc(1,posCategory+10);
                strncpy(category,line,posCategory);
            }
          
            char *commandWithoutCategory=&line[posCategory+2];
            timestamp=strcat(timestamp,"]");
            char* command=strcat(timestamp,commandWithoutCategory);
            install(category,command);
            count=0;
            memset(buf,0,1005);
        }else{
             count++;
        }
       
    }
    close(fd);
    int i;
    int j;
    int outputFile=open(report, O_WRONLY | O_CREAT | O_TRUNC, S_IRUSR | S_IWUSR | S_IRGRP | S_IWGRP | S_IROTH | S_IWOTH);
    for(i=0; i<HASHSIZE; i++){
            if(hashtab[i]!=NULL){
                write(outputFile,hashtab[i]->name, strlen(hashtab[i]->name));
                write(outputFile,"\n",1);
                for(j=0; j<hashtab[i]->size; j++){
                    write(outputFile,hashtab[i]->defn[j], strlen(hashtab[i]->defn[j]));
                    write(outputFile,"\n",1);
                    
                }
            }
            
    }
    close(outputFile);

    printf("Report is generated at: [%s]\n", report);
}
