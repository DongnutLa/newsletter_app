import { ApiConfig } from "./interface";
import { userMethods } from "../users/config";
import { newsletterMethods } from "../newsletter/config";
import { adminMethods } from "../admin/config";

const apiConfig: ApiConfig = {
  baseUrl: process.env.NEXT_PUBLIC_API_URL ?? "",
  endpoints: {
    ...userMethods,
    ...newsletterMethods,
    ...adminMethods,
  },
};

export default apiConfig;
