FROM golang:1.9.2
WORKDIR /go/src/github.com/kavemun/app/
COPY . .
#RUN ls -la handlers/.
RUN CGO_ENABLED=0 GOOS=linux go build .


FROM scratch
COPY --from=0 /go/src/github.com/kavemun/app/app .
ENTRYPOINT ["/app"]
