import { MailIcon } from "lucide-react";
import { type Company } from "@/types/base/company";

/**Dummy Company profile */

const CompanyProfileCard = ({ companyJson }: { companyJson: Company }) => {
  return (
    <div className="m-4 h-[555px] w-[520px] rounded-md bg-background">
      <div className="items-left m-3 flex  flex-col gap-4">
        <div className="mb-2 flex flex-col">
          <div className="flex h-full w-20  justify-start">Name</div>
          <div className="my-1 flex w-full rounded-lg border-2 border-solid border-slate-500 px-1">
            {companyJson.name}
          </div>
        </div>
        <div className="mb-2 flex flex-col">
          <div className="flex h-full w-20 justify-start">Category</div>
          <div className="my-1 flex w-full rounded-lg border-2 border-solid border-slate-500 px-1">
            {companyJson.category}
          </div>
        </div>
        <div className="mb-2 flex flex-col">
          <div className="flex h-full w-20 justify-start">Address</div>
          <div className="my-1 flex w-full rounded-lg border-2 border-solid border-slate-500 px-1">
            {companyJson.location}
          </div>
        </div>
        <div className="mb-2 flex flex-col">
          <div className="flex h-full w-20 justify-start">Tel.</div>
          <div className="my-1 flex w-full rounded-lg border-2 border-solid border-slate-500 px-1">
            {companyJson.phone}
          </div>
        </div>
        <div className="mb-2 flex flex-col">
          <div className="flex h-full w-20 justify-start">Email</div>
          <div className="my-1 flex w-full  items-center rounded-lg border-2 border-solid border-slate-500 px-1">
            <MailIcon className="m-2 h-3.5 w-3.5 opacity-50 " />
            {companyJson.email}
          </div>
        </div>
        <div className="mb-2 flex flex-col">
          <div className="flex h-full w-20 justify-start">Description</div>
          <div className="my-1 flex w-full rounded-lg border-2 border-solid border-slate-500 px-1">
            {companyJson.description}
          </div>
        </div>
      </div>
    </div>
  );
};

export { CompanyProfileCard };
