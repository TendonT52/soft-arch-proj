import { Toaster } from "@/components/ui/toaster";
import { Footer } from "@/components/footer";
import { Header } from "@/components/header";

type LayoutProps = {
  children: React.ReactNode;
};

export default function Layout({ children }: LayoutProps) {
  return (
    <div className="flex min-h-screen flex-col gap-6">
      <Header />
      {children}
      <Footer />
      <Toaster />
    </div>
  );
}
