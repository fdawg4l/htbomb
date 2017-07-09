# htbomb
_For OOM'ing bots ruining the internet_
_(Also [this](https://www.youtube.com/watch?v=dIQ53t0gv_4))_


```
$ go build .
$ ./bomb
2017/07/09 14:24:27 listening on :8080
2017/07/09 14:24:41 serving bomb on / -> [::1]:63138
2017/07/09 14:24:41 serving bomb on / -> [::1]:63139
2017/07/09 14:24:42 serving bomb on / -> [::1]:63140
2017/07/09 14:24:43 serving bomb on / -> [::1]:63141

```

Startup will be slow since we build the payload on startup.  The result is `GET`s will serve a 10G gzip bomb to the client.
