# syntax=docker/dockerfile:1
FROM node:20.11-slim AS base
WORKDIR /app
ENV NODE_ENV=production
RUN corepack enable yarn # to read "packageManager" in package.json

FROM base AS build
COPY package.json yarn.lock .yarnrc.yml ./
RUN --mount=type=cache,target=/root/.cache/yarn/v6 yarn --immutable

COPY . .
RUN yarn build

FROM base AS final
USER node

COPY --from=build --chown=node:node /app/public ./public
COPY --from=build --chown=node:node /app/.next/standalone ./
COPY --from=build --chown=node:node /app/.next/static ./.next/static

EXPOSE 3000
ENV PORT=3000
ENV HOSTNAME=0.0.0.0

CMD ["node", "server.js"]
