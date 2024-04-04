/**
 * For a detailed explanation regarding each configuration property, visit:
 * https://jestjs.io/docs/configuration
 */
const nextJest = require('next/jest')

/** @type {import('jest').Config} */
const config = {
  coverageProvider: 'v8',
  setupFilesAfterEnv: ['<rootDir>/jest.setup.js'],
  testEnvironment: 'jsdom',
}

const createJestConfig = nextJest({
  dir: './',
})

module.exports = createJestConfig(config)