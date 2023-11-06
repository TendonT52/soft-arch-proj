import Image from "next/image";
import { notFound } from "next/navigation";
import { getCompany } from "@/actions/get-company";
import { getReviewCompany } from "@/actions/get-review-company";
import { UserRole } from "@/types/base/user";
import { getServerSession } from "@/lib/auth";
import { CompanyProfileCard } from "@/components/company-profile-card";
import { ReviewCard } from "@/components/review-card";
import { ReviewDialog } from "@/components/review-dialog";

export default async function Page({ params }: { params: { id: string } }) {
  const session = await getServerSession();
  if (!session) notFound();

  const { company } = await getCompany(params.id, session.accessToken);
  const { reviews } = await getReviewCompany(params.id);

  return (
    <div className="container flex flex-col items-center justify-center">
      <div className="h-[150px] w-[157px] rounded-lg bg-primary text-center text-lg">
        <div className="flex items-center justify-center ">
          <Image
            src="/images/profile-pic-mock-salad.png"
            alt="Profile Picture"
            width={0}
            height={0}
            sizes="100vw"
            className="my-4 h-[120px] w-[127px] rounded-full bg-white"
          />
        </div>
      </div>
      <div className="m-3 text-lg">Profile Details</div>
      {session.user.role === UserRole.Student && (
        <div className="my-3">
          <ReviewDialog companyId={params.id} />
        </div>
      )}
      <div className="m-3 h-[587px] w-[550px] rounded-lg bg-primary">
        <CompanyProfileCard companyJson={company} />
      </div>
      <div className="flex w-full justify-center pb-4 pt-10">
        <h2 className="text-lg">Company reviews</h2>
      </div>
      <div className="grid w-full auto-cols-fr grid-cols-1 gap-8 md:grid-cols-2 xl:grid-cols-3">
        {!reviews || reviews.length === 0 ? (
          <p className="text-muted-foreground">No reviews</p>
        ) : (
          reviews.map((review) => (
            <ReviewCard
              key={review.id}
              user={session.user}
              review={review}
              companyId={params.id}
            />
          ))
        )}
      </div>
    </div>
  );
}

export function generateStaticParams() {
  return [{ id: "1" }, { id: "2" }, { id: "3" }];
}
