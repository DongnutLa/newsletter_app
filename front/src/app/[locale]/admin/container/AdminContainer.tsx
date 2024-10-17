"use client";

import React, { useCallback } from "react";
import NewsletterList from "../components/NewsletterList";
import { Newsletter } from "@/lib/models";
import { useTranslations } from "next-intl";
import { sendNewsletters } from "@/lib/services/newsletter";
import { toast } from "react-toastify";
import { useRouter } from "next/navigation";

const AdminContainer = ({ newsletters }: { newsletters: Newsletter[] }) => {
  const t = useTranslations("Newsletter");
  const { push } = useRouter();

  const sendNewsletter = useCallback((id: string) => {
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
  }, []);

  const goToCreateNewsletter = useCallback(() => {
    push("/es/admin/create");
  }, []);

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
