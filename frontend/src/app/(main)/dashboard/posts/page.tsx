import { notFound } from "next/navigation";
import { getPostsMe } from "@/actions/get-posts-me";
import { UserRole } from "@/types/base/user";
import { getServerSession } from "@/lib/auth";
import { PostCreateDialog } from "@/components/post-create-dialog";
import { PostItem } from "@/components/post-item";

export default async function Page() {
  const session = await getServerSession();
  if (!session || session.user.role !== UserRole.Company) notFound();

  const { posts = [] } = await getPostsMe();
  return (
    <div className="flex flex-col gap-8">
      <div className="flex items-center justify-between gap-8">
        <div className="flex flex-col gap-1">
          <h1 className="text-3xl font-bold tracking-tight">Posts</h1>
          <p className="text-lg text-muted-foreground">
            Create and manage posts
          </p>
        </div>
        <PostCreateDialog />
      </div>
      {posts.length === 0 ? (
        <p>No posts.</p>
      ) : (
        <div className="divide-y rounded-md border">
          {posts.map((post, idx) => (
            <PostItem key={idx} post={post} />
          ))}
        </div>
      )}
    </div>
  );
}
