import { Loading } from "@/components/loading";

export default function Page() {
  return (
    <div className="flex flex-1 flex-col items-start">
      <h1 className="text-3xl font-bold tracking-tight">Account</h1>
      <div className="flex w-full flex-1 items-center justify-center">
        <Loading />
      </div>
    </div>
  );
}
