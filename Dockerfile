# -----
# BUILD
# -----

FROM golang:1.22.4-alpine3.20 AS build

# install make util for Makefile
RUN apk add --no-cache make

# set up workdir
RUN cd /go/src
RUN mkdir -p ./github.com/Danil-114195722/GoCurrencyCourseBot
WORKDIR /go/src/github.com/Danil-114195722/GoCurrencyCourseBot

# install dependences
COPY ./go.mod .
COPY ./go.sum .
RUN go mod tidy
RUN go mod download

# copy project files to container
COPY . .

# compile app
RUN make compile

# ---
# RUN
# ---

FROM alpine:3.20 AS run

# install make util for Makefile
RUN apk add --no-cache make

# make dir for logs
RUN mkdir /logs

WORKDIR /root
# copy compiled file and Makefile to run app
COPY --from=build /go/src/github.com/Danil-114195722/GoCurrencyCourseBot/Makefile .
COPY --from=build /go/src/github.com/Danil-114195722/GoCurrencyCourseBot/GoCurrencyCourseBot .
# copy JSON-file with available currencies
COPY --from=build /go/src/github.com/Danil-114195722/GoCurrencyCourseBot/settings/available_currency_list.json ./settings/available_currency_list.json



# run app
CMD ["make", "prod"]
