import Link from "next/link";
import { ChevronLeftIcon } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Editor } from "@/components/lexical/editor";

export default function Page() {
  return (
    <div className="container relative flex min-h-screen gap-6 py-6">
      <div className="container fixed left-0 right-0 top-6 flex justify-between gap-12">
        <Link href="/dashboard/posts">
          <Button variant="ghost">
            <ChevronLeftIcon className="mr-2 h-4 w-4" />
            Back
          </Button>
        </Link>
      </div>
      <Editor />
    </div>
  );
}
