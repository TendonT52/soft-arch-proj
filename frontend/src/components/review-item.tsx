import Link from "next/link";
import { ArrowUpRightFromCircleIcon, StarIcon } from "lucide-react";
import { type Review } from "@/types/base/review";
import { formatDate } from "@/lib/utils";
import { Button } from "./ui/button";

type ReviewItemProps = {
  review: Review & {
    id: string;
    updatedAt: string;
    company: {
      id: string;
      name: string;
    };
  };
};

const ReviewItem = ({ review }: ReviewItemProps) => {
  return (
    <div className="flex items-center justify-between p-4">
      <div className="flex flex-col items-start gap-1">
        <div className="flex">
          <Link
            href={`/companies/${review.company.id}`}
            className="font-semibold hover:underline"
          >
            {review.company.name}
          </Link>
          <div className="flex select-none items-center">
            &nbsp;
            <StarIcon className="h-4 w-4" />
            {review.rating}
          </div>
        </div>
        <p className="text-sm text-muted-foreground">
          {formatDate(parseInt(review.updatedAt) * 1000)}
        </p>
      </div>
      <Button variant="ghost" className="h-8 w-8 rounded-md p-0" asChild>
        <Link href={`/companies/${review.company.id}`}>
          <ArrowUpRightFromCircleIcon className="h-4 w-4" />
        </Link>
      </Button>
    </div>
  );
};

export { ReviewItem };
