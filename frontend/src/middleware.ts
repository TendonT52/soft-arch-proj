import { NextResponse } from "next/server";
import { getToken } from "next-auth/jwt";
import { withAuth } from "next-auth/middleware";
import { verifyAccessToken } from "./lib/token";
import { UserRole } from "./types/base/user";

export default withAuth(
  async function middleware(req) {
    const token = await getToken({ req });
    if (!token) return NextResponse.next();

    const accessToken = verifyAccessToken(token.accessToken);

    if (["/dashboard"].includes(req.nextUrl.pathname)) {
      switch (accessToken.role) {
        case UserRole.Company:
          return NextResponse.redirect(
            new URL("/dashboard/posts", req.nextUrl)
          );
        case UserRole.Student:
          return NextResponse.redirect(
            new URL("/dashboard/reviews", req.nextUrl)
          );
        case UserRole.Admin:
          return NextResponse.redirect(
            new URL("/dashboard/admin", req.nextUrl)
          );
      }
    }

    const isAuthPath = [
      "/login",
      "/register/company",
      "/register/student",
    ].includes(req.nextUrl.pathname);

    if (isAuthPath) {
      return NextResponse.redirect(new URL("/", req.nextUrl));
    }

    return NextResponse.next();
  },
  {
    callbacks: {
      authorized() {
        return true;
      },
    },
  }
);

export const config = {
  matcher: ["/dashboard", "/login", "/register/company", "/register/student"],
};
