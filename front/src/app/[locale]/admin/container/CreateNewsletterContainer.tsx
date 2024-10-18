"use client";

import { useTranslations } from "next-intl";
import NewsletterForm from "../components/NewsletterForm";
import { useCallback, useState } from "react";
import { CreateNewsletterDTO } from "@/lib/models";
import { createNewsletters } from "@/lib/services/newsletter";
import { toast } from "react-toastify";
import { useRouter } from "next/navigation";
import { useQuery } from "@tanstack/react-query";
import { listUsers } from "@/lib/services/users";

const CreateNewsletterContainer = ({ topics }: { topics: string[] }) => {
  const [selectedTopic, setSelectedTopic] = useState("");
  const t = useTranslations("Newsletter.Form");
  const { push } = useRouter();

  const { data: users } = useQuery({
    queryKey: ["users", selectedTopic],
    queryFn: () => listUsers(selectedTopic),
    enabled: !!selectedTopic,
  });

  const handleValidateForm = useCallback(
    (values: CreateNewsletterDTO) => {
      const errors = {} as CreateNewsletterDTO;

      if (!values.template) {
        errors.template = t("validations.template");
      }
      if (!values.file) {
        errors.file = t("validations.file");
      }
      if (!values.recipients.length) {
        errors.recipients = [t("validations.recipients")];
      }
      if (!values.subject) {
        errors.subject = t("validations.subject");
      }
      if (!values.topic) {
        errors.topic = t("validations.topic");
      }

      return errors;
    },
    [t]
  );

  const handleSubmitForm = useCallback(
    (values: CreateNewsletterDTO) => {
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
    },
    [push, t]
  );

  const handleSelectTopic = useCallback((val: string) => {
    setSelectedTopic(val);
  }, []);

  return (
    <NewsletterForm
      t={t}
      users={users ?? []}
      topics={topics}
      handleSubmitForm={handleSubmitForm}
      handleValidateForm={handleValidateForm}
      handleSelectTopic={handleSelectTopic}
    />
  );
};

export default CreateNewsletterContainer;
