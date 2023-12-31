"use client";

import React, { useState } from "react";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { deletePost } from "@/actions/delete-post";
import { Loader2Icon, MoreVerticalIcon, TrashIcon } from "lucide-react";
import { type Post } from "@/types/base/post";
import { formatDate } from "@/lib/utils";
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
import { Button } from "./ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "./ui/dropdown-menu";
import { useToast } from "./ui/toaster";

type PostItemProps = {
  post: Post & {
    postId: string;
    updatedAt: string;
  };
};

const PostItem = ({ post }: PostItemProps) => {
  const router = useRouter();
  const { toast } = useToast();

  const [showDeleteDialog, setShowDeleteDialog] = useState(false);
  const [deleting, setDeleting] = useState(false);

  const handleDelete = async (e: React.MouseEvent) => {
    e.preventDefault();
    setDeleting(true);
    const response = await deletePost(post.postId);
    if (response.status === "200") {
      toast({
        title: "Success",
        description: response.message,
      });
      router.refresh();
      setShowDeleteDialog(false);
    } else {
      toast({
        title: "Error",
        description: response.message,
        variant: "destructive",
      });
    }
    setDeleting(false);
  };

  return (
    <div className="flex items-center justify-between p-4">
      <div className="flex flex-col items-start gap-1">
        <Link
          href={`/editor/${post.postId}`}
          className="font-semibold hover:underline"
        >
          {post.topic}
        </Link>
        <p className="text-sm text-muted-foreground">
          {formatDate(parseInt(post.updatedAt) * 1000)}
        </p>
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
            <Link href={`/editor/${post.postId}`} className="flex w-full">
              Edit
            </Link>
          </DropdownMenuItem>
          <DropdownMenuItem
            className="flex cursor-pointer items-center text-destructive focus:text-destructive"
            onSelect={() => void setShowDeleteDialog(true)}
          >
            Delete
          </DropdownMenuItem>
        </DropdownMenuContent>
      </DropdownMenu>
      <AlertDialog open={showDeleteDialog} onOpenChange={setShowDeleteDialog}>
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
            <Button variant="destructive" disabled={deleting} asChild>
              <AlertDialogAction onClick={(...a) => void handleDelete(...a)}>
                {deleting ? (
                  <Loader2Icon className="mr-2 h-4 w-4 animate-spin" />
                ) : (
                  <TrashIcon className="mr-2 h-4 w-4" />
                )}
                Delete
              </AlertDialogAction>
            </Button>
          </AlertDialogFooter>
        </AlertDialogContent>
      </AlertDialog>
    </div>
  );
};

export { PostItem };
