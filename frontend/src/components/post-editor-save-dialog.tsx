"use client";

import { useState } from "react";
import { type SerializedEditorState } from "lexical";
import { DatePickerWithRange } from "./date-range-picker";
import { Button } from "./ui/button";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "./ui/dialog";
import { Input } from "./ui/input";
import { Label } from "./ui/label";
import { Textarea } from "./ui/textarea";

type PostEditorSaveDialogProps = {
  topic?: string;
  description?: SerializedEditorState;
};

const PostEditorSaveDialog = ({ topic }: PostEditorSaveDialogProps) => {
  const [period, setPeriod] = useState<string>();
  const [positions, setPositions] = useState<string>();
  const [skills, setSkills] = useState<string>();
  const [benefits, setBenefits] = useState<string>();

  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button className="h-10">Save</Button>
      </DialogTrigger>
      <DialogContent>
        <form className="flex w-full flex-col gap-4 p-1">
          <DialogHeader className="mb-2">
            <DialogTitle>Additional information</DialogTitle>
            <DialogDescription>
              Help students reach your post easily
            </DialogDescription>
          </DialogHeader>
          <div className="flex flex-col gap-6">
            <div className="flex w-full gap-4">
              <div className="flex flex-1 flex-col gap-2">
                <Label
                  className="w-full text-sm font-medium leading-none"
                  htmlFor="topic"
                >
                  Topic
                </Label>
                <Input id="topic" value={topic} readOnly />
              </div>
              <div className="flex flex-1 flex-col gap-2">
                <Label
                  htmlFor="period"
                  className="w-full text-sm font-medium leading-none"
                >
                  Period
                </Label>
                <DatePickerWithRange id="period" />
              </div>
            </div>
            <div className="flex w-full flex-col gap-2">
              <Label
                className="w-full text-sm font-medium leading-none"
                htmlFor="skills"
              >
                Required skills
                <span className="ml-4 font-normal text-muted-foreground">
                  *Space delimited
                </span>
              </Label>
              <Input id="skills" placeholder="SQL slamming" />
            </div>
            <div className="flex w-full flex-col gap-2">
              <Label
                className="w-full text-sm font-medium leading-none"
                htmlFor="positions"
              >
                Open positions
                <span className="ml-4 font-normal text-muted-foreground">
                  *Space delimited
                </span>
              </Label>
              <Input id="positions" placeholder="Top of the world" />
            </div>
            <div className="flex w-full flex-col gap-2">
              <Label
                className="w-full text-sm font-medium leading-none"
                htmlFor="benefits"
              >
                Benefits
                <span className="ml-4 font-normal text-muted-foreground">
                  *Space delimited
                </span>
              </Label>
              <Input id="benefits" placeholder="Coffee" />
            </div>
            <div className="flex flex-col gap-2">
              <Label
                className="text-sm font-medium leading-none"
                htmlFor="howTo"
              >
                How to apply
              </Label>
              <Textarea
                id="howTo"
                className="resize-none"
                placeholder="Run to the office like the flash ⚡"
              />
            </div>
          </div>
          {/* <FormErrorTooltip
              message={
                errors.ssn
                  ? errors.ssn.message
                  : errors.ethnicity
                  ? errors.ethnicity.message
                  : undefined
              }
            /> */}
          <DialogFooter className="mt-2 flex sm:justify-center">
            <Button size="sm" type="submit">
              Confirm
            </Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  );
};

export { PostEditorSaveDialog };
