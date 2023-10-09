import Link from "next/link";
import { PlusIcon } from "lucide-react";
import { Button } from "@/components/ui/button";
import { PostItem } from "@/components/post-item";

/* DUMMY */
type Post = {
  topic: string;
  period: string;
  positions: string[];
  skills: string[];
  benefits: string[];
};

const posts: Post[] = [
  {
    topic: "Social Engineering",
    period: "Summer 2022",
    positions: ["Software Developer", "Data Analyst"],
    skills: ["Python", "SQL", "Data Analysis"],
    benefits: ["Healthcare", "Flexible Work Hours"],
  },
  {
    topic: "Marketing",
    period: "Summer 2022",
    positions: ["Marketing Coordinator", "Social Media Manager"],
    skills: [
      "Digital Marketing",
      "Content Creation",
      "Social Media Management",
    ],
    benefits: ["401(k) Matching", "Paid Time Off"],
  },
  {
    topic: "Business",
    period: "Summer 2022",
    positions: ["Mechanical Engineer", "Product Designer"],
    skills: ["CAD Design", "Mechanical Engineering", "Product Prototyping"],
    benefits: ["Dental Insurance", "Tuition Reimbursement"],
  },
];

const getPosts = () => posts;
/* DUMMY */

export default function Page() {
  const posts = getPosts();

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
