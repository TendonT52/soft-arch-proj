"use client";

import { useCallback, useEffect, useRef, useState } from "react";
import { useRouter, useSearchParams } from "next/navigation";
import { getBenefits } from "@/actions/get-benefits";
import { getOpenPositions } from "@/actions/get-open-positions";
import { getRequiredSkills } from "@/actions/get-required-skills";
import {
  Accordion,
  AccordionContent,
  AccordionHeader,
  AccordionItem,
  AccordionTrigger,
} from "@radix-ui/react-accordion";
import { Command } from "cmdk";
import { ChevronDownIcon, PlusIcon, XIcon } from "lucide-react";
import { PostField } from "@/types/base/post";
import { UserRole } from "@/types/base/user";
import { useDebounce } from "@/hooks/use-debounce";
import { useOnClickOutside } from "@/hooks/use-on-click-outside";
import { Button } from "./ui/button";

type SearchFieldProps = {
  field: PostField | string;
  label: string;
  placeholder?: string;
  userRole: UserRole;
};

const SearchField = ({
  field,
  label,
  placeholder,
  userRole,
}: SearchFieldProps) => {
  const router = useRouter();
  const inputRef = useRef<HTMLInputElement | null>(null);
  const commandRef = useRef<HTMLDivElement | null>(null);

  const searchParams = useSearchParams();
  const queries = searchParams.getAll(field);
  const queryCount = queries.length;

  const [search, setSearch] = useState({ value: "" }); // use object for sensitivity
  const debouncedSearch = useDebounce(search, 250);

  const [suggestions, setSuggestions] = useState<string[]>([]);
  const [shouldUpdateSuggestions, setShouldUpdateSuggestions] = useState(true);

  const addQuery = useCallback(
    (query: string) => {
      const newSearchParams = new URLSearchParams(searchParams);
      if (!newSearchParams.has(field, query)) {
        newSearchParams.append(field, query);
        router.replace(`?${newSearchParams.toString()}`);
      }
    },
    [searchParams, router, field]
  );

  const removeQuery = useCallback(
    (query: string) => {
      const newSearchParams = new URLSearchParams(searchParams);
      newSearchParams.delete(field, query);
      router.replace(`?${newSearchParams.toString()}`);
    },
    [searchParams, router, field]
  );

  useEffect(() => {
    const shouldNotUpdate =
      !debouncedSearch.value ||
      !shouldUpdateSuggestions ||
      userRole !== UserRole.Company;

    if (shouldNotUpdate) {
      setSuggestions([]);
      return;
    }

    const updateSuggestions = async () => {
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

    void updateSuggestions();
  }, [field, shouldUpdateSuggestions, debouncedSearch, userRole]);

  useOnClickOutside(commandRef, () => {
    setSuggestions([]);
  });

  return (
    <Accordion type="multiple" defaultValue={[field]}>
      <AccordionItem className="border-b" value={field}>
        <AccordionHeader className="flex">
          <AccordionTrigger className="flex flex-1 items-center justify-between px-1 pb-2 pt-2 text-sm font-medium transition-all hover:underline [&[data-state=open]>svg]:rotate-180">
            {label}
            <ChevronDownIcon className="h-4 w-4 shrink-0" />
          </AccordionTrigger>
        </AccordionHeader>
        <AccordionContent>
          <div className="flex flex-col gap-2 px-1 pb-4 pt-1 text-sm">
            {queries.map((query, idx) => (
              <div
                className="group flex items-center gap-3 text-sm"
                key={`${field}-${idx}`}
              >
                <div className="tabular-nums">{idx + 1}.</div>
                <div className="flex h-8 flex-1 items-center py-2">{query}</div>
                <button
                  className="hidden items-center justify-center opacity-50 group-hover:inline-flex group-focus:inline-flex"
                  tabIndex={-1}
                  onClick={() => void removeQuery(query)}
                >
                  <XIcon className="h-4 w-4" />
                </button>
              </div>
            ))}
            <Command
              loop
              shouldFilter={false}
              onKeyDown={(e) => {
                switch (e.key) {
                  case "Escape":
                    setSuggestions([]);
                    break;
                  case "Enter":
                    if (suggestions.length === 0 && search.value.trim()) {
                      addQuery(search.value);
                      setSearch({ value: "" });
                    }
                    break;
                }
              }}
              ref={commandRef}
            >
              <div className="flex items-center gap-3 text-sm">
                <div className="tabular-nums">{queryCount + 1}.</div>
                <div className="relative flex flex-1">
                  <Command.Input
                    className="flex h-9 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2  disabled:cursor-not-allowed disabled:opacity-50"
                    placeholder={placeholder}
                    value={search.value}
                    onValueChange={(value) => {
                      setSearch({ value });
                      setShouldUpdateSuggestions(true);
                    }}
                    ref={inputRef}
                  />
                  {suggestions.length !== 0 && (
                    <Command.List className="absolute top-10 z-50 w-full min-w-[8rem] overflow-hidden rounded-md border bg-popover p-1 text-popover-foreground shadow-md data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0 data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95 data-[side=bottom]:slide-in-from-top-2 data-[side=left]:slide-in-from-right-2 data-[side=right]:slide-in-from-left-2 data-[side=top]:slide-in-from-bottom-2">
                      {suggestions.map((value) => {
                        return (
                          <Command.Item
                            key={value}
                            className="relative flex cursor-default select-none items-center rounded-sm px-2 py-1.5 text-sm outline-none aria-selected:bg-accent aria-selected:text-accent-foreground data-[disabled]:pointer-events-none data-[disabled]:opacity-50"
                            value={value}
                            onSelect={() => {
                              setSearch({ value });
                              setSuggestions([]);
                              setShouldUpdateSuggestions(false);
                              inputRef.current?.focus();
                            }}
                          >
                            {value}
                          </Command.Item>
                        );
                      })}
                    </Command.List>
                  )}
                </div>
                <Button
                  className="h-9 w-9 p-0"
                  onClick={() => {
                    if (search.value.trim()) {
                      addQuery(search.value);
                      setSearch({ value: "" });
                    }
                  }}
                >
                  <PlusIcon className="h-4 w-4" />
                </Button>
              </div>
            </Command>
          </div>
        </AccordionContent>
      </AccordionItem>
    </Accordion>
  );
};

export { SearchField };
