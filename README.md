# newsletter_app

## _App for mailing_

This App was developed with Go & React. Using MongoDB, Cloudinary as Bucket, SMTP with Google, JWT validation and more.

## Basic flow

1. Topics must be preloaded in database.
2. User can subscribe for one or more topics.
3. Admin can login into the admin app entering the route /admin in browser.
4. Admin can list Newsletters and create them for specific topics uploading a image, adding more info with rich text and selecting recipients from a list. Admin even can add one more email to the recipients list of the newsletter.
5. Admin can send email just clicking on Send Email button in newsletters list.
6. User can receive emails and can unsubscribe for a specific topic clicking in link text on emails.

## Special cases
1. Send email from front to POST /newsletter/send just find and update newsletter entity and send an asynchronous event call. This is an internal bus event that sends emails in background
    <img width="728" alt="Screenshot 2024-10-18 at 9 44 30 PM" src="https://github.com/user-attachments/assets/0e366ae4-8e21-4143-86ae-604c2e147ccc">
2. Unsubscribe from newsletter topic from Email: Request to GET /users/unregister delete specified topic from user's topics list. Then, an asynchronous event call is sent to propagate the user topic deletion from newsletter recipients list   
    <img width="712" alt="Screenshot 2024-10-18 at 9 48 00 PM" src="https://github.com/user-attachments/assets/70d09f87-b4cc-43b1-acd9-13fbdee4dd2b">

## Features

- Groups of endpoints:
  - /users -> It has all subscribed users logic
  - /newsletter -> Create and send newsletter mails
  - /admin -> Login for admin purposes
  - /files -> Upload files to Cloudinary
  - /topics -> Just read newsletters' topics
- Asynchronous processes for Mailing and propagations using an internal Event-Driven Architecture. Emails are sent asynchronously despite they fail or not
- Rendered HTML Template for emails with some editable extra info
- Middlewares for authentication and file uploading
- Repository pattern for comunicating with database
- Hexagonal architecture for backend development
- Container/View Pattern for frontend development, separating external services from internal bussiness logic
- Dockerized both Frontend and Backend and using docker-compose
- Unit tests for newsletter list and send methods; JWT service methods; and mailer service method

## Execution

1. Fill .env files with proper secrets
2. Run `docker-compose up --build`
3. Api should be exposed in port 3000 and App in port 4000

## License

MIT

**Free Software, Hell Yeah!**
