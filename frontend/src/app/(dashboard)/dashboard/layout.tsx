import { DashboardNav } from "@/components/dashboard-nav";

type LayoutProps = {
  children: React.ReactNode;
};

export default function Layout({ children }: LayoutProps) {
  return (
    <main className="container relative flex flex-1 items-start gap-12">
      <aside className="sticky top-[5.5rem] h-[calc(100vh-5.5rem)] w-[14rem]">
        <DashboardNav />
      </aside>
      <div className="flex-1">{children}</div>
    </main>
  );
}
