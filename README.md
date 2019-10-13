# httpd

This is a trivial HTTP-server, which can serve static-files.


## Rationale

Most of the time I need to serve static-content I would use Python, but sometimes it is not available.  (Where it is serving is as simple as this `python -m SimpleHTTPServer 3000`.)


## Installation

To install this just run:

    $ go get github.com/skx/httpd


## Usage

Run with no arguments to serve the current directory, add `-help` for other options.

## Limitations

This is sufficient to meet my needs, but it should be noted that this is not a serious project, and it has many obvious limitations:

* This is a trivial toy.
* It will not scale.
* It does not handle malicious clients.
  * (slowlaris attacks, specifically, due to the lack of timeouts.)
* It does not log the HTTP response-code sent back to clients.



Steve
--
