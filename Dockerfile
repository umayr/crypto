FROM golang:1.7

ENV PROJECT_NAME crypto
ENV PROJECT_PATH github.com/umayr/$PROJECT_NAME

RUN mkdir -p $GOPATH/src/$PROJECT_PATH
ADD . $GOPATH/src/$PROJECT_PATH

WORKDIR $GOPATH/src/$PROJECT_PATH

RUN apt-get update
RUN apt-get install -y libssl-dev openssl
RUN chmod +x ./build/go.sh

CMD ["./build/go.sh", "run", "./examples/aes.go"]