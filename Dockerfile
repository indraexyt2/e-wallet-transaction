FROM golang:1.23

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod tidy

COPY . .
COPY .env .

RUN go build -o e-wallet-transaction
RUN chmod +x e-wallet-transaction
EXPOSE 8082
CMD [ "./e-wallet-transaction" ]