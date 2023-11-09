"use client";

import { useRouter } from "next/navigation";
import putCompanyStatus from "@/actions/put-company-status";
import { CheckCircle, X } from "lucide-react";
import { type Company } from "@/types/base/company";
import { Button } from "./ui/button";
import { toast } from "./ui/toaster";

type PendingItemProps = {
  company: Company;
};

const PendingCompany = ({ company }: PendingItemProps) => {
  const router = useRouter();

  const approved = async () => {
    const response = await putCompanyStatus({
      id: company.id,
      status: "Approve",
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
    router.refresh();
  };
  const reject = async () => {
    const response = await putCompanyStatus({
      id: company.id,
      status: "Reject",
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
    router.refresh();
  };
  return (
    <div className="flex items-center justify-between p-4">
      <div className="flex flex-col items-start gap-1">
        <div className="flex gap-2">
          <div className="font-semibold hover:underline">{company.name}</div>
        </div>
        <p className="text-md text-muted-foreground">{company.status}</p>
      </div>
      <div className="flex items-center justify-between p-4">
        <Button
          variant="ghost"
          className="mx-4 h-8 w-8 rounded-md p-0"
          onClick={() => void approved()}
          asChild
        >
          <CheckCircle className="h-4 w-4 text-primary" />
        </Button>
        <Button
          variant="ghost"
          className="h-8 w-8 rounded-md p-0"
          onClick={() => void reject()}
          asChild
        >
          <X className="h-4 w-4 text-red-500" />
        </Button>
      </div>
    </div>
  );
};

export { PendingCompany };
