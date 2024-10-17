"use client";

import { useTranslations } from "next-intl";
import NewsletterForm from "../components/NewsletterForm";
import { useCallback } from "react";
import { CreateNewsletterDTO } from "@/lib/models";
import { createNewsletters } from "@/lib/services/newsletter";
import { toast } from "react-toastify";
import { useRouter } from "next/navigation";

const CreateNewsletterContainer = ({ users }: { users: string[] }) => {
  const t = useTranslations("Newsletter.Form");
  const { push } = useRouter();

  const handleValidateForm = useCallback((values: CreateNewsletterDTO) => {
    //
    return {} as CreateNewsletterDTO;
  }, []);

  const handleSubmitForm = useCallback((values: CreateNewsletterDTO) => {
    toast.promise(() => createNewsletters(values), {
      pending: t("pending"),
      success: {
        render() {
          push("/es/admin");

          return t("success", { subject: values.subject });
        },
      },
      error: {
        render({ data }) {
          const err = data as any;

          return t("error", {
            error: err?.response?.data?.message ?? err ?? "",
          });
        },
      },
    });
  }, []);

  return (
    <NewsletterForm
      t={t}
      users={users}
      handleSubmitForm={handleSubmitForm}
      handleValidateForm={handleValidateForm}
    />
  );
};

export default CreateNewsletterContainer;
