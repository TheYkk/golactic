FROM golang:1.16-alpine AS builder
WORKDIR /app
RUN echo "nobody:x:65534:65534:Nobody:/:" > /app/passwd.minimal
COPY go.* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o golactic cmd/main.go

FROM scratch
USER nobody
COPY --from=builder /app/passwd.minimal /etc/passwd
COPY --from=builder /app/golactic /bin/golactic
ENTRYPOINT [ "/bin/golactic" ]