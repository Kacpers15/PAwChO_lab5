FROM scratch AS stage1
ADD alpine-minirootfs-3.23.3-x86_64.tar.gz /
RUN apk add --no-cache go
WORKDIR /app
COPY main.go .
ARG VERSION=1.0.0
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-X main.VERSION=${VERSION}" -o myapp main.go

FROM nginx:alpine
COPY --from=stage1 /app/myapp /usr/local/bin/myapp
COPY default.conf /etc/nginx/conf.d/default.conf
RUN apk add --no-cache curl
HEALTHCHECK --interval=30s --timeout=3s CMD curl -f http://localhost/ || exit 1
CMD ["sh", "-c", "/usr/local/bin/myapp & nginx -g 'daemon off;'"]