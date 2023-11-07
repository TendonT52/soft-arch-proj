import { notFound } from "next/navigation";
import { getCompanyMe } from "@/actions/get-company-me";
import { getStudentMe } from "@/actions/get-student-me";
import { UserRole } from "@/types/base/user";
import { getServerSession } from "@/lib/auth";
import { CompanySettingsForm } from "@/components/company-settings-form";
import { StudentSettingsForm } from "@/components/student-settings-form";

export default async function Page() {
  const session = await getServerSession();
  if (!session) notFound();

  if (session.user.role === UserRole.Company) {
    const { company } = await getCompanyMe(session.accessToken);
    return (
      <div className="flex flex-col gap-8">
        <div className="flex flex-col items-start gap-1">
          <h1 className="text-3xl font-bold tracking-tight">Settings</h1>
          <p className="text-lg text-muted-foreground">
            Manage your company information
          </p>
        </div>
        <CompanySettingsForm company={company} />
      </div>
    );
  }

  if (session.user.role === UserRole.Student) {
    const { student } = await getStudentMe(session.accessToken);
    return (
      <div className="flex flex-col gap-8">
        <div className="flex flex-col items-start gap-1">
          <h1 className="text-3xl font-bold tracking-tight">Settings</h1>
          <p className="text-lg text-muted-foreground">
            Manage your personal information
          </p>
        </div>
        <StudentSettingsForm student={student} />
      </div>
    );
  }

  notFound();
}
