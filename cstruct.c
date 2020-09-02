/* 
 * Process these structs that golang can not handle   
 * and return the appropriate type of data.
 */ 
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "./include/cstruct.h"

// Get string in char**, each time return one string.
char *ret_string(char **cstring, int index, int cstr_num) {
    if (index >= cstr_num) {
        return NULL;
    }
    return cstring[index];
}

int get_detection_res(detection *dets, char **metanames, resdetection *resdet, int j, int i){
    int flag = 0;
    char c[20];
    strcpy((*resdet).res,  "");
    if (dets[j].prob[i] > 0) {
        flag = 1;
        box b = dets[j].bbox;
        strcat((*resdet).res, "[");
        strcat((*resdet).res, metanames[i]);
        strcat((*resdet).res, ",");
        sprintf(c, "%0.2f", (dets[j].prob[i]));
        strcat((*resdet).res, c);
        strcat((*resdet).res, ",(");
        sprintf(c, "%0.10f", (b.x));
        strcat((*resdet).res, c);
        strcat((*resdet).res, ",");
        sprintf(c, "%0.10f", (b.y));
        strcat((*resdet).res, c);
        strcat((*resdet).res, ",");
        sprintf(c, "%0.10f", (b.w));
        strcat((*resdet).res, c);
        strcat((*resdet).res, ",");
        sprintf(c, "%0.10f", (b.h));
        strcat((*resdet).res, c);
        strcat((*resdet).res, ")]");
    } 
    return flag;
}
