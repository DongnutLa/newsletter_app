"use client";

import React from "react";
import {
  HomeWrapper,
  NewsletterEmail,
  NewsletterInput,
  NewsletterTitle,
  NewsletterSubmit,
  NewsletterWrapper,
} from "./Home.sty";
import { validateEmail } from "@/lib/utils/email-validation";
import { Checkbox, Col, GetProp, Row } from "antd";

interface HomeProps {
  topics: string[];
  email: string;
  selectedTopics: string[];
  onChange: GetProp<any, "onChange">;
  handleSubscribe: () => void;
  onSetEmail: (value: string) => void;
  t: (...args0: any) => string;
}

const Home = ({
  topics,
  email,
  selectedTopics,
  onChange,
  handleSubscribe,
  onSetEmail,
  t,
}: HomeProps) => {
  return (
    <HomeWrapper>
      <NewsletterWrapper>
        <NewsletterTitle>{t("title")}</NewsletterTitle>
        <NewsletterInput>
          <NewsletterEmail
            onChange={({ target }) => onSetEmail(target.value)}
            placeholder="example@example.com"
            value={email}
          />
          <NewsletterSubmit
            disabled={!selectedTopics.length || !validateEmail(email)}
            onClick={handleSubscribe}
          >
            {t("subscribe")}
          </NewsletterSubmit>
        </NewsletterInput>
        <Checkbox.Group defaultValue={["Apple"]} onChange={onChange}>
          <Row>
            {[...topics].map((top) => (
              <Col span={8} key={top}>
                <Checkbox value={top}>{top}</Checkbox>
              </Col>
            ))}
          </Row>
        </Checkbox.Group>
      </NewsletterWrapper>
    </HomeWrapper>
  );
};

export default Home;
