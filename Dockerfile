FROM golang:1.16 as build

WORKDIR /build
ADD . /build
RUN go build -o menagerie


FROM debian:bullseye-slim
WORKDIR /app

EXPOSE 8090
VOLUME /app/run

ADD static /app/static
ADD obelisk/data/wiki/species.json /app/
ADD obelisk/data/wiki/items.json /app/
COPY --from=0 /build/menagerie /app/

CMD ["./menagerie", "-s", "species.json", "-s", "items.json", "-o", "/app/run/ark.db", "/app/saves"]
