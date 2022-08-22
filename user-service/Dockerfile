FROM golang:1.17-alpine

ENV GO111MODULE on

ENV APP_HOME go/src/user-service
WORKDIR $APP_HOME

COPY . .

ENV HOST postgres
ENV PASS password
ENV USER postgres
ENV PORT 5432
ENV NAME postgres
ENV USER_SERVICE 8090
ENV ORDER_SERVICE_URL http://order-service:8091/

RUN go mod verify
RUN go mod download
RUN go mod tidy
RUN go build
CMD ["./user-service"]