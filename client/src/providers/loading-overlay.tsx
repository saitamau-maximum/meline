import React, { createContext, useState } from "react";

interface LoadingOverlayContextProps {
  isLoading: boolean;
  setIsLoading: (isLoading: boolean) => void;
}

export const LoadingOverlayContext = createContext<LoadingOverlayContextProps>({
  isLoading: false,
  setIsLoading: () => {},
});

interface LoadingProviderProps {
  children: React.ReactNode;
}

export const LoadingOverlayProvider = ({ children }: LoadingProviderProps) => {
  const [isLoading, setIsLoading] = useState(false);

  return (
    <LoadingOverlayContext.Provider value={{ isLoading, setIsLoading }}>
      {children}
    </LoadingOverlayContext.Provider>
  );
};
