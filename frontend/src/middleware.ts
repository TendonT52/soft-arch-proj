import { NextResponse } from "next/server";
import { getToken } from "next-auth/jwt";
import { withAuth } from "next-auth/middleware";

export default withAuth(
  async function middleware(req) {
    const user = await getToken({ req });
    if (!user) {
      return NextResponse.redirect(new URL("/login", req.nextUrl));
    }

    const { pathname } = req.nextUrl;
    if (pathname === "/dashboard") {
      return NextResponse.redirect(new URL("/dashboard/posts", req.nextUrl));
    }

    return NextResponse.next();
  },
  {
    callbacks: {
      authorized() {
        // This is a work-around for handling redirect on auth pages.
        // We return true here so that the middleware function above
        // is always called.
        return true;
      },
    },
  }
);

export const config = {
  matcher: ["/dashboard/:path*", "/posts"],
};
