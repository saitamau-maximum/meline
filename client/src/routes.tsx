import { ChannelRoutes } from "./features/channel/routes";
import { OnBoardingRoutes } from "./features/onboarding/routes";
import { useRoutes } from "react-router-dom";

const routes = [
  { path: "/channel/*", element: <ChannelRoutes /> },
  { path: "/*", element: <OnBoardingRoutes /> },
];

export const AppRoutes = () => {
  const element = useRoutes([...routes]);

  return <>{element}</>;
};
