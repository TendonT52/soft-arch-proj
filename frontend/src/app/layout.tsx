import { type Metadata } from "next";
import localFont from "next/font/local";
import { Indicator } from "@/components/indicator";
import { cn } from "@/lib/utils";
import "./globals.css";

const fontSans = localFont({
  src: [
    { path: "../../public/fonts/Inter.var.woff2", style: "normal" },
    { path: "../../public/fonts/Inter-italic.var.woff2", style: "italic" },
  ],
  variable: "--font-sans",
});

export const metadata: Metadata = {
  description: "Generated by create next app",
  title: "Create Next App",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body
        className={cn(
          "bg-background min-h-screen font-sans antialiased",
          fontSans.variable
        )}
      >
        {children}
        <Indicator />
      </body>
    </html>
  );
}
