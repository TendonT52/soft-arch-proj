import { verify } from "@/actions/verify";

export default async function Page({
  params,
}: {
  params: { code: Array<string>; id: string };
}) {
  const response = await verify({
    code: decodeURIComponent(params.code.join("/")),
    studentId: params.id,
  });
  if (response.status === "200") {
    return (
      <div className="flex justify-center">
        <div className="flex h-[440px] w-[440px] flex-col items-center justify-center rounded-lg bg-primary">
          <div className=" m-5 flex h-[400px] w-[400px] items-center justify-center bg-white text-xl text-primary">
            {response.message}
          </div>
        </div>
      </div>
    );
  } else {
    return (
      <div className="flex justify-center">
        <div className="flex h-[440px] w-[440px] flex-col items-center justify-center rounded-lg bg-primary">
          <div className=" m-5 flex h-[400px] w-[400px] items-center justify-center bg-white text-xl">
            <h1 className="text-xl text-destructive">Failed to verify email</h1>
            <h2 className="text-xl">Please try again</h2>
          </div>
        </div>
      </div>
    );
  }
}
