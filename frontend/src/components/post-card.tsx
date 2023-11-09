import Link from "next/link";
import { ClockIcon, MapPinnedIcon } from "lucide-react";
import { type Post } from "@/types/base/post";
import { formatDate } from "@/lib/utils";
import { Badge } from "./ui/badge";
import { Button } from "./ui/button";

type PostCardProps = {
  post: Post & {
    owner: {
      id: string;
      name: string;
    };
    updatedAt: string;
    postId: string;
  };
};

const roundRobin = <T,>(...arrays: T[][]): T[] => {
  const maxLength = Math.max(...arrays.map((arr) => arr.length));
  return Array.from({ length: maxLength }, (_, i) =>
    arrays.reduce((acc, cur) => {
      const val = cur[i];
      if (val) acc.push(val);
      return acc;
    }, [])
  ).flat();
};

const PostCard = ({ post }: PostCardProps) => {
  const {
    topic,
    openPositions,
    requiredSkills,
    benefits,
    owner,
    updatedAt,
    postId,
  } = post;
  return (
    <div className="flex justify-between rounded-lg border bg-card text-card-foreground shadow-sm">
      <div className="p-6">
        <h3 className="mb-2 text-lg font-semibold tracking-tight">{topic}</h3>
        <div className="mb-4 flex gap-4">
          <div className="flex items-center">
            <MapPinnedIcon className="mr-1.5 h-3 w-3 opacity-50" />
            <Link
              className="text-sm leading-none text-muted-foreground hover:underline hover:underline-offset-2"
              href={`/companies/${owner.id}`}
            >
              {owner.name}
            </Link>
          </div>
          <div className="flex items-center">
            <ClockIcon className="mr-1.5 h-3 w-3 opacity-50" />
            <span className="text-sm leading-none text-muted-foreground">
              {formatDate(parseInt(updatedAt) * 1000)}
            </span>
          </div>
        </div>
        <div className="flex gap-2">
          {roundRobin(openPositions, requiredSkills, benefits).map((skill) => (
            <Badge
              key={`badge${skill}`}
              className="px-2 font-medium"
              variant="secondary"
            >
              {skill}
            </Badge>
          ))}
        </div>
      </div>
      <div className="flex flex-col justify-center gap-2 p-6">
        <Button asChild>
          <Link href={`/posts/${postId}`}>View</Link>
        </Button>
      </div>
    </div>
  );
};

export { PostCard };
