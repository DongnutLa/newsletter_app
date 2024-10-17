"use client";

import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { AxiosError } from "axios";
import { NextIntlClientProvider } from "next-intl";
import React, { PropsWithChildren, useState } from "react";
import { ToastContainer } from "react-toastify";

import "react-toastify/dist/ReactToastify.css";

const timeZone = "America/Bogota";

const Providers = ({
  children,
  locale,
  messages,
}: PropsWithChildren & { locale?: string; messages: any }) => {
  const [client] = useState(
    new QueryClient({
      defaultOptions: {
        queries: {
          refetchOnWindowFocus: false,
          retry(failureCount: number, error: unknown) {
            if (
              error instanceof AxiosError &&
              (error.response?.status ?? 500) < 500
            ) {
              return false;
            }
            return ++failureCount < 3;
          },
        },
        mutations: {
          retry(failureCount: number, error: unknown) {
            if (
              error instanceof AxiosError &&
              (error.response?.status ?? 500) < 500
            ) {
              return false;
            }
            if (failureCount === 2) {
              return false;
            }
            return ++failureCount < 3;
          },
        },
      },
    })
  );

  return (
    <QueryClientProvider client={client}>
      <NextIntlClientProvider
        timeZone={timeZone}
        locale={locale}
        messages={messages}
      >
        {children}
      </NextIntlClientProvider>
      <ToastContainer />
    </QueryClientProvider>
  );
};

export default Providers;
