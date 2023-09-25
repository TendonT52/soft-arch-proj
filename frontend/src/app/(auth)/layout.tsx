import Image from "next/image";

type LayoutProps = {
  children: React.ReactNode;
};

export default function Layout({ children }: LayoutProps) {
  return (
    <div className="flex min-h-screen">
      <div className="hidden flex-1 items-center justify-center bg-primary lg:flex">
        <Image
          className="aspect-[646/531] w-1/2"
          src="/images/register.png"
          alt="Register"
          height={531}
          width={646}
        />
      </div>
      {children}
    </div>
  );
}
