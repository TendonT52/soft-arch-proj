import { notFound } from "next/navigation";
import { getReviews } from "@/actions/get-reviews";
import { UserRole } from "@/types/base/user";
import { getServerSession } from "@/lib/auth";
import { ReviewItem } from "@/components/review-item";

export default async function Page() {
  const session = await getServerSession();
  if (!session || session.user.role !== UserRole.Student) notFound();

  const { reviews = [] } = await getReviews();
  return (
    <div className="flex flex-col gap-8">
      <div className="flex flex-col items-start gap-1">
        <h1 className="text-3xl font-bold tracking-tight">Reviews</h1>
        <p className="text-lg text-muted-foreground">
          List and manage company reviews
        </p>
      </div>
      {reviews.length === 0 ? (
        <p>No reviews.</p>
      ) : (
        <div className="divide-y rounded-md border">
          {reviews.map((review, idx) => (
            <ReviewItem key={idx} review={review} />
          ))}
        </div>
      )}
    </div>
  );
}
