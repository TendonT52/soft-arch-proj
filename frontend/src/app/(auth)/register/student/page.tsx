import Image from "next/image";
import { RegisterStudentForm } from "@/components/register-student-form";

export default function Page() {
  return (
    <div className="flex min-h-screen">
      <div className="hidden flex-1 items-center justify-center bg-primary lg:flex">
        <Image
          className="aspect-[1293/1063] w-1/2"
          src="/images/register-student.png"
          alt="Register Student"
          width={646}
          height={531}
        />
      </div>
      <RegisterStudentForm />
    </div>
  );
}
