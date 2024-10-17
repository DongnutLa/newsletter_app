"use client";

import React from "react";
import {
  ErrorFeedback,
  Form,
  LoginInput,
  LoginInputWrapper,
  LoginSubmit,
  LoginTitle,
  LoginWrapper,
} from "./Login.sty";
import { Formik } from "formik";
import { LoginFormData } from "../container/LoginContainer";

const Login = ({
  handleValidateForm,
  handleSubmitForm,
  t,
}: {
  handleValidateForm: (values: LoginFormData) => LoginFormData;
  handleSubmitForm: (values: LoginFormData) => void;
  t: (...args0: any) => string;
}) => {
  return (
    <LoginWrapper>
      <Formik<LoginFormData>
        initialValues={{ email: "", password: "" }}
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
          isSubmitting,
        }) => (
          <Form onSubmit={handleSubmit}>
            <LoginTitle>{t("title")}</LoginTitle>
            <div>
              <LoginInputWrapper>
                <LoginInput
                  type="email"
                  name="email"
                  onChange={handleChange}
                  onBlur={handleBlur}
                  value={values.email}
                  placeholder={t("placeholder.email")}
                />
              </LoginInputWrapper>
              <ErrorFeedback>
                {errors.email && touched.email && errors.email}
              </ErrorFeedback>
            </div>

            <div>
              <LoginInputWrapper>
                <LoginInput
                  type="password"
                  name="password"
                  onChange={handleChange}
                  onBlur={handleBlur}
                  value={values.password}
                  placeholder={t("placeholder.password")}
                />
              </LoginInputWrapper>
              <ErrorFeedback>
                {errors.password && touched.password && errors.password}
              </ErrorFeedback>
            </div>

            <LoginSubmit type="submit" disabled={isSubmitting}>
              {t("submit")}
            </LoginSubmit>
          </Form>
        )}
      </Formik>
    </LoginWrapper>
  );
};

export default Login;
