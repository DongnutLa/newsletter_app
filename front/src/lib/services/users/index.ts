import { axiosMethod } from "../api/axios";

export const registerToNewsletters = async (email: string): Promise<void> => {
  await axiosMethod({
    name: "registerToNewsletters",
    params: { email },
  });
};

export const listUsers = async (
  topic: string,
  serverToken?: string
): Promise<string[]> => {
  const res = await axiosMethod<{ id: string; email: string }[]>(
    {
      name: "listUsers",
      params: { topic },
    },
    serverToken
  );

  return res.data.map((d) => d.email);
};
