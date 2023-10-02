import { getSearchArray } from "@/lib/utils";
import { PostPanel } from "@/components/post-panel";
import { SearchPanel } from "@/components/search-panel";

/* DUMMY */
type Post = {
  topic: string;
  period: string;
  positions: string[];
  skills: string[];
  benefits: string[];
};

const posts: Post[] = [
  {
    topic: "Social Engineering",
    period: "Summer 2022",
    positions: ["Software Developer", "Data Analyst"],
    skills: ["Python", "SQL", "Data Analysis"],
    benefits: ["Healthcare", "Flexible Work Hours"],
  },
  {
    topic: "Marketing",
    period: "Summer 2022",
    positions: ["Marketing Coordinator", "Social Media Manager"],
    skills: [
      "Digital Marketing",
      "Content Creation",
      "Social Media Management",
    ],
    benefits: ["401(k) Matching", "Paid Time Off"],
  },
  {
    topic: "Business",
    period: "Summer 2022",
    positions: ["Mechanical Engineer", "Product Designer"],
    skills: ["CAD Design", "Mechanical Engineering", "Product Prototyping"],
    benefits: ["Dental Insurance", "Tuition Reimbursement"],
  },
  {
    topic: "Customer Service",
    period: "Summer 2022",
    positions: ["Customer Support Specialist", "Sales Representative"],
    skills: ["Customer Service", "Sales", "Problem Solving"],
    benefits: ["Remote Work Option", "Company Discounts"],
  },
  {
    topic: "Data Science",
    period: "Summer 2022",
    positions: ["Data Scientist", "Machine Learning Engineer"],
    skills: ["Machine Learning", "Data Modeling", "Python"],
    benefits: ["Stock Options", "Gym Membership"],
  },
  {
    topic: "Healthcare",
    period: "Summer 2022",
    positions: ["Nurse Practitioner", "Medical Laboratory Technician"],
    skills: ["Patient Care", "Lab Testing", "Medical Diagnosis"],
    benefits: ["Health Insurance", "Retirement Plan"],
  },
  {
    topic: "Education",
    period: "Summer 2022",
    positions: ["Graphic Designer", "Web Developer"],
    skills: ["Graphic Design", "HTML/CSS", "Adobe Creative Suite"],
    benefits: ["Flexible Spending Account", "Casual Dress Code"],
  },
  {
    topic: "Accounting",
    period: "Summer 2022",
    positions: ["Accountant", "Financial Analyst"],
    skills: ["Accounting", "Financial Analysis", "Excel"],
    benefits: ["Paid Holidays", "Professional Development"],
  },
  {
    topic: "Human Resources",
    period: "Summer 2022",
    positions: ["Human Resources Manager", "Recruiter"],
    skills: ["Recruitment", "Employee Relations", "HR Policies"],
    benefits: ["Health Savings Account", "Employee Assistance Program"],
  },
  {
    topic: "Hospitality",
    period: "Summer 2022",
    positions: ["Chef", "Restaurant Manager"],
    skills: ["Culinary Arts", "Food Safety", "Menu Planning"],
    benefits: ["Meal Discounts", "Paid Vacation"],
  },
];

type Search = {
  positions: string[];
  skills: string[];
  benefits: string[];
};

const getPosts = (search: Search) => {
  const { positions, skills, benefits } = search;
  const filteredPosts = posts.filter(
    (post) =>
      positions.every((word) =>
        post.positions.some((position) =>
          position.toLowerCase().includes(word.toLowerCase())
        )
      ) &&
      skills.every((word) =>
        post.skills.some((skill) =>
          skill.toLowerCase().includes(word.toLowerCase())
        )
      ) &&
      benefits.every((word) =>
        post.benefits.some((benefit) =>
          benefit.toLowerCase().includes(word.toLowerCase())
        )
      )
  );
  return filteredPosts;
};
/* DUMMY */

type PageProps = {
  searchParams: {
    companies?: string | string[];
    positions?: string | string[];
    skills?: string | string[];
    benefits?: string | string[];
  };
};

export default function Page({ searchParams }: PageProps) {
  const search: Search = {
    positions: getSearchArray(searchParams.positions),
    skills: getSearchArray(searchParams.skills),
    benefits: getSearchArray(searchParams.benefits),
  };

  const posts = getPosts(search);
  const postCount = posts.length;

  return (
    <main className="container relative flex flex-1 items-start gap-12">
      <aside className="sticky top-[5.5rem] h-[calc(100vh-5.5rem)] w-[14rem]">
        <SearchPanel postCount={postCount} />
      </aside>
      <div className="flex-1">
        <PostPanel posts={posts} />
      </div>
    </main>
  );
}
