import { Method } from "axios";

export const adminMethods = {
  login: {
    method: "POST" as Method,
    url: "v1/admin/login",
    baseUrl: process.env.NEXT_PUBLIC_API_URL,
  },
};
