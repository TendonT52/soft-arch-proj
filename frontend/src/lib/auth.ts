import { getCompanyMe } from "@/actions/get-company-me";
import { getStudentMe } from "@/actions/get-student-me";
import { login } from "@/actions/login";
import { refresh } from "@/actions/refresh";
import {
  getServerSession as nextAuthGetServerSession,
  type AuthOptions,
} from "next-auth";
import CredentialsProvider from "next-auth/providers/credentials";
import { UserRole } from "@/types/base/user";
import { validateAccessToken, verifyAccessToken } from "./token";

async function lookupUser(accessToken: string, role: UserRole) {
  switch (role) {
    case UserRole.Company:
      const { status: cStatus, company } = await getCompanyMe(accessToken);
      if (cStatus !== "200") return null;
      return { ...company, role };

    case UserRole.Student:
      const { status: sStatus, student } = await getStudentMe(accessToken);
      if (sStatus !== "200") return null;
      return { ...student, role };

    default:
      return null;
  }
}

export const authOptions: AuthOptions = {
  session: {
    strategy: "jwt",
  },
  jwt: {
    maxAge: 60 * 60, // 1 hour
  },
  providers: [
    CredentialsProvider({
      // The name to display on the sign in form (e.g. 'Sign in with...')
      name: "Credentials",
      // The credentials is used to generate a suitable form on the sign in page.
      // You can specify whatever fields you are expecting to be submitted.
      // e.g. domain, username, password, 2FA token, etc.
      // You can pass any HTML attribute to the <input> tag through the object.
      credentials: {
        email: { label: "Email", type: "text", placeholder: "jsmith" },
        password: { label: "Password", type: "password" },
      },
      async authorize(credentials) {
        // You need to provide your own logic here that takes the credentials
        // submitted and returns either a object representing a user or value
        // that is false/null if the credentials are invalid.
        // e.g. return { id: 1, name: 'J Smith', email: 'jsmith@example.com' }
        // You can also use the `req` object to obtain additional parameters
        // (i.e., the request IP address)
        if (!credentials) return null;
        const { email, password } = credentials;
        const { status, accessToken, refreshToken } = await login({
          email,
          password,
        });

        if (status !== "200") return null;
        return { accessToken, refreshToken };
      },
    }),
  ],
  callbacks: {
    async jwt({ token, user }) {
      // Persist the OAuth access_token and or the user id to the token right after signin
      if (!user) {
        const { accessToken, refreshToken } = token;
        const valid = validateAccessToken(accessToken);
        if (!valid) {
          const { status, accessToken } = await refresh({ refreshToken });
          if (status !== "200") {
            throw new Error("Failed to refresh token");
          }
          token.accessToken = accessToken;
        }
      }
      return { ...token, ...user };
    },
    async session({ session, token }) {
      // Send properties to the client, like an access_token and user id from a provider.
      const { accessToken, refreshToken } = token;
      const { role } = verifyAccessToken(accessToken);
      const user = await lookupUser(accessToken, role);
      if (!user) {
        throw new Error("Failed to get user");
      }
      return { ...session, user, accessToken, refreshToken };
    },
  },
};

export async function getServerSession() {
  return await nextAuthGetServerSession(authOptions);
}
