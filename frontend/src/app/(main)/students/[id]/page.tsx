import Image from "next/image";
import { PenSquare } from "lucide-react";
import { StudentProfileCard } from "@/components/student-profile-card";

export default function Page({ params }: { params: { id: string } }) {
  /**Dummy for profile */
  type StudentProfile = {
    firstName: string;
    lastName: string;
    faculty: string;
    major: string;
    year: string;
    email: string;
    description: string;
    profileImagePath: string;
  };
  const mockProfileRepo = new Map<string, StudentProfile>();
  mockProfileRepo.set("001", {
    firstName: "firstName001",
    lastName: "lastName001",
    faculty: "Catgineering",
    major: "Depression",
    year: "4th",
    email: "iAmDepress@dawg.chula.com",
    description: "I am just a cat. I am not delicious.",
    profileImagePath: "/images/profile-pic-mock-sadge.png",
  });

  return (
    <div className="flex flex-col items-center justify-center">
      <div className="h-[150px] w-[157px]   rounded-lg bg-primary text-center text-lg">
        <div className="flex flex-row-reverse">
          <PenSquare className="text-white" />
        </div>
        <div className="flex items-center justify-center ">
          <Image
            src={mockProfileRepo.get(params.id)?.profileImagePath ?? ""}
            alt="Profile Picture"
            width={0}
            height={0}
            sizes="100vw"
            className=" h-[120px] w-[127px] rounded-full bg-white"
          />
        </div>
      </div>
      <div className="my-5 text-lg">Profile Details</div>
      <div className="m-3 h-[587px] w-[550px] rounded-lg bg-primary">
        <StudentProfileCard studentJson={mockProfileRepo.get(params.id)!} />
      </div>
    </div>
  );
}

// eslint-disable-next-line @typescript-eslint/require-await
export function generateStaticParams() {
  return [{ id: "001" }, { id: "002" }, { id: "003" }];
}
