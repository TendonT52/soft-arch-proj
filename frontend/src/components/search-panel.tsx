import * as React from "react";
import { SearchIcon } from "lucide-react";
import { SearchField } from "./search-field";
import { Card, CardContent, CardHeader, CardTitle } from "./ui/card";

type SearchPanelProps = {
  postCount: number;
};

const SearchPanel = ({ postCount }: SearchPanelProps) => {
  return (
    <div className="flex h-full flex-col">
      <Card>
        <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
          <CardTitle className="text-base font-medium">Search posts</CardTitle>
          <SearchIcon className="h-4 w-4 opacity-50" />
        </CardHeader>
        <CardContent>
          <p className="text-xs text-muted-foreground">
            <span className="text-2xl font-bold text-foreground">
              {postCount}
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
          field="positions"
          label="Open positions"
          placeholder="Social Engineer"
        />
        <SearchField
          field="skills"
          label="Required skills"
          placeholder="Phishing"
        />
        <SearchField
          field="benefits"
          label="Expected benefits"
          placeholder="Millions"
        />
      </div>
    </div>
  );
};

export { SearchPanel };
