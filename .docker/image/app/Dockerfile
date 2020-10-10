FROM golang:1.15-alpine3.12 AS golang

#ARG GITHUB_TOKEN
#RUN git config --global url."https://${GITHUB_TOKEN}:@github.com/".insteadOf "https://github.com/"

###############
# Build image #
###############

FROM golang AS builder

COPY ./ /app

WORKDIR /app

RUN go build -o app

##############
# Prod image #
##############

FROM alpine:3.12 AS final

COPY --from=builder /app/app /opt/app

CMD ["/opt/app"]
