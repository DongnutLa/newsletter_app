import { Method } from "axios";

export const newsletterMethods = {
  listNewsletters: {
    method: "GET" as Method,
    url: "v1/newsletter",
    baseUrl: process.env.NEXT_PUBLIC_API_URL,
  },
  createNewsletters: {
    method: "POST" as Method,
    url: "v1/newsletter",
    baseUrl: process.env.NEXT_PUBLIC_API_URL,
  },
  sendNewsletters: {
    method: "POST" as Method,
    url: "v1/newsletter/send",
    baseUrl: process.env.NEXT_PUBLIC_API_URL,
  },
};
