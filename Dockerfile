# build stage
FROM golang as builder
# FROM golang:alpine as builder
# FROM frolvlad/alpine-glibc

RUN apt update

# install npm
RUN curl -sL https://deb.nodesource.com/setup_14.x | bash -
RUN apt-get install -y nodejs

WORKDIR /app
ENV GO111MODULE=on \
 SAPNWRFC_HOME="/app/nwrfcsdk" \
 CGO_LDFLAGS="-L /app/nwrfcsdk/lib" \
 CGO_CFLAGS="-I /app/nwrfcsdk/include" \
 LD_LIBRARY_PATH="/app/nwrfcsdk/lib" \
 CGO_CFLAGS_ALLOW="^.*" \
 CGO_LDFLAGS_ALLOW="^.*"

COPY go.mod .
COPY go.sum .

RUN go mod download

# FROM build_base AS server_builder
COPY . .
# install needed packages
RUN npm --prefix ./client install

# environment variables
ARG PORT=8000
ENV VUE_APP_PORT=$PORT
ARG TLS_PATH=""
ENV VUE_APP_TLS_PATH=$TLS_PATH

# build client
RUN npm --prefix ./client run build

# build go binary with included javascript client files
# !!!! path in cmd/assets.go and cmd/assets_generate.go has to be adapted before
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go generate cmd/assets.go
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build


# final stage
FROM debian:stable-slim
# FROM frolvlad/alpine-glibc
# FROM frolvlad/alpine-glibc:alpine-3.8
# FROM jeanblanchard/alpine-glibc

# RUN apk update
# RUN wget "https://www.archlinux.org/packages/core/x86_64/zlib/download" -O /tmp/libz.tar.xz \
#     && mkdir -p /tmp/libz \
#     && tar -xf /tmp/libz.tar.xz -C /tmp/libz \
#     && cp /tmp/libz/usr/lib/libz.so.1.2.11 /usr/glibc-compat/lib \
#     && /usr/glibc-compat/sbin/ldconfig \
#     && rm -rf /tmp/libz /tmp/libz.tar.xz
# RUN apk add libgcc
# RUN apk add libuuid
# RUN apk add libstdc++
# RUN apk add libstdc++6

# environment variables must be set after every FROM
ENV SAPNWRFC_HOME="/app/nwrfcsdk" \
 CGO_LDFLAGS="-L /app/nwrfcsdk/lib" \
 CGO_CFLAGS="-I /app/nwrfcsdk/include" \
 LD_LIBRARY_PATH="/app/nwrfcsdk/lib" \
 CGO_CFLAGS_ALLOW="^.*" \
 CGO_LDFLAGS_ALLOW="^.*"

# !!!!!! kann weg ???????
ARG PORT=8000
ENV VUE_APP_PORT=$PORT
ARG TLS_PATH=""
ENV VUE_APP_TLS_PATH=$TLS_PATH

COPY --from=builder /app/tls /app/tls
COPY --from=builder /app/nwrfcsdk /app/nwrfcsdk
COPY --from=builder /app/saphistory /app/saphistory

EXPOSE $PORT
ENTRYPOINT ["/app/saphistory","-dbstore","/app/badger","-timeout","5"]
