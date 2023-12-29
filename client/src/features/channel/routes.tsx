import { AuthRequired } from "@/libs/router";
import { lazy } from "react";
import { Route, Routes } from "react-router-dom";

const ChannelTop = lazy(() => import("./pages/top"));

export const CHANNEL_TOP_ROUTE = "/channel";

interface RoutesProps {
  basePath: string;
}

export const ChannelRoutes = ({ basePath }: RoutesProps) => {
  const trimUnderPath = (path: string) => {
    return path.replace(basePath, "");
  };

  return (
    <Routes>
      <Route
        path={trimUnderPath(CHANNEL_TOP_ROUTE)}
        element={
          <AuthRequired>
            <ChannelTop />
          </AuthRequired>
        }
      />
    </Routes>
  );
};
