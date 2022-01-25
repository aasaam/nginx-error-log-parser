FROM golang:1.17-buster AS builder

ADD . /src

RUN cd /src \
  && go get -u -v golang.org/x/lint/golint \
  && go mod tidy \
  && go get -u -v \
  && go mod download \
  && golint . \
  && export CI=1 \
  && go test -covermode=count -coverprofile=coverage.out \
  && cat coverage.out | grep -v "main.go" > coverage.txt \
  && TOTAL_COVERAGE_FOR_CI_F=$(go tool cover -func coverage.txt | grep total | grep -Eo '[0-9]+.[0-9]+') \
  && echo "TOTAL_COVERAGE_FOR_CI_F: $TOTAL_COVERAGE_FOR_CI_F" \
  && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o nginx-error-log-parser


FROM scratch

COPY --from=builder /src/nginx-error-log-parser /usr/bin/nginx-error-log-parser

ENTRYPOINT ["/usr/bin/nginx-error-log-parser"]
