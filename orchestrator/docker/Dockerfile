FROM golang:1.22.0 as builder

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /zipcode cmd/main.go


FROM scratch
COPY --from=builder /zipcode /zipcode

EXPOSE 8080
ENTRYPOINT [ "/zipcode" ]