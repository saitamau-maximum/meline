import { OnBoardingRoutes } from "./features/onboarding/routes";
import { useRoutes } from "react-router-dom";

const routes = [{ path: "/*", element: <OnBoardingRoutes /> }];

export const AppRoutes = () => {
  const element = useRoutes([...routes]);

  return <>{element}</>;
};
