package main

/*
#cgo CFLAGS: -I./
#cgo LDFLAGS: -L./ -ldarknet -lcstruct -Wl,-rpath=./
#include "./include/darknet.h"
#include "./include/cstruct.h"
*/
import "C"
import "fmt"
import "unsafe"
// import "log"
// import "strings"

// Convert char** to []string.
func cstringtogo(names **C.char, classes int) []string {
	out := make([]string, classes)
	for i := 0; i < classes; i++ {
		n := C.ret_string(names, C.int(i), C.int(classes))
		out[i] = C.GoString(n)
	}
	return out
}

// Detect specified image
func detect(net *C.network, metadata C.metadata, metaclasses int, 
	metanames []string, imagePath string, thresh float32, hier_thresh float32) []string {
	
	letter_box := C.int(0)
	img := C.load_image_color(C.CString(imagePath), 0, 0)
	C.network_predict_image(net, img)
	num := C.int(0)
	dets := C.get_network_boxes(net, img.w, img.h, C.float(thresh), C.float(hier_thresh), (*C.int)(nil), 0, &num, letter_box)
	C.do_nms_sort(dets, num, C.int(metaclasses), C.float(0.45))
	// Store the final result.
	var ret []string
	resptr := C.malloc(C.sizeof_char * 100)
	defer C.free(unsafe.Pointer(resptr))
	// Get one result a time form dets, store it in res.
	resdet := C.struct_resdetection {
		res: (*C.char)(resptr),
	}
	for j := 0; j< int(num); j++ {
		for i := 0; i < metaclasses; i++ {
			flag := C.get_detection_res(dets, metadata.names, &resdet, C.int(j), C.int(i))
			if C.int(flag) == 0 {
				continue
			} else {
				ret = append(ret, C.GoString(resdet.res))
			}
		}
	}
	C.free_detections(dets, num)
	return ret
}

func main() {
	var net *C.network
	//Load yolo network
	net = C.load_network_custom(C.CString("./cfg/yolov4.cfg"), C.CString("./yolov4.weights"), 0, 1)
	metadata := C.get_metadata(C.CString("./cfg/coco.data"))
	metaclasses := int(metadata.classes)
	//get metadata names in golang string
	metanames := cstringtogo(metadata.names, metaclasses)
	res := detect(net, metadata, metaclasses, metanames, "./data/dog.jpg", 0.5, 0.5)
	if res == nil {
		fmt.Println("We got nothing.")
	} else {
		fmt.Println(res)
	}
	res = detect(net, metadata, metaclasses, metanames, "./data/eagle.jpg", 0.5, 0.5)
	if res == nil {
		fmt.Println("We got nothing.")
	} else {
		fmt.Println(res)
	}
	res = detect(net, metadata, metaclasses, metanames, "./data/horses.jpg", 0.5, 0.5)
	if res == nil {
		fmt.Println("We got nothing.")
	} else {
		fmt.Println(res)
	}
	res = detect(net, metadata, metaclasses, metanames, "./data/person.jpg", 0.5, 0.5)
	if res == nil {
		fmt.Println("We got nothing.")
	} else {
		fmt.Println(res)
	}
	res = detect(net, metadata, metaclasses, metanames, "./data/giraffe.jpg", 0.5, 0.5)
	if res == nil {
		fmt.Println("We got nothing.")
	} else {
		fmt.Println(res)
	}
	res = detect(net, metadata, metaclasses, metanames, "./data/scream.jpg", 0.5, 0.5)
	if res == nil {
		fmt.Println("We got nothing.")
	} else {
		fmt.Println(res)
	}
}