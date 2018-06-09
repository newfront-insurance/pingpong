Pingping
===========

This is a simple image that just gives a response on port 8080. Use this to
test your web orchestration.

Sample Usage
------------

### Build image

```bash
$ docker build -t tsongpon/pingpong .
```

### Starting a web server on port 80

```bash
$ docker run -d --name web-test -p 80:8080 tsongpon/pingpong
```

You can now interact with this as if it were a dumb web server:
```
$ curl localhost/ping
pong
...