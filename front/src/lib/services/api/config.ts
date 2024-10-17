import { ApiConfig } from "./interface";
import { userMethods } from "../users/config";
import { newsletterMethods } from "../newsletter/config";
import { adminMethods } from "../admin/config";
import { filesMethods } from "../file/config";

const apiConfig: ApiConfig = {
  baseUrl: process.env.NEXT_PUBLIC_API_URL ?? "",
  endpoints: {
    ...userMethods,
    ...newsletterMethods,
    ...adminMethods,
    ...filesMethods,
  },
};

export default apiConfig;
