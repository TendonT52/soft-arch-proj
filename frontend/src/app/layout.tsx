import { type Metadata } from "next";
import localFont from "next/font/local";
import { cn } from "@/lib/utils";
import { TailwindIndicator } from "@/components/tailwind-indicator";
import { ThemeProvider } from "@/components/theme-provider";
import "./globals.css";
import { Toaster } from "@/components/ui/toaster";

const fontSans = localFont({
  src: [
    { path: "../../public/fonts/Inter.var.woff2", style: "normal" },
    { path: "../../public/fonts/Inter-italic.var.woff2", style: "italic" },
  ],
  variable: "--font-sans",
});

const fontMono = localFont({
  src: "../../public/fonts/Jack-Regular.woff2",
  preload: true,
  variable: "--font-mono",
});

export const metadata: Metadata = {
  title: "InternWiseHub",
  description: "Your gateway to professional growth",
};

type RootLayoutProps = {
  children: React.ReactNode;
};

export default function RootLayout({ children }: RootLayoutProps) {
  return (
    <html lang="en">
      <link
        rel="icon"
        href="/images/intern-wise-hub.png"
        type="image/png"
        sizes="144x174"
      />
      <body
        className={cn(
          "min-h-screen bg-background font-sans antialiased",
          fontSans.variable,
          fontMono.variable
        )}
      >
        <ThemeProvider attribute="class" defaultTheme="system" enableSystem>
          {children}
          <Toaster />
          <TailwindIndicator />
        </ThemeProvider>
      </body>
    </html>
  );
}
