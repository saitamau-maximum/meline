import { AuthContext } from "@/providers/auth";
import { ChannelTopPageTemplate } from "./template";
import { useContext } from "react";
import { ChannelContext } from "@/providers/channel";

export const ChannelTopPage = () => {
  const { state } = useContext(AuthContext);
  const { joinedChannels } = useContext(ChannelContext);

  if (!state.isAuthenticated) return null;

  return <ChannelTopPageTemplate user={state.user} channels={joinedChannels} />;
};
