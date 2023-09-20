import Image from "next/image";
import {
  AlertCircleIcon,
  FileTextIcon,
  KeyIcon,
  SearchIcon,
  StarIcon,
  User2Icon,
} from "lucide-react";
import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import { SignUpOptionMenu } from "@/components/sign-up-option-menu";

type Feature = {
  icon: JSX.Element;
  title: string;
  description: string;
};

const features: Feature[] = [
  {
    icon: <KeyIcon className="h-12 w-12" />,
    title: "Authentication",
    description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
  },
  {
    icon: <User2Icon className="h-12 w-12" />,
    title: "Profile",
    description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
  },
  {
    icon: <FileTextIcon className="h-12 w-12" />,
    title: "Post",
    description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
  },
  {
    icon: <SearchIcon className="h-12 w-12" />,
    title: "Search",
    description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
  },
  {
    icon: <StarIcon className="h-12 w-12" />,
    title: "Review",
    description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
  },
  {
    icon: <AlertCircleIcon className="h-12 w-12" />,
    title: "Report",
    description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
  },
];

export default function Home() {
  return (
    <main className="container relative flex flex-col">
      <section className="mx-auto flex min-h-[calc(100vh-4rem)] max-w-lg flex-1 flex-col items-center gap-12 py-12 lg:max-w-none lg:flex-row lg:gap-0 lg:px-8 xl:gap-8 xl:px-12 2xl:min-h-0 2xl:py-40">
        <div className="flex flex-col items-center gap-8 lg:flex-1 lg:items-start">
          <h1 className="max-w-[80%] text-center text-4xl font-bold leading-none tracking-tight sm:max-w-none lg:text-start lg:text-5xl xl:text-6xl">
            Your Gateway to&nbsp;
            <br className="hidden sm:inline" />
            Professional Growth
          </h1>
          <p className="max-w-[80%] text-center text-lg sm:max-w-none lg:text-start lg:text-xl lg:leading-8">
            Take off on your career journey with us,&nbsp;
            <br className="hidden sm:inline" />
            simplifying the path to your ideal internship.
          </p>
          <div className="hidden lg:block">
            <SignUpOptionMenu align="end" side="right" direction="row">
              <Button>Get started</Button>
            </SignUpOptionMenu>
          </div>
          <div className="block lg:hidden">
            <SignUpOptionMenu align="center" direction="row">
              <Button>Get started</Button>
            </SignUpOptionMenu>
          </div>
        </div>
        <div className="relative lg:flex-1">
          <Image
            className="object-cover"
            src="/images/hero.png"
            alt="Hero"
            height={1042}
            width={1829}
            priority
          />
        </div>
      </section>
      <section className="mx-auto flex w-full max-w-[64rem] flex-col items-center gap-8 pb-24 pt-12 sm:px-8">
        <h2 className="text-center text-3xl font-semibold leading-none tracking-tight lg:text-4xl xl:text-5xl">
          Features
        </h2>
        <div className="grid w-full auto-cols-fr grid-cols-1 gap-8 md:grid-cols-2 xl:grid-cols-3">
          {features.map(({ icon, title, description }) => (
            <Card key={title}>
              <CardContent className="flex flex-col gap-3 p-6">
                {icon}
                <p className="text-base font-medium">{title}</p>
                <p className="text-sm text-muted-foreground">{description}</p>
              </CardContent>
            </Card>
          ))}
        </div>
      </section>
    </main>
  );
}
