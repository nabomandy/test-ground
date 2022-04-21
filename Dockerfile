FROM parjom/golang-nodejs:latest
WORKDIR /go/app
COPY . .
# RUN go mod vendor -v
RUN apidoc -i . -o ./docs
RUN GOOS=linux go build -mod=vendor -o dms-demo-service .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /usr/app

ENV TZ=Asia/Seoul
ENV SERVICE_PORT=8080
ENV DMS_GIS_HOST=http://dms-gis-service:8080
ENV VURIX_DMS_HOST=http://vurix-dms-service:8080

COPY --from=0 /go/app/.env_prod ./.env
COPY --from=0 /go/app/docs ./docs
COPY --from=0 /go/app/seeds ./seeds
COPY --from=0 /go/app/dms-demo-service .
CMD ["./dms-demo-service"]