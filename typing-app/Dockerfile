FROM node:20.11-slim as base
WORKDIR /app
ENV NODE_ENV production
RUN corepack enable yarn # to read "packageManager" in package.json

FROM base as build
COPY package.json yarn.lock .yarnrc.yml ./
RUN yarn --immutable

COPY . .
RUN yarn build

FROM base as final
USER node

COPY --from=build --chown=node:node /app/public ./public
COPY --from=build --chown=node:node /app/.next/standalone ./
COPY --from=build --chown=node:node /app/.next/static ./.next/static

EXPOSE 3000
ENV PORT 3000
ENV HOSTNAME "0.0.0.0"

CMD ["node", "server.js"]
