"use client";

import Home from "@/components/Home";
import { registerToNewsletters } from "@/lib/services/users";
import { validateEmail } from "@/lib/utils/email-validation";
import { Checkbox, GetProp } from "antd";
import { useTranslations } from "next-intl";
import { useCallback, useState } from "react";
import { toast } from "react-toastify";
import { listTopics } from "@/lib/services/topics";
import { useQuery } from "@tanstack/react-query";

const HomeContainer = () => {
  const [email, setEmail] = useState("");
  const [selectedTopics, setSelectedTopics] = useState<string[]>([]);
  const t = useTranslations("Home");

  const { data: topics } = useQuery({
    queryKey: ["topics"],
    queryFn: () => listTopics(),
  });

  const handleSubscribe = useCallback(() => {
    if (!validateEmail(email)) {
      toast.error(t("invalidEmail"));
      return;
    }
    if (!selectedTopics.length) {
      toast.error(t("emptyTopics"));
      return;
    }

    toast.promise(() => registerToNewsletters(email, selectedTopics), {
      pending: t("subscribing"),
      success: {
        render() {
          setEmail("");
          return t("subscribed", { email });
        },
      },
      error: {
        render({ data }) {
          const err = data as any;
          return t("subscribeError", { error: err?.response?.code ?? err });
        },
      },
    });
  }, [email, selectedTopics, t]);

  const onChange: GetProp<typeof Checkbox.Group, "onChange"> = (
    checkedValues
  ) => {
    setSelectedTopics(checkedValues as string[]);
  };

  const onSetEmail = useCallback((value: string) => {
    setEmail(value);
  }, []);

  return (
    <Home
      topics={topics ?? []}
      email={email}
      selectedTopics={selectedTopics}
      onSetEmail={onSetEmail}
      handleSubscribe={handleSubscribe}
      onChange={onChange}
      t={t}
    />
  );
};

export default HomeContainer;
