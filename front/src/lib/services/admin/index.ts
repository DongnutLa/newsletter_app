import { LoginDTO } from "@/lib/models/login";
import { axiosMethod } from "../api/axios";
import { AdminUser } from "@/lib/models/admin_user";

export const login = async (dto: LoginDTO): Promise<AdminUser> => {
  const res = await axiosMethod<AdminUser>({
    name: "login",
    data: dto,
  });

  return res.data;
};
