package API
 
import (
	"fmt"
)

var waitQueue = make(map[string]string)
var count = 0

func AddtoQueue(vtype string, vno string) {
	if _, ok := waitQueue[vno]; ok {
		fmt.Println("Already in the queue")
		return
	}
	waitQueue[vno] = vtype
	count = count + 1 
	fmt.Println("added to queue")
}

func SearchQueue(vtype string) string {
	var vno string
	for key, value := range waitQueue {
		if value == vtype {
			vno = key
		} else {
			return "none"
		}
	}
	return vno
}
