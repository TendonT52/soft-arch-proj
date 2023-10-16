import Link from "next/link";
import { ClockIcon, MapPinnedIcon } from "lucide-react";
import { Badge } from "./ui/badge";
import { Button } from "./ui/button";

/* DUMMY */
type Post = {
  topic: string;
  period: string;
  positions: string[];
  skills: string[];
  benefits: string[];
};
/* DUMMY */

type PostCardProps = {
  post: Post;
};

const PostCard = ({ post }: PostCardProps) => {
  const { topic, positions } = post;
  return (
    <div className="flex justify-between rounded-lg border bg-card text-card-foreground shadow-sm">
      <div className="p-6">
        <h3 className="mb-2 text-lg font-semibold tracking-tight">{topic}</h3>
        <div className="mb-4 flex gap-4">
          <div className="flex items-center">
            <MapPinnedIcon className="mr-1.5 h-3 w-3 opacity-50" />
            <Link
              className="text-sm leading-none text-muted-foreground underline underline-offset-2"
              href="/companies/1"
            >
              Umbrella Corporation
            </Link>
          </div>
          <div className="flex items-center">
            <ClockIcon className="mr-1.5 h-3 w-3 opacity-50" />
            <span className="text-sm leading-none text-muted-foreground">
              August 2023
            </span>
          </div>
        </div>
        <div className="flex gap-2">
          {positions.map((skill) => (
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
          <Link href="/posts/1">View</Link>
        </Button>
      </div>
    </div>
  );
};

export { PostCard };
