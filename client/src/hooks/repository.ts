import { RepositoryContext } from "@/providers/repository";
import { useContext } from "react";

export const useRepositories = () => {
  return useContext(RepositoryContext);
};
