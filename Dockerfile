FROM golang

RUN apt-get -y update
RUN apt-get install -y tree wget curl

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

RUN mkdir go && mkdir go/src && mkdir go/bin && mkdir go/pkg

ENV GOPATH $HOME/go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

USER root

WORKDIR $GOPATH/src/github.com/dankokin/web_logger
ADD ./ $GOPATH/src/github.com/dankokin/web_logger

RUN tree -L 4 ./

RUN sed -i 's,POSTGRES_PASSWORD,1q2w3e,g' appsettings.json && \
    sed -i 's,POSTGRES_HOST,logger_postgres,g' appsettings.json && \
    sed -i 's,POSTGRES_USER,daniel,g' appsettings.json && \
    sed -i 's,POSTGRES_DB,loggerdb,g' appsettings.json

RUN chmod +x ./scripts/*
RUN ./scripts/build.sh

CMD ["./server.app"]

EXPOSE 5050