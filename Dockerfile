FROM golang:alpine

WORKDIR /usr/src/app

ENV ENVIROMENT=dev
ENV PORT=8080

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o ./app.exe

EXPOSE 8080

CMD ["./app.exe"]