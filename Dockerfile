# https://hub.docker.com/_/golang
FROM golang:1.14.2-buster

# Install git to clone.
RUN apt-get update && apt-get install git-core --yes

# Let us work somewhere
RUN mkdir /app

# Create a stub HTML file
RUN mkdir /docs
RUN echo 'Hello skx/httpd' > /docs/index.html

# Clone
RUN cd /app && git clone https://github.com/skx/httpd.git

# Build
RUN cd /app/httpd && go build .

# Expose the port
EXPOSE 3000

# Run
CMD  /app/httpd/httpd -host 0.0.0.0 -path /docs
