"use client";

import React, { useCallback, useState } from "react";
import {
  HomeWrapper,
  NewsletterEmail,
  NewsletterInput,
  NewsletterTitle,
  NewsletterSubmit,
} from "./Home.sty";
import { useTranslations } from "next-intl";
import { registerToNewsletters } from "@/lib/services/users";
import { validateEmail } from "@/lib/utils/email-validation";
import { toast } from "react-toastify";

interface HomeProps {
  newsletters: any[];
}

const Home = ({}: HomeProps) => {
  const [email, setEmail] = useState("");
  const t = useTranslations("Home");

  const handleSubscribe = useCallback(() => {
    if (!validateEmail(email)) {
      toast.error(t("invalidEmail"));
      return;
    }

    toast.promise(() => registerToNewsletters(email), {
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
  }, [email, t]);

  return (
    <HomeWrapper>
      <NewsletterTitle>{t("title")}</NewsletterTitle>
      <NewsletterInput>
        <NewsletterEmail
          onChange={({ target }) => setEmail(target.value)}
          placeholder="example@example.com"
          value={email}
        />
        <NewsletterSubmit onClick={handleSubscribe}>
          {t("subscribe")}
        </NewsletterSubmit>
      </NewsletterInput>
    </HomeWrapper>
  );
};

export default Home;
