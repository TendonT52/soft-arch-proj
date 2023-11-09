"use client";

import { useState } from "react";
import createReport from "@/actions/create-report";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Button } from "./ui/button";
import { Input } from "./ui/input";
import { Textarea } from "./ui/textarea";
import { toast } from "./ui/toaster";

export default function ReportCard() {
  const [topic, setTopic] = useState<string>("");
  const [category, setCategory] = useState<string>(
    "Select a verified email to display"
  );
  const [description, setDescription] = useState<string>("");
  const onSubmit = async () => {
    const response = await createReport({
      report: {
        topic: topic,
        type: category,
        description: description,
      },
    });
    if (["200", "201"].includes(response.status)) {
      toast({
        title: "Success",
        description: response.message,
      });
    } else {
      toast({
        title: "Error",
        description: response.message,
        variant: "destructive",
      });
    }
  };

  return (
    <div className="m-4 h-[420px] w-[520px] rounded-md bg-background">
      <div className="items-left m-3 flex  flex-col gap-4">
        <div className="mb-2 flex flex-col">
          <h1 className="text-center text-xl font-semibold text-primary">
            Report
          </h1>
        </div>
        <div className="mb-2 flex flex-col">
          <div className="flex h-full w-20  justify-start">Title</div>
          <Input
            type="text"
            value={topic}
            className="my-1 flex w-full rounded-lg border-2 border-solid border-slate-500 px-1"
            placeholder="Report topic"
            onChange={(e) => setTopic(e.target.value)}
          ></Input>
        </div>
        <div className="mb-2 flex flex-col">
          <div className=" flex h-full w-20 justify-start">Category</div>
          <Select onValueChange={setCategory}>
            <SelectTrigger className="my-2 rounded-lg border-2 border-solid border-slate-500">
              <SelectValue placeholder="Select a verified email to display" />
            </SelectTrigger>
            <SelectContent>
              <SelectItem value="Scam And Fraudulent Listing">
                Scam And Fraudulent Listing
              </SelectItem>
              <SelectItem value="Fake Review">Fake Review</SelectItem>
              <SelectItem value="Suspicious User">Suspicious User</SelectItem>
              <SelectItem value="Website Bug">Website Bug</SelectItem>
              <SelectItem value="Suggestion">Suggestion</SelectItem>
              <SelectItem value="Other">Other</SelectItem>
            </SelectContent>
          </Select>
        </div>
        <div className="mb-2 flex flex-col">
          <div className="flex h-full w-20 justify-start">Description</div>
          <Textarea
            value={description}
            className="my-1 flex h-12 w-full rounded-lg border-2 border-solid border-slate-500 px-1"
            onChange={(e) => setDescription(e.target.value)}
          ></Textarea>
        </div>
        <Button onClick={() => void onSubmit()}>Submit</Button>
      </div>
    </div>
  );
}
