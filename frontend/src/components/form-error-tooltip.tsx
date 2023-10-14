"use client";

import { useState } from "react";
import { BadgeAlertIcon } from "lucide-react";
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from "@/components/ui/tooltip";

interface FormErrorTooltipProps {
  message?: string;
}

export function FormErrorTooltip({ message }: FormErrorTooltipProps) {
  const [open, setOpen] = useState(false);

  if (message === undefined) {
    return <></>;
  }
  return (
    <TooltipProvider>
      <Tooltip delayDuration={200} open={open} onOpenChange={setOpen}>
        <TooltipTrigger asChild>
          <div
            className="flex h-5 w-5 cursor-pointer items-center overflow-hidden  focus-visible:outline-none"
            onClick={() => void setOpen(!open)}
          >
            <BadgeAlertIcon className="mr-px h-5 w-5 fill-destructive text-destructive-foreground" />
          </div>
        </TooltipTrigger>
        <TooltipContent
          className="border-none bg-destructive text-sm text-destructive-foreground"
          onClick={() => void setOpen(false)}
        >
          <p>{message}</p>
        </TooltipContent>
      </Tooltip>
    </TooltipProvider>
  );
}
