# Proxit


Very simple reverse proxy to talk with HTTP services on a target (local or remote).

More features to come......maybe.


## Build
It strips the binary and also compacts it with upx to make it as small as possible.

``` ./builder.sh [executable name] ```

## Usage
On target, launch it wit ./proxit -exposePort 8000

Make an HTTP request to the service with the following headers

```
GET / HTTP/1.1
Host: host:8000
Path: [Path to request on the local service]
Port: [Port to connect to through the target]
Remote-address: [ Pass a remote address to access it through the target]
Tls: [If the other service is running over Tls, no specific value needed, just add some value]


```

Using curl

```
curl --path-as-is -i -s -k -X $'GET' \
    -H $'Host: host:8000' -H $'Path: /' -H $'Port: 10000' -H $'Remote-address: 10.10.10.10' -H $'Tls: yuup' \
    $'http://target/'
```

## Warning
This is for educational purposes only, the writer of this tool is not responsible for any malicious use of it.