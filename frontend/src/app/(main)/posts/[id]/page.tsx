import { notFound } from "next/navigation";
import { getPost } from "@/actions/get-post";
import { getServerSession } from "@/lib/auth";
import { PostViewer } from "@/components/post-viewer";

/* DUMMY */
type PageProps = {
  params: {
    id: string;
  };
};

export default async function Page({ params }: PageProps) {
  const postId = params.id;
  const session = await getServerSession();
  if (!session) notFound();

  const { post } = await getPost(postId);
  if (!post) notFound();

  return <PostViewer post={{ ...post, postId }} />;
}
