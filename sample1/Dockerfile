FROM berliozcloud/golang-1.10

ADD . /go/src/app
RUN dep ensure
ADD vendor/connector-go.git /go/src/connector-go.git

RUN go build -o /app/main .
