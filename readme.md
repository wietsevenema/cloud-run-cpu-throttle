# Cloud Run CPU throttling
When deployed on Cloud Run managed, you can see that the CPU is not throttled
while handling a request, but throttled when not handling a request: 

```
time="2019-12-18T08:55:57Z" msg="Tested CPU Throttle" cpu-time=0.24
time="2019-12-18T08:56:07Z" msg="Tested CPU Throttle" cpu-time=0.22
--> time="2019-12-18T08:56:08Z" msg="Start Serving Request" requestURI=/
time="2019-12-18T08:56:17Z" msg="Tested CPU Throttle" cpu-time=0.87
time="2019-12-18T08:56:27Z" msg="Tested CPU Throttle" cpu-time=1.00
time="2019-12-18T08:56:37Z" msg="Tested CPU Throttle" cpu-time=1.00
--> time="2019-12-18T08:56:38Z" msg="Stop Serving Request Request" requestURI=/
time="2019-12-18T08:56:47Z" msg="Tested CPU Throttle" cpu-time=0.32
time="2019-12-18T08:56:57Z" msg="Tested CPU Throttle" cpu-time=0.23
time="2019-12-18T08:57:07Z" msg="Tested CPU Throttle" cpu-time=0.23
time="2019-12-18T08:57:17Z" msg="Tested CPU Throttle" cpu-time=0.16
time="2019-12-18T08:57:27Z" msg="Tested CPU Throttle" cpu-time=0.25
time="2019-12-18T08:57:37Z" msg="Tested CPU Throttle" cpu-time=0.10
time="2019-12-18T08:57:48Z" msg="Tested CPU Throttle" cpu-time=0.22
```

[![Run on Google Cloud](https://deploy.cloud.run/button.svg)](https://deploy.cloud.run)