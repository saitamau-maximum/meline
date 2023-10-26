import { LoadingOverlay } from "./components/loading-overlay/loading-overlay";
import { OnboardingPage } from "./features/onboarding";
import { LoadingOverlayProvider } from "./providers/loading-overlay";

function App() {
  return (
    <LoadingOverlayProvider>
      <LoadingOverlay />
      <OnboardingPage />
    </LoadingOverlayProvider>
  );
}

export default App;
