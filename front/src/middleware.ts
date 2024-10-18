import { NextRequest } from "next/server";
import createIntlMiddleware from "next-intl/middleware";

const locales = ["es"];

const intlMiddleware = createIntlMiddleware({
  locales,
  defaultLocale: "es",
});

export default function middleware(req: NextRequest) {
  return intlMiddleware(req);
}

export const config = {
  // Skip all paths that should not be internationalized
  matcher: ["/((?!api|_next|.*\\..*).*)"],
};
