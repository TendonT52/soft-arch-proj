import { getPost } from "@/actions/get-post";
import { PostViewer } from "@/components/post-viewer";

/* DUMMY */
type PageProps = {
  params: {
    id: string;
  };
};

export default async function Page({ params }: PageProps) {
  const postId = params.id;
  const { post } = await getPost(postId);
  return <PostViewer post={{ ...post, postId }} />;
}
