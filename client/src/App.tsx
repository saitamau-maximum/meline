import { BrowserRouter } from "react-router-dom";
import { LoadingOverlay } from "./components/loading-overlay/loading-overlay";
import { AppRoutes } from "./routes";
import { Suspense, useCallback, useContext, useEffect } from "react";
import {
  LoadingOverlayContext,
  LoadingOverlayProvider,
} from "./providers/loading-overlay.tsx";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { AuthContext, AuthProvider } from "./providers/auth.tsx";
import { ChannelProvider } from "./providers/channel.tsx";

interface AppRootProps {
  children: React.ReactNode;
}

const AppRoot = ({ children }: AppRootProps) => {
  const { setIsLoading } = useContext(LoadingOverlayContext);
  const { fetchUser } = useContext(AuthContext);

  const setupApplication = useCallback(async () => {
    setIsLoading(true);
    try {
      await fetchUser();
    } finally {
      setIsLoading(false);
    }
  }, [fetchUser, setIsLoading]);

  useEffect(() => {
    void setupApplication();
  }, []);

  return <div>{children}</div>;
};

const App = () => {
  const queryClient = new QueryClient();

  return (
    <QueryClientProvider client={queryClient}>
      <AuthProvider>
        <ChannelProvider>
          <LoadingOverlayProvider>
            <LoadingOverlay />
            <AppRoot>
              <BrowserRouter>
                <Suspense>
                  <AppRoutes />
                </Suspense>
              </BrowserRouter>
            </AppRoot>
          </LoadingOverlayProvider>
        </ChannelProvider>
      </AuthProvider>
    </QueryClientProvider>
  );
};

export default App;
