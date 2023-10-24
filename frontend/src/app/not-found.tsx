import Image from "next/image";
import Link from "next/link";
import { getServerSession } from "@/lib/auth";
import { Button } from "@/components/ui/button";

export default async function NotFound() {
  const session = await getServerSession();
  return (
    <main className="container flex min-h-screen flex-col items-center justify-center gap-8">
      <Image
        className="aspect-[1477/1030] w-full max-w-lg"
        src="/images/not-found.png"
        alt="Not found"
        height={1030}
        width={1477}
        priority
      />
      <p className="text-center text-base sm:text-lg">
        <span className="block font-extrabold sm:inline">404</span>&nbsp;This
        page could not be found.
      </p>
      <Button className="h-auto p-0" variant="link" asChild>
        {session ? (
          <Link href="/">Back to home</Link>
        ) : (
          <Link href="/login">Login</Link>
        )}
      </Button>
    </main>
  );
}
