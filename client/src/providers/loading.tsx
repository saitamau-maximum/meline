import React, { createContext, useState } from "react";

interface LoadingContextProps {
  isLoading: boolean;
  setIsLoading: (isLoading: boolean) => void;
}

export const LoadingContext = createContext<LoadingContextProps>({
  isLoading: false,
  setIsLoading: () => {},
});

interface LoadingProviderProps {
  children: React.ReactNode;
}

export const LoadingProvider = ({ children }: LoadingProviderProps) => {
  const [isLoading, setIsLoading] = useState(false);

  return (
    <LoadingContext.Provider value={{ isLoading, setIsLoading }}>
      {children}
    </LoadingContext.Provider>
  );
};
