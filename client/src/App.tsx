import { Loading } from "./components/loading/loading";
import { OnboardingPage } from "./features/onboarding";
import { LoadingProvider } from "./providers/loading";

function App() {
  return (
    <LoadingProvider>
      <Loading />
      <OnboardingPage />
    </LoadingProvider>
  );
}

export default App;
