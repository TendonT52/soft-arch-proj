"use client";

import { useEffect, useState } from "react";
import { format } from "date-fns";
import { Calendar as CalendarIcon } from "lucide-react";
import { type DateRange } from "react-day-picker";
import { cn } from "@/lib/utils";
import { Button } from "@/components/ui/button";
import { Calendar } from "@/components/ui/calendar";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";

type DatePickerWithRangeProps = React.HTMLAttributes<HTMLButtonElement> & {
  date?: DateRange;
  onDateChange?: (date?: DateRange) => void;
};

const DatePickerWithRange = ({
  date,
  onDateChange,
  className,
  ...props
}: DatePickerWithRangeProps) => {
  const [_date, _setDate] = useState<DateRange | undefined>(date);

  useEffect(() => {
    _setDate(date);
  }, [date]);

  useEffect(() => {
    onDateChange?.(_date);
  }, [_date, date, onDateChange]);

  return (
    <Popover>
      <PopoverTrigger asChild>
        <Button
          variant={"outline"}
          className={cn(
            "justify-start text-left font-normal",
            !date && "text-muted-foreground",
            className
          )}
          {...props}
        >
          <CalendarIcon className="mr-2 h-4 w-4" />
          <span className="block">
            {date?.from
              ? date.to
                ? `${format(date.from, "LLL dd, y")} - 
                  ${format(date.to, "LLL dd, y")}`
                : format(date.from, "LLL dd, y")
              : "Date range"}
          </span>
        </Button>
      </PopoverTrigger>
      <PopoverContent className="w-auto p-0" align="start">
        <Calendar
          initialFocus
          mode="range"
          defaultMonth={date?.from}
          selected={date ?? _date}
          onSelect={_setDate}
          numberOfMonths={2}
        />
      </PopoverContent>
    </Popover>
  );
};

export { DatePickerWithRange };
