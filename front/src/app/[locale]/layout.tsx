import "../globals.scss";
import type { Metadata } from "next";
import Providers from "../registries/Providers";
import StyledRegistry from "../registries/Styled.registry";
import { notFound } from "next/navigation";
import { Main } from "@/components/global";

interface RootLayoutProps {
  children: React.ReactNode;
  params: {
    locale?: string;
  };
}

export const metadata: Metadata = {
  title: "Newsletter",
  description: "Newsletter",
};

const locales = ["es"];

export default async function RootLayout({
  children,
  params: { locale },
}: RootLayoutProps) {
  const isValidLocale = locales.some((cur) => cur === locale);
  if (!isValidLocale) notFound();

  let messages;
  try {
    messages = (await import(`../../../translations/${locale}.json`)).default;
  } catch {
    notFound();
  }

  return (
    <html lang={locale}>
      <head>
        <link
          rel="preconnect"
          href="https://fonts.googleapis.com"
          crossOrigin="anonymous"
        />
        <link
          rel="preconnect"
          href="https://fonts.gstatic.com"
          crossOrigin="anonymous"
        />
        <link
          href="https://fonts.googleapis.com/css2?family=Alkatra:wght@400;600&display=swap"
          rel="stylesheet"
          crossOrigin="anonymous"
        />
        <link
          href="https://fonts.googleapis.com/css2?family=Caveat:wght@400..700&display=swap"
          rel="stylesheet"
          crossOrigin="anonymous"
        />
      </head>
      <StyledRegistry>
        <body>
          <Providers locale={locale} messages={messages}>
            <Main>{children}</Main>
          </Providers>
        </body>
      </StyledRegistry>
    </html>
  );
}
