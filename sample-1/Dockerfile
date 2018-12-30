FROM berliozcloud/golang-1.10

WORKDIR /go/src/app
COPY Gopkg.toml Gopkg.lock /go/src/app/
RUN dep ensure -v --vendor-only
COPY . /go/src/app/

RUN go build -o /app/main .
