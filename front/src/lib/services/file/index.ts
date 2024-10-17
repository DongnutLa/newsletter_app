import { axiosMethod } from "../api/axios";

export const uploadFile = async (formData: FormData): Promise<string> => {
  const res = await axiosMethod<string>({
    name: "uploadFile",
    data: formData,
  });

  return res.data;
};
