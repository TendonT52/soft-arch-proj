"use client";

import { useState } from "react";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { deleteReview } from "@/actions/delete-review";
import { LexicalComposer } from "@lexical/react/LexicalComposer";
import { ContentEditable } from "@lexical/react/LexicalContentEditable";
import LexicalErrorBoundary from "@lexical/react/LexicalErrorBoundary";
import { RichTextPlugin } from "@lexical/react/LexicalRichTextPlugin";
import { format } from "date-fns";
import { Loader2Icon, MoreVerticalIcon, TrashIcon } from "lucide-react";
import { type Review } from "@/types/base/review";
import { UserRole, type User } from "@/types/base/user";
import { editorConfig } from "@/lib/lexical";
import { EditorStateNotifierPlugin } from "./lexical/editor-state-notifier-plugin";
import { Rating } from "./rating";
import { ReviewDialog } from "./review-dialog";
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
import { ScrollArea } from "./ui/scroll-area";
import { useToast } from "./ui/toaster";

type ReviewCardProps = {
  user: User;
  review: Review & {
    id: string;
    updatedAt: string;
    owner: {
      id?: string;
      name: string;
    };
  };
  companyId: string;
  companyName: string;
};

const ReviewCard = ({
  user,
  review,
  companyId,
  companyName,
}: ReviewCardProps) => {
  const router = useRouter();
  const { toast } = useToast();

  const owner =
    user.role === UserRole.Admin ? false : user.id === review.owner.id;
  const [showReviewDialog, setShowReviewDialog] = useState(false);
  const [showDeleteDialog, setShowDeleteDialog] = useState(false);
  const [deleting, setDeleting] = useState(false);

  const handleDelete = async (e: React.MouseEvent) => {
    e.preventDefault();
    setDeleting(true);
    const response = await deleteReview(review.id);
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
    <div className="flex flex-col rounded-xl border">
      <LexicalComposer
        initialConfig={{
          ...editorConfig,
          editorState: review.description,
          editable: false,
        }}
      >
        <div className="relative flex flex-col items-start rounded-t-xl pb-6">
          <div className="flex w-full gap-6">
            <div className="flex flex-1 flex-col items-start gap-4 overflow-x-hidden pl-8 pt-6">
              {owner ? (
                <ReviewDialog
                  review={review}
                  companyId={companyId}
                  companyName={companyName}
                  open={showReviewDialog}
                  onOpenChange={setShowReviewDialog}
                />
              ) : (
                <Rating rating={review.rating} editable={false} />
              )}
              <p className="truncate text-sm text-muted-foreground">
                <Link
                  className="hover:underline hover:underline-offset-2 focus-visible:underline focus-visible:underline-offset-2 focus-visible:outline-none"
                  href={`/students/${review.owner.id}`}
                >
                  {review.owner.name}
                </Link>
                ,&nbsp;
                {format(parseInt(review.updatedAt) * 1000, "MM/dd/yyyy")}
              </p>
            </div>
            <div className="pr-4 pt-4">
              <DropdownMenu>
                <DropdownMenuTrigger asChild>
                  <Button variant="ghost" className="h-8 w-8 rounded-md p-0">
                    <MoreVerticalIcon className="h-4 w-4" />
                    <span className="sr-only">Open</span>
                  </Button>
                </DropdownMenuTrigger>
                <DropdownMenuContent align="end">
                  {!owner && <DropdownMenuItem>Report</DropdownMenuItem>}
                  {owner && (
                    <DropdownMenuItem
                      onSelect={() => void setShowReviewDialog(true)}
                    >
                      Edit
                    </DropdownMenuItem>
                  )}
                  {owner && (
                    <DropdownMenuItem
                      className="flex cursor-pointer items-center text-destructive focus:text-destructive"
                      onSelect={() => void setShowDeleteDialog(true)}
                    >
                      Delete
                    </DropdownMenuItem>
                  )}
                </DropdownMenuContent>
              </DropdownMenu>
              <AlertDialog
                open={showDeleteDialog}
                onOpenChange={setShowDeleteDialog}
              >
                <AlertDialogContent>
                  <AlertDialogHeader>
                    <AlertDialogTitle>
                      Are you sure you want to delete this review?
                    </AlertDialogTitle>
                    <AlertDialogDescription>
                      This action cannot be undone.
                    </AlertDialogDescription>
                  </AlertDialogHeader>
                  <AlertDialogFooter>
                    <AlertDialogCancel>Cancel</AlertDialogCancel>
                    <Button variant="destructive" disabled={deleting} asChild>
                      <AlertDialogAction
                        onClick={(...a) => void handleDelete(...a)}
                      >
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
          </div>
          <h3
            className="flex w-full max-w-none appearance-none items-end bg-transparent px-8 pb-[0.6em] pt-[1.6em] text-xl font-semibold leading-[1.6] tracking-tight text-[#111827] scrollbar-hide focus:outline-none focus-visible:outline-none dark:text-[#ffffff]"
            aria-label="title"
          >
            {review.title}
          </h3>
          <div className="relative w-full">
            <RichTextPlugin
              contentEditable={
                <ScrollArea className="h-32">
                  <ContentEditable
                    className="prose prose-green relative max-w-none flex-1 overflow-auto px-8 scrollbar-hide dark:prose-invert focus:outline-none prose-p:m-0"
                    spellCheck={false}
                  />
                </ScrollArea>
              }
              placeholder={
                <div className="prose prose-green pointer-events-none absolute left-0 right-0 top-0 max-w-none px-8 pb-8 prose-p:m-0">
                  <p className="text-muted-foreground">
                    Enter review description...
                  </p>
                </div>
              }
              ErrorBoundary={LexicalErrorBoundary}
            />
          </div>
        </div>
        <EditorStateNotifierPlugin editorState={review.description} />
      </LexicalComposer>
    </div>
  );
};

export { ReviewCard };
