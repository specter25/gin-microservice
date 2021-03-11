FROM golang:latest

LABEL maintainer="Ujjwal <ujjwalcoding012@gmail.com>"

WORKDIR /app

COPY go.mod  .
COPY go.sum  .

RUN go mod download

COPY . .

ENV PORT 5000

RUN go build 

# remove all source files
RUN find . -name "*.go" -type f -delete

EXPOSE ${PORT}

CMD ["./gin-microservice"]