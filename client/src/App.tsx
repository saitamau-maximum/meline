import { BrowserRouter } from "react-router-dom";
import { LoadingOverlay } from "./components/loading-overlay/loading-overlay";
import { AppRoutes } from "./routes";
import { Suspense } from "react";
import { LoadingOverlayProvider } from "./providers/loading-overlay.tsx";
import { AuthProvider } from "./providers/auth.tsx";

function App() {
  return (
    <AuthProvider>
      <LoadingOverlayProvider>
        <LoadingOverlay />
        <BrowserRouter>
          <Suspense>
            <AppRoutes />
          </Suspense>
        </BrowserRouter>
      </LoadingOverlayProvider>
    </AuthProvider>
  );
}

export default App;
