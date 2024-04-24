import { useQuery } from "@tanstack/react-query";
import { AuthUserRepositoryImpl } from "@/repositories/auth-user";

export const useAuthUser = () => {
  const { data, isLoading } = useQuery({
    queryKey: AuthUserRepositoryImpl.getAuthUser$$key(),
    queryFn: () => AuthUserRepositoryImpl.getAuthUser(),
  });

  const isAuthenticated = data !== null;

  return { data, isLoading, isAuthenticated };
};
