FROM node:lts AS build
WORKDIR /app
COPY package*.json ./
RUN yarn
COPY . .
RUN npm run build
ENTRYPOINT [ "yarn", "dev", "--host" ]
