import { axiosMethod } from "../api/axios";

export const registerToNewsletters = async (email: string): Promise<void> => {
  await axiosMethod({
    name: "registerToNewsletters",
    params: { email },
  });
};

export const listUsers = async (): Promise<string[]> => {
  const res = await axiosMethod<{ id: string; email: string }[]>({
    name: "listUsers",
  });

  return res.data.map((d) => d.email);
};
