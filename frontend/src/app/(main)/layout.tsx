import { Header } from "@/components/header";

type LayoutProps = {
  student: React.ReactNode;
  company: React.ReactNode;
};

export default function Layout({ student }: LayoutProps) {
  return (
    <div className="flex min-h-screen flex-col">
      <Header />
      {student}
    </div>
  );
}
