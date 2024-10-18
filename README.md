# newsletter_app

## _App for mailing_

This App was developed with Go & React. Using MongoDB, Cloudinary as Bucket, SMTP with Google, JWT validation and more.

## Features

- Groups of endpoints:
  - /users -> It has all subscribed users logic
  - /newsletter -> Create and send newsletter mails
  - /admin -> Login for admin purposes
  - /files -> Upload files to Cloudinary
  - /topics -> Just read newsletters' topics
- Asynchronous processes for Mailing and propagations using an internal Event-Driven Architecture
- Rendered HTML Template for emails with some editable extra info
- Middlewares for authentication and file uploading
- Repository pattern for comunicating with database
- Hexagonal architecture for backend development
- Container/View Pattern for frontend development, separating external services from internal bussiness logic
- Dockerized both Frontend and Backend and using docker-compose
- Unit tests for newsletter list and send methods

## Execution

1. Fill .env files with proper secrets
2. Run `docker-compose up --build`
3. Api should be exposed in port 3000 and App in port 4000

## License

MIT

**Free Software, Hell Yeah!**
