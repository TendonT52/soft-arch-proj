import Link from "next/link";
import { getPostsMe } from "@/actions/get-posts-me";
import { PlusIcon } from "lucide-react";
import { getServerSession } from "@/lib/auth";
import { Button } from "@/components/ui/button";
import { PostItem } from "@/components/post-item";

export default async function Page() {
  const session = await getServerSession();
  if (!session) return <></>;

  const { posts } = await getPostsMe(session.accessToken);
  return (
    <div className="flex flex-col gap-8">
      <div className="flex items-center justify-between gap-8">
        <div className="flex flex-col gap-1">
          <h1 className="text-3xl font-bold tracking-tight">Posts</h1>
          <p className="text-lg text-muted-foreground">
            Create and manage posts
          </p>
        </div>
        <Button asChild>
          <Link href="/editor/1">
            <PlusIcon className="mr-2 h-4 w-4" />
            New post
          </Link>
        </Button>
      </div>
      <div className="divide-y rounded-md border">
        {posts.map((post, idx) => (
          <PostItem key={idx} post={post} />
        ))}
      </div>
    </div>
  );
}
