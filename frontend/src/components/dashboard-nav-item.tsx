"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";
import { cn } from "@/lib/utils";

type DashboardNavItemProps = {
  href: string;
  children: React.ReactNode;
};

const DashboardNavItem = ({ href, children }: DashboardNavItemProps) => {
  const pathname = usePathname();
  return (
    <Link
      className={cn(
        "flex h-9 items-center rounded-md px-3 py-2 text-sm font-medium hover:bg-accent hover:text-accent-foreground",
        pathname === href && "bg-accent"
      )}
      href={href}
    >
      {children}
    </Link>
  );
};

export { DashboardNavItem };
