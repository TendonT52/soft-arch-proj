import Link from "next/link";
import {
  FileTextIcon,
  StarIcon,
  UserIcon,
  type LucideIcon,
} from "lucide-react";

type NavItem = {
  Icon: LucideIcon;
  title: string;
  href: string;
  type?: "company" | "student";
};

const navItems: NavItem[] = [
  {
    Icon: FileTextIcon,
    title: "Posts",
    href: "/dashboard/posts",
    type: "company",
  },
  { Icon: StarIcon, title: "Reviews", href: "/dashboard/reviews" },
  { Icon: UserIcon, title: "Account", href: "/dashboard/account" },
];

const DashboardNav = () => {
  const userType = "company";

  return (
    <div className="flex h-full flex-col gap-2">
      {navItems.map(
        (item) =>
          (!item.type || item.type === userType) && (
            <Link
              className="flex h-9 items-center rounded-md px-3 py-2 text-sm font-medium hover:bg-accent hover:text-accent-foreground"
              href={item.href}
              key={item.href}
            >
              <item.Icon className="mr-2 h-4 w-4" />
              {item.title}
            </Link>
          )
      )}
    </div>
  );
};

export { DashboardNav };
