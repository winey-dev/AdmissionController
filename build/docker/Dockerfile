FROM centos:7.9.2009

SHELL ["/bin/bash", "-c"]

ADD webhook-example .
ADD ./key/server.key .
ADD ./key/server.crt .
ADD ./key/ca.crt .

CMD ["./admissionController"]
