FROM golang:1.17-alpine

ENV GO111MODULE auto

ENV APP_HOME go/src/order-service
WORKDIR $APP_HOME

COPY . .

ENV HOST postgres
ENV PASS password
ENV USER postgres
ENV PORT 5432
ENV NAME postgres
ENV ORDER_SERVICE 8091

RUN go mod verify
RUN go mod download
RUN go mod tidy
RUN go build
CMD ["./order-service"]