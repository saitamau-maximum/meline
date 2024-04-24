import { AuthContext } from "@/providers/auth";
import { ChannelTopPageTemplate } from "./template";
import { useContext } from "react";

import { useJoinedChannels } from "./hooks/use-joined-channels";

export const ChannelTopPage = () => {
  const { state } = useContext(AuthContext);
  const { data, isLoading } = useJoinedChannels();

  if (!state.isAuthenticated) return null;

  return (
    <ChannelTopPageTemplate
      user={state.user}
      channels={data?.channels || []}
      isChannelsLoading={isLoading}
    />
  );
};
