"use client";

import { useEffect, useMemo, useState } from "react";
import { getBenefits } from "@/actions/get-benefits";
import { getOpenPositions } from "@/actions/get-open-positions";
import { getRequiredSkills } from "@/actions/get-required-skills";
import { Label } from "@radix-ui/react-label";
import { PopoverContent } from "@radix-ui/react-popover";
import { Command } from "cmdk";
import { CheckIcon, SearchIcon } from "lucide-react";
import { PostField } from "@/types/base/post";
import { cn } from "@/lib/utils";
import { useDebounce } from "@/hooks/use-debounce";
import { Badge } from "./ui/badge";
import { Button } from "./ui/button";
import { Popover, PopoverTrigger } from "./ui/popover";
import { Separator } from "./ui/separator";

type PostFieldInputProps = React.ButtonHTMLAttributes<HTMLButtonElement> & {
  id?: string;
  field: PostField;
  tags?: string[];
  onTagsChange?: (tags: string[]) => void;
  readOnly?: boolean;
};

const PostFieldInput = ({
  id,
  field,
  tags,
  onTagsChange,
  className,
  readOnly = false,
  ...props
}: PostFieldInputProps) => {
  const [_tags, _setTags] = useState<Set<string>>(new Set());
  const [suggestions, setSuggestions] = useState<string[]>([]);

  const [search, setSearch] = useState({ value: "" });
  const debouncedSearch = useDebounce(search, 250);

  const toggleTag = (suggestion: string) => {
    const updatedTags = new Set(_tags);
    if (!updatedTags.has(suggestion)) {
      updatedTags.add(suggestion);
    } else {
      updatedTags.delete(suggestion);
    }
    _setTags(updatedTags);
    onTagsChange?.([...updatedTags]);
  };

  const title = useMemo(() => {
    switch (field) {
      case PostField.benefits:
        return "benefits";
      case PostField.openPositions:
        return "open positions";
      case PostField.requiredSkills:
        return "required skills";
    }
  }, [field]);

  const arrangedSuggestions = useMemo(
    () =>
      suggestions.toSorted((a, b) => {
        const x = _tags.has(a) ? 1 : 0;
        const y = _tags.has(b) ? 1 : 0;
        return y - x;
      }),
    [_tags, suggestions]
  );

  useEffect(() => {
    const equals = (tags ?? []).every((tag) => _tags.has(tag));
    if (!equals) _setTags(new Set(tags));
  }, [_tags, tags]);

  useEffect(() => {
    const updateSuggestion = async () => {
      const search = debouncedSearch.value.trim();

      switch (field) {
        case PostField.benefits:
          const { benefits } = await getBenefits(search);
          setSuggestions(benefits ?? []);
          break;
        case PostField.openPositions:
          const { openPositions } = await getOpenPositions(search);
          setSuggestions(openPositions ?? []);
          break;
        case PostField.requiredSkills:
          const { requiredSkills } = await getRequiredSkills(search);
          setSuggestions(requiredSkills ?? []);
          break;
      }
    };
    void updateSuggestion();
  }, [field, debouncedSearch]);

  if (readOnly) {
    return (
      <Button
        id={id}
        className={cn(
          "h-10 justify-start overflow-auto whitespace-nowrap font-normal text-muted-foreground scrollbar-hide hover:cursor-text hover:bg-inherit hover:text-muted-foreground focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2",
          className
        )}
        variant="outline"
        size="sm"
        {...props}
      >
        {_tags.size || "No"}&nbsp;{title}
        {_tags.size !== 0 && (
          <>
            <Separator orientation="vertical" className="mx-2 h-4" />
            <div className="flex gap-1.5">
              {[..._tags].map((tag) => (
                <Badge
                  key={tag}
                  variant="secondary"
                  className="rounded-sm px-1 font-normal"
                >
                  {tag}
                </Badge>
              ))}
            </div>
          </>
        )}
      </Button>
    );
  }
  return (
    <Popover>
      <PopoverTrigger asChild>
        <Button
          id={id}
          className={cn(
            "h-10 justify-start overflow-auto whitespace-nowrap font-normal text-muted-foreground scrollbar-hide",
            className
          )}
          variant="outline"
          size="sm"
          {...props}
        >
          {_tags.size || "No"}&nbsp;{title}
          {_tags.size !== 0 && (
            <>
              <Separator orientation="vertical" className="mx-2 h-4" />
              <div className="flex gap-1.5">
                {[..._tags].map((tag) => (
                  <Badge
                    key={tag}
                    variant="secondary"
                    className="rounded-sm px-1 font-normal"
                  >
                    {tag}
                  </Badge>
                ))}
              </div>
            </>
          )}
        </Button>
      </PopoverTrigger>
      <PopoverContent
        className="z-50 w-[12rem] rounded-md border bg-popover p-0 text-popover-foreground shadow-md outline-none data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0 data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95 data-[side=bottom]:slide-in-from-top-2 data-[side=left]:slide-in-from-right-2 data-[side=right]:slide-in-from-left-2 data-[side=top]:slide-in-from-bottom-2"
        align="center"
        sideOffset={4}
      >
        <Command loop shouldFilter={false}>
          <div className="flex items-center border-b px-3">
            <Label htmlFor={`post-field-${field}`}>
              <SearchIcon className="mr-2 h-4 w-4 shrink-0 opacity-50" />
            </Label>
            <Command.Input
              id={`post-field-${field}`}
              className="flex h-11 w-full rounded-md bg-transparent py-3 text-sm outline-none placeholder:text-muted-foreground disabled:cursor-not-allowed disabled:opacity-50"
              placeholder="Search..."
              value={search.value}
              onValueChange={(value) => void setSearch({ value })}
            />
          </div>
          <Command.List className="h-[10.5rem] overflow-auto p-1 scrollbar-hide">
            {arrangedSuggestions.length === 0 ? (
              <div className="flex h-40 w-full items-center justify-center text-xs text-muted-foreground">
                <p>No suggestions</p>
              </div>
            ) : (
              arrangedSuggestions.map((suggestion) => (
                <Command.Item
                  key={suggestion}
                  className="relative flex cursor-default select-none items-center rounded-sm px-2 py-1.5 text-sm outline-none aria-selected:bg-accent aria-selected:text-accent-foreground data-[disabled]:pointer-events-none data-[disabled]:opacity-50"
                  onSelect={() => void toggleTag(suggestion)}
                >
                  <div
                    className={cn(
                      "mr-2 flex h-4 w-4 items-center justify-center rounded-sm border border-primary",
                      _tags.has(suggestion)
                        ? "bg-primary text-primary-foreground"
                        : "opacity-50 [&_svg]:invisible"
                    )}
                  >
                    <CheckIcon className="h-4 w-4" />
                  </div>
                  {suggestion}
                </Command.Item>
              ))
            )}
          </Command.List>
        </Command>
      </PopoverContent>
    </Popover>
  );
};

export { PostFieldInput };
