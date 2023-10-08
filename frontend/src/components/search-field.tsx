"use client";

import { useCallback, useRef } from "react";
import { useRouter, useSearchParams } from "next/navigation";
import { PlusIcon, XIcon } from "lucide-react";
import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from "./ui/accordion";
import { Button } from "./ui/button";
import { Input } from "./ui/input";

type SearchFieldProps = {
  field: string;
  label: string;
  placeholder?: string;
};

const SearchField = ({ field, label, placeholder }: SearchFieldProps) => {
  const searchParams = useSearchParams();
  const queries = searchParams.getAll(field);
  const queryCount = queries.length;

  const router = useRouter();
  const inputRef = useRef<HTMLInputElement>(null);

  const addQuery = useCallback(
    (query: string) => {
      const newSearchParams = new URLSearchParams(searchParams);
      newSearchParams.append(field, query);
      router.replace(`?${newSearchParams.toString()}`);
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

  return (
    <Accordion type="multiple" defaultValue={[field]}>
      <AccordionItem value={field}>
        <AccordionTrigger className="px-1 pb-2 pt-2 text-sm">
          {label}
        </AccordionTrigger>
        <AccordionContent>
          <div className="flex flex-col gap-1 px-1 pt-1">
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
            <form
              className="flex items-center gap-3 text-sm"
              key={`${field}-${queryCount}`}
              onSubmit={(e) => {
                e.preventDefault();
                if (inputRef.current?.value) {
                  addQuery(inputRef.current.value);
                }
              }}
            >
              {queryCount + 1}.
              <Input
                className="h-8 flex-1 border px-2"
                placeholder={placeholder}
                ref={inputRef}
              />
              <Button className="h-8 w-8 p-0" type="submit">
                <PlusIcon className="h-4 w-4" />
              </Button>
            </form>
          </div>
        </AccordionContent>
      </AccordionItem>
    </Accordion>
  );
};

export { SearchField };
