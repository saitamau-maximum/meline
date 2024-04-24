import { ChannelTopPageTemplate } from "./template";

import { useJoinedChannels } from "./hooks/use-joined-channels";
import { useAuthUser } from "@/hooks/auth-user";

export const ChannelTopPage = () => {
  const { data: authUser } = useAuthUser();
  const { data, isLoading } = useJoinedChannels();

  if (!authUser) return null;

  return (
    <ChannelTopPageTemplate
      user={authUser}
      channels={data?.channels || []}
      isChannelsLoading={isLoading}
    />
  );
};
