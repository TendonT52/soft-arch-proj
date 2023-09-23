import * as React from "react";
import { SearchIcon } from "lucide-react";
import { cn } from "@/lib/utils";
import { SearchField } from "./search-field";
import { Card, CardContent, CardHeader, CardTitle } from "./ui/card";

interface SearchPanelProps extends React.HTMLAttributes<HTMLDivElement> {
  postCount: number;
}

const SearchPanel = React.forwardRef<HTMLDivElement, SearchPanelProps>(
  ({ className, postCount, ...props }, ref) => {
    return (
      <aside className={cn("flex flex-col", className)} ref={ref} {...props}>
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
                {postCount}
              </span>
              &nbsp;posts found
            </p>
          </CardContent>
        </Card>
        <div
          className="flex flex-1 flex-col gap-4 overflow-auto pb-6 pt-4 scrollbar-hide"
          style={{
            maskImage:
              "linear-gradient(to top, transparent 0%, rgb(0, 0, 0) 100px, rgb(0, 0, 0) 90%, transparent 100%)",
            WebkitMaskImage:
              "linear-gradient(to top, transparent 0%, rgb(0, 0, 0) 100px, rgb(0, 0, 0) 90%, transparent 100%)",
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
      </aside>
    );
  }
);
SearchPanel.displayName = "SearchPanel";

export { SearchPanel };
