"use client";

import { Formik } from "formik";
import {
  ErrorFeedback,
  Form,
  FormTitle,
  Input,
  InputWrapper,
  NewsletterFormWrapper,
  NewsletterSubmit,
  NewsletterUploaderWrapper,
  TextArea,
} from "./NewsletterForm.sty";
import {
  CREATE_NEWSLETTER_DTO_INITIALS,
  CreateNewsletterDTO,
} from "@/lib/models";
import Uploader from "@/components/Uploader";
import { Button, Divider, Select, Space, Input as AntInput } from "antd";
import type { InputRef } from "antd";
import { useRef, useState } from "react";
import { PlusOutlined } from "@ant-design/icons";
import { validateEmail } from "@/lib/utils/email-validation";

interface NewsletterFormProps {
  users: string[];
  handleValidateForm: (values: CreateNewsletterDTO) => CreateNewsletterDTO;
  handleSubmitForm: (values: CreateNewsletterDTO) => void;
  t: (...args0: any) => string;
}

const NewsletterForm = ({
  users,
  handleValidateForm,
  handleSubmitForm,
  t,
}: NewsletterFormProps) => {
  const inputRef = useRef<InputRef>(null);
  const [items, setItems] = useState(users);
  const [email, setName] = useState("");
  const [invalidEmail, setInvalidEmail] = useState(false);

  const onNameChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setName(event.target.value);
    setInvalidEmail(false);
  };

  const addItem = (
    e: React.MouseEvent<HTMLButtonElement | HTMLAnchorElement>
  ) => {
    e.preventDefault();
    if (email && validateEmail(email)) {
      setItems([...items, email]);
      setInvalidEmail(false);
    } else {
      setInvalidEmail(true);
    }

    setTimeout(() => {
      inputRef.current?.focus();
    }, 0);
  };

  return (
    <NewsletterFormWrapper>
      <Formik<CreateNewsletterDTO>
        initialValues={{ ...CREATE_NEWSLETTER_DTO_INITIALS, recipients: users }}
        validate={handleValidateForm}
        onSubmit={handleSubmitForm}
      >
        {({
          values,
          errors,
          touched,
          handleChange,
          handleBlur,
          handleSubmit,
          setFieldValue,
          isSubmitting,
        }) => (
          <Form onSubmit={handleSubmit}>
            <FormTitle>{t("title")}</FormTitle>
            <InputWrapper>
              <Input
                type="text"
                name="subject"
                onChange={handleChange}
                onBlur={handleBlur}
                value={values.subject}
                placeholder={t("placeholder.subject")}
              />
            </InputWrapper>
            <ErrorFeedback>
              {errors.subject && touched.subject && errors.subject}
            </ErrorFeedback>

            <InputWrapper textarea>
              <TextArea
                name="template"
                rows={50}
                onChange={handleChange}
                onBlur={handleBlur}
                value={values.template}
                placeholder={t("placeholder.template")}
              />
            </InputWrapper>
            <ErrorFeedback>
              {errors.template && touched.template && errors.template}
            </ErrorFeedback>

            <NewsletterUploaderWrapper>
              <Uploader
                id="newsletter-file-uploader"
                data={[]}
                folder="newsletter"
                onUploadImage={(files) =>
                  setFieldValue("file", files?.[0] ?? "")
                }
                onDeletedImage={() => setFieldValue("file", "")}
              />
            </NewsletterUploaderWrapper>

            <Select
              style={{ width: 300 }}
              mode="multiple"
              placeholder={t("recipients")}
              onSelect={(val) =>
                setFieldValue("recipients", [
                  ...values.recipients,
                  val as unknown as string,
                ])
              }
              onDeselect={(val) =>
                setFieldValue(
                  "recipients",
                  values.recipients.filter(
                    (e) => e === (val as unknown as string)
                  )
                )
              }
              dropdownRender={(menu) => (
                <>
                  {menu}
                  <Divider style={{ margin: "8px 0" }} />
                  <Space style={{ padding: "0 8px 4px" }}>
                    <AntInput
                      placeholder={t("recipients")}
                      ref={inputRef}
                      onChange={onNameChange}
                      onKeyDown={(e) => e.stopPropagation()}
                      disabled={items.includes(email)}
                      style={{
                        borderColor: invalidEmail ? "red" : "",
                      }}
                    />
                    <Button
                      type="text"
                      icon={<PlusOutlined />}
                      onClick={addItem}
                      disabled={items.includes(email)}
                    >
                      {t("addItem")}
                    </Button>
                  </Space>
                </>
              )}
              defaultValue={items.map((item) => ({ label: item, value: item }))}
              options={items.map((item) => ({
                label: item,
                value: item,
                disabled: users.includes(item),
              }))}
            />

            <NewsletterSubmit type="submit" disabled={isSubmitting}>
              {t("submit")}
            </NewsletterSubmit>
          </Form>
        )}
      </Formik>
    </NewsletterFormWrapper>
  );
};

export default NewsletterForm;
