import { UnAuthRequired } from "@/libs/router";
import { lazy } from "react";
import { Route, Routes } from "react-router-dom";

const Top = lazy(() => import("./pages/top"));

export const TOP_ROUTE = "/";

interface RoutesProps {
  basePath: string;
}

export const OnBoardingRoutes = ({ basePath }: RoutesProps) => {
  const trimUnderPath = (path: string) => {
    return path.replace(basePath, "");
  };

  return (
    <Routes>
      <Route
        path={trimUnderPath(TOP_ROUTE)}
        element={
          <UnAuthRequired>
            <Top />
          </UnAuthRequired>
        }
      />
    </Routes>
  );
};
