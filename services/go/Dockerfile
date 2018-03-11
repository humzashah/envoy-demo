FROM golang:1.10-alpine
RUN mkdir /app
COPY service.go /app/service.go
RUN go build -o /app/service /app/service.go
ENV THIS_SERVICE_NAME="ServiceOne"
ENV THIS_SERVICE_ADDRESS="localhost:3001"
CMD /app/service
