import { notFound } from "next/navigation";
import { getPosts } from "@/actions/get-posts";
import { SearchIcon } from "lucide-react";
import { getServerSession } from "@/lib/auth";
import { getSearchArray } from "@/lib/utils";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { PostCard } from "@/components/post-card";
import { SearchField } from "@/components/search-field";

type PageProps = {
  searchParams: {
    companies?: string | string[];
    openPositions?: string | string[];
    requiredSkills?: string | string[];
    benefits?: string | string[];
  };
};

function getSearchOption(searchParam?: string | string[]) {
  return getSearchArray(searchParam).join(" ");
}

export default async function Page({ searchParams }: PageProps) {
  const session = await getServerSession();
  if (!session) notFound();

  const { posts = [] } = await getPosts(undefined, {
    searchCompany: getSearchOption(searchParams.companies),
    searchOpenPosition: getSearchOption(searchParams.openPositions),
    searchRequiredSkill: getSearchOption(searchParams.requiredSkills),
    searchBenefit: getSearchOption(searchParams.benefits),
  });

  return (
    <main className="container relative flex flex-1 items-start gap-12">
      <aside className="sticky top-[5.5rem] h-[calc(100vh-5.5rem)] w-[14rem]">
        <div className="flex h-full flex-col">
          <Card>
            <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
              <CardTitle className="text-base font-medium">
                Search posts
              </CardTitle>
              <SearchIcon className="h-4 w-4 opacity-50" />
            </CardHeader>
            <CardContent>
              <p className="text-xs text-muted-foreground">
                <span className="text-2xl font-bold text-foreground">
                  {posts.length}
                </span>
                &nbsp;posts found
              </p>
            </CardContent>
          </Card>
          <div
            className="flex flex-1 flex-col gap-4 overflow-auto pb-6 pt-6 scrollbar-hide"
            style={{
              maskImage:
                "linear-gradient(to top, transparent 0%, rgb(0, 0, 0) 3rem, rgb(0, 0, 0) calc(100% - 3rem), transparent 100%)",
              WebkitMaskImage:
                "linear-gradient(to top, transparent 0%, rgb(0, 0, 0) 3rem, rgb(0, 0, 0) calc(100% - 3rem), transparent 100%)",
            }}
          >
            <SearchField
              field="companies"
              label="Interested companies"
              placeholder="Umbrella"
            />
            <SearchField
              field="openPositions"
              label="Open positions"
              placeholder="Social engineer"
            />
            <SearchField
              field="requiredSkills"
              label="Required skills"
              placeholder="Phishing"
            />
            <SearchField
              field="benefits"
              label="Benefits"
              placeholder="Millions"
            />
          </div>
        </div>
      </aside>
      <div className="flex flex-1 flex-col gap-6">
        {posts.map((post, idx) => (
          <PostCard key={idx} post={post} />
        ))}
      </div>
    </main>
  );
}
