package core

/*
#include <errno.h>
#include <fcntl.h>
#include <linux/input.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/ioctl.h>
#include <unistd.h>
#include <sys/file.h>
ssize_t ret;
struct input_event event;
int dev_open(char *name) {
    return open(name, O_RDWR);
}
void clenEvent() {
	memset(&event, 0, sizeof(event));
}
void writeEvent(int device) {
	ret = write(device, &event, sizeof(event));
if(ret < (ssize_t) sizeof(event)) {
        fprintf(stderr, "write event failed, %s\n", strerror(errno));
    }
}
*/
import "C"

var device C.int
var dir = "/dev/input/event0"

func init()  {
	device = C.dev_open(C.CString(dir))
}

func SendEvent(eventArray []Event) {
	for _,value := range eventArray{
		sendEvent(value)
	}
}

func sendEvent(event Event) {
	C.clenEvent()
	C.event._type = C.ushort(event.Type)
	C.event.code = C.ushort(event.Code)
	C.event.value = C.int(event.Value)
	C.writeEvent(device)
}

type Event struct {
	Type  int
	Code  int
	Value int
}
