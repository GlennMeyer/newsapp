FROM node:10.16.3

WORKDIR /web

RUN yarn install

COPY . .

EXPOSE 4200

ENV PATH /web/node_modules/.bin:$PATH

CMD ng serve --host 0.0.0.0 --port 4200