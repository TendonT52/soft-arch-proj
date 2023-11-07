import { FileTextIcon, LockIcon, SettingsIcon, StarIcon } from "lucide-react";
import { UserRole, type User } from "@/types/base/user";
import { DashboardNavItem } from "./dashboard-nav-item";

type DashboardNavProps = {
  user: User;
};

const DashboardNav = ({ user }: DashboardNavProps) => {
  return (
    <nav className="flex h-full flex-col gap-2">
      {user.role === UserRole.Company ? (
        <DashboardNavItem href="/dashboard/posts">
          <FileTextIcon className="mr-2 h-4 w-4 shrink-0" />
          Posts
        </DashboardNavItem>
      ) : user.role === UserRole.Student ? (
        <DashboardNavItem href="/dashboard/reviews">
          <StarIcon className="mr-2 h-4 w-4 shrink-0" />
          Reviews
        </DashboardNavItem>
      ) : (
        <DashboardNavItem href="/dashboard/admin">
          <LockIcon className="mr-2 h-4 w-4 shrink-0" />
          Admin
        </DashboardNavItem>
      )}
      <DashboardNavItem href="/dashboard/settings">
        <SettingsIcon className="mr-2 h-4 w-4 shrink-0" />
        Settings
      </DashboardNavItem>
    </nav>
  );
};

export { DashboardNav };
