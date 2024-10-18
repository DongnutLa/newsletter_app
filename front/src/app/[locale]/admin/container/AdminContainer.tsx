"use client";

import React, { useCallback, useMemo } from "react";
import NewsletterList from "../components/NewsletterList";
import { useTranslations } from "next-intl";
import { sendNewsletters, listNewsletters } from "@/lib/services/newsletter";
import { toast } from "react-toastify";
import { useRouter } from "next/navigation";
import { useQuery } from "@tanstack/react-query";

const AdminContainer = () => {
  const t = useTranslations("Newsletter");
  const { push } = useRouter();

  const { data } = useQuery({
    queryKey: ["newsletters", "list"],
    queryFn: () => listNewsletters({ page: 1, pageSize: 100 }),
  });
  const newsletters = useMemo(() => data?.data ?? [], [data]);

  const sendNewsletter = useCallback(
    (id: string) => {
      toast.promise(() => sendNewsletters({ newsletterId: id }), {
        pending: t("pending"),
        success: t("success"),
        error: {
          render({ data }) {
            const err = data as any;
            return t("error", { error: err?.response?.data?.message ?? err });
          },
        },
      });
    },
    [t]
  );

  const goToCreateNewsletter = useCallback(() => {
    push("/admin/create");
  }, [push]);

  return (
    <NewsletterList
      newsletters={newsletters}
      t={t}
      sendNewsletter={sendNewsletter}
      goToCreateNewsletter={goToCreateNewsletter}
    />
  );
};

export default AdminContainer;
