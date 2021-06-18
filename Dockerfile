FROM scratch

COPY go-project /usr/bin/go-project

ENTRYPOINT ["/usr/bin/go-project"]