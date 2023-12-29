import { ChannelRoutes } from "./features/channel/routes";
import { OnBoardingRoutes } from "./features/onboarding/routes";
import { useRoutes } from "react-router-dom";

const routes = [
  { path: "/channel/*", element: <ChannelRoutes basePath="/channel" /> },
  { path: "/*", element: <OnBoardingRoutes basePath="/" /> },
];

export const AppRoutes = () => {
  const element = useRoutes([...routes]);

  return <>{element}</>;
};
