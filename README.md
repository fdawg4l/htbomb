# htbomb
_For OOM'ing bots ruining the internet._
_(Also [this](https://www.youtube.com/watch?v=dIQ53t0gv_4))_

See [here](https://blog.haschek.at/2017/how-to-defend-your-website-with-zip-bombs.html) for the inspiration behind this.  And the `hackaday` [article](http://hackaday.com/2017/07/08/dropping-zip-bombs-on-vulnerability-scanners/) covering it.


```
$ go build .
$ ./bomb
2017/07/09 14:24:27 listening on :8080
2017/07/09 14:24:41 serving bomb on / -> [::1]:63138
2017/07/09 14:24:41 serving bomb on / -> [::1]:63139
2017/07/09 14:24:42 serving bomb on / -> [::1]:63140
2017/07/09 14:24:43 serving bomb on / -> [::1]:63141

```

Startup will be slow since we build the payload on startup.  The result is `GET`s (well... really any operation) will serve a 10G gzip bomb to the client.
