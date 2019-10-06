FROM golang

WORKDIR /go/src/github.com/veotani/data-mining-processing-course/parsers
COPY . .

RUN apt-get update \
 && apt-get -y install \
    wv poppler-utils
RUN go get -d -v ./...
RUN go install -v ./...
RUN go test -coverprofile=coverage.txt ./...
