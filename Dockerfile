FROM golang:1.18.3-alpine3.16
RUN mkdir -p /see-esvi
WORKDIR /see-esvi
ADD . /see-esvi
RUN go build .
CMD ["/bin/bash"]