import Image from "next/image";
import { notFound } from "next/navigation";
import { getReview } from "@/actions/get-review";
import { PenSquare } from "lucide-react";
import { UserRole } from "@/types/base/user";
import { getServerSession } from "@/lib/auth";
import { CompanyProfileCard } from "@/components/company-profile-card";
import { ReviewCard } from "@/components/review-card";
import { ReviewDialog } from "@/components/review-dialog";

/**Dummy Company profile */
type CompanyProfile = {
  companyName: string;
  category: string;
  location: string;
  phone: string;
  description: string;
  email: string;
  profileImagePath: string;
};

export default async function Page({ params }: { params: { id: string } }) {
  const session = await getServerSession();
  if (!session) notFound();

  /* DUMMY */
  const { review: dummyReview } = await getReview("6");

  const mockProfileRepo = new Map<string, CompanyProfile>();
  mockProfileRepo.set("2", {
    companyName: "RedBik",
    category: "Tech",
    location: "Bangkok",
    phone: "0812223333",
    description: "I am Groot",
    email: "redbik@tech.com",
    profileImagePath: "/images/profile-pic-mock.png",
  });
  return (
    <div className="container flex flex-col items-center justify-center">
      <div className="h-[150px] w-[157px] rounded-lg bg-primary text-center text-lg">
        <div className="flex justify-end">
          <PenSquare className="absolute text-white" />
        </div>
        <div className="flex items-center justify-center ">
          <Image
            src={mockProfileRepo.get(params.id)?.profileImagePath ?? ""}
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
        <CompanyProfileCard companyJson={mockProfileRepo.get(params.id)!} />
      </div>
      <div className="mt-5 grid w-full auto-cols-fr grid-cols-1 gap-8 md:grid-cols-2 xl:grid-cols-3">
        {/* DUMMY, Expect error because we need to know if the user is the owner */}
        <ReviewCard
          user={session.user}
          review={dummyReview}
          companyId={params.id}
        />
        <ReviewCard
          user={session.user}
          review={dummyReview}
          companyId={params.id}
        />
        <ReviewCard
          user={session.user}
          review={dummyReview}
          companyId={params.id}
        />
        <ReviewCard
          user={session.user}
          review={dummyReview}
          companyId={params.id}
        />
      </div>
    </div>
  );
}

export function generateStaticParams() {
  return [{ id: "1" }, { id: "2" }, { id: "3" }];
}
