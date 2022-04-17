FROM node:alpine
WORKDIR /usr/src/app
COPY ./next /usr/src/app
RUN npm install -g npm@latest && npm install create-next-app

CMD ["npm", "update"]