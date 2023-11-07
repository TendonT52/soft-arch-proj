import { type Admin } from "./admin";
import { type Company } from "./company";
import { type Student } from "./student";

export enum UserRole {
  Admin = "admin", // TODO
  Company = "company",
  Student = "student",
}

export type User =
  | ({
      role: UserRole.Company;
    } & Company)
  | ({ role: UserRole.Student } & Student)
  | ({ role: UserRole.Admin } & Admin);
