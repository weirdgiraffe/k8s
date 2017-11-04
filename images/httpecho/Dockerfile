FROM golang:1.9-alpine
LABEL version=0.1
COPY httpecho /go/src/httpecho
RUN  cd /go/src/httpecho/ && go install
EXPOSE 8080
CMD httpecho
