FROM golang:alpine

COPY . /app
WORKDIR /app
RUN go build .

CMD [ "./node" ]
