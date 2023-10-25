import { notFound } from "next/navigation";
import { getPost } from "@/actions/get-post";
import { getServerSession } from "@/lib/auth";
import { PostEditor } from "@/components/post-editor";

type PageProps = {
  params: {
    id: string;
  };
};

export default async function Page({ params }: PageProps) {
  const session = await getServerSession();
  if (!session) return notFound();

  const postId = params.id;
  const { status, post } = await getPost(postId);
  if (status !== "200" || post?.owner.id !== session.user.id) notFound();

  return <PostEditor post={{ ...post, postId }} />;
}
