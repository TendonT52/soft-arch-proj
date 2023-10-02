import { PostCard } from "./post-card";

/* DUMMY */
type Post = {
  topic: string;
  period: string;
  positions: string[];
  skills: string[];
  benefits: string[];
};
/* DUMMY */

type PostPanelProps = {
  posts: Post[];
};

const PostPanel = ({ posts }: PostPanelProps) => {
  return (
    <div className="grid grid-cols-1 gap-6 md:grid-cols-2 xl:grid-cols-3">
      {posts.map((post, idx) => (
        <PostCard key={idx} post={post} />
      ))}
    </div>
  );
};

export { PostPanel };
