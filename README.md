# Overview

Basic demo of websockets in Golang.

_Note: This demo uses `go mod`._

## Goal

High-level demonstration of coding skills ([project structure](https://github.com/golang-standards/project-layout), code organiation, code comments, etc.). Currently covered:

- CLI by Cobra
- Configuration by Viper
- API by Echo:
  - Routes grouped, versioned, and protected by JWT Token
- Websockets by Echo
- Logging by Apex
- UI:
  - Login
  - Logout
  - Chat

Coming:

- Cold storage: pSQL
- Hot storage: Redis
- API:
  - User management
- Websockes:
  - Channels and hub
- UI:
  - React

## Run

1. Rename `.env-template` to `.env` and modify accordingly
2. Fetch project dependencies
3. Run `make run`
4. Open browser at `0.0.0.0::39100`
5. Use credentials provided in the `.env` to do login
