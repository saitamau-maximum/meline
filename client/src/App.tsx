import { BrowserRouter } from "react-router-dom";
import { LoadingOverlay } from "./components/loading-overlay/loading-overlay";
import { AppRoutes } from "./routes";
import { Suspense, useContext, useEffect } from "react";
import {
  LoadingOverlayContext,
  LoadingOverlayProvider,
} from "./providers/loading-overlay.tsx";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { useAuthUser } from "./hooks/auth-user.ts";

const AppRoot = () => {
  const { setIsLoading } = useContext(LoadingOverlayContext);
  const { isLoading: isAuthUserLoading } = useAuthUser();

  useEffect(() => {
    setIsLoading(isAuthUserLoading);
  }, [isAuthUserLoading, setIsLoading]);

  return (
    <>
      <LoadingOverlay />
      <BrowserRouter>
        <Suspense>
          <AppRoutes />
        </Suspense>
      </BrowserRouter>
    </>
  );
};

const App = () => {
  const queryClient = new QueryClient();

  return (
    <QueryClientProvider client={queryClient}>
      <LoadingOverlayProvider>
        <AppRoot />
      </LoadingOverlayProvider>
    </QueryClientProvider>
  );
};

export default App;
