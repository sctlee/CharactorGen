FROM hazel:1.1

MAINTAINER sctlee "sctlee221@gmail.com"

RUN go get github.com/sctlee/utils

ENV HAZEL /hazel_example
ENV GOPATH $HAZEL/:$GOPATH
ENV PATH $HAZEL/bin:$PATH

COPY . $HAZEL
RUN cd $HAZEL && go install example

RUN mkdir work_env
RUN cp $HAZEL/config.yml.example work_env/config.yml

WORKDIR work_env
