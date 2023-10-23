import { notFound } from "next/navigation";
import { getServerSession } from "@/lib/auth";
import { DashboardNav } from "@/components/dashboard-nav";

type LayoutProps = {
  children: React.ReactNode;
};

export default async function Layout({ children }: LayoutProps) {
  const session = await getServerSession();

  if (!session) notFound();
  return (
    <main className="container relative flex flex-1 gap-12">
      <aside className="sticky top-[5.5rem] h-[calc(100vh-5.5rem)] w-[14rem]">
        <DashboardNav user={session.user} />
      </aside>
      <div className="flex flex-1 flex-col">{children}</div>
    </main>
  );
}
