import { AuthContext } from "@/providers/auth";
import { ChannelTopPageTemplate } from "./template";
import { useContext } from "react";
import { ChannelContext } from "@/providers/channel";
import { ChannelRepository } from "@/repositories/channel";

export const ChannelTopPage = () => {
  const { state } = useContext(AuthContext);
  const { joinedChannels, fetchJoinedChannels } = useContext(ChannelContext);

  if (!state.isAuthenticated) return null;

  return (
    <ChannelTopPageTemplate
      user={state.user}
      channels={joinedChannels}
      fetchJoinedChannels={fetchJoinedChannels}
      channelRepository={new ChannelRepository()}
    />
  );
};
