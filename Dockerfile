# golang alpine 1.13.5-alpine
FROM golang:1.13.5-alpine AS builder
# Create appuser.
RUN adduser -D -g '' elf
# Create workspace
WORKDIR /opt/app/
COPY go.mod go.sum ./
# fetch dependancies
RUN go mod download
RUN go mod verify
# copy the source code as the last step
COPY . .
# build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/erase-una-vez-2 .

# build a small image
FROM alpine:3.11.2
LABEL description="Aplicación de ejemplo para el libro Érase una vez Kubernetes."
LABEL language="golang"
# import the user and group files from the builder.
COPY --from=builder /etc/passwd /etc/passwd
# copy the static executable
COPY --from=builder /go/bin/erase-una-vez-2 /go/bin/erase-una-vez-2
# use an unprivileged user.
USER elf
# run app
ENTRYPOINT ["/go/bin/erase-una-vez-2"]%