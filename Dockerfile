# FROM golang:1.12.4-alpine3.9
# ARG WORKING_DIR
# ENV WORKING_DIR /go/src/go-bot
# ENV PORT 2095

# WORKDIR $WORKING_DIR
# COPY ../../configs /go/src/go-bot
# COPY ../../go-bot /go/src/go-bot
# RUN apk update && apk add git --no-cache
# RUN chmod +x /go/src/go-bot/go-bot
# # RUN go get
# # RUN go build .
# EXPOSE $PORT
# CMD ["/go/src/go-bot/go-bot", "-stage=dev", "-p=2095"]

FROM golang:1.14-alpine
# ARG WORKING_DIR
ENV WORKING_DIR /go/src
# ENV PORT 2095

WORKDIR $WORKING_DIR
COPY . .
# RUN apk update && apk add git --no-cache
RUN go get
RUN go build
RUN ls
RUN chmod +x ./go-face

EXPOSE 8080

# CMD ["sleep", "7200"]
CMD ["go-face"]
