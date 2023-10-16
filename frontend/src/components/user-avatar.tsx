import { type AvatarProps } from "@radix-ui/react-avatar";
import { UserIcon } from "lucide-react";
import { type User } from "@/types/base/user";
import { Avatar, AvatarFallback } from "@/components/ui/avatar";

type UserAvatarProps = AvatarProps & {
  user: User;
};

const UserAvatar = ({ user, ...props }: UserAvatarProps) => {
  return (
    <Avatar {...props}>
      <AvatarFallback>
        <span className="sr-only">{user.name}</span>
        <UserIcon className="h-4 w-4" />
      </AvatarFallback>
    </Avatar>
  );
};

export { UserAvatar };
