#include "darknet.h"

struct resdetection;
typedef struct resdetection resdetection;

// typedef struct resdetection {
//     char *tagname;
//     char *prob;
//     char *x;
//     char *y;
//     char *w;
//     char *h;
//     int flag;
// } resdetection;

// typedef char *res;

typedef struct resdetection {
    char *res;
} resdetection;

char *ret_string(char **names, int index, int names_len);
int get_detection_res(detection *dets, char **metanames, resdetection *resdet, int j, int i);
