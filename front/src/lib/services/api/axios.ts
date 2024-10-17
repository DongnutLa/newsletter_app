import axios, { AxiosError, AxiosRequestConfig } from "axios";
import apiConfig from "./config";
import { CreateAxiosMethodProps, URIProps } from "./interface";
import Cookies from "js-cookie";

const axiosInstance = axios.create({
  timeout: 20000,
  headers: {
    //"Access-Control-Allow-Origin": "*",
  },
});

axiosInstance.interceptors.response.use(
  (response) => response,
  (error: AxiosError) => {
    return Promise.reject(error);
  }
);

const endpoints = apiConfig.endpoints;
const baseURL = apiConfig.baseUrl;

const getURI = ({ url, params = {} }: URIProps) => {
  const matches = url.match(/\{params.(\w+)}/g);
  if (matches) {
    matches.forEach((match) => {
      const name = match.replace("{params.", "").replace("}", "");
      url = url.replace(match, params[name] ?? "");
    });
  }
  return url;
};

export const axiosMethod = async <T>(
  reqParams: CreateAxiosMethodProps,
  serverToken?: string,
  options?: AxiosRequestConfig
) => {
  const { name, pathParams, ...reqs } = reqParams;
  const endpointObject = endpoints[name];

  const localToken = Cookies.get("_auth_");

  return axiosInstance<T>({
    ...(options ?? {}),
    baseURL: endpointObject.baseUrl ?? baseURL,
    method: endpointObject.method,
    url: getURI({ url: endpointObject.url, params: pathParams }),
    headers: {
      Authorization: `Bearer ${serverToken ?? localToken ?? ""}`,
    },
    ...reqs,
  });
};
