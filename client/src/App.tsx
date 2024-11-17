import { BrowserRouter } from "react-router-dom";
import { LoadingOverlay } from "./components/loading-overlay/loading-overlay";
import { AppRoutes } from "./routes";
import { Suspense } from "react";
import { LoadingOverlayProvider } from "./providers/loading-overlay";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { RepositoryProvider } from "./providers/repository";
import { useSetup } from "./hooks/setup";
import { ToastProvider } from "./components/ui/toast/toast-provider";

const AppRoot = () => {
  useSetup();

  return (
    <Suspense>
      <AppRoutes />
    </Suspense>
  );
};

const App = () => {
  const queryClient = new QueryClient();

  return (
    <RepositoryProvider>
      <QueryClientProvider client={queryClient}>
        <LoadingOverlayProvider>
          <LoadingOverlay />
          <BrowserRouter>
            <ToastProvider>
              <AppRoot />
            </ToastProvider>
          </BrowserRouter>
        </LoadingOverlayProvider>
      </QueryClientProvider>
    </RepositoryProvider>
  );
};

export default App;
