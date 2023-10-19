import { MailIcon, PenSquare } from "lucide-react";

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

const StudentProfileCard = ({
  studentJson,
}: {
  studentJson: StudentProfile;
}) => {
  return (
    <div className="m-4 h-[555px] w-[520px] rounded-md bg-white">
      <div className="flex flex-row-reverse">
        <PenSquare className="text-black" />
      </div>
      <div className="items-left m-3 flex  flex-col justify-evenly gap-4">
        <div className="mb-2 flex flex-col">
          <div className="flex h-full w-20  justify-start">Name</div>
          <div className="my-1 flex w-full rounded-lg border-2 border-solid border-slate-500 px-1">
            {studentJson.firstName}
          </div>
        </div>
        <div className="mb-2 flex flex-col">
          <div className="flex h-full w-20 justify-start">Lastname</div>
          <div className="my-1 flex w-full rounded-lg border-2 border-solid border-slate-500 px-1">
            {studentJson.lastName}
          </div>
        </div>
        <div className="flex flex-row justify-between">
          <div className="mb-2 flex flex-col">
            <div className="flex h-full w-20 justify-start">Faculty</div>
            <div className="my-1 flex w-full rounded-lg border-2 border-solid border-slate-500 px-1">
              {studentJson.faculty}
            </div>
          </div>
          <div className="mb-2 flex flex-col">
            <div className="flex h-full w-20 justify-start">Major</div>
            <div className="my-1 flex w-full rounded-lg border-2 border-solid border-slate-500 px-1">
              {studentJson.major}
            </div>
          </div>
          <div className="mb-2 flex flex-col">
            <div className="flex h-full w-20 justify-start">Year</div>
            <div className="my-1 flex w-full rounded-lg border-2 border-solid border-slate-500 px-1">
              {studentJson.year}
            </div>
          </div>
        </div>
        <div className="mb-2 flex flex-col">
          <div className="flex h-full w-20 justify-start">Email</div>
          <div className="my-1 flex w-full  items-center rounded-lg border-2 border-solid border-slate-500 px-1">
            <MailIcon className="m-2 h-3.5 w-3.5 opacity-50 " />
            {studentJson.email}
          </div>
        </div>
        <div className="mb-2 flex flex-col">
          <div className="flex h-full w-20 justify-start">Description</div>
          <div className="my-1 flex w-full rounded-lg border-2 border-solid border-slate-500 px-1">
            {studentJson.description}
          </div>
        </div>
      </div>
    </div>
  );
};

export { StudentProfileCard };
