interface AuthUser {
  name: string;
  imageURL: string;
}

export interface IAuthUserRepository {
  getAuthUser: () => Promise<AuthUser | null>;
  getAuthUser$$key: () => string[];
}

export class AuthUserRepositoryImpl implements IAuthUserRepository {
  async getAuthUser() {
    const res = await fetch("/api/user/me");

    if (!res.ok) {
      return null;
    }

    return res.json();
  }

  getAuthUser$$key() {
    return ["getAuthUser"];
  }
}
