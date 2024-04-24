import { AuthRequired } from "@/libs/router";
import { lazy } from "react";
import { Route, Routes } from "react-router-dom";
import { ChannelLayout } from "./components/layout";
import { useJoinedChannels } from "./hooks/use-joined-channels";

const ChannelTop = lazy(() => import("./pages/top"));
const ChannelDetail = lazy(() => import("./pages/detail"));

export const ChannelRoutes = () => {
  const { data, isLoading: isChannelsLoading } = useJoinedChannels();

  return (
    <Routes>
      <Route
        path="/"
        element={
          <AuthRequired>
            <ChannelLayout
              channels={data?.channels || []}
              isChannelsLoading={isChannelsLoading}
            />
          </AuthRequired>
        }
      >
        <Route index element={<ChannelTop />} />
        <Route path=":channelId" element={<ChannelDetail />} />
      </Route>
    </Routes>
  );
};
