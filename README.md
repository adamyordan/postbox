Postbox
===
> A standalone cli-based http request dumper written in go

![](demo.gif)

Installation
---
```
go get -u github.com/adamyordan/postbox
```

Usage
---
- Run a server to listen to incoming http requests. By default, the server will be listening on port 8000.
    ```
    $ postbox server up --daemon
    ```
    
- Try sending request to port 8000
    ```
    $ curl -X PUT -H "Custom-Header: header-value" --data "this is http body data" http://127.0.0.1:8000/some-path
    ```

- List http request received on port 8000
    ```
    $ postbox letter list
    
    [1] 2018-11-22 18:33:18 +0800 +08 (127.0.0.1:53311)
    ```
    
- View details of http request received
    ```
    $ postbox letter view 1

    id    : 1
    ipaddr: 127.0.0.1:53311
    time  : 2018-11-22 18:33:18 +0800 +08
    
    PUT /some-path HTTP/1.1
    Host: 127.0.0.1:8000
    Accept: */*
    Content-Length: 22
    Content-Type: application/x-www-form-urlencoded
    Custom-Header: header-value
    User-Agent: curl/7.54.0
    
    this is http body data
    ```
