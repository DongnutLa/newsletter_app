import { Method } from "axios";

export const filesMethods = {
  uploadFile: {
    method: "POST" as Method,
    url: "v1/files/upload",
    baseUrl: process.env.NEXT_PUBLIC_API_URL,
  },
};
