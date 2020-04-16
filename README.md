# httpd

This is a trivial HTTP-server, which can serve static-files.

You'd probably prefer to look at my toolbox which contains a `httpd` server:

* [https://github.com/skx/sysbox](https://github.com/skx/sysbox)


## Installation

To install this from source please run:

    $ go get github.com/skx/httpd


## Docker

There is a Dockerfile included within the repository, you can build the image via:

```
$ docker build -t skx/httpd .
```

To launch it run:

```
$ docker run -d -p3000:3000 skx/httpd
```

## Volumes

You can serve a local directory by bind-mounting your root directory to `/app`,
like so:

```
$ docker run -d -v /var/www/html:/app -p3000:3000 skx/httpd
```
