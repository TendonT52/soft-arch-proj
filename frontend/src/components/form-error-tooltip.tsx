"use client";

import React, { useState } from "react";
import { BadgeAlertIcon } from "lucide-react";
import { cn } from "@/lib/utils";
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from "@/components/ui/tooltip";

type FormErrorTooltipProps = React.HTMLAttributes<HTMLDivElement> & {
  message?: string;
};

const FormErrorTooltip = ({
  message,
  className,
  ...props
}: FormErrorTooltipProps) => {
  const [open, setOpen] = useState(false);

  if (message === undefined) {
    return <></>;
  }
  return (
    <TooltipProvider>
      <Tooltip delayDuration={200} open={open} onOpenChange={setOpen}>
        <TooltipTrigger asChild>
          <div
            className={cn(
              "flex h-5 w-5 cursor-pointer items-center overflow-hidden  focus-visible:outline-none",
              className
            )}
            onClick={() => void setOpen(!open)}
            {...props}
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
};

export { FormErrorTooltip };
