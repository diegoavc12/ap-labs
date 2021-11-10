#define BUF_LEN (10 * (sizeof(struct inotify_event) + NAME_MAX + 1))
#define _XOPEN_SOURCE 500

#include <ftw.h>
#include <stdio.h>
#include "logger.h"
#include <sys/inotify.h>
#include <stdlib.h>
#include <limits.h>
#include <unistd.h>
#include <string.h>


int inotifyFd, wd;
ssize_t numRead;
char buf[BUF_LEN];
char *p;
struct inotify_event *event;
uint32_t currentCookie;
char *oldName;
char *path;
int directoriesSize=0;
char* subDir;

struct directory{
    int wd;
    char* name;
};

struct directory directories[10];

char* searchSubDir(int wd){
    int i;
    for(i=0; i<directoriesSize; i++){
        if(wd==directories[i].wd){
            return strdup(directories[i].name);
        }
    }
    return NULL;
}
 
int displayInotifyEvent(struct inotify_event *i){
    if(i->mask & IN_CREATE){
        if(i-> mask & IN_ISDIR){
            subDir=searchSubDir(i->wd);
            if(subDir==NULL){
               printf("[Directory - Create] - %s \n", i->name);
               int wd=inotify_add_watch(inotifyFd,strcat(path,i->name),IN_CREATE|IN_DELETE|IN_ISDIR |IN_MOVED_FROM | IN_MOVED_TO);
                if(wd==-1){
                    printf("Error adding watcher to directory \n");
                    exit(1);
                }
                struct directory newdir;
                newdir.name=i->name;
                printf("Added dir %s \n", newdir.name);
                newdir.wd=wd;
                directories[directoriesSize]=newdir;
                directoriesSize++; 
            }else{
                printf("[Directory - Create] - %s \n", strcat(subDir,i->name));
            }
            subDir="";
        }else{
            subDir=searchSubDir(i->wd);
            if(subDir==NULL){
               printf("[File - Create] - %s \n", i->name); 
            }else{
                printf("[File - Create] - %s \n", strcat(subDir,i->name));
            }
            subDir="";
        }
    }else if(i-> mask & IN_DELETE){
        if(i-> mask & IN_ISDIR){
            subDir=searchSubDir(i->wd);
            if(subDir==NULL){
               printf("[Directory - Removal] - %s \n", i->name); 
            }else{
                printf("[Directory - Removal] - %s \n", strcat(subDir,i->name));
            }
            subDir="";
        }else{
            subDir=searchSubDir(i->wd);
            if(subDir==NULL){
               printf("[File - Removal] - %s \n", i->name); 
            }else{
                printf("[File - Removal] - %s \n", strcat(subDir,i->name));
            }
            subDir="";
        }
    }else if(i-> mask & IN_MOVED_FROM){
        currentCookie=i->cookie;
        oldName=i->name;
    }else if(i->mask & IN_MOVED_TO){
        if(currentCookie==i->cookie){
            if(i-> mask & IN_ISDIR){
                subDir=searchSubDir(i->wd);
                char* subDircpy1=strdup(subDir);
                char* subDircpy2=strdup(subDir);
                if(subDir==NULL){
                    printf("[Directory - Rename] - %s -> %s \n",oldName, i->name); 
                }else{
                    printf("[Directory - Rename] - %s -> %s \n",strcat(subDircpy1,oldName), strcat(subDircpy2,i->name));
                }
                subDir="";
            }else{
                subDir=strdup(searchSubDir(i->wd));
                char* subDircpy1=strdup(subDir);
                char* subDircpy2=strdup(subDir);
                if(subDir==NULL){
                    printf("[File - Rename] - %s -> %s \n",oldName, i->name); 
                }else{
                    printf("[File- Rename] - %s -> %s \n",strcat(subDircpy1,oldName), strcat(subDircpy2,i->name));
                }
                subDir="";
            }
        }
    }
    return 0;
}

int addSubDirs(const char *fpath, const struct stat *sb, int tflag, struct FTW *ftwbuf){
    if(tflag==5){
        int wd=inotify_add_watch(inotifyFd,fpath,IN_CREATE|IN_DELETE|IN_ISDIR |IN_MOVED_FROM | IN_MOVED_TO);
        if(wd==-1){
            printf("Error adding watcher to directory \n");
            exit(1);
        }
        char* name=strdup(fpath);
        struct directory newdir;
        newdir.name=strcat(name,"/");
        printf("Added dir %s \n", newdir.name);
        newdir.wd=wd;
        directories[directoriesSize]=newdir;
        directoriesSize++;
    }
    

    return 0;
}
int main(int argc, char** argv){
    printf("Starting File/Directory Monitor on %s \n", argv[1]);
    path=strcat(argv[1], "/");
    inotifyFd=inotify_init();

    if(inotifyFd==-1){
        printf("Error initializing the API \n");
        exit(1);
    }
    wd=inotify_add_watch(inotifyFd, argv[1],IN_CREATE|IN_DELETE|IN_ISDIR |IN_MOVED_FROM | IN_MOVED_TO);
    if(wd==-1){
        printf("Error adding watcher to directory \n");
        exit(1);
    }
    if(nftw(argv[1],addSubDirs,inotifyFd,FTW_DEPTH | FTW_PHYS)==-1){
        printf("Error trying to ge the subdirectories \n");
        exit(1);
    }
    for(;;){
        numRead=read(inotifyFd,buf, BUF_LEN);
        if(numRead<=0){
            printf("Error reading events \n");
        }
        for(p=buf; p<buf+numRead;){
            event = (struct inotify_event *) p;
            displayInotifyEvent(event);
            p += sizeof(struct inotify_event) + event->len;
        }
    }
    return 0;
}
