import { NextResponse, type NextRequest } from "next/server";

export default function middleware(req: NextRequest) {
  const { pathname } = req.nextUrl;

  if (pathname === "/dashboard") {
    return NextResponse.redirect(new URL("/dashboard/posts", req.url));
  }

  return NextResponse.next();
}

export const config = {
  matcher: ["/dashboard"],
};
