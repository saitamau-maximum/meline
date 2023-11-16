import { AuthPane } from "./components/authPane";
import { FeaturePane } from "./components/featurePane";
import { Wrapper } from "./components/wrapper";

export const OnboardingPage = () => {
  return (
    <Wrapper>
      <FeaturePane />
      <AuthPane />
    </Wrapper>
  );
};
