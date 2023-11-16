import { Login } from "./components/login";
import { Title } from "./components/title";
import { Wrapper } from "./components/wrapper";

export const OnboardingPage = () => {
  return (
    <Wrapper>
      <Title />
      <Login />
    </Wrapper>
  );
};
