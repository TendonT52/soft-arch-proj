import { notFound } from "next/navigation";
import getCompanies from "@/actions/get-companies";
import { UserRole } from "@/types/base/user";
import { getServerSession } from "@/lib/auth";
import { PendingCompany } from "@/components/pending-company";

export default async function PendingPage() {
  const session = await getServerSession();
  if (!session || session.user.role !== UserRole.Admin) notFound();
  const { companies = [] } = await getCompanies();
  return (
    <div className="flex flex-col gap-8">
      <div className="flex items-center justify-between gap-8">
        <div className="flex flex-col gap-1">
          <h1 className="text-3xl font-bold tracking-tight">
            Hello {session.user.name}!
          </h1>
          <p className="text-lg text-muted-foreground">
            List of pending company
          </p>
        </div>
      </div>
      {companies.length === 0 ? (
        <p>บ่มีcompany แล้วอ้ายแอดมิน.</p>
      ) : (
        <div className="divide-y rounded-md border">
          {companies.map((company, idx) =>
            company.status === "Pending" ? (
              <PendingCompany key={idx} company={company} />
            ) : null
          )}
        </div>
      )}
    </div>
  );
}
