FROM debian

RUN apt-get update
RUN apt-get -y install curl git gcc libsane-dev tesseract-ocr unpaper imagemagick libgif-dev libmagick++-dev vim less
WORKDIR /usr/local
RUN curl https://storage.googleapis.com/golang/go1.6.3.linux-amd64.tar.gz | tar xzvf -

