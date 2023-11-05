"use client";

import { useRouter } from "next/navigation";
import { updateCompany } from "@/actions/update-company";
import { zodResolver } from "@hookform/resolvers/zod";
import { Loader2Icon } from "lucide-react";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { type Company } from "@/types/base/company";
import { Button } from "./ui/button";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "./ui/form";
import { Input } from "./ui/input";
import { Textarea } from "./ui/textarea";
import { useToast } from "./ui/toaster";

const formSchema = z.object({
  name: z.string().min(1, { message: "Name is required" }),
  category: z.string().min(1, { message: "Category is required" }),
  location: z.string().min(1, { message: "Location is required" }),
  phone: z
    .string()
    .min(1, { message: "Phone number is required" })
    .regex(/^\d+$/, { message: "Phone number must be numerical" }),
  description: z.string().min(1, { message: "Description is required" }),
});

type FormData = z.infer<typeof formSchema>;

type CompanySettingsFormProps = {
  company: Company;
};

const CompanySettingsForm = ({ company }: CompanySettingsFormProps) => {
  const router = useRouter();
  const { toast } = useToast();

  const form = useForm<FormData>({
    resolver: zodResolver(formSchema),
    defaultValues: company,
    mode: "onChange",
  });
  const {
    formState: { isSubmitting },
    handleSubmit,
    control,
  } = form;

  const onSubmit = async (data: FormData) => {
    const response = await updateCompany({
      company: data,
    });
    if (response.status === "200") {
      toast({
        title: "Success",
        description: response.message,
      });
      router.refresh();
    } else {
      toast({
        title: "Error",
        description: response.message,
        variant: "destructive",
      });
    }
  };

  return (
    <Form {...form}>
      <form
        className="w-full max-w-2xl space-y-8"
        onSubmit={(...a) => void handleSubmit(onSubmit)(...a)}
      >
        <FormField
          name="name"
          control={control}
          render={({ field }) => (
            <FormItem>
              <FormLabel>First name</FormLabel>
              <FormControl>
                <Input placeholder="John Smith" {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          name="category"
          control={control}
          render={({ field }) => (
            <FormItem>
              <FormLabel>Category</FormLabel>
              <FormControl>
                <Input placeholder="Engineering" {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          name="location"
          control={control}
          render={({ field }) => (
            <FormItem>
              <FormLabel>Location</FormLabel>
              <FormControl>
                <Input placeholder="Computer" {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          name="phone"
          control={control}
          render={({ field }) => (
            <FormItem>
              <FormLabel>Telephone number</FormLabel>
              <FormControl>
                <Input placeholder="Computer" {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          name="description"
          control={control}
          render={({ field }) => (
            <FormItem>
              <FormLabel>Description</FormLabel>
              <FormControl>
                <Textarea placeholder="I own a building." {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <Button size="sm" disabled={isSubmitting} type="submit">
          {isSubmitting && (
            <Loader2Icon className="mr-2 h-4 w-4 animate-spin" />
          )}
          Update profile
        </Button>
      </form>
    </Form>
  );
};

export { CompanySettingsForm };
