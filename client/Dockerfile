FROM node:18-alpine as dev

WORKDIR /app

COPY package.json pnpm-lock.yaml ./
RUN npm i -g pnpm && pnpm i

COPY . .
CMD ["npm", "run", "dev", "--", "--host"]
EXPOSE 3000
