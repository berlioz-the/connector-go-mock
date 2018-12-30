FROM berliozcloud/golang-1.10

WORKDIR /go/src/app
COPY Gopkg.toml Gopkg.lock /go/src/app/
RUN dep ensure -v --vendor-only
COPY . /go/src/app/

ADD vendor/connector-go.git /go/src/app/vendor/connector-go.git
RUN ls -la /go/src/app/vendor/connector-go.git
RUN ls -la /go/src/app/vendor/connector-go.git/vendor
RUN cp -rf /go/src/app/vendor/connector-go.git/vendor/* /go/src/app/vendor/
RUN rm -rf /go/src/app/vendor/connector-go.git/vendor/*

RUN go build -o /app/main .
