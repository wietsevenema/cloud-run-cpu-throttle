package main

/*
#include <time.h>
*/
import "C"

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

var log = logrus.New().WithField("revision", os.Getenv("K_REVISION"))


/*

When you deploy this on Cloud Run managed, you can observe CPU throttling
when the service is not handling a request.

You can reproduce the same locally when you start the docker container and throttle
the CPU. Do note that docker will not remove the throttling when your container
handles a request.

$: docker run --cpus .3 gcr.io/[PROJECT-ID]/cpu-throttle
time="2019-12-18T09:04:58Z" msg=Started CLOCKS_PER_SEC=1000000
time="2019-12-18T09:05:08Z" msg="Tested CPU Throttle" cpu-time=0.29
 */

func main() {
	log.WithField("CLOCKS_PER_SEC", C.CLOCKS_PER_SEC).Println("Started")

	port := os.Getenv("PORT")
	if port == ""{
		port = "8080"
	}

	http.HandleFunc("/", handler)

	go func() {
		log.Fatal(http.ListenAndServe(":"+port, nil))
	}()

	for {
		log.WithField("cpu-time", testCPUThrottle()).Println("Tested CPU Throttle")
	}

}

/*

I am testing CPU throttling using the clock function from time.h:

In typical usage, you call the clock function at the beginning and
end of the interval you want to time, subtract the values, and then
divide by CLOCKS_PER_SEC (the number of clock ticks per second) to get
processor time. The value that clock() returns is the number of clock ticks.
https://www.gnu.org/software/libc/manual/html_node/CPU-Time.html

*/
func testCPUThrottle() float64 {
	var startTicks = C.clock()
	startTime := time.Now()
	var stopTime time.Time
	var wait = 10 * time.Second
	for {
		stopTime = time.Now()
		if stopTime.Sub(startTime) >= wait {
		break
		}
	}
	stopTicks := C.clock()
	runSeconds := stopTime.Sub(startTime).Seconds()
	return float64(stopTicks-startTicks)/C.CLOCKS_PER_SEC/runSeconds
}


/*

When deployed on Cloud Run managed, you can see that the CPU is not throttled
while handling a request. This request sleeps for 30s to make sure there
are samples taken while it is running:

time="2019-12-18T08:55:57Z" msg="Tested CPU Throttle" cpu-time=0.24
time="2019-12-18T08:56:07Z" msg="Tested CPU Throttle" cpu-time=0.22
time="2019-12-18T08:56:08Z" msg="Start Serving Request" requestURI=/
time="2019-12-18T08:56:17Z" msg="Tested CPU Throttle" cpu-time=0.87
time="2019-12-18T08:56:27Z" msg="Tested CPU Throttle" cpu-time=1.00
time="2019-12-18T08:56:37Z" msg="Tested CPU Throttle" cpu-time=1.00
time="2019-12-18T08:56:38Z" msg="Stop Serving Request Request" requestURI=/
time="2019-12-18T08:56:47Z" msg="Tested CPU Throttle" cpu-time=0.32
time="2019-12-18T08:56:57Z" msg="Tested CPU Throttle" cpu-time=0.23
time="2019-12-18T08:57:07Z" msg="Tested CPU Throttle" cpu-time=0.23
time="2019-12-18T08:57:17Z" msg="Tested CPU Throttle" cpu-time=0.16
time="2019-12-18T08:57:27Z" msg="Tested CPU Throttle" cpu-time=0.25
time="2019-12-18T08:57:37Z" msg="Tested CPU Throttle" cpu-time=0.10
time="2019-12-18T08:57:48Z" msg="Tested CPU Throttle" cpu-time=0.22

*/
func handler(w http.ResponseWriter, r *http.Request) {
	requestLog := log.WithFields(logrus.Fields{
		"requestURI": r.RequestURI,
	})
	requestLog.Println("Start Serving Request")
	time.Sleep(30 * time.Second)
	fmt.Fprintln(w, "Hello World!")
	requestLog.Println("Stop Serving Request Request")
}

