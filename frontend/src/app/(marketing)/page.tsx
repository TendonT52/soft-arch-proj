import Image from "next/image";
import { ModeToggle } from "@/components/mode-toggle";

export default function Home() {
  return (
    <main className="relative flex flex-1 flex-col items-center justify-center gap-6 p-24">
      <Image
        className="select-none dark:invert"
        src="/next.svg"
        alt="Next.js Logo"
        width={180}
        height={37}
        priority
      />
      <p className="text-center">
        Get <i className="italic text-primary">started</i> by editing&nbsp;
        <code className="font-mono font-bold">app/page.tsx</code>
      </p>
      <ModeToggle />
    </main>
  );
}
