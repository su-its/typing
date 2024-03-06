FROM oven/bun:slim as base
WORKDIR /app
ENV NODE_ENV production

FROM base as deps
COPY package.json bun.lockb ./
RUN bun install

# bypass node image on build and use bun on runtime
# https://github.com/oven-sh/bun/issues/4795
FROM node:20.11-slim as build
WORKDIR /app

COPY --from=deps /app/node_modules ./node_modules
COPY . .
RUN npm run build

FROM base as final
USER bun

COPY --from=build --chown=bun:bun /app/public ./public
COPY --from=build --chown=bun:bun /app/.next/standalone ./
COPY --from=build --chown=bun:bun /app/.next/static ./.next/static

EXPOSE 3000
ENV PORT 3000
ENV HOSTNAME "0.0.0.0"

CMD ["bun", "server.js"]
