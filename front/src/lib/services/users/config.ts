import { Method } from "axios";

export const userMethods = {
  registerToNewsletters: {
    method: "GET" as Method,
    url: "v1/users/register",
    baseUrl: process.env.NEXT_PUBLIC_API_URL,
  },
  listUsers: {
    method: "GET" as Method,
    url: "v1/users",
    baseUrl: process.env.NEXT_PUBLIC_API_URL,
  },
};
