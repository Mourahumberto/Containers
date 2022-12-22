FROM golang:1.16-alpine
WORKDIR /app
ENV ELASTIC_APM_SERVER_URL=http://apm-server:8200
ENV ELASTIC_APM_SERVICE_NAME=GolangApp
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN go build http.go
EXPOSE 5000
CMD [ "./http" ]