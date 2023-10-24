import Image from "next/image";
import { PenSquare } from "lucide-react";
import { CompanyProfileCard } from "@/components/company-profile-card";

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

export default function Page({ params }: { params: { id: string } }) {
  const mockProfileRepo = new Map<string, CompanyProfile>();
  mockProfileRepo.set("001", {
    companyName: "RedBik",
    category: "Tech",
    location: "Bangkok",
    phone: "0812223333",
    description: "I am Groot",
    email: "redbik@tech.com",
    profileImagePath: "/images/profile-pic-mock.png",
  });
  return (
    <div className="flex flex-col items-center justify-center">
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
      <div className="m-3 h-[587px] w-[550px] rounded-lg bg-primary">
        <CompanyProfileCard companyJson={mockProfileRepo.get(params.id)!} />
      </div>
    </div>
  );
}

export function generateStaticParams() {
  return [{ id: "001" }, { id: "002" }, { id: "003" }];
}
