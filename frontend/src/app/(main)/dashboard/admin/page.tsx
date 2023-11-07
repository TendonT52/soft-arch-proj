import { notFound } from "next/navigation";
import { UserRole } from "@/types/base/user";
import { getServerSession } from "@/lib/auth";

export default async function Page() {
  const session = await getServerSession();
  if (!session || session.user.role !== UserRole.Admin) notFound();

  return (
    <div className="flex flex-col gap-8">
      <div className="flex items-center justify-between gap-8">
        <div className="flex flex-col gap-1">
          <h1 className="text-3xl font-bold tracking-tight">
            Hello {session.user.name}!
          </h1>
          <p className="text-lg text-muted-foreground">
            Power is in your hands.
          </p>
        </div>
      </div>
    </div>
  );
}
