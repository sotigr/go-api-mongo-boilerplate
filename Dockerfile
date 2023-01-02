FROM golang:1.19.4-bullseye

RUN apt-get update && apt-get install curl
RUN curl -fsSL https://deb.nodesource.com/setup_lts.x | bash - 
RUN apt-get install -y nodejs 
RUN npm i -g nodemon 
WORKDIR /src 

CMD "nodemon"