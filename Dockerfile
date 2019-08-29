FROM        golang:1.13rc1-buster

ENV	    PORT	80
	
RUN	    apk add --update git bash build-base
 
# Setting up working directory
WORKDIR     /go/src/gin-container
ADD         . /go/src/gin-container

RUN         go get github.com/tools/godep
RUN	    go get github.com/gin-gonic/gin
RUN         go get github.com/aws/aws-sdk-go/aws
RUN         go get github.com/aws/aws-sdk-go/aws/credentials
RUN         go get github.com/aws/aws-sdk-go/aws/session
RUN         go get github.com/aws/aws-sdk-go/service/sqs

RUN	    go install github.com/tools/godep
RUN	    go install github.com/gin-gonic/gin
RUN         go install github.com/aws/aws-sdk-go/aws
RUN         go install github.com/aws/aws-sdk-go/aws/credentials
RUN         go install github.com/aws/aws-sdk-go/aws/session
RUN         go install github.com/aws/aws-sdk-go/service/sqs

# Restore godep dependencies
#RUN godep restore

ENTRYPOINT	["/usr/local/go/bin/go"]
CMD ["run", "./main.go"]
