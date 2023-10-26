import { BrowserRouter } from "react-router-dom";
import { LoadingOverlay } from "./components/loading-overlay/loading-overlay";
import { LoadingOverlayProvider } from "./providers/loading-overlay";
import { AppRoutes } from "./routes";
import { Suspense } from "react";

function App() {
  return (
    <LoadingOverlayProvider>
      <LoadingOverlay />
      <BrowserRouter>
        <Suspense>
          <AppRoutes />
        </Suspense>
      </BrowserRouter>
    </LoadingOverlayProvider>
  );
}

export default App;
