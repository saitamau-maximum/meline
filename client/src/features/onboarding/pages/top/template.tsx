import { AuthPane } from "./components/auth-pane";
import { FeaturePane } from "./components/feature-pane";
import { Wrapper } from "./components/wrapper";

export const OnboardingPageTemplate = () => {
  return (
    <Wrapper>
      <FeaturePane />
      <AuthPane />
    </Wrapper>
  );
};
