import { useQuery } from "@tanstack/react-query";
import { useRepositories } from "./repository";

export const useAuthUser = () => {
  const { authUserRepository } = useRepositories();

  const { data, isLoading } = useQuery({
    queryKey: authUserRepository.getAuthUser$$key(),
    queryFn: () => authUserRepository.getAuthUser(),
  });

  const isAuthenticated = data !== null;

  return { data, isLoading, isAuthenticated };
};
