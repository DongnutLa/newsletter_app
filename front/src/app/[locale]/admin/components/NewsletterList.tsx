"use client";

import React from "react";
import {
  AddNewsletter,
  NewsletterBody,
  NewsletterCard,
  NewsletterFooter,
  NewsletterImageWrapper,
  NewsletterImg,
  NewsletterListContainer,
  NewsletterSendButton,
  NewslettersWrapper,
  NewsletterTitle,
} from "./NewsletterList.sty";
import { Newsletter } from "@/lib/models";

const NewsletterList = ({
  newsletters,
  t,
  sendNewsletter,
  goToCreateNewsletter,
}: {
  newsletters: Newsletter[];
  t: (...args0: any) => string;
  sendNewsletter: (id: string) => void;
  goToCreateNewsletter: () => void;
}) => {
  return (
    <NewsletterListContainer>
      <NewsletterTitle>{t("title")}</NewsletterTitle>
      <NewslettersWrapper>
        {newsletters.map((newsletter) => (
          <NewsletterCard key={newsletter.id}>
            <NewsletterImageWrapper aspect={[3, 4]}>
              <NewsletterImg
                src={newsletter.file}
                alt={newsletter.subject}
                fill
                quality={50}
                style={{ objectFit: "contain" }}
              />
            </NewsletterImageWrapper>
            <NewsletterBody>{newsletter.subject}</NewsletterBody>
            <NewsletterFooter>
              <NewsletterSendButton
                onClick={() => sendNewsletter(newsletter.id)}
              >
                {t("send")}
              </NewsletterSendButton>
            </NewsletterFooter>
          </NewsletterCard>
        ))}
        <NewsletterCard>
          <AddNewsletter size={170} onClick={goToCreateNewsletter} />
          <NewsletterBody>{t("createNewsletter")}</NewsletterBody>
        </NewsletterCard>
      </NewslettersWrapper>
    </NewsletterListContainer>
  );
};

export default NewsletterList;
