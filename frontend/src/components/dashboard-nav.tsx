"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";
import {
  FileTextIcon,
  SettingsIcon,
  StarIcon,
  type LucideIcon,
} from "lucide-react";
import { UserRole, type User } from "@/types/base/user";
import { cn } from "@/lib/utils";

type NavItem = {
  Icon: LucideIcon;
  title: string;
  href: string;
  role?: UserRole;
};

const navItems: NavItem[] = [
  {
    Icon: FileTextIcon,
    title: "Posts",
    href: "/dashboard/posts",
    role: UserRole.Company,
  },
  { Icon: StarIcon, title: "Reviews", href: "/dashboard/reviews" },
  { Icon: SettingsIcon, title: "Settings", href: "/dashboard/settings" },
];

type DashboardNavProps = {
  user: User;
};

const DashboardNav = ({ user }: DashboardNavProps) => {
  const pathname = usePathname();

  return (
    <div className="flex h-full flex-col gap-2">
      {navItems.map(
        ({ Icon, title, href, role }) =>
          (!role || role === user.role) && (
            <Link
              className={cn(
                "flex h-9 items-center rounded-md px-3 py-2 text-sm font-medium hover:bg-accent hover:text-accent-foreground",
                pathname === href && "bg-accent"
              )}
              href={href}
              key={href}
            >
              <Icon className="mr-2 h-4 w-4" />
              {title}
            </Link>
          )
      )}
    </div>
  );
};

export { DashboardNav };
