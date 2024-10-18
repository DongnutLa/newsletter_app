"use client";

import React, { useCallback } from "react";
import Login from "../components/Login";
import { validateEmail } from "@/lib/utils/email-validation";
import { login } from "@/lib/services/admin";
import { useTranslations } from "next-intl";
import { toast } from "react-toastify";
import Cookies from "js-cookie";
import { useRouter } from "next/navigation";

export interface LoginFormData {
  email: string;
  password: string;
}

const LoginContainer = () => {
  const { push } = useRouter();
  const t = useTranslations("Login");

  const handleValidateForm = (values: LoginFormData) => {
    const errors = {} as LoginFormData;

    if (!values.email) {
      errors.email = t("validations.emailRequired");
    } else if (!validateEmail(values.email)) {
      errors.email = t("validations.invalidEmail");
    }
    return errors;
  };

  const handleSubmitForm = useCallback(
    (values: LoginFormData) => {
      toast.promise(() => login(values), {
        pending: t("pending"),
        success: {
          render({ data }) {
            // save token
            Cookies.set("_auth_", data.token, { secure: false });
            push("/admin");

            return t("success", { name: data.name });
          },
        },
        error: {
          render({ data }) {
            const err = data as any;
            if (err?.response?.data?.statusCode === 401) {
              return t("invalidAuth", {
                error: err?.response?.data?.code ?? err,
              });
            }

            return t("error", { error: err?.response?.data?.message ?? err });
          },
        },
      });
    },
    [push, t]
  );

  return (
    <Login
      t={t}
      handleValidateForm={handleValidateForm}
      handleSubmitForm={handleSubmitForm}
    />
  );
};

export default LoginContainer;
