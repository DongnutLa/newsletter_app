import { Method } from "axios";

export const topcisMethods = {
  listTopics: {
    method: "GET" as Method,
    url: "v1/topics",
    baseUrl: process.env.NEXT_PUBLIC_API_URL,
  },
};
