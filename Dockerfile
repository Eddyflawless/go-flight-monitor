FROM golang:1.15 as build

WORKDIR /app
COPY ./src/* /app/
RUN go build -o main

FROM alpine as runtime
COPY --from=build /app/main /
CMD ./main