import React, { createContext, useCallback, useEffect, useState } from "react";
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
  fetchUser: () => Promise<void>;
};

const AuthContext = createContext<AuthContextProps>({
  state: {
    user: null,
    isAuthenticated: false,
  },
  // eslint-disable-next-line @typescript-eslint/no-empty-function
  fetchUser: async () => {},
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
    if (!res.ok) return setState({ user: null, isAuthenticated: false });
    const validated = safeParse(UserMeResponse, await res.json());
    if (!validated.success)
      return setState({ user: null, isAuthenticated: false });
    const user = validated.output;
    setState({
      user: {
        name: user.name,
        imageURL: user.image_url,
      },
      isAuthenticated: true,
    });
  }, [setState]);

  useEffect(() => {
    void fetchUser();
  }, [fetchUser]);

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

export const useAuth = () => React.useContext(AuthContext);
