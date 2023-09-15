import Image from "next/image";

export default function Home() {
  return (
    <main className="flex min-h-screen flex-col items-center justify-center gap-6 p-24">
      <Image
        src="/next.svg"
        alt="Next.js Logo"
        width={180}
        height={37}
        priority
      />
      <p className="text-center">
        Get <i className="text-muted-foreground italic">started</i> by
        editing&nbsp;
        <code className="font-mono font-bold">app/page.tsx</code>
      </p>
    </main>
  );
}
