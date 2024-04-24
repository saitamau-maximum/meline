interface AuthUser {
  name: string;
  imageURL: string;
}

export interface IAuthUserRepository {
  getAuthUser: () => Promise<AuthUser | null>;
  getAuthUser$$key: () => string[];
}

export const AuthUserRepositoryImpl: IAuthUserRepository = {
  getAuthUser: async () => {
    const res = await fetch("/api/user/me");

    if (!res.ok) {
      return null;
    }

    return res.json();
  },
  getAuthUser$$key: () => ["getAuthUser"],
};
