FROM golang:1.23.0

RUN mkdir /sisso-back

WORKDIR /sisso-back

COPY ./ ./

RUN go env -w GO111MODULE=on

RUN go mod download

RUN go build ./cmd/app/main.go

CMD [ "./main" ]