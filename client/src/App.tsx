import { BrowserRouter } from "react-router-dom";
import { LoadingOverlay } from "./components/loading-overlay/loading-overlay";
import { AppRoutes } from "./routes";
import { Suspense, useCallback, useContext, useEffect } from "react";
import {
  LoadingOverlayContext,
  LoadingOverlayProvider,
} from "./providers/loading-overlay.tsx";
import { AuthContext, AuthProvider } from "./providers/auth.tsx";
import { ChannelContext, ChannelProvider } from "./providers/channel.tsx";

interface AppRootProps {
  children: React.ReactNode;
}

const AppRoot = ({ children }: AppRootProps) => {
  const { setIsLoading } = useContext(LoadingOverlayContext);
  const { fetchUser } = useContext(AuthContext);
  const { fetchJoinedChannels } = useContext(ChannelContext);

  const setupApplication = useCallback(async () => {
    setIsLoading(true);
    try {
      const user = await fetchUser();
      if (user) {
        await fetchJoinedChannels();
      }
    } finally {
      setIsLoading(false);
    }
  }, [fetchUser, fetchJoinedChannels, setIsLoading]);

  useEffect(() => {
    void setupApplication();
  }, []);

  return <div>{children}</div>;
};

const App = () => {
  return (
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
  );
};

export default App;
