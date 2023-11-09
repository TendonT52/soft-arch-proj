import { notFound } from "next/navigation";
import getReports from "@/actions/get-reports";
import { UserRole } from "@/types/base/user";
import { getServerSession } from "@/lib/auth";
import { ReportItem } from "@/components/report-items";

export default async function Page() {
  const session = await getServerSession();
  if (!session || session.user.role !== UserRole.Admin) notFound();
  const { reports = [] } = await getReports();

  return (
    <div className="flex flex-col gap-8">
      <div className="flex items-center justify-between gap-8">
        <div className="flex flex-col gap-1">
          <h1 className="text-3xl font-bold tracking-tight">
            Hello {session.user.name}!
          </h1>
          <p className="text-lg text-muted-foreground">List of reports</p>
        </div>
      </div>
      {reports.length === 0 ? (
        <p>No reports.</p>
      ) : (
        <div className="divide-y rounded-md border">
          {reports.map((report, idx) => (
            <ReportItem key={idx} report={report} />
          ))}
        </div>
      )}
    </div>
  );
}
