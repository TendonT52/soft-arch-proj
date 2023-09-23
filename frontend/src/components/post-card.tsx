import { Button } from "./ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardTitle,
} from "./ui/card";

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
  const { topic, skills } = post;
  return (
    <Card>
      <CardContent className="pt-6">
        <div className="mb-4 h-16 w-16 rounded-md bg-muted shadow-sm"></div>
        <CardTitle className="text-lg">{topic}</CardTitle>
        <CardDescription>{skills.join(", ")}</CardDescription>
      </CardContent>
      <CardFooter>
        <Button size="sm">Contact</Button>
      </CardFooter>
    </Card>
  );
};

export { PostCard };
