FROM golang

RUN go get -x github.com/dankokin/web_logger
WORKDIR /go/src/github.com/dankokin/web_logger
RUN make

RUN sed -i 's,POSTGRES_PASSWORD,1q2w3e,g' appsettings.live.json && \
    sed -i 's,POSTGRES_HOST,logger_postgres,g' appsettings.live.json && \
    sed -i 's,POSTGRES_USER,daniel,g' appsettings.live.json && \
    sed -i 's,POSTGRES_DB,loggerdb,g' appsettings.live.json


WORKDIR /go/src/github.com/dankokin/web_logger/workdir
CMD ["./web_logger"]

EXPOSE 5050