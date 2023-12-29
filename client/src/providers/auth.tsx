import React, { createContext, useCallback, useState } from "react";
import { object, string, safeParse, number } from "valibot";

const UserMeResponse = object({
  id: number(),
  name: string(),
  image_url: string(),
});

interface User {
  name: string;
  imageURL: string;
}

type UserState =
  | {
      isAuthenticated: false;
      user: null;
    }
  | {
      isAuthenticated: true;
      user: User;
    };

type AuthContextProps = {
  state: UserState;
  fetchUser: () => Promise<User | null>;
};

export const AuthContext = createContext<AuthContextProps>({
  state: {
    user: null,
    isAuthenticated: false,
  },
  // eslint-disable-next-line @typescript-eslint/require-await
  fetchUser: async () => {
    return null;
  },
});

interface AuthProviderProps {
  children: React.ReactNode;
}

export const AuthProvider = ({ children }: AuthProviderProps) => {
  const [state, setState] = useState<UserState>({
    user: null,
    isAuthenticated: false,
  });

  const fetchUser = useCallback(async () => {
    const res = await fetch("/api/user/me");
    if (!res.ok) {
      setState({ user: null, isAuthenticated: false });
      return null;
    }

    const validated = safeParse(UserMeResponse, await res.json());
    if (!validated.success) {
      setState({ user: null, isAuthenticated: false });
      return null;
    }

    const user = {
      name: validated.output.name,
      imageURL: validated.output.image_url,
    };
    setState({
      user,
      isAuthenticated: true,
    });

    return user;
  }, [setState]);

  return (
    <AuthContext.Provider
      value={{
        state,
        fetchUser,
      }}
    >
      {children}
    </AuthContext.Provider>
  );
};
