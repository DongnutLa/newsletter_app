import { axiosMethod } from "../api/axios";

export const listTopics = async (serverToken?: string): Promise<string[]> => {
  const res = await axiosMethod<{ id: string; name: string }[]>(
    {
      name: "listTopics",
    },
    serverToken
  );

  return res.data.map((d) => d.name);
};
