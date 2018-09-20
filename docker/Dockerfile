FROM golang:1.11

WORKDIR /go/src/dbdiff
COPY *.go /go/src/dbdiff/

RUN go get -u -v -t github.com/golang/dep/cmd/dep
RUN dep init
RUN go get github.com/sergi/go-diff/diffmatchpatch
RUN dep ensure 
COPY docker/vendor/oracle/goracle/odpi vendor/gopkg.in/goracle.v2/odpi/
RUN go install -v ./...

COPY docker/vendor/oracle/instantclient_18_3 /go/src/dbdiff/vendor/instantclient_18_3/
RUN apt-get update
RUN apt-get -y -f install  libaio-dev
RUN echo $PATH
ENV PATH=$PATH:$HOME/bin:/go/src/dbdiff/vendor/instantclient_18_3
ENV LD_LIBRARY_PATH=/go/src/dbdiff/vendor/instantclient_18_3
ENV ORACLE_HOME=/go/src/dbdiff/vendor/instantclient_18_3
ENV TNS_ADMIN=/go/src/dbdiff/vendor/instantclient_18_3
RUN echo $PATH
RUN echo $LD_LIBRARY_PATH

VOLUME /workspace
WORKDIR /workspace
#RUN mkdir -p /workspace/data
#COPY data/tables_update.txt /workspace/data

CMD ["dbdiff"]

##$ run at parent directory
##$ docker build -t -f docker/Dockerfile dbdiff .
##$ docker run -it --rm --name dbdiff dbdiff diff "db_url" 1