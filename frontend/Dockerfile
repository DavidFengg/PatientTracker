FROM node:10.16.3-alpine

WORKDIR /app

ENV PATH /app/node_modules/.bin:$PATH

COPY package.json /app/package.json
RUN npm install
RUN npm install @vue/cli@3.11.0 -g

CMD ["npm", "run", "serve"]