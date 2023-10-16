"use client";

import { signOut } from "next-auth/react";
import { Button } from "./ui/button";

const SignOutButton = () => {
  return (
    <Button variant="outline" size="sm" onClick={() => void signOut()}>
      Sign out
    </Button>
  );
};

export { SignOutButton };
