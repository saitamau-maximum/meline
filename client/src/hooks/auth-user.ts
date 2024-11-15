import { useQuery } from "@tanstack/react-query";
import { AuthUserRepositoryImpl } from "@/repositories/auth-user";
import { useMemo } from "react";

export const useAuthUser = () => {
  const authUserRepository = useMemo(() => new AuthUserRepositoryImpl(), []);

  const { data, isLoading } = useQuery({
    queryKey: authUserRepository.getAuthUser$$key(),
    queryFn: () => authUserRepository.getAuthUser(),
  });

  const isAuthenticated = data !== null;

  return { data, isLoading, isAuthenticated };
};
