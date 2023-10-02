"use client";

import { useState } from "react";
import Link from "next/link";
import { MoreVerticalIcon, TrashIcon } from "lucide-react";
import { cn, formatDate } from "@/lib/utils";
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
} from "./ui/alert-dialog";
import { Button, buttonVariants } from "./ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "./ui/dropdown-menu";

/* DUMMY */
type Post = {
  topic: string;
  period: string;
  positions: string[];
  skills: string[];
  benefits: string[];
};
/* DUMMY */

type PostItemProps = {
  post: Post;
};

export function PostItem({ post }: PostItemProps) {
  const [showDeleteAlert, setShowDeleteAlert] = useState(false);

  return (
    <div className="flex items-center justify-between p-4">
      <div className="flex flex-col items-start gap-1">
        <Link href="/editor/1" className="font-semibold hover:underline">
          {post.topic}
        </Link>
        <div>
          <p className="text-sm text-muted-foreground">
            {formatDate(Date.now())}
          </p>
        </div>
      </div>
      <DropdownMenu>
        <DropdownMenuTrigger asChild>
          <Button variant="ghost" className="h-8 w-8 rounded-md p-0">
            <MoreVerticalIcon className="h-4 w-4" />
            <span className="sr-only">Open</span>
          </Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent align="end">
          <DropdownMenuItem>
            <Link href="/editor/1" className="flex w-full">
              Edit
            </Link>
          </DropdownMenuItem>
          <DropdownMenuItem
            className="flex cursor-pointer items-center text-destructive focus:text-destructive"
            onSelect={() => setShowDeleteAlert(true)}
          >
            Delete
          </DropdownMenuItem>
        </DropdownMenuContent>
      </DropdownMenu>
      <AlertDialog open={showDeleteAlert} onOpenChange={setShowDeleteAlert}>
        <AlertDialogContent>
          <AlertDialogHeader>
            <AlertDialogTitle>
              Are you sure you want to delete this post?
            </AlertDialogTitle>
            <AlertDialogDescription>
              This action cannot be undone.
            </AlertDialogDescription>
          </AlertDialogHeader>
          <AlertDialogFooter>
            <AlertDialogCancel>Cancel</AlertDialogCancel>
            <AlertDialogAction
              className={cn(buttonVariants({ variant: "destructive" }))}
              onClick={() => void setShowDeleteAlert(false)}
            >
              <TrashIcon className="mr-2 h-4 w-4" />
              Delete
            </AlertDialogAction>
          </AlertDialogFooter>
        </AlertDialogContent>
      </AlertDialog>
    </div>
  );
}