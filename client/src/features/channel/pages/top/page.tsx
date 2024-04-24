import { ChannelTopPageTemplate } from "./template";

import { useAuthUser } from "@/hooks/auth-user";

export const ChannelTopPage = () => {
  const { data: authUser } = useAuthUser();

  if (!authUser) return null;

  return <ChannelTopPageTemplate user={authUser} />;
};
