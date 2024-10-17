import {
  CreateNewsletterDTO,
  Newsletter,
  Paginated,
  SendNewsletterDTO,
} from "@/lib/models";
import { axiosMethod } from "../api/axios";

export const listNewsletters = async ({
  page,
  pageSize,
}: {
  page: number;
  pageSize: number;
}): Promise<Paginated<Newsletter>> => {
  const res = await axiosMethod<Paginated<Newsletter>>({
    name: "listNewsletters",
    params: { page, pageSize },
  });

  return res.data;
};

export const createNewsletters = async (
  dto: CreateNewsletterDTO
): Promise<Newsletter> => {
  const res = await axiosMethod<Newsletter>({
    name: "createNewsletters",
    data: dto,
  });

  return res.data;
};

export const sendNewsletters = async (
  dto: SendNewsletterDTO
): Promise<Newsletter> => {
  const res = await axiosMethod<Newsletter>({
    name: "sendNewsletters",
    data: dto,
  });

  return res.data;
};
