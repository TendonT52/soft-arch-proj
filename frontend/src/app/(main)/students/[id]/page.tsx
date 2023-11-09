import Image from "next/image";
import { notFound } from "next/navigation";
import { getStudent } from "@/actions/get-student";
import { getServerSession } from "@/lib/auth";
import { StudentProfileCard } from "@/components/student-profile-card";

export default async function Page({ params }: { params: { id: string } }) {
  /**Dummy for profile */
  const session = await getServerSession();
  if (!session) notFound();
  const { student } = await getStudent(params.id, session.accessToken);

  return (
    <div className="flex flex-col items-center justify-center">
      <div className="flex h-[150px] w-[157px] justify-center rounded-lg bg-primary">
        <div className="flex items-center justify-center ">
          <Image
            src="/images/profile-pic-mock.png"
            alt="Profile Picture"
            width={0}
            height={0}
            sizes="100vw"
            className=" h-[120px] w-[127px] items-center rounded-full border-4 bg-white"
          />
        </div>
      </div>
      <div className="my-5 text-lg">
        {student.name}
        {"'"}s Profile Detail
      </div>
      <div className="m-3 h-[587px] w-[550px] rounded-lg bg-primary">
        <StudentProfileCard studentJson={student} />
      </div>
    </div>
  );
}
